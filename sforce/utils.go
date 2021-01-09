package sforce

import (
	"io/ioutil"
	"log"
	"regexp"
)

// GetSessionID get salesforce session id
func (s *Sforce) GetSessionID() *Sforce {
	body, err := ioutil.ReadAll(s.Res.Body)
	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile("<sessionId>(.*?)</sessionId>")
	if err != nil {
		log.Fatal(err)
	}

	result := r.FindAllStringSubmatch(string(body), -1)
	s.SessionID = result[0][1]

	return s
}

// GetJobID get salesforce job id
func (s *Sforce) GetJobID() *Sforce {
	body, err := ioutil.ReadAll(s.Res.Body)
	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile("<id>(.*?)</id>")
	if err != nil {
		log.Fatal(err)
	}
	result := r.FindAllStringSubmatch(string(body), -1)
	s.JobID = result[0][1]
	return s
}

// GetBodyString get salesforce response body string
func (s *Sforce) GetBodyString() *Sforce {
	body, err := ioutil.ReadAll(s.Res.Body)
	if err != nil {
		log.Fatal(err)
	}
	s.BodyString = string(body)
	return s
}
