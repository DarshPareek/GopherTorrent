package torrentfile

import (
	"bufio"
	"crypto/sha1"
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

func (m *MetainfoFile) SetData(fname string) {
	buf, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(buf)
	data, err := bencodeparser.Decode(r)
	if err != nil {
		panic(err)
	}
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if key == "announce" {
				switch v2 := value.(type) {
				case string:
					m.Announce = v2
				}
			} else if key == "info" {
				bencodedInfo, err := bencodeparser.GetEncodedInfo(value)
				if err != nil {
					panic(err)
				}
				infohash := sha1.Sum(bencodedInfo)
				m.InfoHash = infohash
			}
		}
	}
}
