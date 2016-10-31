package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("affairs.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	//fmt.Println(str)
	str_arr := strings.Split(str, "\n")
	fmt.Println(str_arr)
}
