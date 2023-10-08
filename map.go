package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Map(c *config) {
	var res *http.Response
	var err error
	if c.Next != "" {
		res, err = http.Get(c.Next)
	} else {
		res, err = http.Get(BASE_URL + "location-area")
	}
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var ParsedBody struct {
		Count    int
		Url      string
		Next     *string
		Previous *string
		Results  []struct {
			Name string
			Url  string
		}
	}
	err = json.Unmarshal(body, &ParsedBody)
	if err != nil {
		log.Fatal(err)
	}
	if ParsedBody.Next != nil {
		c.Next = *ParsedBody.Next
	}
	if ParsedBody.Previous != nil {
		c.Previous = *ParsedBody.Previous
	}
	for _, result := range ParsedBody.Results {
		fmt.Println(result.Name)
	}
}
