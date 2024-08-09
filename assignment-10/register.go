package main

import (
	"fmt"
)

type Pupil struct {
	fullName    FullName
	dateOfBirth string
	age         int
}

type FullName struct {
	firstName  string
	middleName string
	lastName   string
}

var pupils = []Pupil{
	{
		fullName: FullName{
			firstName:  "Obi",
			middleName: "Wan",
			lastName:   "Kenobi",
		},
		dateOfBirth: "07/04/2015",
		age:         9,
	},
	{
		fullName: FullName{
			firstName:  "Qui",
			middleName: "Gon",
			lastName:   "Jin",
		},
		dateOfBirth: "18/07/2015",
		age:         9,
	},
	{
		fullName: FullName{
			firstName:  "Luke",
			middleName: "",
			lastName:   "Skywalker",
		},
		dateOfBirth: "20/11/2014",
		age:         10,
	},
	{
		fullName: FullName{
			firstName:  "Leia",
			middleName: "",
			lastName:   "Skywalker",
		},
		dateOfBirth: "20/11/2014",
		age:         10,
	},
	{
		fullName: FullName{
			firstName:  "Llando",
			middleName: "",
			lastName:   "Calrissian",
		},
		dateOfBirth: "15/10/2014",
		age:         10,
	},
	{
		fullName: FullName{
			firstName:  "Jean",
			middleName: "Luc",
			lastName:   "Picard",
		},
		dateOfBirth: "02/03/2015",
		age:         10,
	},
	{
		fullName: FullName{
			firstName:  "Geordie",
			middleName: "La",
			lastName:   "Ford",
		},
		dateOfBirth: "06/07/2015",
		age:         10,
	},
	{
		fullName: FullName{
			firstName:  "Kira",
			middleName: "",
			lastName:   "Nerys",
		},
		dateOfBirth: "22/04/2015",
		age:         10,
	},
	{
		fullName: FullName{
			firstName:  "Jadzia",
			middleName: "",
			lastName:   "Dax",
		},
		dateOfBirth: "10/09/2015",
		age:         10,
	},
	{
		fullName: FullName{
			firstName:  "Miles",
			middleName: "",
			lastName:   "O'Brien",
		},
		dateOfBirth: "28/10/2015",
		age:         10,
	},
}

func main() {
	printHeader()
	for _, pupil := range pupils {
		fmt.Println(pupil)
	}
	printBar()
}

func printHeader() {
	printBar()
	fmt.Println(titles())
	printBar()
}

func titles() string {
	return fmt.Sprintf("| %-20s | %-20s | %-20s | %-20s | %-20s |", "First Name", "Middle Name", "Last Name", "Date of Birth", "Age")
}

func printBar() {
	fmt.Println("|----------------------|----------------------|----------------------|----------------------|----------------------|")
}

func (p Pupil) String() string {
	return fmt.Sprintf("| %-20s | %-20s | %-20s | %-20s | %-20d |", p.fullName.firstName, p.fullName.middleName, p.fullName.lastName, p.dateOfBirth, p.age)
}
