package torrentfile

type torrentInfo struct {
	length      int
	name        string
	pieceLength int
	pieces      []byte
}

type TorrentFile struct {
	announce     any
	comment      any
	createdBy    any
	creationDate any
	Info         torrentInfo
}

func organizeData(data interface{}) TorrentFile {
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
			for k, v := range value {
				switch k {
				case "length":
					d.length = v
				case "name":
					d.name = v
				case "piece lenght":
					d.pieceLength = v
				case "pieces":
					d.pieces = v
				}
			}
			t.Info = d
		}
	}
	return t
}
