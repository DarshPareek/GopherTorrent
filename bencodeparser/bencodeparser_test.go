package bencodeparser

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	torrent, err := os.ReadFile("bencodeparser/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(torrent))
	assert.Equal(t, string(torrent), "45:Arch Linux 2024.05.01 <https://archlinux.org>")
}
