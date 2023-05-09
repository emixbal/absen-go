package requests

import "github.com/gookit/validate"

type AddClassAttendanceRemark struct {
	Date         string `json:"date" xml:"date" form:"date" validate:"required"`
	Code         string `json:"code" xml:"code" form:"code" validate:"required"`
	RemarkTypeID int    `json:"remark_type_id" xml:"remark_type_id" form:"remark_type_id" validate:"required"`
	Remark       string `json:"remark" xml:"remark" form:"remark" validate:"required"`
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
		"Date":         "date",
		"Code":         "code",
		"RemarkTypeID": "remark_type_id",
		"Remark":       "remark",
	}
}
