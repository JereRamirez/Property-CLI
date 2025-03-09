package geo

import (
	"testing"

	"property-cli/geo"
)

func TestParseCoordinates(t *testing.T) {
	tests := []struct {
		input       string
		expected    [2]float64
		expectError bool
	}{
		{"40.7128,-74.0060", [2]float64{40.7128, -74.0060}, false},
		{" 40.7128 , -74.0060 ", [2]float64{40.7128, -74.0060}, false},
		{"90.0000,180.0000", [2]float64{90.0000, 180.0000}, false},
		{"-90.0000,-180.0000", [2]float64{-90.0000, -180.0000}, false},
		{"invalid,coordinates", [2]float64{}, true},
		{"40.7128", [2]float64{}, true},
		{"40.7128,-74.0060,100", [2]float64{}, true},
		{"", [2]float64{}, true},
	}

	for _, test := range tests {
		result, err := geo.ParseCoordinates(test.input)

		if test.expectError {
			if err == nil {
				t.Errorf("Expected error for input '%s', but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input '%s': %v", test.input, err)
			} else if result != test.expected {
				t.Errorf("Expected %v for input '%s', but got %v", test.expected, test.input, result)
			}
		}
	}
}
