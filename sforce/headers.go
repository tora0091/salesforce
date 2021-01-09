package sforce

import "log"

// SetHeaderContentTypeXML is 'Content-Type: text/xml; charset=UTF-8'
func (s *Sforce) SetHeaderContentTypeXML() *Sforce {
	s.Req.Header.Set("Content-Type", "text/xml; charset=UTF-8")
	return s
}

// SetHeaderContentTypeTextCsv is 'Content-type: text/csv; charset=UTF-8'
func (s *Sforce) SetHeaderContentTypeTextCsv() *Sforce {
	s.Req.Header.Set("Content-Type", "text/csv; charset=UTF-8")
	return s
}

// SetHeaderContentTypeApplicationXML is 'Content-Type: application/xml; charset=UTF-8'
func (s *Sforce) SetHeaderContentTypeApplicationXML() *Sforce {
	s.Req.Header.Set("Content-Type", "application/xml; charset=UTF-8")
	return s
}

// SetHeaderSOAPAction is 'SOAPAction: login'
func (s *Sforce) SetHeaderSOAPAction() *Sforce {
	s.Req.Header.Set("SOAPAction", "login")
	return s
}

// SetHeaderBearer is 'Authorization: Bearer [your session id]
func (s *Sforce) SetHeaderBearer() *Sforce {
	if s.SessionID == "" {
		log.Fatal("Error: session id not found.")
	}
	s.Req.Header.Set("Authorization", "Bearer "+s.SessionID)
	return s
}

// SetHeaderXPrettyPrint is 'X-PrettyPrint:1'
func (s *Sforce) SetHeaderXPrettyPrint() *Sforce {
	s.Req.Header.Set("X-PrettyPrint", "1")
	return s
}

// SetHeaderXSFDCSession is 'X-SFDC-Session: [your session id]'
func (s *Sforce) SetHeaderXSFDCSession() *Sforce {
	if s.SessionID == "" {
		log.Fatal("Error: session id not found.")
	}
	s.Req.Header.Set("X-SFDC-Session", s.SessionID)
	return s
}
