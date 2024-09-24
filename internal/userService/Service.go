package userService

type UserService struct {
	repo UsersRepository
}

func NewUserService(service UsersRepository) *UserService {
	return &UserService{repo: service}
}

func (r *UserService) CreateUser(user User) (User, error) {
	return r.repo.CreateUser(user)
}

func (r *UserService) GetAllUsers() ([]User, error) {
	return r.repo.GetAllUsers()
}

func (r *UserService) UpdateUserByID(id int, user User) (User, error) {
	return r.repo.UpdateUserByID(id, user)
}

func (r *UserService) DeleteUserByID(id int) error {
	return r.repo.DeleteUserByID(id)
}
