package private_users

type Employee struct {
	age int
}

type Customer struct {
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
	case Employee:
		return t.age
	case Customer:
		return t.age
	default:
		return 0
	}
}
