package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	filePath, err := filepath.Abs("assets")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Building assets at: %v\n", filePath)

	os.MkdirAll(filePath, os.FileMode(int(0777)))

	data := []byte("hello world")
	err = ioutil.WriteFile(filePath+"/file.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nFinished!")
}
