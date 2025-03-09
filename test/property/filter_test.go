package property

import (
	"testing"

	"property-cli/models"
	"property-cli/property"
)

func TestValidFilterProperties(t *testing.T) {
	properties := []models.Property{
		{Price: 100000, SquareFootage: 1200, Rooms: 3, Bathrooms: 2, Amenities: map[string]bool{"garage": true}, Description: "Cozy home with a garage", Location: [2]float64{40.7128, -74.0060}},
		{Price: 150000, SquareFootage: 1400, Rooms: 4, Bathrooms: 3, Amenities: map[string]bool{"pool": true}, Description: "Luxury house with a pool", Location: [2]float64{40.7306, -73.9352}},
		{Price: 200000, SquareFootage: 1600, Rooms: 5, Bathrooms: 4, Amenities: map[string]bool{"garage": true, "pool": true}, Description: "Spacious house with garage and pool", Location: [2]float64{40.7488, -73.9857}},
	}

	tests := []struct {
		name        string
		filterKey   string
		filterType  string
		filterValue string
		expectedLen int
	}{
		{"Filter by price equal", "price", "equal", "150000", 1},
		{"Filter by price lessThan", "price", "lessThan", "200000", 2},
		{"Filter by price greaterThan", "price", "greaterThan", "120000", 2},
		{"Filter by rooms equal", "rooms", "equal", "3", 1},
		{"Filter by bathrooms greaterThan", "bathrooms", "greaterThan", "2", 2},
		{"Filter by inclusion (garage)", "amenities", "include", "garage", 2},
		{"Filter by matching (description contains 'pool')", "description", "match", "pool", 2},
		{"Filter by distance (within 5km of 40.7128,-74.0060)", "location", "distance", "40.7128,-74.0060,5", 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := property.FilterProperties(properties, test.filterKey, test.filterType, test.filterValue)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(result) != test.expectedLen {
				t.Errorf("Expected %d properties, got %d", test.expectedLen, len(result))
			}
		})
	}
}

func TestInvalidFilterProperties(t *testing.T) {
	properties := []models.Property{
		{Price: 100000, SquareFootage: 1200, Rooms: 3, Bathrooms: 2, Amenities: map[string]bool{"garage": true}, Description: "Cozy home with a garage", Location: [2]float64{40.7128, -74.0060}},
		{Price: 150000, SquareFootage: 1400, Rooms: 4, Bathrooms: 3, Amenities: map[string]bool{"pool": true}, Description: "Luxury house with a pool", Location: [2]float64{40.7306, -73.9352}},
	}

	tests := []struct {
		name        string
		filterKey   string
		filterType  string
		filterValue string
	}{
		{"Invalid amenities filter type", "amenities", "equal", "garage"},
		{"Invalid description filter type", "description", "contains", "pool"},
		{"Invalid filter key", "invalidKey", "equal", "100"},
		{"Invalid filter type", "price", "approximately", "150000"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := property.FilterProperties(properties, test.filterKey, test.filterType, test.filterValue)
			if err == nil {
				t.Errorf("Expected an error but got none for filterKey=%s, filterType=%s", test.filterKey, test.filterType)
			}
		})
	}
}
