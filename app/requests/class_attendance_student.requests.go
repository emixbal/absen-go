package requests

import "github.com/gookit/validate"

type AddClassAttendanceStudent struct {
	Code string `json:"code" xml:"code" form:"code" validate:"required"`
}

// Messages you can custom validator error messages.
func (f AddClassAttendanceStudent) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f AddClassAttendanceStudent) Translates() map[string]string {
	return validate.MS{
		"Code": "code",
	}
}
