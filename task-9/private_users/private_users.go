package private_users

type PrivateEmployee struct {
	age int
}

type PrivateCustomer struct {
	age int
}

func Oldest(users ...any) any {

	var m any

	for _, u := range users {
		if age(u) > age(m) {
			m = u
		}
	}
	return m
}

func age(u any) int {
	switch t := u.(type) {
	case PrivateEmployee:
		return t.age
	case PrivateCustomer:
		return t.age
	default:
		return 0
	}
}
