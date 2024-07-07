package bencodeparser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
)

func Decode(r *bufio.Reader) (interface{}, error) {
	ch, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	switch ch {
	case 'i':
		intBuf, err := readTill(r, 'e')
		if err != nil {
			return nil, err
		}
		intBuf = intBuf[:len(intBuf)-1]
		integer, err := strconv.ParseInt(string(intBuf), 10, 64)
		if err != nil {
			return nil, err
		}
		return integer, nil
	case 'l':
		list := []interface{}{}
		for {
			c, err2 := r.ReadByte()
			if err2 == nil {
				if c == 'e' {
					return list, nil
				} else {
					r.UnreadByte()
				}
			}
			value, err := Decode(r)
			if err != nil {
				return nil, err
			}
			list = append(list, value)

		}
	case 'd':
		dict := map[string]interface{}{}
		for {
			c, err2 := r.ReadByte()
			if err2 == nil {
				if c == 'e' {
					return dict, nil
				} else {
					r.UnreadByte()
				}
			}
			value, err := Decode(r)
			if err != nil {
				return nil, err
			}
			key, ok := value.(string)
			if !ok {
				return nil, errors.New("bencode: non-string dicts key")
			}
			value, err = Decode(r)
			if err != nil {
				return nil, err
			}
			dict[key] = value

		}
	default:
		r.UnreadByte()
		strLen, err := readTill(r, ':')
		if err != nil {
			return nil, err
		}
		strLen = strLen[:len(strLen)-1]
		length, err := strconv.ParseInt(string(strLen), 10, 64)
		if err != nil {
			return nil, err
		}
		buf := make([]byte, length)
		_, err = read(r, buf, int(length))
		return string(buf), err
	}
}

func readTill(r *bufio.Reader, delim byte) ([]byte, error) {
	buffered := r.Buffered()
	var buff []byte
	var err error
	if buff, err = r.Peek(buffered); err != nil {
		return nil, err
	}
	if i := bytes.IndexByte(buff, delim); i >= 0 {
		return r.ReadSlice(delim)
	}
	return r.ReadSlice(delim)
}

func read(r *bufio.Reader, buf []byte, l int) (int, error) {
	if len(buf) < l {
		return 0, io.ErrShortBuffer
	}
	var n int
	var err error
	for n < l && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}
	if n >= l {
		err = nil
	} else if n > 0 && err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return n, err
}

func GetEncodedInfo(data interface{}) ([]byte, error) {
	var encoded []byte
	encoded = append(encoded, 'd')
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			// encoded = encoded + strconv.Itoa(len(key)) + ":" + key
			// fmt.Println(key, value, reflect.TypeOf(key), reflect.TypeOf(value))
			encoded = append(encoded, []byte(strconv.Itoa(len(key)))...)
			encoded = append(encoded, ':')
			encoded = append(encoded, []byte(key)...)
			switch v := value.(type) {
			case int64:
				// encoded = encoded + "i" + strconv.Itoa(v) + "e"
				encoded = append(encoded, 'i')
				encoded = append(encoded, []byte(strconv.FormatInt(v, 10))...)
				encoded = append(encoded, 'e')
			case string:
				// encoded = encoded + strconv.Itoa(len(v)) + ":" + v
				encoded = append(encoded, []byte(strconv.Itoa(len(v)))...)
				encoded = append(encoded, ':')
				encoded = append(encoded, []byte(v)...)
			}
			// fmt.Println(key, value)
		}
	}
	encoded = append(encoded, 'e')
	return encoded, nil
}
