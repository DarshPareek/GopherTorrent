package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/DarshPareek/GopherTorrent/client"
	"github.com/DarshPareek/GopherTorrent/torrentfile"
	"github.com/DarshPareek/GopherTorrent/tracker"
)

func main() {
	peerId := make([]byte, 20)
	rand.Read(peerId)
	var file torrentfile.MetainfoFile
	file.SetData("bencodeparser/main.torrent")
	resp, err := tracker.MakeRequest(file, [20]byte(peerId))
	// resp, err := os.ReadFile("trackerResponse.txt")
	if err != nil {
		panic(err)
	}
	tp := tracker.ParseResponse(string(resp))
	for _, peer := range tp.Peers {
		data, err := client.New(peer, [20]byte(peerId), file.InfoHash)
		if err != nil {
			log.Printf("couldn't handshake with %s \n", peer.IP)
		}
		fmt.Println(data)
	}
}
