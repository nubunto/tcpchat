package server

import (
	"bufio"
	"encoding/json"
	"github.com/nubunto/tcpchat/types"
	"log"
	"net"
	"fmt"
)

var connections *ConnectionList

func Start(ip string) error {
	log.Println("Starting server on", ip)
	ln, err := net.Listen("tcp", ip)
	if err != nil {
		return err
	}
	defer ln.Close()
	messages, errs := make(chan string), make(chan error)
	connections = new(ConnectionList)

	go connections.Broadcast(messages)
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		var user Connection

		jsonDecoder := json.NewDecoder(conn)

		if err = jsonDecoder.Decode(&user); err != nil {
			log.Println("Handshake err:", err, " -- dropping connection")
			continue
		}

		user.conn = conn
		jsonEncoder := json.NewEncoder(conn)
		hsresult := new(types.HandshakeResult)
		contains := connections.Contains(user.Name)
		if contains {
			hsresult.Message = "User already exists on the server"
			hsresult.Code = types.AlreadyExists
			if err = jsonEncoder.Encode(hsresult); err != nil {
				log.Println("Something went wrong communicating a handshake error. Dropping to next connection.")
			}
			continue
		} else {
			hsresult.Message = "OK"
			hsresult.Code = types.Ok
			if err = jsonEncoder.Encode(hsresult); err != nil {
				log.Println("Something went wrong communicating a handshake success. Dropping to next connection.")
				continue
			}
		}

		*connections = append(*connections, user)
		go handle(user, messages, errs)
		go logErrs(errs)
	}
	close(messages)

	return nil
}

func logErrs(errors chan error) {
	for err := range errors {
		log.Println("[SERVER ERROR]", err)
	}
}

func handle(serverConn Connection, messages chan<- string, errors chan<- error) {
	conn := serverConn.conn
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			errors <- err
			break
		} else {
			messages <- fmt.Sprintf("%s-> %s", serverConn.Name, message)
		}
		log.Println("Got message:", message)
	}
	log.Println("dropping connection from", serverConn.Name)
	connections.Remove(serverConn.Name)
}
