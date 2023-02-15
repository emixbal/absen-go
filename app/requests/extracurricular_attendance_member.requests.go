package requests

import "github.com/gookit/validate"

type AddExtracurricularAttendanceMember struct {
	Name string `json:"name" xml:"string" form:"name" validate:"required"`
}

// Messages you can custom validator error messages.
func (f AddExtracurricularAttendanceMember) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f AddExtracurricularAttendanceMember) Translates() map[string]string {
	return validate.MS{
		"Name": "name",
	}
}
