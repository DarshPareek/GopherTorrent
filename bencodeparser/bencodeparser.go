package bencodeparser

import (
	"log"
	"os"
)

func Test() string {
	content, err := os.ReadFile("bencodeparser/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	return parse(content)
}
func parse(data []byte) string {
	return "Hello"
}
