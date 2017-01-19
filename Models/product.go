package Models

type Product struct {
	Id          int64    `json:"id"`
	Title       string   `json:"title"`
	OriPrice    float32  `json:"oriPrice"`
	Price       float32  `json:"price"`
	Description string   `json:"description"`
	Imgs        []string `json:"imgs"`
}
