package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// URL of the image you want to download
	imageUrl := "https://www.collinsdictionary.com/images/full/apple_158989157.jpg"

	// Create an HTTP GET request
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return
	}

	// Create a new file to save the image
	outputFile, err := os.Create("testpic.png")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return
	}

	fmt.Println("Image downloaded and saved as 'testpic.png'")
}
