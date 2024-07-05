package torrentfile

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DarshPareek/GopherTorrent/bencodeparser"
)

type MetainfoFile struct {
	Announce string
	InfoHash [20]byte
	Info     TorrentInfo
}

type TorrentInfo struct {
	pieceLength int
	pieces      [][20]byte
}

func (m MetainfoFile) SetData(fname string) {
	buf, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(buf)
	data, err := bencodeparser.Decode(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
