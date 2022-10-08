package main

import (
	"fmt"
	"github.com/thinkeridea/go-extend"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func main() {
	var fileSrc string
	fileSrc = "G:/pornhub"
	//execRar := "C:/Program Files/WinRAR/WinRAR.exe"
	files, _ := ioutil.ReadDir(fileSrc)
	for _, f := range files {
		if f.IsDir() {
			continue
		} else {
			if filepath.Ext(f.Name())[1:] == ".rar" {
				continue
			}
			fmt.Println(f.Name()[:])

			//cmdLord = execRar + " a –ibck –m5 –p xxxgp " + f.Name()+"/" +" "+ fileSrc + "/"
			//cmd := exec.Command("ls", "-a")
		}
	}
}
func SubStrRuneSubString(s string, length int) string {
	return exutf8.RuneSubString(s, 0, length)
}

func BenchmarkSubStrRuneSubString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrRuneSubString(benchmarkSubString, benchmarkSubStringLength)
	}
}
