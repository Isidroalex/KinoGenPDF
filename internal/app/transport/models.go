package transport

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DocumentCenter struct {
	Name         string
	Email        string
	Mobile       string
	Federation   string
	FolderNumber string
	StumpCode    string
	Address      string
}

type Mating struct {
	Identification string
	First          string
	Control        string
	Address        string
	Male           *Dog
	Female         *Dog
	DCenter        *DocumentCenter
	Personal       *Personal
}

type Litter struct {
	Stump            string
	ActDate          string
	Birthday         string
	Control          string
	Address          string
	BurnTotal        int
	BurnMale         int
	BurnFemale       int
	ClaimTotal       int
	ClaimTotalMale   int
	ClaimTotalFemale int
	Defected         int
	Good             int

	Male     *Dog
	Female   *Dog
	Puppy    []*Puppy
	DCenter  *DocumentCenter
	Personal *Personal
	Claim    *Personal
	Chief    *Personal
}

type Personal struct {
	FIO     string
	Address string
}

type Owner struct {
	FIO     string
	Contact string
	Email   string
}

type Breeder struct {
	FIO          string
	Name         string
	NameEng      string
	NamePosition string
	Contact      string
	Email        string
}

type Dog struct {
	Type        string
	Nickname    string
	NicknameEng string
	Sex         string
	Stamp       string
	Pedigree    string
	Father      *Dog
	Mother      *Dog
	Breeder     *Breeder
	Owner       *Owner
}

type Puppy struct {
	Dog
	Color       string
	Birthday    string
	WoolType    string
	Description struct {
		Comment  string
		Defect   string
		Revision string
	}
	IncreaseIndex func(a int) int
}

type PuppyCard struct {
	Puppy
	DCenter DocumentCenter
}

func (l Litter) ShowPuppyCount(sex string) int {
	var counter int

	if sex != "" {
		for _, v := range l.Puppy {
			if v.Sex == sex {
				counter++
			}
		}
		return counter
	} else {
		return len(l.Puppy)
	}

}

func (l Litter) ShowPuppyStumped(sex string) int {
	var counter int
	if sex != "" {
		for _, v := range l.Puppy {
			if v.Stamp != "" && v.Sex == sex {
				counter++
			}
		}
		return counter
	} else {
		for _, v := range l.Puppy {
			if v.Stamp != "" {
				counter++
			}
		}
		return counter
	}

}

func (l Litter) ShowPuppyDefected(checkbox bool) int {
	var counter int
	if checkbox {
		for _, v := range l.Puppy {
			if v.Description.Defect == "defect" {
				counter++
			}
		}
	} else {
		for _, v := range l.Puppy {
			if v.Description.Defect != "defect" {
				counter++
			}
		}
	}
	return counter
}

func (l Litter) ShowPuppyRevision(checkbox bool) int {
	var counter int
	if checkbox {
		for _, v := range l.Puppy {
			if v.Description.Defect == "revision" {
				counter++
			}
		}
	} else {
		for _, v := range l.Puppy {
			if v.Description.Defect != "revision" {
				counter++
			}
		}
	}
	return counter
}

func (l Litter) RevisionDate() string {

	return l.Puppy[0].Description.Revision

}

func (p Puppy) ShowStumpNumber() int {

	number := strings.Split(p.Stamp, " ")
	if len(number) == 2 {
		number, _ := strconv.Atoi(number[1])
		return number
	}

	return 0
}

func (p Puppy) ShowWoolType() string {

	if p.WoolType != "" {
		return fmt.Sprintf("(%s)", p.WoolType)
	}
	return ""
}

func (p Puppy) ShowFullNameRus() string {
	var result []string
	//Приставка определяется по значению поля NamePosition. False - спереди, true - сзади
	if p.Breeder.NamePosition == "false" {
		result = append(result, p.Breeder.Name, p.Nickname)
	} else {
		result = append(result, p.Nickname, p.Breeder.Name)
	}
	return strings.Join(result, " ")
}

func (p Puppy) ShowFullNameEng() string {
	var result []string
	//Приставка определяется по значению поля NamePosition. False - спереди, true - сзади.
	//Если введено значение с клавиатуры, используется оно, если нет - конвертируется по правилам.
	if p.NicknameEng == "" {
		if p.Breeder.NamePosition == "false" {
			result = append(result, p.Breeder.NameEng, convertToLatin(p.Nickname))
		} else {
			result = append(result, convertToLatin(p.Nickname), p.Breeder.NameEng)
		}
		return strings.Join(result, " ")
	} else {
		if p.Breeder.NamePosition == "false" {
			result = append(result, p.Breeder.NameEng, p.NicknameEng)
		} else {
			result = append(result, p.NicknameEng, p.Breeder.NameEng)
		}
		return strings.Join(result, " ")
	}
}

func (p Puppy) LongSex() string {

	if p.Sex == "к" {
		return "кобель"
	} else {
		return "сука"
	}
}

func (p Puppy) ShortDisription() string {

	switch p.Description.Defect {
	case "revision":
		if p.WoolType != "" {
			return fmt.Sprintf("Необходим пересмотр (%s) тип шерсти: %s", p.Description.Comment, p.WoolType)
		} else {
			return fmt.Sprintf("Необходим пересмотр (%s)", p.Description.Comment)
		}
	case "defect":
		if p.WoolType != "" {
			return fmt.Sprintf("Отбракован (%s) тип шерсти: %s", p.Description.Comment, p.WoolType)
		} else {
			return fmt.Sprintf("Отбракован (%s)", p.Description.Comment)
		}
	default:
		if p.WoolType != "" {
			return fmt.Sprintf("Без замечаний, тип шерсти: %s", p.WoolType)
		} else {
			return fmt.Sprintf("Без замечаний")
		}
	}

}

func (p Puppy) ShortDisruptionPuppyCard(place string) string {
	switch place {
	case "defect":
		if p.Description.Defect == "defect" {
			return p.Description.Comment
		} else {
			return ""
		}
	case "revision":
		if p.Description.Defect == "revision" {
			return p.Description.Comment
		} else {
			return ""
		}
	case "month":
		if p.Description.Defect == "revision" {
			return p.Description.Revision
		} else {
			return ""
		}
	default:
		return ""
	}
}

func (p Puppy) ShowNameEng() string {
	if p.NicknameEng != "" {
		return p.NicknameEng
	}
	return convertToLatin(p.Nickname)
}

type Input struct {
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

func (m *Mating) Construct(i *DTO) {
	*m = Mating{
		Identification: FormatDate(i.Identification),
		First:          FormatDate(i.First),
		Control:        FormatDate(i.Control),
		Address:        i.MatingAddress,
		Male: &Dog{
			Type:     i.MotherType,
			Nickname: i.FatherName,
			Sex:      "к",
			Stamp:    i.FatherStump,
			Pedigree: i.FatherPedigree,
			Owner: &Owner{
				FIO:     i.FatherOwnerFIO,
				Contact: i.FatherAddress,
			},
		},
		Female: &Dog{
			Type:     i.MotherType,
			Nickname: i.MotherName,
			Sex:      "с",
			Stamp:    i.MotherStump,
			Pedigree: i.MotherPedigree,
			Breeder: &Breeder{
				FIO:          i.BreederFIO,
				Name:         i.BreederName,
				NameEng:      i.BreederNameEnd,
				NamePosition: i.Suffix,
				Contact:      i.BreederContact,
				Email:        i.BreederEmail,
			},
			Owner: &Owner{
				FIO:     i.BreederFIO,
				Contact: i.BreederContact,
				Email:   i.BreederEmail,
			},
		},
		DCenter: &DocumentCenter{
			Name:         i.NameDC,
			Email:        i.EmailDC,
			Mobile:       i.Mobile,
			Federation:   i.Federation,
			FolderNumber: i.FolderNumber,
			StumpCode:    i.StumpCode,
			Address:      i.AddressDC,
		},
		Personal: &Personal{
			FIO:     i.PersonalFIO,
			Address: i.PersonalAddress,
		},
	}
}

func (l *Litter) Construct(i Input) {
	*l = Litter{
		Stump:    i.StumpCode,
		ActDate:  FormatDate(i.LitterSurvey),
		Birthday: FormatDate(i.LitterBirthday),
		Control:  FormatDate(i.Control),
		Address:  i.MatingAddress,
		Male: &Dog{
			Type:     i.MotherType,
			Nickname: i.FatherName,
			Sex:      "к",
			Stamp:    i.FatherStump,
			Pedigree: i.FatherPedigree,
			Owner: &Owner{
				FIO:     i.FatherOwnerFIO,
				Contact: i.FatherAddress,
			},
		},
		Female: &Dog{
			Type:     i.MotherType,
			Nickname: i.MotherName,
			Sex:      "с",
			Stamp:    i.MotherStump,
			Pedigree: i.MotherPedigree,
			Breeder: &Breeder{
				FIO:          i.BreederFIO,
				Name:         i.BreederName,
				NameEng:      i.BreederNameEnd,
				NamePosition: i.Suffix,
				Contact:      i.BreederContact,
				Email:        i.BreederEmail,
			},
			Owner: &Owner{
				FIO:     i.BreederFIO,
				Contact: i.BreederContact,
				Email:   i.BreederEmail,
			},
		},
		Puppy: nil,
		DCenter: &DocumentCenter{
			Name:         i.NameDC,
			Email:        i.EmailDC,
			Mobile:       i.Mobile,
			Federation:   i.Federation,
			FolderNumber: i.FolderNumber,
			StumpCode:    i.StumpCode,
			Address:      i.AddressDC,
		},
		Personal: &Personal{
			FIO:     i.PersonalFIO,
			Address: i.PersonalAddress,
		},
		Claim: &Personal{
			FIO:     i.ClaimFIO,
			Address: i.ClaimAddress,
		},
		Chief: &Personal{
			FIO:     i.ChiefFIO,
			Address: i.ChiefAddress,
		},
	}

	for _, v := range i.Puppy {
		tmp := Puppy{
			Dog: Dog{
				Type:        i.MotherType,
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
			Birthday: FormatDate(i.LitterBirthday),
			WoolType: v.WoolType,
			Description: struct {
				Comment  string
				Defect   string
				Revision string
			}{
				v.StatusComment,
				v.StatusPuppy,
				i.RevisionPeriod,
			},
			IncreaseIndex: func(a int) int {
				return a + 1
			},
		}
		l.Puppy = append(l.Puppy, &tmp)
	}

}

func convertToLatin(input string) string {
	convertMap := map[rune]string{
		'А': "A",
		'Б': "B",
		'В': "V",
		'Г': "G",
		'Д': "D",
		'Е': "E",
		'Ё': "YO",
		'Ж': "ZH",
		'З': "Z",
		'И': "I",
		'Й': "Y",
		'К': "K",
		'Л': "L",
		'М': "M",
		'Н': "N",
		'О': "O",
		'П': "P",
		'Р': "R",
		'С': "S",
		'Т': "T",
		'У': "U",
		'Ф': "F",
		'Х': "KH",
		'Ц': "TS",
		'Ч': "CH",
		'Ш': "SH",
		'Щ': "SHCH",
		'Ъ': "",
		'Ы': "Y",
		'Ь': "",
		'Э': "E",
		'Ю': "YU",
		'Я': "YA",
		'а': "a",
		'б': "b",
		'в': "v",
		'г': "g",
		'д': "d",
		'е': "e",
		'ё': "yo",
		'ж': "zh",
		'з': "z",
		'и': "i",
		'й': "y",
		'к': "k",
		'л': "l",
		'м': "m",
		'н': "n",
		'о': "o",
		'п': "p",
		'р': "r",
		'с': "s",
		'т': "t",
		'у': "u",
		'ф': "f",
		'х': "kh",
		'ц': "ts",
		'ч': "ch",
		'ш': "sh",
		'щ': "shch",
		'ъ': "",
		'ы': "y",
		'ь': "",
		'э': "e",
		'ю': "yu",
		'я': "ya",
	}

	var result []string

	for _, char := range input {
		if converted, ok := convertMap[char]; ok {
			result = append(result, converted)
		} else {
			result = append(result, string(char))
		}
	}

	return strings.Join(result, "")
}

func FormatDate(input string) string {
	// Парсим входную строку в формате YYYY-MM-DD
	date, err := time.Parse("2006-01-02", input)
	if err != nil {
		return ""
	}

	var monthNames = map[time.Month]string{
		time.January:   "января",
		time.February:  "февраля",
		time.March:     "марта",
		time.April:     "апреля",
		time.May:       "мая",
		time.June:      "июня",
		time.July:      "июля",
		time.August:    "августа",
		time.September: "сентября",
		time.October:   "октября",
		time.November:  "ноября",
		time.December:  "декабря",
	}

	// Получаем день, месяц и год из даты
	day := date.Day()
	month := date.Month()
	year := date.Year()

	// Получаем название месяца на русском языке
	monthName := monthNames[month]

	// Формируем итоговую строку
	formattedDate := fmt.Sprintf("%d %s %d", day, monthName, year)

	return formattedDate
}
