package main

import (
	"crypto/rand"
	"log"
	"os"

	"github.com/DarshPareek/GopherTorrent/p2p"
	"github.com/DarshPareek/GopherTorrent/torrentfile"
	"github.com/DarshPareek/GopherTorrent/tracker"
)

func main() {
	var peerID [20]byte
	rand.Read(peerID[:])
	if len(os.Args) < 2 {
		log.Println("Invalid Command Line Arguments")
		return
	}
	torrentPath := os.Args[1]
	saveName := os.Args[2]
	var tfile torrentfile.MetainfoFile
	tfile.SetData(torrentPath)
	resp, err := tracker.MakeRequest(tfile, peerID)
	if err != nil {
		log.Println(err)
		return
	}
	tres := tracker.ParseResponse(resp)
	torrent := p2p.Torrent{
		Peers:       tres.Peers,
		PeerID:      [20]byte(peerID),
		InfoHash:    [20]byte(tfile.InfoHash),
		PieceHashes: tfile.Info.Pieces,
		PieceLength: int(tfile.Info.PieceLength),
		Length:      int(tfile.Info.Length),
	}
	data, err := torrent.Download()
	if err != nil {
		log.Println(err)
		return
	}
	outFile, err := os.Create("./Downloaded/" + saveName)
	if err != nil {
		log.Println(err)
		return
	}
	defer outFile.Close()
	_, err = outFile.Write(data)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("File Downloaded Successfully at ./Downloaded/" + saveName)
}
