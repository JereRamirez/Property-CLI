package property

import (
	"testing"

	"property-cli/models"
	"property-cli/property"
)

func TestPaginateProperties(t *testing.T) {
	properties := []models.Property{
		{Price: 100000}, {Price: 200000}, {Price: 300000},
		{Price: 400000}, {Price: 500000}, {Price: 600000},
	}

	// Page 1, PageSize 2
	page1 := property.PaginateProperties(properties, 1, 2)
	if len(page1) != 2 || page1[0].Price != 100000 || page1[1].Price != 200000 {
		t.Errorf("Unexpected results for page 1: %+v", page1)
	}

	// Page 2, PageSize 2
	page2 := property.PaginateProperties(properties, 2, 2)
	if len(page2) != 2 || page2[0].Price != 300000 || page2[1].Price != 400000 {
		t.Errorf("Unexpected results for page 2: %+v", page2)
	}

	// Last page
	lastPage := property.PaginateProperties(properties, 3, 2)
	if len(lastPage) != 2 || lastPage[0].Price != 500000 || lastPage[1].Price != 600000 {
		t.Errorf("Unexpected results for last page: %+v", lastPage)
	}
}
