package main

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Lady soul", Artist: "Aretha Franklin", Price: 87.99},
	{ID: "2", Title: "Dream a little dream of me", Artist: "Ella Fitzgerald", Price: 87.80},
	{ID: "3", Title: "Solitude", Artist: "Billie Holiday", Price: 87.00},
}
