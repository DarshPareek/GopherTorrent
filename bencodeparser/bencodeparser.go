package bencodeparser

import (
	"log"
	"os"
	"strconv"
)

func Test() string {
	content, err := os.ReadFile("bencodeparser/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	return parse(content)
}
func stringParse(data []byte) string {
	var colonIndex int
	for i := 0; i < len(data); i++ {
		if string(data[i]) == ":" {
			colonIndex = i
			break
		}
	}
	numberStr := data[:colonIndex]
	number, err := strconv.Atoi(string(numberStr))
	if err != nil {
		log.Fatal(err)
	}
	return string(data[colonIndex+1 : colonIndex+1+number])
}
func parse(data []byte) string {
	//stringParse(data)
	intParse(data)
	return string(data)
}
func intParse(data []byte) string {
	var eIndex int
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "e" {
			eIndex = i
			break
		}
	}
	return string(data[1:eIndex])

}
