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
					PrivateCustomer{age: 99},
					PrivateEmployee{age: 15},
					PrivateCustomer{age: 26},
				}},
			want: PrivateCustomer{age: 99},
		},
		{
			name: "Works for employee",
			args: args{
				users: []any{
					PrivateCustomer{age: 14},
					PrivateEmployee{age: 102},
					PrivateCustomer{age: 32},
				}},
			want: PrivateEmployee{age: 102},
		},
		{
			name: "Works for one",
			args: args{
				users: []any{
					PrivateCustomer{age: 14},
				}},
			want: PrivateCustomer{age: 14},
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
