package models

import (
	"errors"
	"strings"
)

type Company struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Company) Prepare() error {
	if err := c.validate(); err != nil {
		return err
	}

	c.format()
	return nil
}

func (c *Company) validate() error {
	if c.Name == "" {
		return errors.New("name is mandatory")
	}

	return nil
}

func (c *Company) format() {
	c.Name = strings.TrimSpace(c.Name)
}
