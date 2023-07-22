package private_users

import "fmt"

type PrivateEmployee struct {
	age int
}

type PrivateCustomer struct {
	age int
}

func Oldest(users ...any) any {

	if len(users) == 0 {
		return nil
	}

	if len(users) == 1 {
		return users[0]
	}

	maxIndex := 0
	maxAge := 0

	fmt.Println('m')

	for i, u := range users {
		currentAge := 0
		switch t := u.(type) {
		case PrivateEmployee:
			currentAge = t.age
		case PrivateCustomer:
			currentAge = t.age
		default:
			currentAge = 0
		}
		if currentAge > maxAge {
			maxAge, maxIndex = currentAge, i
		}
	}
	return users[maxIndex]
}
