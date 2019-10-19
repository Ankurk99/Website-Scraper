package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Enter the url first! \n")
		os.Exit(1)
	}
	retrieve(args[0])
}

func retrieve(url string) {

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("failed creating file  %s", err)
	}
	defer file.Close()
	file.WriteString(string(body))
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	} else {
		fmt.Printf("Serving the file at : http://localhost:3000 \n")
		http.Handle("/", http.FileServer(http.Dir("./")))
		http.ListenAndServe(":3000", nil)
	}
}
