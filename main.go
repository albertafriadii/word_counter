package main

import (
	"encoding/json"
	"fmt"
	"log"
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
			w.WriteHeader(http.StatusBadRequest)
			resp := make(map[string]string)
			resp["message"] = "Bad Request"
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}
			w.Write(jsonResp)
		}
		w.Write(jsonResult)

		// _ = json.NewDecoder(r.Body).Decode(result)
		// json.NewEncoder(w).Encode(result)
		// fmt.Println(string(b))
		// fmt.Printf("%s: %d\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", wordCounter)

	fmt.Printf("Starting server at port 8000\n")
	http.ListenAndServe(":8000", nil)
}
