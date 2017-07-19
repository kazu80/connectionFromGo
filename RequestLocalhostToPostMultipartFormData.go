package main

import (
	"bytes"
	"mime/multipart"
	"os"
	"io"
	"net/http"
	"log"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}

	defer readFile.Close();

	io.Copy(fileWriter, readFile);
	writer.Close()

	resp, err := http.Post("http://localhost:8000/upload/movie", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
