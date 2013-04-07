package user

// var instantiated *user = nil

type user struct {
	Name  string
	Age   int
	Email string
}

// func GoUser() *user {
// 	if instantiated == nil {
// 		instantiated = new(user)
// 	}
// 	return instantiated
// }

func GetUser() user {
	return user{}
}
