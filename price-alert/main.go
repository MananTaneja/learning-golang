package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	priceAlert "priceAlert/internal"
)

type RequestBody struct {
	PageUrl string `json:"pageUrl"`
}

type Response struct {
	Price string `json:"price"`
}

func amazonHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var body RequestBody

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			fmt.Println("unable to parse body")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		price, err := priceAlert.Amazon(body.PageUrl)

		if err != nil {
			fmt.Println("Error while parsing price from the given url")
		}

		w.Header().Set("Content-Type", "application/json")

		response := Response{
			Price: price,
		}

		json.NewEncoder(w).Encode(response)

	}

}

// func flipkartHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Flipkart api handler")

// 		var body RequestBody

// 		err := json.NewDecoder(r.Body).Decode(&body)

// 		if err != nil {
// 			log.Fatal("Unable to decode request body")
// 		}

// 		defer r.Body.Close()

// 		price := priceAlert.FlipkartFetch(body.PageUrl)

// 		// if err != nil {
// 		// 	log.Fatal("Error unable to parse price")
// 		// }

// 		w.Header().Set("Content-Type", "application/json")

// 		response := Response{
// 			Price: price,
// 		}

// 		json.NewDecoder(w).Encode(response)
// 	}
// }

func main() {
	fmt.Println("Running server on 8080")

	http.HandleFunc("/api/v1/amazon", amazonHandler())

	// http.HandleFunc("api/v1/flipkart", flipkartHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))

}
