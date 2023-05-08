package main

import (
	pdf "Breeding/pdf"
	"fmt"
)

type Data struct {
	Name string
}

func main() {

	r := pdf.NewRequestPdf("")

	father := Dog{
		Type:     "среднеазиатская овчарка",
		Nickname: "Dastan",
		Sex:      "к",
		Stamp:    "2139",
		Pedigree: "UKU.R.0047773",
		Owner: &Owner{
			FIO:     "Кузнецов А.А.",
			Contact: "Донецк, Флотская, 68",
		},
	}

	mother := Dog{
		Type:     "среднеазиатская овчарка",
		Nickname: "Анзурат Майсун Ф-Айпери",
		Sex:      "c",
		Stamp:    "CKF 4707",
		Pedigree: "5582067 Р",
		Owner: &Owner{
			FIO:     "Каратаева Е.В",
			Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
		},
	}

	data := Mating{
		Identification: "31 декабря 2022",
		First:          "01 января 2023",
		Control:        "02 января 2023",
		Address:        "Донецк, Флотская, 68",
		Male:           &father,
		Female:         &mother,
		DCenter: &DocumentCenter{
			Name:         "ТООО «Кинологический центр «СКИФ» г. Тула",
			Email:        "tula-skif@mail.ru",
			Mobile:       "+7 910 701-55-99",
			Federation:   "РФСС",
			FolderNumber: "7107",
		},
		Personal: &Instructor{
			"Девятиярова Е.А.",
			"Тульская обл, Ленинский р-н, Нижняя Китаевка, 24",
		},
	}

	templatePath := "blankMating.html"
	//templatePath := "blank4.html"
	outputPath := fmt.Sprintf("%s_акт_вязки_%s.pdf", data.Personal.FIO, data.Female.Nickname)
	if err := r.ParseTemplate(templatePath, data); err == nil {
		args := []string{"no-pdf-compression"}
		ok, _ := r.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}

	templatePath = "blankAgreement.html"
	//templatePath := "blank4.html"
	outputPath = fmt.Sprintf("%s_согласие_ПД_%s.pdf", data.Personal.FIO, data.Female.Nickname)
	if err := r.ParseTemplate(templatePath, data); err == nil {
		args := []string{"no-pdf-compression"}
		ok, _ := r.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}

}

type DocumentCenter struct {
	Name         string
	Email        string
	Mobile       string
	Federation   string
	FolderNumber string
}

type Mating struct {
	Identification string
	First          string
	Control        string
	Address        string
	Male           *Dog
	Female         *Dog
	DCenter        *DocumentCenter
	Personal       *Instructor
}

type Instructor struct {
	FIO     string
	Address string
}

type Owner struct {
	FIO     string
	Contact string
	Email   string
}

type Breeder struct {
	FIO     string
	Name    string
	Contact string
	Email   string
}

type Dog struct {
	Type        string
	Nickname    string
	Birthday    string
	Sex         string
	Color       string
	Stamp       string
	Pedigree    string
	WoolType    string
	Description struct {
		Comment  string
		Defect   bool
		Revision int
	}
	Father  *Dog
	Mother  *Dog
	Breeder *Breeder
	Owner   *Owner
}
