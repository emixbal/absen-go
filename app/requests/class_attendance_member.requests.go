package requests

import "github.com/gookit/validate"

type AddClassAttendanceMember struct {
	Code string `json:"code" xml:"code" form:"code" validate:"required"`
}

// Messages you can custom validator error messages.
func (f AddClassAttendanceMember) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f AddClassAttendanceMember) Translates() map[string]string {
	return validate.MS{
		"Code": "code",
	}
}
