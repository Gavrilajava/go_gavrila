package geo

import "testing"

func TestPoint_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		a            Point
		b            Point
		wantDistance float64
	}{
		{
			name:         "#1 Tokio to LA",
			a:            Point{Lat: 35.757522810596406, Lon: 139.9600793929037},
			b:            Point{Lat: 34.07328060599402, Lon: -118.3307173185072},
			wantDistance: 8783.34291575311,
		},
		{
			name:         "#2 Melbourne to Chicago",
			a:            Point{Lat: -37.815018, Lon: 44.946014},
			b:            Point{Lat: 41.881832, Lon: -87.623177},
			wantDistance: 15992.390541906436,
		},
		{
			name:         "#3 0 to 0",
			a:            Point{Lat: 0, Lon: 0},
			b:            Point{Lat: 0, Lon: 0},
			wantDistance: 0,
		},
		{
			name:         "#4 unreal to unreal",
			a:            Point{Lat: -37815018, Lon: 44946014},
			b:            Point{Lat: 41881832, Lon: -87623177},
			wantDistance: 6944.886206055825,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := tt.a.Distance(tt.b); gotDistance != tt.wantDistance {
				t.Errorf("Point.Distance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
