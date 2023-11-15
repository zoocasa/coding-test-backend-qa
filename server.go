package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type listing struct {
	Id   int
	City string
}

var listings = []listing{
	{
		Id:   1,
		City: "Toronto",
	},
	{
		Id:   2,
		City: "Burlington",
	},
}

func getListings(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filteredListings := listings

	if req.URL.Query().Has("city") {
		cityFilter := req.URL.Query().Get("city")

		filteredListings = make([]listing, 0)

		for _, listing := range listings {
			if listing.City == cityFilter {
				filteredListings = append(filteredListings, listing)
			}
		}
	}

	listingsJson, err := json.Marshal(filteredListings)
	if err != nil {
		http.Error(w, "Could not serialize listings", http.StatusInternalServerError)
		return
	}

	w.Write(listingsJson)
}

func StartServer() {
	http.HandleFunc("/listings", getListings)

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
