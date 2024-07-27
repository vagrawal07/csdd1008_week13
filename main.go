package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func randomColor() string {
	colors := []string{"#ffcccc", "#ccffcc", "#ccccff", "#ffffcc", "#ccffff", "#ffccff"}
	rand.Seed(time.Now().UnixNano())
	return colors[rand.Intn(len(colors))]
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}

	bgColor := randomColor()

	// HTML content with background color
	htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Hostname</title>
		<style>
			body {
				background-color: ` + bgColor + `;
				font-family: Arial, sans-serif;
				text-align: center;
				padding: 50px;
			}
			h1 {
				color: #333;
			}
		</style>
	</head>
	<body>
		<h1>Hostname: ` + hostname + `</h1>
	</body>
	</html>
	`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, htmlContent)
}

func main() {
	http.HandleFunc("/", hostnameHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
