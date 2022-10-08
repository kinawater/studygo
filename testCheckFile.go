package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var fileSrc string
	fileSrc = "G:/TG VIDEO DOWLAND/20220807"
	filePath := fileSrc + "/mulu.txt"
	putFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	checkAndPut(fileSrc, putFile)
}

func checkAndPut(fileSrc string, putFile *os.File) {
	files, _ := ioutil.ReadDir(fileSrc)
	for _, f := range files {
		if f.IsDir() {
			fmt.Println(f.Name())
			_, _ = putFile.Write([]byte("---" + f.Name() + "\r\n"))
			checkAndPut(fileSrc+"/"+f.Name(), putFile)
		} else {
			_, _ = putFile.Write([]byte("    |___" + f.Name() + "\r\n"))
		}
	}
}
