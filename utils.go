/*

Binary tree search
*/

package kademlia

import (
	"io"
	"bytes"
	"math/rand"
	"crypto/sha1"
	"time"
	)

// Generate random string of characters of a particular length
func RandomString(length int) (string){

	var buffer bytes.Buffer

	for i:=0; i<length; i++ {
		buffer.WriteByte(byte(rand.Intn(255)));
	}

	return buffer.String()
}

// Random hash string
func RandomID() (string){
	h := sha1.New()
	io.WriteString(h, RandomString(20))

	return string(h.Sum(nil))
}

func HashToInt(s string) (int){
	return 10
}

func init(){
	rand.Seed(time.Now().UTC().UnixNano())
}