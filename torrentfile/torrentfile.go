package torrentfile

import (
	"bufio"
	"bytes"
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
	PieceLength int64
	Pieces      [][20]byte
	Length      int64
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
				switch v := value.(type) {
				case map[string]interface{}:
					for key, value := range v {
						if key == "length" {
							switch v := value.(type) {
							case int64:
								m.Info.Length = v
							}
						} else if key == "pieces" {
							// fmt.Println(reflect.TypeOf(value))
							switch v := value.(type) {
							case string:
								temp := bytes.NewBufferString(v)
								for i := 0; i < len(v)/20; i += 1 {
									var val [20]byte
									temp.Read(val[:])
									m.Info.Pieces = append(m.Info.Pieces, val)
								}
							}
						} else if key == "piece length" {
							// fmt.Println(reflect.TypeOf(value))
							switch v := value.(type) {
							case int64:
								m.Info.PieceLength = v
							}
						}
					}
				}
			}
		}
	}
}
