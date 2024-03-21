package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func updateJSON() {
	for {
		d := data{
			Water: rand.Intn(100) + 1,
			Wind:  rand.Intn(100) + 1,
		}
		jsonData, err := json.Marshal(d)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			continue
		}
		err = ioutil.WriteFile("data.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing JSON to file:", err)
			continue
		}
		time.Sleep(15 * time.Second)
	}
}

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()
	go updateJSON()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile("data.json")
		if err != nil {
			http.Error(w, "Error reading JSON file", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Write(b)
	})
	fmt.Printf("Listening on port %d...\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
