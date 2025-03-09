package property

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"property-cli/models"
)

func LoadProperties(source string) ([]models.Property, error) {
	var data []byte
	var err error

	if strings.HasPrefix(source, "http") {
		resp, err := http.Get(source)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = os.ReadFile(source)
		if err != nil {
			return nil, err
		}
	}

	var properties []models.Property
	err = json.Unmarshal(data, &properties)
	if err != nil {
		return nil, err
	}

	return properties, nil
}
