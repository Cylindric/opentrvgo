package opentrvgo

import "testing"

func TestParseSensorReport(t *testing.T) {
	input := []byte(`{"@":"8CD3878ACDCE86BB","+":8,"H|%":23,"O":56,"L":34,"v|%":89,"B|cV":1200,"tT|C":67,"T|C16":64,"vac|h":78}`)

	sample, err := ParseSensorReport(input)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if sample["sequence"] != 8 {
		t.Error("Expected Sequence 8, got ", sample["Sequence"])
	}

	if sample["Battery"] != 12 {
		t.Error("Expected Battery 12, got ", sample["Battery"])
	}

	if sample["Device"] != "8CD3878ACDCE86BB" {
		t.Error("Expected Device 8CD3878ACDCE86BB, got ", sample["Device"])
	}

	if sample["Humidity"] != 23 {
		t.Error("Expected Humidity 23, got ", sample["Humidity"])
	}

	if sample["Light"] != 34 {
		t.Error("Expected Light 34, got ", sample["Light"])
	}

	if sample["Occupancy"] != 56 {
		t.Error("Expected Occupancy 56, got ", sample["Occupancy"])
	}

	if sample["TargetTemperature"] != 67 {
		t.Error("Expected Target Temperature 67, got ", sample["TargetTemperature"])
	}

	if sample["Temperature"] != 4 {
		t.Error("Expected Temperature 4, got ", sample["Temperature"])
	}

	if sample["Vacancy"] != 78 {
		t.Error("Expected Vacancy 78, got ", sample["Vacancy"])
	}

	if sample["Valve"] != 89 {
		t.Error("Expected Valve 89, got ", sample["Valve"])
	}

}

func TestParseSensorReportFailsWhenMissingSerial(t *testing.T) {
	input := []byte(`{"+":8,"H|%":48,"O":1,"L":6,"v|%":0}`)

	_, err := ParseSensorReport(input)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestParsedSensorReportReturnsErrorOnInvalidJSON(t *testing.T) {
	input := []byte("invalid")

	_, err := ParseSensorReport(input)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
