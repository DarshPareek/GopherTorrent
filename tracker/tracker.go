package tracker

import (
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/DarshPareek/GopherTorrent/torrentfile"
)

func MakeRequest(file torrentfile.MetainfoFile) (string, error) {
	peerId := make([]byte, 20)
	rand.Read(peerId)
	port := 6881
	base, err := url.Parse(file.Announce)
	if err != nil {
		return "", err
	}
	params := url.Values{
		"info_hash":  []string{string(file.InfoHash[:])},
		"peer_id":    []string{string(peerId[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(661651456)},
	}
	base.RawQuery = params.Encode()
	resp, err := http.Get(base.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
func ParseResponse(resp string) {
	// reader := bufio.NewReader(strings.NewReader(resp))
	// t, _ := reader.Peek(100)
	fmt.Println(resp)
	strings.NewReader(resp)
	// fmt.Println(t)
	// data, _ := bencodeparser.Decode(reader)
	// fmt.Println(data)
}
