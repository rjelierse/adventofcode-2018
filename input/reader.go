package input

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
)

// ReadAsInts gathers the input data and returns the integer values for each line
func ReadAsInts(path string) []int {
	data := read(path)
	var ints []int

	for _, item := range data {
		value, err := strconv.Atoi(string(item))
		if err != nil {
			log.Fatal("Cannot parse number:", err)
		}
		ints = append(ints, value)
	}

	return ints
}

func ReadAsStrings(path string) []string {
	data := read(path)
	var strings []string
	for _, item := range data {
		strings = append(strings, string(item))
	}
	return strings
}

func read(path string) [][]byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cannot read file:", err)
	}
	data = bytes.TrimSpace(data)
	return bytes.Split(data, []byte{'\n'})
}
