package primary

import "CPaaS/pkg/domains/calls"

type CLIPrimaryAdapter struct {
	service calls.PrimaryPort
}

func NewCLIPrimaryAdapter(s calls.PrimaryPort) *CLIPrimaryAdapter {
	return &CLIPrimaryAdapter{
		s,
	}
}
