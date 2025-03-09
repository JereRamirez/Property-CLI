package property

import (
	"testing"

	"property-cli/models"
	"property-cli/property"
)

func TestSortProperties(t *testing.T) {
	properties := []models.Property{
		{Price: 300000, SquareFootage: 1500},
		{Price: 200000, SquareFootage: 1200},
		{Price: 400000, SquareFootage: 1800},
	}

	sortedAsc := property.SortProperties(properties, "price", false)
	if sortedAsc[0].Price != 200000 || sortedAsc[1].Price != 300000 || sortedAsc[2].Price != 400000 {
		t.Errorf("Expected ascending order by price, got %+v", sortedAsc)
	}

	sortedDesc := property.SortProperties(properties, "price", true)
	if sortedDesc[0].Price != 400000 || sortedDesc[1].Price != 300000 || sortedDesc[2].Price != 200000 {
		t.Errorf("Expected descending order by price, got %+v", sortedDesc)
	}
}
