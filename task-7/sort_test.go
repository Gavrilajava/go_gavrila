package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestSort_Ints(t *testing.T) {
	got := []int{3, 2, 1}
	sort.Ints(got)
	want := []int{1, 2, 3}
	if reflect.DeepEqual(got, want) == false {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestSort_Strings(t *testing.T) {

	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "#1 even",
			got:  []string{"b", "c", "a"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "#2 odd",
			got:  []string{"b", "d", "a", "c"},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "#3 empty",
			got:  []string{},
			want: []string{},
		},
		{
			name: "#4 EQUAL",
			got:  []string{"b", "b", "b", "b"},
			want: []string{"b", "b", "b", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sort.Strings(tt.got); reflect.DeepEqual(tt.got, tt.want) == false {
				t.Errorf("reverse() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func sampleInt() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	return data
}

func sampleFloat() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}

	return data
}

func BenchmarkSort_Int(b *testing.B) {
	data := sampleInt()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sort.Ints(data)
		_ = data
	}
}

func BenchmarkSort_Float64(b *testing.B) {
	data := sampleFloat()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		sort.Float64s(data)
		_ = data
	}
}
