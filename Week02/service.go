package Week02

type UserService struct {
}

func (u *UserService) GetNameWithMessage(name string) (*User, error) {
	return GetNameWithMessage(name)
}

func (u *UserService) GetNameNotCare(name string) (*User, error) {
	return GetNameNotCare(name)
}
