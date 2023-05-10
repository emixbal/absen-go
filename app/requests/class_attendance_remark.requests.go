package requests

import "github.com/gookit/validate"

type AddClassAttendanceRemark struct {
	Date        string `json:"date" xml:"date" form:"date" validate:"required"`
	Code        string `json:"code" xml:"code" form:"code" validate:"required"`
	IsSick      bool   `json:"is_sick" xml:"is_sick" form:"is_sick"`
	Description string `json:"description" xml:"description" form:"description" validate:"required"`
}

// Messages you can custom validator error messages.
func (f AddClassAttendanceRemark) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
	}
}

// Translates you can custom field translates.
func (f AddClassAttendanceRemark) Translates() map[string]string {
	return validate.MS{
		"Date":        "date",
		"Code":        "code",
		"IsSick":      "is_sick",
		"Description": "description",
	}
}
