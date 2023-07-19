package geo

import (
	"math"
)

const earthRadius = 6371.0

type Point struct {
	Lat, Lon float64
}

// converts decimal degrees to radians
func rad(c float64) float64 {
	return float64(math.Pi * c / 180)
}

// returns distance between two points in kilometers
func (a Point) Distance(b Point) float64 {

	d := math.Pow(math.Sin(rad(a.Lat-b.Lat)/2), 2) +
		math.Cos(rad(a.Lat))*math.Cos(rad(b.Lat))*
			math.Pow(math.Sin(rad(a.Lon-b.Lon)/2), 2)

	d = 2 * math.Asin(math.Sqrt(d))

	return d * earthRadius

}
