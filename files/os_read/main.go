package main

import (
	"log"
	"os"
)

func main() {
	dir := "/Users/sergey.ionin/GolandProjects/GoStudyProject/files/os_read"
	//dir := "temp/withEnvs/docker_test"
	// list all files in the directory dir
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	for _, file := range files {
		log.Println("filename = ", file.Name())
	}
	f.Close()
}
