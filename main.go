package main

import (
	"fmt"
	"os"

	"github.com/DarshPareek/GopherTorrent/torrentfile"
	"github.com/DarshPareek/GopherTorrent/tracker"
)

func main() {
	var file torrentfile.MetainfoFile
	file.SetData("bencodeparser/main.torrent")
	resp, err := os.ReadFile("trackerResponse.txt")
	if err != nil {
		panic(err)
	}
	tp := tracker.ParseResponse(string(resp))
	fmt.Println(tp.Peers)
}
