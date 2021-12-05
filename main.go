package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type DiscordData struct {
	Content string         `json:"content"`
	Embeds  []DiscordEmbed `json:"embeds"`
}

type DiscordField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DiscordEmbed struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Color       int            `json:"color"`
	Fields      []DiscordField `json:"fields,omitempty"`
}

func postWebhook(webhook string, data DiscordData) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("json.Marshal %v\n", err)
		return
	}

	req, err := http.NewRequest(
		"POST",
		webhook,
		bytes.NewBuffer(dataBytes),
	)

	if err != nil {
		fmt.Printf("http.NewRequest %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	c := &http.Client{}
	res, err := c.Do(req)

	if err != nil {
		fmt.Printf("http.Do %v\n", err)
		return
	}

	defer res.Body.Close()
}

func main() {
	var name, webhook string

	flag.StringVar(&webhook, "w", "", "Specify discord webhook")
	flag.StringVar(&name, "n", "", "Specify torrent name.")
	flag.Parse()

	if name != "" {
		postWebhook(webhook,
			DiscordData{
				Embeds: []DiscordEmbed{
					{
						Title:       "New torrent downloaded",
						Description: name,
						Color:       4251719,
					},
				},
			})
	}
}
