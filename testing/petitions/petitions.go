package petitions

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func marshalBody(body *any) (*[]byte) {
    jsonBody, err := json.Marshal(&body)
    if err != nil {
	log.Fatalf("Could not marshal book body: %v", err)
    }

    return &jsonBody
}

func BodyRequest(httpMethod string, postURL url.URL, body any) {
    jsonBody := marshalBody(&body)
    bodyReader := bytes.NewReader(*jsonBody)

    request, err  := http.NewRequest(httpMethod, postURL.String(), bodyReader)
    if err != nil {
	log.Fatalf("Could not create %s request: %v", httpMethod, err)
    }

    if httpMethod != http.MethodPut || httpMethod != http.MethodPost {
	request.Header.Set("Content-Type", "application/json")
    }


    resp, err := http.DefaultClient.Do(request);
    handleRequest(httpMethod, *resp, err)
}

func SimpleRequest(postURL url.URL) {
    resp, err  := http.Get(postURL.String())
    handleRequest("GET", *resp, err)
}

func handleRequest(httpMethod string, resp http.Response, err error) {
    if err != nil {
	log.Fatalf("Failed to execute %s request: %v", httpMethod, err)
    }

    defer func() {
	if resp.Body.Close() != nil {
	    log.Fatalf("Could not close body: %v", err)
	}
    }()

    respBody, err  := io.ReadAll(resp.Body)
    log.Print(resp.Status)
    log.Println("Response:\n", string(respBody))
}
