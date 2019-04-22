package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
	"math/rand"
)

var lines []string
var numquotes int

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Marvin received a request")
	fmt.Fprintf(w, lines[rand.Intn(numquotes)])
}

func main() {
	log.Print("Marvin started up")

	fileBytes, err := ioutil.ReadFile("quotes.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines = strings.Split(string(fileBytes), "\n")
	lines = lines[:len(lines)-1]
	numquotes = len(lines)
	log.Print("Read ", numquotes, " quotes into memory")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
