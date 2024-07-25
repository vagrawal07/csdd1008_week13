package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Could not get hostname", http.StatusInternalServerError)
		return
	}

	colors := []string{"#ffcccc", "#ccffcc", "#ccccff"}
	instanceNumStr := os.Getenv("INSTANCE_NUM")
	instanceNum, err := strconv.Atoi(instanceNumStr)
	if err != nil || instanceNum < 0 || instanceNum >= len(colors) {
		instanceNum = 0 // default to the first color if there's an error or invalid instance number
	}
	bgColor := colors[instanceNum]

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
