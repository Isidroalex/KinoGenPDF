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

	http.HandleFunc("/get_pdf", handleRequestPDF)
	http.HandleFunc("/", handleRequestIndex)
	fmt.Println("Server listening on port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
	fmt.Println("Sorry")
}

func handleRequestPDF(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		return
	}
	defer r.Body.Close()

	fmt.Printf("Receive request %s, %s", r.Method, r.Body)

	if r.Method == "POST" {
		t := time.Now().Unix()
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "Акт вязки.pdf")
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "СогласиеПД.pdf")
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "Обследование.pdf")
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "Регистрация.pdf")
		defer os.Remove(strconv.FormatInt(t, 10) + "/" + "Щенячка.pdf")

		Data := Input{}
		requestPdf := pdf.NewRequestPdf("")
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
		if err := requestPdf.ParseTemplate(templatePath, mating); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := requestPdf.GeneratePDF(outputPath, args)
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
		if err := requestPdf.ParseTemplate(templatePath, mating); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := requestPdf.GeneratePDF(outputPath, args)
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
		if err := requestPdf.ParseTemplate(templatePath, litter); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := requestPdf.GeneratePDF(outputPath, args)
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
		if err := requestPdf.ParseTemplate(templatePath, litter); err == nil {
			args := []string{"no-pdf-compression"}
			ok, _ := requestPdf.GeneratePDF(outputPath, args)
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

		var puppyCards []PuppyCard
		for _, puppy := range Data.Puppy {

			owner := Owner{
				FIO:     "",
				Contact: "",
				Email:   "",
			}
			dogStruct := Dog{
				Type:     Data.MotherType,
				Nickname: puppy.Nickname,
				Sex:      puppy.SexPuppy,
				Stamp:    puppy.PuppyStump,
				Father:   mating.Male,
				Mother:   mating.Female,
				Breeder:  mating.Female.Breeder,
				Owner:    &owner,
			}

			puppyStruct := Puppy{

				Color:    puppy.Color,
				Birthday: formatDate(Data.LitterBirthday),
				WoolType: puppy.WoolType,
				Description: struct {
					Comment  string
					Defect   string
					Revision string
				}{
					Comment:  puppy.StatusComment,
					Defect:   puppy.StatusPuppy,
					Revision: Data.RevisionPeriod,
				},
				IncreaseIndex: func(a int) int {
					return a + 1
				},
			}
			puppyStruct.Dog = dogStruct

			puppyCard := PuppyCard{
				Puppy:   puppyStruct,
				DCenter: *litter.DCenter,
			}
			puppyCards = append(puppyCards, puppyCard)
		}

		for i, v := range puppyCards {
			templatePath = "blankPuppyCard.html"
			outputPath = fmt.Sprintf(strconv.FormatInt(t, 10) + "/" + "Щенячка.pdf")
			if err := requestPdf.ParseTemplate(templatePath, v); err == nil {
				args := []string{"no-pdf-compression"}
				ok, _ := requestPdf.GeneratePDF(outputPath, args)
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

			fileArhiveName := fmt.Sprintf("Щенячка %d.pdf", i)

			zipEntry, err = zipWriter.Create(fileArhiveName)
			if err != nil {
				http.Error(w, "Internal server error", 500)
				return
			}
			_, err = io.Copy(zipEntry, pdfFile)
			if err != nil {
				http.Error(w, "Internal server error", 500)
				return
			}

		}

		zipWriter.Close()

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=archive.zip")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		fileInfo, err := zipFile.Stat()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, "archive.zip", fileInfo.ModTime(), zipFile)
	} else {
		// Отправка заголовков CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // Здесь можно указать конкретный источник
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		//Отправка ответа клиенту
		w.WriteHeader(http.StatusOK)
	}

}

func handleRequestIndex(w http.ResponseWriter, r *http.Request) {
	// читаем файл index.html
	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		// если файл не найден, отдаем ошибку 404
		http.Error(w, "File not found", 404)
		return
	}
	// отдаем содержимое файла
	fmt.Fprintf(w, "%s", content)
}
