package location

import "math/rand/v2"

type Location struct {
	Latitude  float64
	Longitude float64
}

func GenerateLocations(count int) []Location {
	locations := make([]Location, count)
	for i := range locations {
		locations[i] = Location{
			Latitude:  rand.Float64()*180 - 90,
			Longitude: rand.Float64()*360 - 180,
		}
	}
	return locations
}
