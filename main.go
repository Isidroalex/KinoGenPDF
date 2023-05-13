package main

import (
	pdf "Breeding/pdf"
	"fmt"
)

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
		Breeder: &Breeder{
			FIO:     "Каратаева Е.В",
			Name:    "Анзурат Майсун",
			Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
			Email:   "breeder@email.com",
		},
	}

	gaplan := Puppy{
		Dog: Dog{
			Type:     "среднеазиатская овчарка",
			Nickname: "Гаплан",
			Sex:      "к",
			Stamp:    "CKF 6059",
			Pedigree: "",
			Owner: &Owner{
				FIO:     " ",
				Contact: " ",
			},
			Breeder: &Breeder{
				FIO:     "Каратаева Е.В",
				Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
				Name:    "Анзурат Майсун",
				Email:   "tula-skif@mail.ru",
			},
			Father: &father,
			Mother: &mother,
		},
		Color:    "бел-чер",
		Birthday: "28.03.2023",
		WoolType: "",
		Description: struct {
			Comment  string
			Defect   bool
			Revision int
		}{"Замечаний нет", false, 0},
		IncreaseIndex: func(a int) int {
			return a + 1
		},
	}

	galsan := Puppy{
		Dog: Dog{
			Type:     "среднеазиатская овчарка",
			Nickname: "Гаплан",
			Sex:      "к",
			Stamp:    "CKF 6060",
			Pedigree: "",
			Owner: &Owner{
				FIO:     "Каратаева Е.В",
				Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
			},
		},
		Color:    "бел-чер",
		Birthday: "28 марта 2023",
		WoolType: "",
		Description: struct {
			Comment  string
			Defect   bool
			Revision int
		}{"Замечаний нет", false, 0},
		IncreaseIndex: func(a int) int {
			return a + 1
		},
	}

	ganzhin := Puppy{
		Dog: Dog{
			Type:     "среднеазиатская овчарка",
			Nickname: "Гаплан",
			Sex:      "с",
			Stamp:    "CKF 6064",
			Pedigree: "",
			Owner: &Owner{
				FIO:     "Каратаева Е.В",
				Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
			},
		},
		Color:    "бел-чер",
		Birthday: "28 марта 2023",
		WoolType: "",
		Description: struct {
			Comment  string
			Defect   bool
			Revision int
		}{"Замечаний нет", false, 0},
		IncreaseIndex: func(a int) int {
			return a + 1
		},
	}

	skiff := DocumentCenter{
		Name:         "ТООО «Кинологический центр «СКИФ» г. Тула",
		Email:        "tula-skif@mail.ru",
		Mobile:       "+7 910 701-55-99",
		Federation:   "РФСС",
		FolderNumber: "7107",
		Address:      "300025 Тула, проспект Ленина, 104-201",
	}

	instructor := Personal{
		"Девятиярова Е.А.",
		"Тульская обл, Ленинский р-н, Нижняя Китаевка, 24",
	}

	stumpMan := Personal{"Аргунова Л.А.", ""}

	data := Mating{
		Identification: "31 декабря 2022",
		First:          "01 января 2023",
		Control:        "02 января 2023",
		Address:        "Донецк, Флотская, 68",
		Male:           &father,
		Female:         &mother,
		DCenter:        &skiff,
		Personal:       &instructor,
	}

	litter1 := Litter{
		Stump:    "CKF",
		ActDate:  "20 августа 1944",
		Birthday: "16 декабря 2016",
		Control:  "",
		Address:  "",
		Male:     &father,
		Female:   &mother,
		Puppy:    []*Puppy{&gaplan, &galsan, &ganzhin},
		DCenter:  &skiff,
		Personal: &instructor,
		Claim:    &stumpMan,
		Chief:    &instructor,
	}

	gaplanPC := PuppyCard{
		Puppy:   gaplan,
		DCenter: skiff,
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

	templatePath = "blankSurvey.html"
	//templatePath := "blank4.html"
	outputPath = fmt.Sprintf("%s_обследование_%s.pdf", litter1.Personal.FIO, litter1.Birthday)
	if err := r.ParseTemplate(templatePath, litter1); err == nil {
		args := []string{"no-pdf-compression"}
		ok, _ := r.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}

	templatePath = "blankRegister.html"
	//templatePath := "blank4.html"
	outputPath = fmt.Sprintf("%s_регистрация_%s.pdf", litter1.Personal.FIO, litter1.Birthday)
	if err := r.ParseTemplate(templatePath, litter1); err == nil {
		args := []string{"no-pdf-compression"}
		ok, _ := r.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}

	templatePath = "blankPuppyCard.html"
	//templatePath := "blank4.html"
	outputPath = fmt.Sprintf("Щенячка_%s.pdf", gaplanPC.Nickname)
	if err := r.ParseTemplate(templatePath, gaplanPC); err == nil {
		args := []string{"no-pdf-compression"}
		ok, _ := r.GeneratePDF(outputPath, args)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}

}
