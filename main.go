package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const DEFAULT_DELAY = "1"
const DEFAULT_PORT = "8080"

func main() {

	var delayStr = os.Getenv("DELAY")
	if delayStr == "" {
		delayStr = DEFAULT_DELAY
	}
	delayInt, err := strconv.Atoi(delayStr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		time.Sleep(time.Duration(delayInt) * time.Second)
		fmt.Fprintln(w, "Final response after ", delayInt, " seconds")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	fmt.Println("Started, serving at ", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
