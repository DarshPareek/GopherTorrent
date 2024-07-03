package bencodeparser

import (
	"log"
	"os"
	"strconv"
	"unicode"
)

func Test() string {
	content, err := os.ReadFile("bencodeparser/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	return parse(content)
}
func stringParse(data []byte) (string, int) {
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
	return string(data[colonIndex+1 : colonIndex+1+number]), (colonIndex + 1 + number)
}
func parse(data []byte) string {
	stringParse([]byte("4:helo"))
	intParse([]byte("i123e"))
	listParse([]byte("li1e3:twoi525e5:helloe"))
	return string(data)
}
func intParse(data []byte) (string, int) {
	var eIndex int
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "e" {
			eIndex = i
			break
		}
	}
	return string(data[1:eIndex]), eIndex

}
func listParse(data []byte) []string {
	data = data[1:]
	startIndex := 0
	var temp string
	res := make([]string, 0)
	for startIndex != len(data)-1 {
		startIndex = 0
		s := string(data[startIndex])
		if s == "i" {
			temp, startIndex = intParse(data[startIndex:])
			res = append(res, temp)
			startIndex = startIndex + 1
			data = data[startIndex:]
		} else if unicode.IsDigit([]rune(s)[0]) {
			temp, startIndex = stringParse(data[startIndex:])
			res = append(res, temp)
			data = data[startIndex:]
		}
	}
	return res
}
