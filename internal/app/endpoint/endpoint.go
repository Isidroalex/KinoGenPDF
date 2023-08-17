package endpoint

import (
	"Breeding/internal/app/transport"
	"Breeding/internal/app/useCase"
	"fmt"
	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	IndexPage string
}

func New() *Endpoint {
	return &Endpoint{
		IndexPage: "internal/app/endpoint/index.html",
	}
}

func (e *Endpoint) Index(ctx echo.Context) error {
	return ctx.File(e.IndexPage)
}

func (e *Endpoint) Post(ctx echo.Context) error {
	//Fill up DTO
	dto := transport.New()
	dto.Parse(ctx)

	//Fill up useCase
	uc := useCase.New(dto)

	if err := uc.AddMating(); err != nil {
		return err
	}
	if err := uc.AddFolderName(); err != nil {
		return err
	}
	if err := uc.AddLitter(); err != nil {
		return err
	}
	if err := uc.AddPuppyCards(); err != nil {
		return err
	}

	// creating PDFs in ZIP archive
	if err := uc.WritePDF(useCase.BlackAgreement, useCase.Agreement, uc.Mating); err != nil {
		return err
	}
	if err := uc.WritePDF(useCase.BlackMating, useCase.Mating, uc.Mating); err != nil {
		return err
	}
	if err := uc.WritePDF(useCase.BlackRegister, useCase.Register, uc.Litter); err != nil {
		return err
	}
	if err := uc.WritePDF(useCase.BlackSurvey, useCase.Survey, uc.Litter); err != nil {
		return err
	}

	//Add Puppy slice to ZIP archive
	for i, v := range uc.PappyCards {
		filename := fmt.Sprint(useCase.PuppyCard, " ", i, ".pdf")
		if err := uc.WritePDF(useCase.BlackPuppyCard, filename, v); err != nil {
			return err
		}
	}

	if err := uc.WriteZip(); err != nil {
		return err
	}

	// change default content-type to zip
	ctx.Response().Header().Set("content-type", "application/zip")

	return ctx.Attachment("output.zip", uc.FolderName+".zip")
}
