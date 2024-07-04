package bencodeparser

import (
	"fmt"
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
	assert.Equal(t, string(torrent), "li1e3:twoi525e5:helloe")
}

func TestString(t *testing.T) {
	ans, i := stringParse([]byte("45:Arch Linux 2024.05.01 <https://archlinux.org>"))
	log.Print(i)
	assert.Equal(t, ans, "Arch Linux 2024.05.01 <https://archlinux.org>")
}
func TestInt(t *testing.T) {
	ans, i := intParse([]byte("i5454e"))
	log.Print(i)
	assert.Equal(t, ans, "5454")
}
func TestList(t *testing.T) {
	s, _ := listParse([]byte("li1e3:twoi525e5:helloe"))
	q := []string{"1", "two", "525", "hello"}
	fmt.Println(q)
	assert.Equal(t, s, q)
}
func TestDict(t *testing.T) {
	s, _ := dictParse([]byte("d4:spaml1:a1:bee"))
	q := make(map[any]any)
	q["spam"] = []string{"a", "b"}
	fmt.Println(q)
	assert.Equal(t, s, q)
}
