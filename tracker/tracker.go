package tracker

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/DarshPareek/GopherTorrent/bencodeparser"
	"github.com/DarshPareek/GopherTorrent/torrentfile"
)

type Trackerresponse struct {
	Interval int64
	Peers    []Peer
}
type Peer struct {
	IP   net.IP
	Port uint16
}

func MakeRequest(file torrentfile.MetainfoFile, peerId [20]byte) (string, error) {
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
	fmt.Println(string(bodyBytes))
	return string(bodyBytes), nil
}
func ParseResponse(resp string) Trackerresponse {
	reader := bufio.NewReader(strings.NewReader(resp))
	data, _ := bencodeparser.Decode(reader)
	var r Trackerresponse
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if key == "interval" {
				switch v := value.(type) {
				case int64:
					r.Interval = v
				}
			} else if key == "peers" {
				switch v := value.(type) {
				case string:
					temp, err := SavePeers(v)
					if err != nil {
						panic(err)
					}
					r.Peers = temp
				}
			}
		}
	}
	return r
}
func SavePeers(peers string) ([]Peer, error) {
	peersBin := []byte(peers)
	const psize = 6
	numPeers := len(peersBin) / psize
	if len(peersBin)%psize != 0 {
		err := fmt.Errorf("r5ecieved malformed peers ")
		return nil, err
	}
	peer := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * psize
		peer[i].IP = net.IP(peersBin[offset : offset+4])
		peer[i].Port = binary.BigEndian.Uint16(peersBin[offset+4 : offset+6])
	}
	return peer, nil
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
