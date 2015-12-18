package server

func (cl *ConnectionList) Broadcast(messages chan string) {
	for m := range messages {
		for _, c := range *cl {
			c.conn.Write([]byte(m))
		}
	}
}

func (cl *ConnectionList) Remove(name string) {
	concrete := *cl
	for i, l := 0, len(*cl); i < l; i++ {
		cur := concrete[i]
		if cur.Name == name {
			*cl = append(concrete[:i], concrete[i+1:]...)
			return
		}
	}
}

func (cl ConnectionList) Contains(name string) bool {
	for i, l := 0, len(cl); i < l; i++ {
		if cl[i].Name == name {
			return true
		}
	}
	return false
}
