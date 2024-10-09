package endpoints

import (
	"api/testing/petitions"
	"net/url"
)


type Patron struct {
	Name              string  `json:"name"`
	Email             string  `json:"email"`
}

// e.POST("/patron", patronHandler.CreatePatron)
// e.GET("/patron", patronHandler.GetAllPatrons)

type PatronTest struct {
	address url.URL
}

func (p PatronTest) post(){
	body := Patron{
		Name:             "Juan",
		Email:            "Lima@lima.com",
	}

	petitions.BodyRequest("POST", p.address, body)
}

func (p PatronTest) getAll(){
	petitions.SimpleRequest(p.address)
}

func InitPatron(){
	address := url.URL{
		Scheme:      "http",
		Host:        BASE_URL,
		Path:        "patron",
	}

	patronTest := PatronTest{
		address: address,
	}

	patronTest.post()
	patronTest.getAll()
}
