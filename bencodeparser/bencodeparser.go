package bencodeparser

import (
	"log"
	"strconv"
	"unicode"
)

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
func Parse(data []byte) map[any]any {
	res, _ := dictParse(data)
	return res
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
func listParse(data []byte) ([]string, int) {
	var processed int
	data = data[1:]
	startIndex := 0
	var temp string
	res := make([]string, 0)
	for string(data[startIndex]) != "e" {
		s := string(data[startIndex])
		if s == "i" {
			temp, startIndex = intParse(data[startIndex:])
			res = append(res, temp)
			startIndex = startIndex + 1
			processed += startIndex + 1
			data = data[startIndex:]
		} else if unicode.IsDigit([]rune(s)[0]) {
			temp, startIndex = stringParse(data[startIndex:])
			res = append(res, temp)
			processed += startIndex
			data = data[startIndex:]
		}
		startIndex = 0
	}
	return res, processed + 1
}
func parsePart(data []byte) ([]any, int) {
	startIndex := 0
	s := string(data[startIndex])
	index := 0
	var res interface{}
	var resi interface{}
	var rest interface{}
	switch s {
	case "l":
		res, index = listParse(data[startIndex:])
		index = index + 1
	case "i":
		resi, index = intParse(data[startIndex:])
		index = index + 1
	case "d":
		rest, index = dictParse(data[startIndex:])
	}
	if unicode.IsDigit([]rune(s)[0]) {
		resi, index = stringParse(data[startIndex:])
	}
	return []any{res, resi, rest}, index
}
func dictParse(data []byte) (map[any]any, int) {
	res := make(map[any]any)
	var p1, p2 []any
	var q int
	startIndex := 1
	processed := 1
	for string(data[startIndex]) != "e" {
		p1, q = parsePart(data[startIndex:])
		data = data[q:]
		processed += q
		p2, q = parsePart(data[startIndex:])
		data = data[q:]
		processed += q
		if p2[0] != nil {
			res[p1[1]] = p2[0]
		} else if p2[1] != nil {
			res[p1[1]] = p2[1]
		} else {
			res[p1[1]] = p2[2]
		}
	}
	return res, processed
}
