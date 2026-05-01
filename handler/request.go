package handler

import (
	"fmt"
	"strings"
)

type CreateOpiningrequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpiningrequest) Validate() error {
	if r.Role == "" || strings.TrimSpace(r.Role) == "" {
		return errParamsIsRequired("role", "string")
	}
	if r.Company == "" || strings.TrimSpace(r.Company) == "" {
		return errParamsIsRequired("company", "string")
	}
	if r.Location == "" || strings.TrimSpace(r.Location) == "" {
		return errParamsIsRequired("location", "string")
	}
	if r.Link == "" || strings.TrimSpace(r.Link) == "" {
		return errParamsIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamsIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamsIsRequired("salary", "int64")
	}
	return nil
}

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
