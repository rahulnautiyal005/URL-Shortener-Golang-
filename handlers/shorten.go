package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/database"
	"url-shortener/models"
	"url-shortener/utils"
)

type ShortenRequest struct {
	URL            string `json:"url"`
	ExpiresInHours int    `json:"expires_in_hours,omitempty"`
}

type ShortenResponse struct {
	ShortURL  string    `json:"short_url"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}

func ShortenHandler(db database.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		id, err := db.GetNextID()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		shortCode := utils.Encode(id)
		var expiresAt time.Time
		if req.ExpiresInHours > 0 {
			expiresAt = time.Now().Add(time.Duration(req.ExpiresInHours) * time.Hour)
		}

		urlModel := &models.URL{
			ID:        id,
			LongURL:   req.URL,
			ShortCode: shortCode,
			ExpiresAt: expiresAt,
		}

		if err := db.Save(urlModel); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		host := r.Host
		if host == "" {
			host = "localhost:8080"
		}

		resp := ShortenResponse{
			ShortURL:  fmt.Sprintf("http://%s/%s", host, shortCode),
			ExpiresAt: expiresAt,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
