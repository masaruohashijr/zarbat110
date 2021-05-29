package primary

import "CPaaS/pkg/domains/calls"

type HTTPPrimaryAdapter struct {
	service calls.PrimaryPort
}

func NewHTTPPrimaryAdapter(s calls.PrimaryPort) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{
		s,
	}
}
