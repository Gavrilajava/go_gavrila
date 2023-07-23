package private_users

type PrivateEmployee struct {
	age int
}

type PrivateCustomer struct {
	age int
}

func Oldest(users ...any) any {
	var oldest any
	max := 0
	for _, user := range users {
		if age := age(user); age > max {
			oldest, max = user, age
		}
	}
	return oldest
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
