package rest

import (
	"net/http"
)

const (
	GET  string = "GET"
	POST string = "POST"
)

type urlDescription struct {
	URL        url                 `json:"url"`
	Methods    []methodDescription `json:"methods"`
	urlPattern string
	handler    http.HandlerFunc
}

type methodDescription struct {
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

var urlData []urlDescription = []urlDescription{
	{
		URL: url("/"),
		Methods: []methodDescription{
			{
				Method:      GET,
				Description: "See Documentation",
			},
		},
		urlPattern: "/",
		handler:    documentation,
	},
	{
		URL:        url("/blocks"),
		urlPattern: "/blocks",
		Methods: []methodDescription{
			{
				Method:      GET,
				Description: "See All Blocks",
			},
			{
				Method:      POST,
				Description: "Add A Block",
				Payload:     "data:string",
			},
		},
		handler: blocks,
	},
	{
		URL:        url("/blocks/{hash}"),
		urlPattern: "/blocks/{hash:[a-f0-9]+}",
		Methods: []methodDescription{
			{
				Method:      GET,
				Description: "See A Block",
			},
		},
		handler: block,
	},
	{
		URL:        url("/status"),
		urlPattern: "/status",
		Methods: []methodDescription{
			{
				Method:      GET,
				Description: "See A Block",
			},
		},
		handler: status,
	},
}
