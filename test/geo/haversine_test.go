package geo

import (
	"testing"

	"property-cli/geo"
)

func TestHaversine(t *testing.T) {
	baseLocation := [2]float64{40.7128, -74.0060}
	closerLocation := [2]float64{40.7150, -74.020}
	furtherLocation := [2]float64{40.7306, -73.9352}

	furtherDistance := geo.Haversine(baseLocation, furtherLocation)

	if furtherDistance < 5.0 {
		t.Errorf("Expected distance >= 5.0 km, got %.2f km", furtherDistance)
	}

	closerDistance := geo.Haversine(baseLocation, closerLocation)

	if closerDistance > 5.0 {
		t.Errorf("Expected distance <= 5.0 km, got %.2f km", furtherDistance)
	}
}
