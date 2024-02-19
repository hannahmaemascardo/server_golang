package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Define the IP address and port of the server
	ipAddress := "192.168.1.33" //
	port := 8080

	// Construct the URL with the server's IP address and port
	url := fmt.Sprintf("http://%s:%d/hello", ipAddress, port)

	// Send HTTP GET request to the server
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the server's response
	fmt.Println("Server Response:", string(body))
}
