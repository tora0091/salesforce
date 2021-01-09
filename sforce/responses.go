package sforce

import (
	"log"
	"net/http"
)

// StartLogin make salesforce http response client
func (s *Sforce) StartLogin() *Sforce {
	return s.StartClient()
}

// StartClient make http response client
func (s *Sforce) StartClient() *Sforce {
	client := new(http.Client)
	resp, err := client.Do(s.Req)
	if err != nil {
		log.Fatal(err)
	}
	s.Res = resp
	return s
}
