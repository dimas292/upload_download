package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const PATH = "public/download"

func main() {
	url := "https://picsum.photos/200/300"

	dst, err := os.Create(fmt.Sprintf("%s/image_%d.%s", PATH, time.Now().Unix(), "jpg"))
	if err != nil {
		log.Println(err)
	}

	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	_ ,err = io.Copy(dst, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("download success")

}