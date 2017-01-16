package main

type Product struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	OriPrice    float32  `json:"oriPrice"`
	Price       float32  `json:"price"`
	Description string   `json:"description"`
	Imgs        []string `json:"imgs"`
}
