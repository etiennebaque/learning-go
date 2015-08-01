package main

type Ad struct {
	Id int `json: "id"`
	Title string `json: "title"`
	Description string `json: "description"`
}

type Ads []Ad