package requests

import "github.com/gookit/validate"

type KambingFeedingAddDataForm struct {
	KambingID string `json:"kambing_id" xml:"kambing_id" form:"kambing_id" validate:"required|min:1"`
	FeedID    string `json:"feed_id" xml:"feed_id" form:"feed_id" validate:"required|min:1"`
}

// Messages you can custom validator error messages.
func (f KambingFeedingAddDataForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
		"int":      "{field} must int",
	}
}

// Translates you can custom field translates.
func (f KambingFeedingAddDataForm) Translates() map[string]string {
	return validate.MS{
		"KambingID": "Kambing ID",
		"FeedID":    "Feed ID",
	}
}

type KambingMedecineAddDataForm struct {
	KambingID  string `json:"kambing_id" xml:"kambing_id" form:"kambing_id" validate:"required|min:1"`
	MedicineID string `json:"medicine_id" xml:"medicine_id" form:"medicine_id" validate:"required|min:1"`
}

// Messages you can custom validator error messages.
func (f KambingMedecineAddDataForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
		"int":      "{field} must int",
	}
}

// Translates you can custom field translates.
func (f KambingMedecineAddDataForm) Translates() map[string]string {
	return validate.MS{
		"KambingID":  "Kambing ID",
		"MedicineID": "Medicine ID",
	}
}

type KambingWeighingAddDataForm struct {
	KambingID string `json:"kambing_id" xml:"kambing_id" form:"kambing_id" validate:"required|min:1"`
	Weight    string `json:"weight" xml:"weight" form:"weight" validate:"required|min:1"`
}

// Messages you can custom validator error messages.
func (f KambingWeighingAddDataForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
		"int":      "{field} must int",
	}
}

// Translates you can custom field translates.
func (f KambingWeighingAddDataForm) Translates() map[string]string {
	return validate.MS{
		"KambingID": "Kambing ID",
		"Weight":    "Weight",
	}
}

type KambingWeighingDeleteForm struct {
	WeightID string `json:"weight_id" xml:"weight_id" form:"weight_id" validate:"required|min:1"`
}

// Messages you can custom validator error messages.
func (f KambingWeighingDeleteForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
		"int":      "{field} must int",
	}
}

// Translates you can custom field translates.
func (f KambingWeighingDeleteForm) Translates() map[string]string {
	return validate.MS{
		"WeightID": "Weight ID",
	}
}

type KambingFeedingDeleteForm struct {
	FeedID string `json:"feed_id" xml:"feed_id" form:"feed_id" validate:"required|min:1"`
}

// Messages you can custom validator error messages.
func (f KambingFeedingDeleteForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
		"int":      "{field} must int",
	}
}

// Translates you can custom field translates.
func (f KambingFeedingDeleteForm) Translates() map[string]string {
	return validate.MS{
		"FeedID": "Feed ID",
	}
}

type KambingMedicineDeleteForm struct {
	MedicineID string `json:"medicine_id" xml:"medicine_id" form:"medicine_id" validate:"required|min:1"`
}

// Messages you can custom validator error messages.
func (f KambingMedicineDeleteForm) Messages() map[string]string {
	return validate.MS{
		"required": "{field} is required.",
		"int":      "{field} must int",
	}
}

// Translates you can custom field translates.
func (f KambingMedicineDeleteForm) Translates() map[string]string {
	return validate.MS{
		"MedicineID": "Medicine ID",
	}
}
