package main

import "fmt"

func main() {
	baseClient := &ClientConfig{
		BaseURL: "http://localhost:8080/",
	}

	secondaryClient := baseClient.Clone().(*ClientConfig)
	baseUrl := secondaryClient.GetBaseUrl()
	fmt.Println(baseUrl)
}
