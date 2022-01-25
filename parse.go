package fsspgoparser

import (
	"fmt"
	"strconv"
	"strings"
)

func IsPerson(s string) bool {
	return personName.MatchString(s)
}

func ParsePersonName(s string) (lastName string, firstName string, secondName string, birthDate string, birtPlace string) {
	data := personName.FindAllStringSubmatch(s, 1)

	lastName = titleWord(data[0][1])
	firstName = titleWord(data[0][2])
	secondName = titleWord(data[0][3])
	birthDate = formatDate(data[0][4])
	birtPlace = data[0][5]

	return lastName, firstName, secondName, birthDate, birtPlace
}

func titleWord(s string) string {
	return strings.Title(strings.ToLower(s))
}

func formatDate(s string) string {
	date := strings.Split(s, ".")

	return fmt.Sprintf("%s-%s-%s", date[2], date[1], date[0])
}

func ParseLegalName(s string) (name string, address string) {
	data := legalName.FindAllStringSubmatch(s, 1)

	name = data[0][1]
	address = data[0][2]

	return name, address
}

func ParseProduction(s string) (mainProductionNumber string, startProduction string, addonProductionNumber string) {
	data := production.FindAllStringSubmatch(s, 1)

	mainProductionNumber = data[0][1]
	startProduction = formatDate(data[0][2])
	addonProductionNumber = data[0][3]

	return mainProductionNumber, startProduction, addonProductionNumber
}

func ParseDocument(s string) (documentName string, documentDate string, documentNumber string, documentWorker string) {
	data := document.FindAllStringSubmatch(s, 1)

	documentName = data[0][1]
	documentDate = formatDate(data[0][2])
	documentNumber = data[0][3]
	documentWorker = data[0][4]

	return documentName, documentDate, documentNumber, documentWorker
}

func ParseSubject(s string) (themeSubject string, mainSum float64, addonSum float64) {
	data := subject.FindAllStringSubmatch(s, 2)

	if len(data) == 2 {
		themeSubject = data[0][1]
		if data[0][2] != "" {
			mainSum, _ = strconv.ParseFloat(data[0][2], 10)
		}
		if data[1][2] != "" {
			addonSum, _ = strconv.ParseFloat(data[1][2], 10)
		}
	}

	if len(data) == 1 {
		themeSubject = data[0][1]
		if data[0][2] != "" {
			mainSum, _ = strconv.ParseFloat(data[0][2], 10)
		}
	}

	if len(data) == 0 {
		themeSubject = s
	}

	return themeSubject, mainSum, addonSum
}

func ParseDepartment(s string) (departmentName string) {
	if !department.MatchString(s) {
		return s
	}

	data := department.FindAllStringSubmatch(s, 1)

	departmentName = data[0][1]

	return departmentName
}

func ParseBailiff(s string) (bailiffName string, bailiffPhone string) {
	data := bailiff.FindAllStringSubmatch(s, 1)

	bailiffName = fmt.Sprintf("%s %s%s", titleWord(data[0][1]), data[0][2], data[0][3])
	bailiffPhone = data[0][4]

	return bailiffName, bailiffPhone
}

func ParseIPEnd(s string) (date string, cause string) {
	if !ipEnd.MatchString(s) {
		return date, cause
	}

	data := ipEnd.FindAllStringSubmatch(s, 1)

	date = data[0][1]
	cause = fmt.Sprintf("п. %s ч. %s ст. %s", data[0][4], data[0][3], data[0][2])

	return date, cause
}
