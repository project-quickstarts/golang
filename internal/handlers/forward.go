package handlers

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Forward(w http.ResponseWriter, r *http.Request) {
	// Get params
	externalUrl := r.URL.Query().Get("url")
	if externalUrl == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	// Validate the URL
	_, err := url.ParseRequestURI(externalUrl)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	log.Println("Forwarding to", externalUrl)

	// Create a new HTTP request to the external API
	req, err := http.NewRequest(r.Method, externalUrl, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy all headers from the incoming request to the new request
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Make the HTTP request to the external API
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil { // Handle errors
		http.Error(w, "Failed to make request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response headers from the external API to the client
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set the status code
	w.WriteHeader(resp.StatusCode)

	// Copy the response body from the external API to the client
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to write response body", http.StatusInternalServerError)
		return
	}
}
