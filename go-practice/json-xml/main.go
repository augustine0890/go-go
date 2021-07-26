package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

/*
Encoding is the process of transforming a piece of data into a 'coded form'
'coded form' can be transformed back into the original message by decodeing.
*/
type MyJson struct {
	Cat `json:"cat"`
}

type Cat struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type Product struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	SKU  string `json:"sku"`
	Cat  Category
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Price struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type Item struct {
	Comment string `xml:",comment"`
	Price   Price  `xml:"price"`
	Name    string `xml:"name"`
}

func main() {
	myJson := []byte(`{"cat": { "name": "Joey", "age": 8 }}`)
	c := MyJson{}
	err := json.Unmarshal(myJson, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)

	p := Product{ID: 42, Name: "T-Shirt", SKU: "TS33", Cat: Category{ID: 2, Name: "Cloth"}}
	b, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	i := Item{Comment: "Looks good", Price: Price{Text: "4500", Currency: "USD"}, Name: "Dish"}
	x, err := xml.MarshalIndent(i, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(x))
}
