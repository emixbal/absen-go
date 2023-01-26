package requests

import "github.com/gookit/validate"

type AddClassAttendanceStudent struct {
	StudentID int `json:"student_id" xml:"student_id" form:"student_id" validate:"required"`
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
		"StudentID": "student_id",
	}
}
