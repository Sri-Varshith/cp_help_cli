package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

func main() {
	userptr := flag.String("rating", "0", "a string")
	flag.Parse()

	resp, err := http.Get("https://codeforces.com/api/user.info?handles=" + *userptr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	json.Unmarshal(body, )

	log.Println(user)

}
