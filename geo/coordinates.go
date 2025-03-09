package geo

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseCoordinates(coordStr string) ([2]float64, error) {
	parts := strings.Split(coordStr, ",")
	if len(parts) != 2 {
		return [2]float64{}, fmt.Errorf("expected format lat,lon")
	}

	lat, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	lon, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)

	if err1 != nil || err2 != nil {
		return [2]float64{}, fmt.Errorf("invalid coordinate values")
	}

	return [2]float64{lat, lon}, nil
}
