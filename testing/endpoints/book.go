package endpoints

import (
	"api/testing/petitions"
	"fmt"
	"net/url"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// e.POST("/books", bookHandler.CreateBook)
// e.GET("/books", bookHandler.GetAllBooks)
// e.GET("/books/:id", bookHandler.GetBookByID)
// e.PUT("/books/:id", bookHandler.UpdateBook)
// e.DELETE("/books/:id", bookHandler.DeleteBook)

const BASE_URL = "localhost:8080"
const SERVER_PORT = "8080"

type BookTest struct {
	address url.URL
}

func (b BookTest) post() {
	body := Book{
		ID:     0,
		Title:  "The pragmatic programmer",
		Author: "Andrew Hunt",
		ISBN:   "0-201-61622-X",
	}

	petitions.BodyRequest("POST", b.address, body)
}

func (b BookTest) put(id string) {
	body := Book{
		ID:     0,
		Title:  "",
		Author: "Mr bananas",
		ISBN:   "0-101-0000-X",
	}

	b.address.Path = fmt.Sprintf("/books/%s", id)
	petitions.BodyRequest("PUT", b.address, body)
}

func (b BookTest) getAll() {
	petitions.SimpleRequest(b.address)
}

func (b BookTest) getById(id string) {
	b.address.Path = fmt.Sprintf("/books/%s", id)
	petitions.SimpleRequest(b.address)
}

func (b BookTest) remove(id string) {
	b.address.Path = fmt.Sprintf("/books/%s", id)
	petitions.BodyRequest("DELETE", b.address, nil)
}

func InitBook() {
	address := url.URL{
		Scheme:      "http",
		Host:        BASE_URL,
		Path:        "books",
	}

	bookTest := BookTest{
		address: address,
	}

	bookTest.post()
	bookTest.getById("0")
	bookTest.put("0")

	bookTest.post()
	bookTest.remove("0")
	bookTest.getById("0")
}

