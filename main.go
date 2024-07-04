package main

import (
	"log"
	"os"

	"github.com/DarshPareek/GopherTorrent/bencodeparser"
	"github.com/DarshPareek/GopherTorrent/torrentfile"
)

func main() {

	torrent, err := os.ReadFile("bencodeparser/main.torrent")
	if err != nil {
		log.Fatal(err)
	}
	data := bencodeparser.Parse(torrent)
	// fmt.Println(data)
	clean := torrentfile.OrganizeData(data)
	clean.ShowData()
}
