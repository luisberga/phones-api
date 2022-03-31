package models

import (
	"errors"
	"strings"
)

type Phone struct {
	ID        uint64 `json:"id,omitempty"`
	CompanyID uint64 `json:"company_id,omitempty"`
	Number    string `json:"phone_number,omitempty"`
}

func (p *Phone) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}

	p.format()
	return nil
}

func (p *Phone) validate() error {
	if p.Number == "" {
		return errors.New("name is mandatory")
	}

	return nil
}

func (p *Phone) format() {
	p.Number = strings.TrimSpace(p.Number)
}
