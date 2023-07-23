package private_users

import (
	"reflect"
	"testing"
)

func TestOldest(t *testing.T) {
	type args struct {
		users []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Works for customer",
			args: args{
				users: []any{
					Customer{age: 99},
					Employee{age: 15},
					Customer{age: 26},
				}},
			want: Customer{age: 99},
		},
		{
			name: "Works for employee",
			args: args{
				users: []any{
					Customer{age: 14},
					Employee{age: 102},
					Customer{age: 32},
				}},
			want: Employee{age: 102},
		},
		{
			name: "Works for one",
			args: args{
				users: []any{
					Customer{age: 14},
				}},
			want: Customer{age: 14},
		},
		{
			name: "Works for empty",
			args: args{
				users: []any{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Oldest(tt.args.users...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Oldest() = %v, want %v", got, tt.want)
			}
		})
	}
}
