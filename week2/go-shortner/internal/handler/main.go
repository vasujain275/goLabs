package handler

import (
	"encoding/json"
	"goLabs/go-shortner/internal/store"
	"goLabs/go-shortner/internal/types"
	"net/http"
)

func RedirectHandler(d store.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortCode := r.PathValue("short_code")

		url, err := d.GetShortCode(shortCode)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	}
}

func ShortenHandler(d store.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
			return
		}

		var body types.ShortenRequest

		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			http.Error(w,"Internal Server Error!", http.StatusInternalServerError)
			return
		}

		shortCode := d.CreateShortCode(body.Url)

		resp := types.ShortenResponse{
			ShortCode: shortCode,
		}

		respJSON, err := json.Marshal(resp)

		if err != nil {
			http.Error(w,"Internal Server Error!", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type","application/json")

		w.Write(respJSON)
	}
}
