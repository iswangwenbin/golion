package randx

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"github.com/pborman/uuid"
)

var encoding = base32.NewEncoding("hCNa0qENbRVFlN2O52gDfQpmPU7ZMQIn")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}

func NewRandomString(length int) string {
	var b bytes.Buffer
	str := make([]byte, length+8)
	rand.Read(str)
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(str)
	encoder.Close()
	b.Truncate(length) // removes the '==' padding
	return b.String()
}
