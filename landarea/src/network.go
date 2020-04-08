package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"
)

var mu sync.Mutex


func lock() {
	mu.Lock()
	log.Printf("lock")
}

func unlock() {
	mu.Unlock()
	log.Printf("unlock")
}

func foo() int {
	lock()

	func() {
		log.Printf("entry inner")
		defer unlock()
		log.Printf("exit inner")
	}()

	time.Sleep(1 * time.Second)
	log.Printf("return")
	return 0;
}

func main() {
	r := foo()
	log.Println("r=",r)
	return

	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	fileWriter, _ := bodyWriter.CreateFormFile("files", "file.txt")

	file, _ := os.Open("file.txt")
	defer file.Close()

	io.Copy(fileWriter, file)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, _ := http.Post("http://10.67.76.16:43855/upload", contentType, bodyBuffer)
	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(resp.Status)
	log.Println(string(resp_body))
}