package kademlia

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
)

const (
	idLength = 200
)

type Node struct {
	NodeId [20]byte
}

func (n *Node) Distance(other *Node) (ret [20]byte) {
	for i := 0; i < 20; i++ {
		ret[i] = n.NodeId[i] ^ other.NodeId[i]
	}
	return
}

func (n *Node) String() string {
	myStr := []string{}
	for i := 0; i < 20; i++ {
		myStr = append(myStr, strconv.FormatInt(int64(n.NodeId[i]), 2))
	}
	return strings.Join(myStr, "-")
}

func (n *Node) makeId() {
	num := make([]byte, idLength)
	_, err := rand.Read(num)
	if err != nil {
		fmt.Println(err)
	}
	n.NodeId = sha1.Sum(num)
}

func MakeNode() *Node {
	n := &Node{}
	n.makeId()

	return n
}
