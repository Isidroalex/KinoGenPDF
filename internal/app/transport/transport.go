package transport

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
)

type DTO struct {
	NameDC          string `json:"NameDC"`
	AddressDC       string `json:"AddressDC"`
	EmailDC         string `json:"EmailDC"`
	Mobile          string `json:"Mobile"`
	Federation      string `json:"Federation"`
	FolderNumber    string `json:"FolderNumber"`
	StumpCode       string `json:"StumpCode"`
	ChiefFIO        string `json:"ChiefFIO"`
	ChiefAddress    string `json:"ChiefAddress"`
	PersonalFIO     string `json:"PersonalFIO"`
	PersonalAddress string `json:"PersonalAddress"`
	ClaimFIO        string `json:"ClaimFIO"`
	ClaimAddress    string `json:"ClaimAddress"`
	Identification  string `json:"Identification"`
	First           string `json:"First"`
	Control         string `json:"Control"`
	MatingAddress   string `json:"MatingAddress"`
	LitterBirthday  string `json:"LitterBirthday"`
	LitterSurvey    string `json:"LitterSurvey"`
	RevisionPeriod  string `json:"RevisionPeriod"`
	FatherName      string `json:"FatherName"`
	FatherPedigree  string `json:"FatherPedigree"`
	FatherStump     string `json:"FatherStump"`
	FatherOwnerFIO  string `json:"FatherOwnerFIO"`
	FatherAddress   string `json:"FatherAddress"`
	MotherName      string `json:"MotherName"`
	MotherPedigree  string `json:"MotherPedigree"`
	MotherStump     string `json:"MotherStump"`
	MotherType      string `json:"MotherType"`
	Suffix          string `json:"Suffix"`
	BreederFIO      string `json:"BreederFIO"`
	BreederContact  string `json:"BreederContact"`
	BreederEmail    string `json:"BreederEmail"`
	BreederName     string `json:"BreederName"`
	BreederNameEnd  string `json:"BreederNameEnd"`
	Puppy           []struct {
		Nickname      string `json:"Nickname"`
		NicknameEng   string `json:"NicknameEng"`
		Color         string `json:"Color"`
		WoolType      string `json:"WoolType"`
		PuppyStump    string `json:"PuppyStump"`
		SexPuppy      string `json:"SexPuppy"`
		StatusPuppy   string `json:"StatusPuppy"`
		StatusComment string `json:"StatusComment"`
	} `json:"Puppy"`
}

func New() *DTO {
	return &DTO{}
}

func (r *DTO) Parse(ctx echo.Context) error {
	b, _ := io.ReadAll(ctx.Request().Body)
	defer ctx.Request().Body.Close()

	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}
	return nil
}

func (r *DTO) SetMating() (*Mating, error) {
	return &Mating{
		Identification: FormatDate(r.Identification),
		First:          FormatDate(r.First),
		Control:        FormatDate(r.Control),
		Address:        r.MatingAddress,
		Male: &Dog{
			Type:     r.MotherType,
			Nickname: r.FatherName,
			Sex:      "к",
			Stamp:    r.FatherStump,
			Pedigree: r.FatherPedigree,
			Owner: &Owner{
				FIO:     r.FatherOwnerFIO,
				Contact: r.FatherAddress,
			},
		},
		Female: &Dog{
			Type:     r.MotherType,
			Nickname: r.MotherName,
			Sex:      "с",
			Stamp:    r.MotherStump,
			Pedigree: r.MotherPedigree,
			Breeder: &Breeder{
				FIO:          r.BreederFIO,
				Name:         r.BreederName,
				NameEng:      r.BreederNameEnd,
				NamePosition: r.Suffix,
				Contact:      r.BreederContact,
				Email:        r.BreederEmail,
			},
			Owner: &Owner{
				FIO:     r.BreederFIO,
				Contact: r.BreederContact,
				Email:   r.BreederEmail,
			},
		},
		DCenter: &DocumentCenter{
			Name:         r.NameDC,
			Email:        r.EmailDC,
			Mobile:       r.Mobile,
			Federation:   r.Federation,
			FolderNumber: r.FolderNumber,
			StumpCode:    r.StumpCode,
			Address:      r.AddressDC,
		},
		Personal: &Personal{
			FIO:     r.PersonalFIO,
			Address: r.PersonalAddress,
		},
	}, nil
}

func (r *DTO) SetLitter() (*Litter, error) {

	l := &Litter{
		Stump:    r.StumpCode,
		ActDate:  FormatDate(r.LitterSurvey),
		Birthday: FormatDate(r.LitterBirthday),
		Control:  FormatDate(r.Control),
		Address:  r.MatingAddress,
		Male: &Dog{
			Type:     r.MotherType,
			Nickname: r.FatherName,
			Sex:      "к",
			Stamp:    r.FatherStump,
			Pedigree: r.FatherPedigree,
			Owner: &Owner{
				FIO:     r.FatherOwnerFIO,
				Contact: r.FatherAddress,
			},
		},
		Female: &Dog{
			Type:     r.MotherType,
			Nickname: r.MotherName,
			Sex:      "с",
			Stamp:    r.MotherStump,
			Pedigree: r.MotherPedigree,
			Breeder: &Breeder{
				FIO:          r.BreederFIO,
				Name:         r.BreederName,
				NameEng:      r.BreederNameEnd,
				NamePosition: r.Suffix,
				Contact:      r.BreederContact,
				Email:        r.BreederEmail,
			},
			Owner: &Owner{
				FIO:     r.BreederFIO,
				Contact: r.BreederContact,
				Email:   r.BreederEmail,
			},
		},
		Puppy: nil,
		DCenter: &DocumentCenter{
			Name:         r.NameDC,
			Email:        r.EmailDC,
			Mobile:       r.Mobile,
			Federation:   r.Federation,
			FolderNumber: r.FolderNumber,
			StumpCode:    r.StumpCode,
			Address:      r.AddressDC,
		},
		Personal: &Personal{
			FIO:     r.PersonalFIO,
			Address: r.PersonalAddress,
		},
		Claim: &Personal{
			FIO:     r.ClaimFIO,
			Address: r.ClaimAddress,
		},
		Chief: &Personal{
			FIO:     r.ChiefFIO,
			Address: r.ChiefAddress,
		},
	}

	for _, v := range r.Puppy {
		tmp := Puppy{
			Dog: Dog{
				Type:        r.MotherType,
				Nickname:    v.Nickname,
				NicknameEng: v.NicknameEng,
				Sex:         v.SexPuppy,
				Stamp:       v.PuppyStump,
				Pedigree:    "",
				Father:      l.Male,
				Mother:      l.Female,
				Breeder:     l.Female.Breeder,
				Owner:       nil,
			},
			Color:    v.Color,
			Birthday: FormatDate(r.LitterBirthday),
			WoolType: v.WoolType,
			Description: struct {
				Comment  string
				Defect   string
				Revision string
			}{
				v.StatusComment,
				v.StatusPuppy,
				r.RevisionPeriod,
			},
			IncreaseIndex: func(a int) int {
				return a + 1
			},
		}
		l.Puppy = append(l.Puppy, &tmp)
	}

	return l, nil
}

func (r *DTO) SetPuppyCards() ([]*PuppyCard, error) {

	var puppyCards []*PuppyCard

	for _, puppy := range r.Puppy {

		owner := Owner{
			FIO:     "",
			Contact: "",
			Email:   "",
		}
		dogStruct := Dog{
			Type:        r.MotherType,
			Nickname:    puppy.Nickname,
			NicknameEng: puppy.NicknameEng,
			Sex:         puppy.SexPuppy,
			Stamp:       puppy.PuppyStump,
			Father: &Dog{
				Type:     r.MotherType,
				Nickname: r.FatherName,
				Sex:      "к",
				Stamp:    r.FatherStump,
				Pedigree: r.FatherPedigree,
				Owner: &Owner{
					FIO:     r.FatherOwnerFIO,
					Contact: r.FatherAddress,
				},
			},
			Mother: &Dog{
				Type:     r.MotherType,
				Nickname: r.MotherName,
				Sex:      "с",
				Stamp:    r.MotherStump,
				Pedigree: r.MotherPedigree,
				Breeder: &Breeder{
					FIO:          r.BreederFIO,
					Name:         r.BreederName,
					NameEng:      r.BreederNameEnd,
					NamePosition: r.Suffix,
					Contact:      r.BreederContact,
					Email:        r.BreederEmail,
				}},
			Breeder: &Breeder{
				FIO:          r.BreederFIO,
				Name:         r.BreederName,
				NameEng:      r.BreederNameEnd,
				NamePosition: r.Suffix,
				Contact:      r.BreederContact,
				Email:        r.BreederEmail,
			},
			Owner: &owner,
		}
		puppyStruct := Puppy{

			Color:    puppy.Color,
			Birthday: FormatDate(r.LitterBirthday),
			WoolType: puppy.WoolType,
			Description: struct {
				Comment  string
				Defect   string
				Revision string
			}{
				Comment:  puppy.StatusComment,
				Defect:   puppy.StatusPuppy,
				Revision: r.RevisionPeriod,
			},
			IncreaseIndex: func(a int) int {
				return a + 1
			},
		}
		puppyStruct.Dog = dogStruct

		puppyCard := PuppyCard{
			Puppy: puppyStruct,
			DCenter: DocumentCenter{
				Name:         r.NameDC,
				Email:        r.EmailDC,
				Mobile:       r.Mobile,
				Federation:   r.Federation,
				FolderNumber: r.FolderNumber,
				StumpCode:    r.StumpCode,
				Address:      r.AddressDC,
			},
		}
		puppyCards = append(puppyCards, &puppyCard)
	}
	return puppyCards, nil
}
