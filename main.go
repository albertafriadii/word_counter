package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"word_counter/utils"
)

type Word struct {
	Key   string
	Value int
}

func wordCounter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	input := r.FormValue("input")

	for k, v := range utils.WordCount(input) {
		result := &Word{Key: k, Value: v}
		jsonResult, err := json.Marshal(result)

		if err != nil {
			http.Error(w, "Error", http.StatusBadRequest)
		}
		w.Write(jsonResult)

		// _ = json.NewDecoder(r.Body).Decode(result)
		// json.NewEncoder(w).Encode(result)
		fmt.Println(string(jsonResult))
		// fmt.Printf("%s: %d\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", wordCounter)

	fmt.Printf("Starting server at port 8000\n")
	http.ListenAndServe(":8000", nil)
}
