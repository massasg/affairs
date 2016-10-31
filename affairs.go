package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	//"time"
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
	str_arr := strings.Split(str, "\n")

	fmt.Println(str_arr)
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand_affair_num := rand.Intn(len(str_arr) - 1)
	fmt.Println(str_arr)
	fmt.Println(str_arr[rand_affair_num])
}
