package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type InputEvent struct {
	Button string // json:button
	Action string // json:action
}

// logic for button presses
func inputHandler(w http.ResponseWriter, r *http.Request) {
	var event InputEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: add the ability to toggle debug statements
	log.Printf("Input: %s %s\n", event.Button, event.Action)

	w.WriteHeader(http.StatusOK)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/input", inputHandler)
	log.Println("Starting http server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
