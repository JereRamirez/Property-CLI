package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"property-cli/models"
	"property-cli/property"
)

func main() {
	args := os.Args[1:]

	if len(args) < 4 {
		fmt.Println("Usage: go run main.go <source> <filterKey> <filterType> <filterValue> [-sortBy key] [-desc true/false] [-page N] [-pageSize N]")
		return
	}

	source := args[0]
	filterKey := args[1]
	filterType := args[2]
	filterValue := args[3]

	var sortBy string
	descending := false
	page, pageSize := 1, 10

	for i := 4; i < len(args); i++ {
		switch args[i] {
		case "-sortBy":
			if i+1 < len(args) {
				sortBy = args[i+1]
				i++
			}
		case "-desc":
			if i+1 < len(args) {
				descending, _ = strconv.ParseBool(args[i+1])
				i++
			}
		case "-page":
			if i+1 < len(args) {
				page, _ = strconv.Atoi(args[i+1])
				i++
			}
		case "-pageSize":
			if i+1 < len(args) {
				pageSize, _ = strconv.Atoi(args[i+1])
				i++
			}
		}
	}

	properties, err := property.LoadProperties(source)
	if err != nil {
		fmt.Println("Error loading properties:", err)
		return
	}

	properties, err = property.FilterProperties(properties, filterKey, filterType, filterValue)
	if err != nil {
		fmt.Println("Error filtering properties:", err)
		return
	}

	totalResults := len(properties)
	totalPages := (totalResults + pageSize - 1) / pageSize

	if sortBy != "" {
		properties = property.SortProperties(properties, sortBy, descending)
	}

	properties = property.PaginateProperties(properties, page, pageSize)

	response := models.Response{
		Page:         page,
		PageSize:     pageSize,
		TotalResults: totalResults,
		TotalPages:   totalPages,
		Properties:   properties,
	}

	result, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling results:", err)
		return
	}

	fmt.Println(string(result))
}
