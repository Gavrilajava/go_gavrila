package users

import (
	"testing"
)

func TestEmployee_Age(t *testing.T) {
	type fields struct {
		age int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns age",
			fields: fields{
				age: 32,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Employee{
				age: tt.fields.age,
			}
			if got := e.Age(); got != tt.want {
				t.Errorf("Employee.Age() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomer_Age(t *testing.T) {
	type fields struct {
		age int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "returns age",
			fields: fields{
				age: 32,
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Customer{
				age: tt.fields.age,
			}
			if got := c.Age(); got != tt.want {
				t.Errorf("Customer.Age() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxAge(t *testing.T) {
	type args struct {
		users []hasAge
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Works for customer",
			args: args{
				users: []hasAge{
					Customer{age: 99},
					Employee{age: 15},
					Customer{age: 26},
				}},
			want: 99,
		},
		{
			name: "Works for employee",
			args: args{
				users: []hasAge{
					Customer{age: 14},
					Employee{age: 102},
					Customer{age: 32},
				}},
			want: 102,
		},
		{
			name: "Works for empty",
			args: args{
				users: []hasAge{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.args.users...); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
