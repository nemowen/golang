package rpctest

type User struct {
	Id   int
	Name string
	Age  int
	Job  string
}

var users map[int]*User = make(map[int]*User, 2)

func init() {
	users[2] = &User{2, "Nora", 27, "销售"}
	users[1] = &User{1, "Nemo", 27, "软件工程师"}
}

func (user *User) GetUserById(args *int, reply *User) error {
	*reply = *users[*args]
	return nil
}

func (user *User) GetInfo(args *int, reply *string) error {
	*reply = "Nemo and Nora"
	return nil
}
