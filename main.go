package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

//Defining struct for the CSV header
type HeaderRow struct {
	Month       string
	PaymentDate string
	BonusDate   string
}

//Adding helper so struct can be converted to array
func (r HeaderRow) ToSlice() []string {
	return []string{r.Month, r.PaymentDate, r.BonusDate}
}

//End of struct for CSV header

//Defining struct for CSV rows
type CSVRow struct {
	Month       string
	PaymentDate int
	BonusDate   int
}

//Helper to convert CSV rows into arrays
func (r CSVRow) ToSlice() []string {
	return []string{r.Month, strconv.Itoa(r.PaymentDate), strconv.Itoa(r.BonusDate)}
}

//End of struct for CSV rows

//start of consts
var months = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

var year = 2022

// End of consts

func main() {

	paymentData := ProcessDates(year)

	WorkWithCSV(paymentData)
}

//Start of dates logic
func ProcessDates(year int) []CSVRow {
	processedDates := []CSVRow{}
	for i := 1; i <= 12; i++ {
		processedDates = append(processedDates, CheckMonthDates(year, i))
	}
	return processedDates
}

func CheckMonthDates(year int, month int) CSVRow {
	paymentDay := GetPaymentDay(month)
	bonusDay := GetBonusDay(month)
	return CSVRow{months[month], paymentDay.Day(), bonusDay.Day()}
}

func GetPaymentDay(month int) time.Time {
	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	switch lastOfMonth.Weekday() {
	case 0:
		return lastOfMonth.AddDate(0, 0, -2)
	case 6:
		return lastOfMonth.AddDate(0, 0, -1)
	default:
		return lastOfMonth
	}
}

func GetBonusDay(month int) time.Time {
	bonusDay := time.Date(year, time.Month(month), 15, 0, 0, 0, 0, time.UTC)
	switch bonusDay.Weekday() {
	case 0:
		return bonusDay.AddDate(0, 0, 3)
	case 6:
		return bonusDay.AddDate(0, 0, 4)
	default:
		return bonusDay
	}
}

//End of dates logic

//Start of CSV logic
func WorkWithCSV(paymentData []CSVRow) {

	RemoveDuplicateCSV()

	csvFile := CreateEmptyCSV()

	csvwriter := csv.NewWriter(csvFile)

	csvwriter.Write(HeaderRow{"month", "payment date", "bonus date"}.ToSlice())
	for _, empRow := range paymentData {
		_ = csvwriter.Write(empRow.ToSlice())
	}
	csvwriter.Flush()
	csvFile.Close()
}

func RemoveDuplicateCSV() {
	if _, err := os.Stat("payments.csv"); err == nil {
		err := os.Remove("payments.csv")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CreateEmptyCSV() *os.File {
	csvFile, err := os.Create("payments.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	return csvFile
}

//End of CSV Logic
