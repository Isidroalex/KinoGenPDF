package main

import (
	pdf "Breeding/pdf"
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	//r := pdf.NewRequestPdf("")

	//father := Dog{
	//	Type:     "среднеазиатская овчарка",
	//	Nickname: "Dastan",
	//	Sex:      "к",
	//	Stamp:    "2139",
	//	Pedigree: "UKU.R.0047773",
	//	Owner: &Owner{
	//		FIO:     "Кузнецов А.А.",
	//		Contact: "Донецк, Флотская, 68",
	//	},
	//}
	//
	//mother := Dog{
	//	Type:     "среднеазиатская овчарка",
	//	Nickname: "Анзурат Майсун Ф-Айпери",
	//	Sex:      "c",
	//	Stamp:    "CKF 4707",
	//	Pedigree: "5582067 Р",
	//	Owner: &Owner{
	//		FIO:     "Каратаева Е.В",
	//		Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
	//	},
	//	Breeder: &Breeder{
	//		FIO:     "Каратаева Е.В",
	//		Name:    "Анзурат Майсун",
	//		Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
	//		Email:   "breeder@email.com",
	//	},
	//}
	//
	//gaplan := Puppy{
	//	Dog: Dog{
	//		Type:     "среднеазиатская овчарка",
	//		Nickname: "Гаплан",
	//		Sex:      "к",
	//		Stamp:    "CKF 6059",
	//		Pedigree: "",
	//		Owner: &Owner{
	//			FIO:     " ",
	//			Contact: " ",
	//		},
	//		Breeder: &Breeder{
	//			FIO:     "Каратаева Е.В",
	//			Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
	//			Name:    "Анзурат Майсун",
	//			Email:   "tula-skif@mail.ru",
	//		},
	//		Father: &father,
	//		Mother: &mother,
	//	},
	//	Color:    "бел-чер",
	//	Birthday: "28.03.2023",
	//	WoolType: "",
	//	Description: struct {
	//		Comment  string
	//		Defect   bool
	//		Revision int
	//	}{"Замечаний нет", false, 0},
	//	IncreaseIndex: func(a int) int {
	//		return a + 1
	//	},
	//}
	//
	//galsan := Puppy{
	//	Dog: Dog{
	//		Type:     "среднеазиатская овчарка",
	//		Nickname: "Гаплан",
	//		Sex:      "к",
	//		Stamp:    "CKF 6060",
	//		Pedigree: "",
	//		Owner: &Owner{
	//			FIO:     "Каратаева Е.В",
	//			Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
	//		},
	//	},
	//	Color:    "бел-чер",
	//	Birthday: "28 марта 2023",
	//	WoolType: "",
	//	Description: struct {
	//		Comment  string
	//		Defect   bool
	//		Revision int
	//	}{"Замечаний нет", false, 0},
	//	IncreaseIndex: func(a int) int {
	//		return a + 1
	//	},
	//}
	//
	//ganzhin := Puppy{
	//	Dog: Dog{
	//		Type:     "среднеазиатская овчарка",
	//		Nickname: "Гаплан",
	//		Sex:      "с",
	//		Stamp:    "CKF 6064",
	//		Pedigree: "",
	//		Owner: &Owner{
	//			FIO:     "Каратаева Е.В",
	//			Contact: "142660 МО, Орехово-Зуевский р-н, Савостьяново, 66",
	//		},
	//	},
	//	Color:    "бел-чер",
	//	Birthday: "28 марта 2023",
	//	WoolType: "",
	//	Description: struct {
	//		Comment  string
	//		Defect   bool
	//		Revision int
	//	}{"Замечаний нет", false, 0},
	//	IncreaseIndex: func(a int) int {
	//		return a + 1
	//	},
	//}
	//
	//skiff := DocumentCenter{
	//	Name:         "ТООО «Кинологический центр «СКИФ» г. Тула",
	//	Email:        "tula-skif@mail.ru",
	//	Mobile:       "+7 910 701-55-99",
	//	Federation:   "РФСС",
	//	FolderNumber: "7107",
	//	Address:      "300025 Тула, проспект Ленина, 104-201",
	//	StumpCode:    "CKF",
	//}
	//
	//instructor := Personal{
	//	"Девятиярова Е.А.",
	//	"Тульская обл, Ленинский р-н, Нижняя Китаевка, 24",
	//}
	//
	//stumpMan := Personal{"Аргунова Л.А.", ""}
	//
	//data := Mating{
	//	Identification: "31 декабря 2022",
	//	First:          "01 января 2023",
	//	Control:        "02 января 2023",
	//	Address:        "Донецк, Флотская, 68",
	//	Male:           &father,
	//	Female:         &mother,
	//	DCenter:        &skiff,
	//	Personal:       &instructor,
	//}
	//
	//litter1 := Litter{
	//	Stump:    "CKF",
	//	ActDate:  "20 августа 1944",
	//	Birthday: "16 декабря 2016",
	//	Control:  "",
	//	Address:  "",
	//	Male:     &father,
	//	Female:   &mother,
	//	Puppy:    []*Puppy{&gaplan, &galsan, &ganzhin},
	//	DCenter:  &skiff,
	//	Personal: &instructor,
	//	Claim:    &stumpMan,
	//	Chief:    &instructor,
	//}
	//
	//gaplanPC := PuppyCard{
	//	Puppy:   gaplan,
	//	DCenter: skiff,
	//}

	//templatePath := "blankMating.html"
	////templatePath := "blank4.html"
	//outputPath := fmt.Sprintf("%s_акт_вязки_%s.pdf", data.Personal.FIO, data.Female.Nickname)
	//if err := r.ParseTemplate(templatePath, data); err == nil {
	//	args := []string{"no-pdf-compression"}
	//	ok, _ := r.GeneratePDF(outputPath, args)
	//	fmt.Println(ok, "pdf generated successfully")
	//} else {
	//	fmt.Println(err)
	//}
	//
	//templatePath = "blankAgreement.html"
	////templatePath := "blank4.html"
	//outputPath = fmt.Sprintf("%s_согласие_ПД_%s.pdf", data.Personal.FIO, data.Female.Nickname)
	//if err := r.ParseTemplate(templatePath, data); err == nil {
	//	args := []string{"no-pdf-compression"}
	//
	//	ok, _ := r.GeneratePDF(outputPath, args)
	//	fmt.Println(ok, "pdf generated successfully")
	//} else {
	//	fmt.Println(err)
	//}
	//
	//templatePath = "blankSurvey.html"
	////templatePath := "blank4.html"
	//outputPath = fmt.Sprintf("%s_обследование_%s.pdf", litter1.Personal.FIO, litter1.Birthday)
	//if err := r.ParseTemplate(templatePath, litter1); err == nil {
	//	args := []string{"no-pdf-compression"}
	//	ok, _ := r.GeneratePDF(outputPath, args)
	//	fmt.Println(ok, "pdf generated successfully")
	//} else {
	//	fmt.Println(err)
	//}
	//
	//templatePath = "blankRegister.html"
	////templatePath := "blank4.html"
	//outputPath = fmt.Sprintf("%s_регистрация_%s.pdf", litter1.Personal.FIO, litter1.Birthday)
	//if err := r.ParseTemplate(templatePath, litter1); err == nil {
	//	args := []string{"no-pdf-compression"}
	//	ok, _ := r.GeneratePDF(outputPath, args)
	//	fmt.Println(ok, "pdf generated successfully")
	//} else {
	//	fmt.Println(err)
	//}
	//
	//templatePath = "blankPuppyCard.html"
	////templatePath := "blank4.html"
	//outputPath = fmt.Sprintf("Щенячка_%s.pdf", gaplanPC.Nickname)
	//if err := r.ParseTemplate(templatePath, gaplanPC); err == nil {
	//	args := []string{"no-pdf-compression"}
	//	ok, _ := r.GeneratePDF(outputPath, args)
	//	fmt.Println(ok, "pdf generated successfully")
	//} else {
	//	fmt.Println(err)
	//}

	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on port 3333")
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		return
	}
	defer r.Body.Close()

	if r.Method == "POST" {
		t := time.Now().Unix()
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "Акт вязки.pdf")
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "СогласиеПД.pdf")
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "Обследование.pdf")
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "Регистрация.pdf")

		Data := Input{}
		r := pdf.NewRequestPdf("")
		if err := json.Unmarshal(body, &Data); err != nil {
			fmt.Fprintf(w, "Received error: %s", err)
			w.WriteHeader(http.StatusNotAcceptable)
		}
		mating := Mating{}
		mating.Construct(Data)
		litter := Litter{}
		litter.Construct(Data)

		if _, err := os.Stat(strconv.FormatInt(t, 10) + "/"); os.IsNotExist(err) {
			errDir := os.Mkdir(strconv.FormatInt(t, 10)+"/", 0777)
			if errDir != nil {
				http.Error(w, "Internal server error", 500)
				return
			}
		}

		zipFile, err := os.Create(strconv.FormatInt(t, 10) + "/" + "archive.zip")
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer zipFile.Close()

		zipWriter := zip.NewWriter(zipFile)
		defer zipWriter.Close()

		templatePath := "blankMating.html"
		outputPath := fmt.Sprintf(strconv.FormatInt(t, 10) + "/" + "Акт вязки.pdf")
		if err := r.ParseTemplate(templatePath, mating); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := r.GeneratePDF(outputPath, args)
			fmt.Println(ok, "pdf generated successfully")
		} else {
			fmt.Fprintf(w, "Received error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		pdfFile, err := os.Open(outputPath)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer pdfFile.Close()

		zipEntry, err := zipWriter.Create("Акт вязки.pdf")
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		_, err = io.Copy(zipEntry, pdfFile)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}

		templatePath = "blankAgreement.html"
		outputPath = fmt.Sprintf(strconv.FormatInt(t, 10) + "/" + "СогласиеПД.pdf")
		if err := r.ParseTemplate(templatePath, mating); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := r.GeneratePDF(outputPath, args)
			fmt.Println(ok, "pdf generated successfully")
		} else {
			fmt.Fprintf(w, "Received error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		pdfFile, err = os.Open(outputPath)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer pdfFile.Close()

		zipEntry, err = zipWriter.Create("СогласиеПД.pdf")
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		_, err = io.Copy(zipEntry, pdfFile)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}

		templatePath = "blankSurvey.html"
		outputPath = fmt.Sprintf(strconv.FormatInt(t, 10) + "/" + "Обследование.pdf")
		if err := r.ParseTemplate(templatePath, litter); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := r.GeneratePDF(outputPath, args)
			fmt.Println(ok, "pdf generated successfully")
		} else {
			fmt.Fprintf(w, "Received error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		pdfFile, err = os.Open(outputPath)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer pdfFile.Close()

		zipEntry, err = zipWriter.Create("Обследование.pdf")
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		_, err = io.Copy(zipEntry, pdfFile)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}

		templatePath = "blankRegister.html"
		outputPath = fmt.Sprintf(strconv.FormatInt(t, 10) + "/" + "Регистрация.pdf")
		if err := r.ParseTemplate(templatePath, litter); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := r.GeneratePDF(outputPath, args)
			fmt.Println(ok, "pdf generated successfully")
		} else {
			fmt.Fprintf(w, "Received error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		pdfFile, err = os.Open(outputPath)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer pdfFile.Close()

		zipEntry, err = zipWriter.Create("Регистрация.pdf")
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		_, err = io.Copy(zipEntry, pdfFile)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=matingAct.pdf")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		_, err = io.Copy(w, zipFile)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		fmt.Println("pdf copied successfully")

	} else {
		// Отправка заголовков CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Здесь можно указать конкретный источник
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//Отправка ответа клиенту
		w.WriteHeader(http.StatusOK)
	}

}
