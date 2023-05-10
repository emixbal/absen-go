package requests

import "github.com/gookit/validate"

type OffSiteAssignment struct {
	Code        string `json:"code" xml:"code" form:"code" validate:"required"`
	Description string `json:"description" xml:"description" form:"description" validate:"required"`
}

// Messages you can custom validator error messages.
func (f OffSiteAssignment) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f OffSiteAssignment) Translates() map[string]string {
	return validate.MS{
		"Code":        "code",
		"Description": "description",
	}
}
