//go-spf is implemnentation of Youtube's SPF.js response spec.
package spf

import (
	"bytes"
	"encoding/json"
)

//SPF spec type
type SPF struct {
	URL        string                            `json:"url"`
	Title      string                            `json:"title"`
	Header     string                            `json:"head"`
	Body       map[string]interface{}            `json:"body"`
	Attributes map[string]map[string]interface{} `json:"attr"`
	Footer     string                            `json:"foot"`
}

//New returns SPF type with initialized maps
func New() SPF {
	return SPF{Body: makeElements(), Attributes: makeAttributes()}
}

//SetURL sets spf url
func (s *SPF) SetURL(url string) {
	s.URL = url
}

//SetTitle sets spf title
func (s *SPF) SetTitle(title string) {
	s.Title = title
}

//SetHeader sets spf header
func (s *SPF) SetHeader(header string) {
	s.Header = header
}

//SetBody sets spf body element content by ID
func (s *SPF) SetBody(domid, v string) {
	s.Body[domid] = v
}

//SetFooter sets spf footer
func (s *SPF) SetFooter(footer string) {
	s.Footer = footer
}

//SetAttribute creates a map of elements that will have their desired attributes changed
func (s *SPF) SetAttribute(domid, attr, v string) {
	mm, ok := s.Attributes[domid]
	if !ok {
		child := makeElements()
		child[attr] = v
		s.Attributes[domid] = child
	} else {
		mm[attr] = v
		s.Attributes[domid] = mm
	}
}

//EncodeJSON encode json
func (s *SPF) EncodeJSON() (buf *bytes.Buffer, err error) {
	buf = new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err = encoder.Encode(s)
	return
}

//make Elements map
func makeElements() map[string]interface{} {
	return make(map[string]interface{})
}

//make Attributes map
func makeAttributes() map[string]map[string]interface{} {
	return make(map[string]map[string]interface{})
}
