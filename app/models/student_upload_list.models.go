package models

import (
	"absen-go/config"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func StudentsUploadList(class_id string) Response {
	var res Response
	var students []Student
	var student Student

	db := config.GetDBInstance()

	int_class_id, _ := strconv.Atoi(class_id)
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
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	for _, val := range records {
		nisn_nis := strings.Split(val[1], "/")

		student.Name = val[0]
		student.NIS = strings.Trim(nisn_nis[0], " ")
		student.NISN = strings.Trim(nisn_nis[1], " ")
		student.ClassID = int_class_id

		students = append(students, student)
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
