package main

import (
	book "api/testing/endpoints"
	"log"
)

const BASE_URL = "localhost:8080"

func main(){
    log.SetFlags(log.Llongfile)
    book.Init()
}

