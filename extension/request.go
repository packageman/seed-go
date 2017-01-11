package extension

import (
	request "github.com/parnurzeal/gorequest"
)

func NewRestClient() *request.SuperAgent {
	return request.New()
}
