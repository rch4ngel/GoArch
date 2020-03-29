package user

// The user model will include all the basic properties of a user.
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Username  string
	Password  string
	Age       int64
}

type Users []*User

var xu Users

func init() {

	u1 := User{
		ID:        1,
		FirstName: "Jubie",
		LastName:  "Kibagami",
		Username:  "Vagabond",
		Password:  "Samurai",
		Age:       32,
	}
	u2 := User{
		ID:        2,
		FirstName: "Ryu",
		LastName:  "",
		Username:  "Dragon",
		Password:  "Haduken",
		Age:       32,
	}
	u3 := User{
		ID:        3,
		FirstName: "Michael",
		LastName:  "",
		Username:  "ArchAngel",
		Password:  "Ascend",
		Age:       32,
	}
	xu = append(xu, &u1, &u2, &u3)
}

func GetUsers() Users {
	return xu
}
