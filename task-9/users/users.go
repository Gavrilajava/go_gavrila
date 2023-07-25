package users

type Employee struct {
	age int
}

type Customer struct {
	age int
}

type hasAge interface {
	Age() int
}

func (e Employee) Age() int {
	return e.age
}

func (c Customer) Age() int {
	return c.age
}

func MaxAge(users ...hasAge) int {
	m := 0
	for _, u := range users {
		if u.Age() > m {
			m = u.Age()
		}
	}
	return m
}
