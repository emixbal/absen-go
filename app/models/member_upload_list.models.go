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

func MembersUploadList(class_id string) Response {
	var res Response
	var members []Member
	var member Member
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

	if result := db.Where("class_id = ?", int_class_id).Find(&members); result.Error != nil {
		log.Print("error check class is empty")
		log.Print(result.Error)

		res.Status = 500
		res.Message = "error Check class is empty"
		return res
	}

	if len(members) > 0 {
		res.Status = 400
		res.Message = "kelas tidak kosong"
		return res
	}

	file, err := os.Open("./files/members_files_temp/" + class_id + ".csv")
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

	var codeMap []string
	for index, val := range records {
		if index != 0 {
			if len(val) < 4 {
				res.Status = 400
				res.Message = "Format kolom salah. Harus berformat:Nama,NIS/NISN,NomerKartu,NBM"
				return res
			}
			if len([]rune(val[0])) < 1 {
				res.Status = 400
				res.Message = "Ada nama yang kosong"
				return res
			}
			if len([]rune(val[2])) < 1 {
				res.Status = 400
				res.Message = "Ada nomer kartu yang kosong"
				return res
			}

			nisn_nis := strings.Split(val[1], "/")

			member.Name = val[0]
			if len(nisn_nis) == 2 {
				member.NIS = strings.Trim(nisn_nis[0], " ")
				member.NISN = strings.Trim(nisn_nis[1], " ")
			}
			member.Code = strings.Trim(val[2], " ")
			member.NBM = strings.Trim(val[3], " ")
			member.ClassID = int_class_id

			members = append(members, member)

			// cek apakah nomor kartu sudah ada pada map
			for _, value := range codeMap {
				if value == member.Code {
					res.Status = 400
					res.Message = "Nomor kartu " + value + " duplikat"
					return res
				}
			}
			// tambahkan nomor kartu pada map
			codeMap = append(codeMap, member.Code)
		}
	}

	if len(members) == 0 {
		res.Status = 400
		res.Message = "File kosong"
		return res
	}

	if result := db.Create(&members); result.Error != nil {
		log.Print("error create batch members")
		log.Print(result.Error)

		res.Status = 500
		res.Message = "error create batch members"
		return res
	}

	res.Status = 200
	res.Message = "ok"
	res.Data = members

	return res
}
