package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mutex sync.Mutex

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		mutex.Lock()
		type Response struct {
			Data string `json:"data"`
		}

		resData := Response{Data: "This is a message from Go."}

		jsonRes, _ := json.MarshalIndent(resData, "", "  ")

		// _, err := res.Write([]byte("Kenny"))

		// res.Header().Add("Content-Type", "application/json")
		_, err := res.Write(append(jsonRes, '\n'))

		if err != nil {
			fmt.Println("Unable to send response")
			os.Exit(1)
		}
		mutex.Unlock()
	})

	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
