package bencodeparser

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	torrent, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, string(torrent), "i5454e")
}

//	func TestString(t *testing.T) {
//		torrent, err := os.ReadFile("file.txt")
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Print(string(torrent))
//		assert.Equal(t, stringParse([]byte("45:Arch Linux 2024.05.01 <https://archlinux.org>")), "Arch Linux 2024.05.01 <https://archlinux.org>")
//	}
func TestInt(t *testing.T) {
	torrent, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, stringParse(torrent), "5454")
}
