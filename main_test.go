package main

import (
	"reflect"
	"testing"
)

func TestProcessDates(t *testing.T) {
	mockDateJanuary := CSVRow{
		"January", 31, 19,
	}
	mockDateFebruary := CSVRow{
		"February", 28, 15,
	}
	mockDateMarch := CSVRow{
		"March", 31, 15,
	}
	mockDateApril := CSVRow{
		"April", 29, 15,
	}
	mockDateMay := CSVRow{
		"May", 31, 18,
	}
	mockDateJune := CSVRow{
		"June", 30, 15,
	}
	mockDateJuly := CSVRow{
		"July", 29, 15,
	}
	mockDateAugust := CSVRow{
		"August", 31, 15,
	}
	mockDateSeptember := CSVRow{
		"September", 30, 15,
	}
	mockDateOctober := CSVRow{
		"October", 31, 19,
	}
	mockDateNovember := CSVRow{
		"November", 30, 15,
	}
	mockDateDecember := CSVRow{
		"December", 30, 15,
	}

	mockDatesArray := []CSVRow{
		mockDateJanuary,
		mockDateFebruary,
		mockDateMarch,
		mockDateApril,
		mockDateMay,
		mockDateJune,
		mockDateJuly,
		mockDateAugust,
		mockDateSeptember,
		mockDateOctober,
		mockDateNovember,
		mockDateDecember,
	}

	got := ProcessDates(year)
	want := mockDatesArray

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v but wanted %+v", got, want)
	}
}
