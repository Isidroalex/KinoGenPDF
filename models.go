package main

import (
	"strconv"
	"strings"
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
	FIO     string
	Name    string
	Contact string
	Email   string
}

type Dog struct {
	Type     string
	Nickname string
	Sex      string
	Stamp    string
	Pedigree string
	Father   *Dog
	Mother   *Dog
	Breeder  *Breeder
	Owner    *Owner
}

type Puppy struct {
	Dog
	Color       string
	Birthday    string
	WoolType    string
	Description struct {
		Comment  string
		Defect   bool
		Revision int
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

	for _, v := range l.Puppy {
		if v.Description.Defect == checkbox {
			counter++
		}
	}
	return counter

}

func (p Puppy) ShowStumpNumber() int {

	number := strings.Split(p.Stamp, " ")
	if len(number) == 2 {
		number, _ := strconv.Atoi(number[1])
		return number
	}

	return 0
}
