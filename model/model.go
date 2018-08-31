package model

type User string

func (u User) String() string {
	return string(u)
}

func ConvertToInterfaceArary(users ...User) []interface{}{
	var ret []interface{}

	for _, item := range users {
		ret = append(ret, item )
	}

	return ret
}

func ConvertToStringArray(users []User) []string{
	var _ret []string
	for _, user := range users {
		_ret = append(_ret, string(user))
	}
	return _ret
}