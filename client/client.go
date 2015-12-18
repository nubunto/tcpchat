package client

import (
	"bufio"
	"encoding/json"
	"github.com/nubunto/tcpchat/types"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func Connect(ip, name string) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}

	err = handshake(conn, name)
	if err != nil {
		log.Fatal(err)
	}
	intercept(os.Stdin, conn)
}

func handshake(conn net.Conn, name string) error {
	user := types.User{Name: name}

	jsonEncoder := json.NewEncoder(conn)
	jsonDecoder := json.NewDecoder(conn)

	var hsresult types.HandshakeResult
	if err := jsonEncoder.Encode(user); err != nil {
		return err
	}

	if err := jsonDecoder.Decode(&hsresult); err != nil {
		return err
	}

	if hsresult.Code == types.Ok {
		return nil
	}

	return hsresult
}

func intercept(reader io.Reader, conn net.Conn) {
	interfaceReader := bufio.NewReader(reader)
	go handleConn(conn)
	for {
		typed, err := interfaceReader.ReadString('\n')
		if err != nil {
			log.Println("[CLIENT ERROR]", err)
			continue
		}
		conn.Write([]byte(typed))
		time.Sleep(100 * time.Millisecond)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			log.Println("[CLIENT ERROR]", err)
			break
		}

		log.Println(message)
		time.Sleep(100 * time.Millisecond)
	}
}
