package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DarshPareek/GopherTorrent/bencodeparser"
)

func main() {

	torrent, err := os.ReadFile("bencodeparser/main.torrent")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bencodeparser.Parse(torrent))
}
