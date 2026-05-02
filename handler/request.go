package handler

import (
	"fmt"
	"strings"

	"github.com/BrunoNeves1995/api-portunities/schemas"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
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

// UpdateOpening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	//If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	//If none of the fields were provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")

}

func applyOpeningUpdate(opening *schemas.Opening, request UpdateOpeningRequest) {
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
}

func errParamsIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
