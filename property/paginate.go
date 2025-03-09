package property

import "property-cli/models"

func PaginateProperties(properties []models.Property, page, pageSize int) []models.Property {
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > len(properties) {
		return []models.Property{}
	}
	if end > len(properties) {
		end = len(properties)
	}
	return properties[start:end]
}
