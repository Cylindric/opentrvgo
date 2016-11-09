package opentrvgo

import "testing"

func TestParseSensorReport(t *testing.T) {
	var sample Sample
	input := []byte(`{"@":"8CD3878ACDCE86BB","+":8,"H|%":48,"O":1,"L":6,"v|%":0}`)

	valid, err := ParseSensorReport(input, sample)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if valid == false {
		t.Error("Expected valid==true, got ", valid)
	}
}

func TestParsedSensorReportReturnsErrorOnInvalidJSON(t *testing.T) {
	var sample Sample
	input := []byte("invalid")

	valid, err := ParseSensorReport(input, sample)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
	if valid {
		t.Error("Expected valid==false, got ", valid)
	}
}
