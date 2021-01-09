package sforce

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../config"
)

// SetObjectRequest set salesforce object request url
func (s *Sforce) SetObjectRequest() *Sforce {
	s.SetRequest(config.C.URL.Base)
	return s
}

// SetObjectDescribeRequest set salesforce object describe request url
func (s *Sforce) SetObjectDescribeRequest(objectName string) *Sforce {
	s.SetRequest(strings.Replace(config.C.URL.Describe, "[salesforce_object]", objectName, 1))
	return s
}

// SetJobStatusRequest set salesforce job id request url
func (s *Sforce) SetJobStatusRequest(jobID string) *Sforce {
	s.SetRequest(strings.Replace(config.C.URL.Batch, "[salesforce_job_id]", jobID, 1))
	return s
}

// SetJobResultRequest set salesforce job result request url
func (s *Sforce) SetJobResultRequest(jobID, batchID string) *Sforce {
	text := strings.Replace(config.C.URL.Result, "[salesforce_job_id]", jobID, 1)
	s.SetRequest(strings.Replace(text, "[salesforce_batch_id]", batchID, 1))
	return s
}

// SetRequest is http request. method is GET
func (s *Sforce) SetRequest(url string) *Sforce {
	var err error
	s.Req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// SetLoginRequest is login request, set username and password, request login to salesforce, method is POST
func (s *Sforce) SetLoginRequest() *Sforce {
	bytes, err := ioutil.ReadFile(config.C.Login.XML)
	if err != nil {
		log.Fatal(err)
	}
	tmp := strings.Replace(string(bytes), "[salesforce_username]", config.C.Username, 1)
	xml := strings.Replace(string(tmp), "[salesforce_password]", config.C.Password, 1)

	s.Req, err = http.NewRequest(http.MethodPost, config.C.Login.URL, strings.NewReader(xml))
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// SetInsertJobRequest is insert job request
func (s *Sforce) SetInsertJobRequest(objectName string) *Sforce {
	return s.SetJobRequest(objectName, config.C.XML.Insert, "")
}

// SetUpdateJobRequest is update job request
func (s *Sforce) SetUpdateJobRequest(objectName, externalIDFieldName string) *Sforce {
	return s.SetJobRequest(objectName, config.C.XML.Update, externalIDFieldName)
}

// SetDeleteJobRequest is delete job request
func (s *Sforce) SetDeleteJobRequest(objectName string) *Sforce {
	return s.SetJobRequest(objectName, config.C.XML.Delete, "")
}

// SetJobRequest is job request
func (s *Sforce) SetJobRequest(objectName, jobXML, externalIDFieldName string) *Sforce {
	bytes, err := ioutil.ReadFile(jobXML)
	if err != nil {
		log.Fatal(err)
	}

	xml := strings.Replace(string(bytes), "[salesforce_object]", objectName, 1)
	if len(externalIDFieldName) != 0 {
		xml = strings.Replace(xml, "[external_id_field_name]", externalIDFieldName, 1)
	}
	readerXML := strings.NewReader(xml)

	s.Req, err = http.NewRequest(http.MethodPost, config.C.URL.Job, readerXML)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// SetJobCloseRequest is job close request
func (s *Sforce) SetJobCloseRequest(jobID string) *Sforce {
	closeURL := strings.Replace(config.C.URL.Close, "[salesforce_job_id]", jobID, 1)

	bytes, err := ioutil.ReadFile(config.C.XML.Close)
	if err != nil {
		log.Fatal(err)
	}

	s.Req, err = http.NewRequest(http.MethodPost, closeURL, strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// SetBatchRequest is batch request
func (s *Sforce) SetBatchRequest(csvFilePath string) *Sforce {
	bytes, err := ioutil.ReadFile(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}

	batchURL := strings.Replace(config.C.URL.Batch, "[salesforce_job_id]", s.JobID, 1)
	s.Req, err = http.NewRequest(http.MethodPost, batchURL, strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatal(err)
	}
	return s
}
