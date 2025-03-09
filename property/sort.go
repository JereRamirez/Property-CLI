package property

import (
	"property-cli/models"
	"sort"
)

func SortProperties(properties []models.Property, sortBy string, descending bool) []models.Property {
	switch sortBy {
	case "price":
		sort.SliceStable(properties, func(i, j int) bool {
			if descending {
				return properties[i].Price > properties[j].Price
			}
			return properties[i].Price < properties[j].Price
		})
	case "squareFootage":
		sort.SliceStable(properties, func(i, j int) bool {
			if descending {
				return properties[i].SquareFootage > properties[j].SquareFootage
			}
			return properties[i].SquareFootage < properties[j].SquareFootage
		})
	}
	return properties
}
