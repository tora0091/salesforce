package sforce

import (
	"net/http"
)

// Sforce struct, salesforce object, set reqest, response and session id...
type Sforce struct {
	Req        *http.Request
	Res        *http.Response
	SessionID  string
	JobID      string
	BodyString string
}

// NewSforce return salesforce object
func NewSforce() *Sforce {
	return &Sforce{}
}

// Login is salesforce login process
func (s *Sforce) Login() *Sforce {
	s.SetLoginRequest().
		SetHeaderContentTypeXML().
		SetHeaderSOAPAction().
		StartLogin().
		GetSessionID()
	defer s.Res.Body.Close()
	return s
}

// GetObjectList get object list to salesforce
func (s *Sforce) GetObjectList() *Sforce {
	s.SetObjectRequest().
		SetHeaderBearer().
		SetHeaderXPrettyPrint().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}

// GetObjectDescribe get object describe to salesforce
func (s *Sforce) GetObjectDescribe(objectName string) *Sforce {
	s.SetObjectDescribeRequest(objectName).
		SetHeaderBearer().
		SetHeaderXPrettyPrint().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}

// GetJobStatus get job status
func (s *Sforce) GetJobStatus(jobID string) *Sforce {
	s.SetJobStatusRequest(jobID).
		SetHeaderXSFDCSession().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}

// GetJobResult get job/batch result
func (s *Sforce) GetJobResult(jobID, batchID string) *Sforce {
	s.SetJobResultRequest(jobID, batchID).
		SetHeaderXSFDCSession().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}

// GetJobClose is job closed
func (s *Sforce) GetJobClose(jobID string) *Sforce {
	s.SetJobCloseRequest(jobID).
		SetHeaderXSFDCSession().
		SetHeaderContentTypeApplicationXML().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}

// BulkInsert is bulk insert process. you necessary object name ant csv file
func (s *Sforce) BulkInsert(objectName, csvFilePath string) *Sforce {
	// create job
	s.SetInsertJobRequest(objectName).
		SetHeaderXSFDCSession().
		SetHeaderContentTypeApplicationXML().
		StartClient()
	defer s.Res.Body.Close()

	// bulk insert
	s.GetJobID().
		SetBatchRequest(csvFilePath).
		SetHeaderXSFDCSession().
		SetHeaderContentTypeTextCsv().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}

// BulkUpdate is bulk update process. you necessary object name ant csv file, external id field name
func (s *Sforce) BulkUpdate(objectName, externalIDFieldName, csvFilePath string) *Sforce {
	// create job
	s.SetUpdateJobRequest(objectName, externalIDFieldName).
		SetHeaderXSFDCSession().
		SetHeaderContentTypeApplicationXML().
		StartClient()
	defer s.Res.Body.Close()

	// bulk update
	s.GetJobID().
		SetBatchRequest(csvFilePath).
		SetHeaderXSFDCSession().
		SetHeaderContentTypeTextCsv().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}

// BulkDelete is bulk delete process. you necessary object name ant csv file
func (s *Sforce) BulkDelete(objectName, csvFilePath string) *Sforce {
	// create job
	s.SetDeleteJobRequest(objectName).
		SetHeaderXSFDCSession().
		SetHeaderContentTypeApplicationXML().
		StartClient()
	defer s.Res.Body.Close()

	// bulk delete
	s.GetJobID().
		SetBatchRequest(csvFilePath).
		SetHeaderXSFDCSession().
		SetHeaderContentTypeTextCsv().
		StartClient().
		GetBodyString()
	defer s.Res.Body.Close()
	return s
}
