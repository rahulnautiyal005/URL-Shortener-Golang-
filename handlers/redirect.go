package handlers

import (
	"net/http"
	"strings"
	"time"
	"url-shortener/database"
)

func RedirectHandler(db database.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract code from path (e.g., /abc123 -> abc123)
		code := strings.TrimPrefix(r.URL.Path, "/")
		if code == "" || code == "shorten" {
			return
		}

		url, err := db.GetByCode(code)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		// Check for expiration
		if !url.ExpiresAt.IsZero() && time.Now().After(url.ExpiresAt) {
			http.Error(w, "URL has expired", http.StatusGone)
			return
		}

		// Track click
		_ = db.IncrementClick(code)

		http.Redirect(w, r, url.LongURL, http.StatusMovedPermanently)
	}
}
