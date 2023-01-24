package requests

import "github.com/gookit/validate"

type AddExtracurricularAttendanceStudent struct {
	Name string `json:"name" xml:"string" form:"name" validate:"required"`
}

// Messages you can custom validator error messages.
func (f AddExtracurricularAttendanceStudent) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f AddExtracurricularAttendanceStudent) Translates() map[string]string {
	return validate.MS{
		"Name": "name",
	}
}
