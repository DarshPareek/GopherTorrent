package torrentfile

import "fmt"

type torrentInfo struct {
	length      any
	name        any
	pieceLength any
	pieces      any
}

type TorrentFile struct {
	announce     any
	comment      any
	createdBy    any
	creationDate any
	Info         torrentInfo
}

func OrganizeData(data map[any]any) TorrentFile {
	var t TorrentFile
	var d torrentInfo
	for key, value := range data {
		switch key {
		case "announce":
			t.announce = value
		case "comment":
			t.comment = value
		case "created by":
			t.createdBy = value
		case "creation date":
			t.creationDate = value
		case "info":
			switch v := value.(type) {
			case map[any]any:
				for k, va := range v {
					switch k {
					case "length":
						d.length = va
					case "name":
						d.name = va
					case "piece length":
						d.pieceLength = va
					case "pieces":
						d.pieces = va
					}
				}
			}
			t.Info = d
		}
	}
	return t
}

func (t TorrentFile) ShowData() {
	fmt.Println(t.announce)
	fmt.Println(t.comment)
	fmt.Println(t.createdBy)
	fmt.Println(t.creationDate)
	fmt.Println(t.Info.length)
	fmt.Println(t.Info.name)
	fmt.Println(t.Info.pieceLength)
	fmt.Println(t.Info.pieces)
}
