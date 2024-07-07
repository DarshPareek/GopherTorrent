package main

import (
	"os"

	"github.com/DarshPareek/GopherTorrent/tracker"
)

func main() {
	// var file torrentfile.MetainfoFile
	// file.SetData("bencodeparser/main.torrent")
	// resp, err := tracker.MakeRequest(file)
	// if err != nil {
	// 	panic(err)
	// }
	// f, err := os.Create("trackerResponse.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// os.WriteFile("trackerResponse.txt", []byte(resp), 0666)
	// peerId := make([]byte, 20)
	// rand.Read(peerId)
	// port := 6881
	// base, err := url.Parse(file.Announce)
	// if err != nil {
	// 	panic(err)
	// }
	// params := url.Values{
	// 	"info_hash":  []string{string(file.InfoHash[:])},
	// 	"peer_id":    []string{string(peerId[:])},
	// 	"port":       []string{strconv.Itoa(int(port))},
	// 	"uploaded":   []string{"0"},
	// 	"downloaded": []string{"0"},
	// 	"compact":    []string{"1"},
	// 	"left":       []string{strconv.Itoa(661651456)},
	// }
	// base.RawQuery = params.Encode()
	// resp, err := http.Get(base.String())
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(bodyBytes))
	resp, err := os.Open("trackerResponse.txt")
	if err != nil {
		panic(err)
	}

	var data []byte
	resp.Read(data)
	tracker.ParseResponse(string(data))
}
