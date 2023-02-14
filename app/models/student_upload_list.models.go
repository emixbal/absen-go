package models

import (
	"absen-go/config"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func StudentsUploadList(class_id string) Response {
	var res Response
	var students []Student
	var student Student
	var class Class

	db := config.GetDBInstance()

	int_class_id, _ := strconv.Atoi(class_id)

	if result := db.First(&class, int_class_id); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = 400
			res.Message = "class not exist"
			return res
		}

		log.Print("error check class is exist")
		log.Print(result.Error)

		res.Status = 500
		res.Message = "error check class is exist"
		return res
	}

	if result := db.Where("class_id = ?", int_class_id).Find(&students); result.Error != nil {
		log.Print("error check class is empty")
		log.Print(result.Error)

		res.Status = 500
		res.Message = "error Check class is empty"
		return res
	}

	if len(students) > 0 {
		res.Status = 400
		res.Message = "kelas tidak kosong"
		return res
	}

	file, err := os.Open("./files/students_files_temp/" + class_id + ".csv")
	if err != nil {
		log.Println(err)
		res.Status = 500
		res.Message = "err os.Open"
		return res
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
		res.Status = 500
		res.Message = "Unable to parse file as CSV"
		return res
	}

	for _, val := range records {

		if len([]rune(val[0])) < 1 {
			res.Status = 400
			res.Message = "Ada nama yang kosong"
			return res
		}
		if len([]rune(val[1])) < 1 {
			res.Status = 400
			res.Message = "Ada nisn/nis yang kosong"
			return res
		}
		if len([]rune(val[2])) < 1 {
			res.Status = 400
			res.Message = "Ada nisn/nis yang kosong"
			return res
		}

		nisn_nis := strings.Split(val[1], "/")

		if len(nisn_nis) != 2 {
			res.Status = 400
			res.Message = "format nis/nisn salah " + val[1]
			return res
		}

		student.Name = val[0]
		student.NIS = strings.Trim(nisn_nis[0], " ")
		student.NISN = strings.Trim(nisn_nis[1], " ")
		student.Code = strings.Trim(val[2], " ")
		student.ClassID = int_class_id

		students = append(students, student)
	}

	if len(students) == 0 {
		res.Status = 400
		res.Message = "File kosong"
		return res
	}

	if result := db.Create(&students); result.Error != nil {
		log.Print("error create batch students")
		log.Print(result.Error)

		res.Status = 500
		res.Message = "error create batch students"
		return res
	}

	res.Status = 200
	res.Message = "ok"
	res.Data = students

	return res
}
