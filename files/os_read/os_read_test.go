package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOSread(t *testing.T) {

	homeDir, _ := os.Getwd()
	dir := fmt.Sprintf("%s/%s", homeDir, "temp")

	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	for _, file := range files {
		fmt.Println("filename = ", file.Name())
	}
	f.Close()
	assert.NoError(t, err)
}
