package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	fileUrl := "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4"
	parts := strings.Split(fileUrl, "/")
	fileName := parts[len(parts)-1]
	fmt.Println(fileName)
	err := DownloadFile(fileName, fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
