package bencodeparser

import (
	"fmt"
	"log"
	"os"
)

func Test() {
	content, err := os.ReadFile("~/devel/GopherTorrent/bencodeparser/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
