package kademlia

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	node1 := MakeNode()
	node2 := MakeNode()
	node3 := MakeNode()

	distance1 := node1.Distance(node2)
	distance2 := node2.Distance(node3)
	distance3 := node3.Distance(node1)

	dist1int1, _ := binary.Uvarint(distance1[:]) //cast array to slice
	dist1int2, _ := binary.Uvarint(distance2[:]) //cast array to slice
	dist1int3, _ := binary.Uvarint(distance3[:]) //cast array to slice

	fmt.Printf("Node1 is %x", node1.NodeId)
	fmt.Printf("Node2 Hash is %x", node2.NodeId)
	fmt.Printf("Node3 Hash is %x", node3.NodeId)
	fmt.Printf("\nXOR of them is %x", distance1)
	fmt.Printf("\nXOR of them is %x\n", distance2)

	fmt.Println("H1:", node1)
	fmt.Println("H2:", node2)
	fmt.Println("H3:", node3)

	fmt.Println("D1:", distance1)
	fmt.Println("D2:", distance2)
	fmt.Println("D3:", distance3)

	fmt.Println("H1 and H2 are: ", dist1int1, " away")
	fmt.Println("H2 and H3 are: ", dist1int2, " away")
	fmt.Println("H1 and H3 are: ", dist1int3, " away")
}
