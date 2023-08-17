package useCase

import (
	"Breeding/internal/app/pdf"
	transport "Breeding/internal/app/transport"
	"archive/zip"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	BlackAgreement = "internal/app/pdf/blankAgreement.html"
	Agreement      = "СогласиеПД.pdf"
	BlackMating    = "internal/app/pdf/blankMating.html"
	Mating         = "Акт вязки.pdf"
	BlackPuppyCard = "internal/app/pdf/blankPuppyCard.html"
	PuppyCard      = "Щенячка "
	BlackRegister  = "internal/app/pdf/blankRegister.html"
	Register       = "Регистрация.pdf"
	BlackSurvey    = "internal/app/pdf/blankSurvey.html"
	Survey         = "Обследование.pdf"
)

type UseCase struct {
	dto        DataObject
	ctx        echo.Context
	Mating     *transport.Mating
	Litter     *transport.Litter
	PappyCards []*transport.PuppyCard
	ZipFile    string
	FolderName string
}

type DataObject interface {
	Parse(ctx echo.Context) error
	SetMating() (*transport.Mating, error)
	SetLitter() (*transport.Litter, error)
	SetPuppyCards() ([]*transport.PuppyCard, error)
}

func New(dto DataObject) *UseCase {
	return &UseCase{
		dto: dto,
	}
}

func (u *UseCase) AddMating() error {
	var err error
	u.Mating, err = u.dto.SetMating()
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) AddFolderName() error {
	u.FolderName = fmt.Sprintf(strconv.FormatInt(time.Now().Unix(), 10) + "/")
	return nil
}

func (u *UseCase) AddLitter() error {
	var err error
	u.Litter, err = u.dto.SetLitter()
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) AddPuppyCards() error {
	var err error
	u.PappyCards, err = u.dto.SetPuppyCards()
	if err != nil {
		return err
	}
	return nil
}

func (u *UseCase) WritePDF(tmplPath, name string, data interface{}) error {

	path := u.FolderName
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	outputPath := u.FolderName + name
	requestPdf := pdf.NewRequestPdf("")
	if err := requestPdf.ParseTemplate(tmplPath, data); err == nil {
		args := []string{"no-pdf-compression"}
		_, err := requestPdf.GeneratePDF(outputPath, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *UseCase) WriteZip() error {
	file, err := os.Create("output.zip")
	if err != nil {
		return err
	}
	defer file.Close()
	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}

	err = filepath.Walk(u.FolderName, walker)
	if err != nil {
		return err
	}

	return nil
}
