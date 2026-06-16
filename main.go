package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func main() {
	userptr := flag.String("rating", "0", "a string")
	flag.Parse()

	resp, err := http.Get("https://codeforces.com/api/user.rating?handle=" + *userptr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

}
