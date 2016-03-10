package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

func XOR(hash1 [20]byte, hash2 [20]byte) (ret [20]byte) {
	for i := 0; i < 20; i++ {
		ret[i] = hash1[i] ^ hash2[i]
	}
	return
}

func Binarize(hash [20]byte) string {
	myStr := []string{}
	for i := 0; i < 20; i++ {
		myStr = append(myStr, strconv.FormatInt(int64(hash[i]), 2))
	}
	return strings.Join(myStr, "-")
}

type NodeId [20]byte

func main() {
	c := 200
	num := make([]byte, c)
	num2 := make([]byte, c)
	num3 := make([]byte, c)
	_, err := rand.Read(num)
	_, err = rand.Read(num2)
	_, err = rand.Read(num3)
	if err != nil {
		fmt.Println(err)
	}

	hash1 := sha1.Sum(num)
	hash2 := sha1.Sum(num2)
	hash3 := sha1.Sum(num3)

	distance1 := XOR(hash1, hash2)
	distance2 := XOR(hash2, hash3)
	distance3 := XOR(hash3, hash1)

	dist1int1, _ := binary.Uvarint(distance1[:]) //cast array to slice
	dist1int2, _ := binary.Uvarint(distance2[:]) //cast array to slice
	dist1int3, _ := binary.Uvarint(distance3[:]) //cast array to slice

	fmt.Printf("Sha1 Hash is %x", hash1)
	fmt.Printf("\nSha1 Hash is %x", hash2)
	fmt.Printf("\nSha1 Hash is %x", hash3)
	fmt.Printf("\nXOR of them is %x", distance1)
	fmt.Printf("\nXOR of them is %x\n", distance2)

	fmt.Printf("\nH1 binary: %s", Binarize(hash1))
	fmt.Printf("\nH2 binary: %s", Binarize(hash2))
	fmt.Printf("\nH2 binary: %s", Binarize(hash3))
	fmt.Printf("\nD1 binary: %s", Binarize(distance1))
	fmt.Printf("\nD2 binary: %s", Binarize(distance2))
	fmt.Printf("\nD3 binary: %s\n\b", Binarize(distance2))

	fmt.Println("H1:", hash1)
	fmt.Println("H2:", hash2)
	fmt.Println("H3:", hash3)
	fmt.Println("D1:", distance1)
	fmt.Println("D2:", distance2)
	fmt.Println("D3:", distance3)

	fmt.Println("H1 and H2 are: ", dist1int1, " away")
	fmt.Println("H2 and H3 are: ", dist1int2, " away")
	fmt.Println("H1 and H3 are: ", dist1int3, " away")
}
