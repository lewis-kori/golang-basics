package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	// a book representation
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales    int  `json:"book_sales"`
	Age      int  `json:"age,omitempty"`
	Youtuber bool `json:"is_youtuber"`
}

type sensorReading struct {
	Name     string `json:"name"`
	Capacity string `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	author := Author{Sales: 20000, Youtuber: true}
	book := Book{Title: "New thinking", Author: author}
	byteArray, err := json.MarshalIndent(book, "", " ")

	if err != nil {
		fmt.Println("error is", err)
	}
	fmt.Println(string(byteArray))
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading map[string]interface{}

	err = json.Unmarshal([]byte(jsonString), &reading)

	fmt.Printf("%+v\n", reading["capacity"])

}
