package endpoints

import (
	"api/testing/petitions"
	"fmt"
	"log"
	"net/url"
	"strconv"
)


type Borrow struct {
	PatronId  string  `json:"patronId"`
	BookId    int  `json:"bookId"`
	DueDate   string  `json:"dueDate"`
}

type borrowTest struct {
	address url.URL
}

func (b borrowTest) borrowBook(patronId string, bookId string){
	id, err := strconv.Atoi(bookId)
	if err != nil {
		log.Fatalf("Could not cast to int")
	}
	

	body := Borrow{
		PatronId: patronId,
		BookId:  id,
		DueDate:  "2024-10-08",
	}
	petitions.BodyRequest("POST", b.address, body)
}

func (b borrowTest) checkIsBorrowed(bookId string){
	b.address.Path = fmt.Sprintf("borrow/availability/%s", bookId)
	petitions.SimpleRequest(b.address)
}

func (b borrowTest) getAll(){
	petitions.SimpleRequest(b.address)
}

func (b borrowTest) returnBook(bookId string){
	b.address.Path = fmt.Sprintf("borrow/return/%s", bookId)
	petitions.SimpleRequest(b.address)
}

func InitBorrow(){
	address := url.URL{
		Scheme:      "http",
		Host:        BASE_URL,
		Path:        "borrow",
	}

	borrowTest := borrowTest{
		address: address,
	}

	bookId := "10"
	borrowTest.borrowBook("d2be2fe1-1a63-472f-bbce-414586b88b3e", bookId)
	borrowTest.checkIsBorrowed(bookId)

	borrowTest.returnBook(bookId)
	borrowTest.checkIsBorrowed(bookId)

	borrowTest.getAll()
}

