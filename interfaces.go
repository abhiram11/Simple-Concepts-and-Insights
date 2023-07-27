package main

import "fmt"

type Book struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type Drinks struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (d Drinks) printInfo() {
	fmt.Printf("Drink : %v , and Prince %vn", d.Name, d.Price)
}

func (b Book) printInfo() {
	fmt.Printf("Title: %v, and Price: %vn", b.Title, b.Price)
}

func main() {

	bha := Book{
		Title: "Bhagvatgita",
		Price: 10.23,
	}

	coffee := Drinks{
		Name:  "Vietnamese ICed Tea",
		Price: 11.22,
	}

	bha.printInfo()
	coffee.printInfo()

}
