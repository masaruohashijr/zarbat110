package primary

import "CPaaS/pkg/domains/calls"

type port struct {
	driven calls.SecondaryPort
}

func NewService(driven calls.SecondaryPort) calls.PrimaryPort {
	return &port{
		driven,
	}
}

func (p *port) MakeCall() error {
	err := p.driven.MakeCall()
	return err
}
