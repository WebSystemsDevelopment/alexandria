package main

import (
	endpoint "api/testing/endpoints"
	"log"
)

func main(){
    log.SetFlags(log.Lshortfile)
    // endpoint.InitPatron()
    // endpoint.InitBook()
    endpoint.InitBorrow()
}

