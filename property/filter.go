package property

import (
	"fmt"
	"strconv"
	"strings"

	"property-cli/geo"
	"property-cli/models"
)

func FilterProperties(properties []models.Property, filterKey, filterType, filterValue string) ([]models.Property, error) {
	var filtered []models.Property

	// Numeric filters
	numericFields := map[string]func(p models.Property) int{
		"price":         func(p models.Property) int { return p.Price },
		"squareFootage": func(p models.Property) int { return p.SquareFootage },
		"rooms":         func(p models.Property) int { return p.Rooms },
		"bathrooms":     func(p models.Property) int { return p.Bathrooms },
	}

	validNumericFilters := map[string]bool{
		"equal":       true,
		"lessThan":    true,
		"greaterThan": true,
	}

	if numExtractor, exists := numericFields[filterKey]; exists {
		if !validNumericFilters[filterType] {
			return nil, fmt.Errorf("invalid filter type for numeric field: %s", filterType)
		}

		value, err := strconv.Atoi(filterValue)
		if err != nil {
			return nil, fmt.Errorf("invalid numeric value: %v", err)
		}

		for _, property := range properties {
			if ApplyFilter(numExtractor(property), value, filterType) {
				filtered = append(filtered, property)
			}
		}
		return filtered, nil
	}

	// Amenities filter
	if filterKey == "amenities" {
		if filterType != "include" {
			return nil, fmt.Errorf("invalid filter type for amenities, expected 'include'")
		}
		for _, property := range properties {
			if included, exists := property.Amenities[filterValue]; exists && included {
				filtered = append(filtered, property)
			}
		}
		return filtered, nil
	}

	// Description filter
	if filterKey == "description" {
		if filterType != "match" {
			return nil, fmt.Errorf("invalid filter type for description, expected 'match'")
		}
		for _, property := range properties {
			if strings.Contains(strings.ToLower(property.Description), strings.ToLower(filterValue)) {
				filtered = append(filtered, property)
			}
		}
		return filtered, nil
	}

	// Location filter
	if filterKey == "location" {
		parts := strings.Split(filterValue, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("expected format lat,lon,radius")
		}

		refCoords, err := geo.ParseCoordinates(parts[0] + "," + parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid reference coordinates: %v", err)
		}

		radius, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid radius value: %v", err)
		}

		for _, property := range properties {
			distance := geo.Haversine(property.Location, refCoords)
			if distance <= radius {
				filtered = append(filtered, property)
			}
		}
		return filtered, nil
	}

	return nil, fmt.Errorf("unsupported filter key: %s", filterKey)
}

func ApplyFilter(propertyValue, filterValue int, filterType string) bool {
	switch filterType {
	case "equal":
		return propertyValue == filterValue
	case "lessThan":
		return propertyValue < filterValue
	case "greaterThan":
		return propertyValue > filterValue
	default:
		return false
	}
}
