package users

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

/*
	Add crud operations
	// handler -> service -> repo
	we get createrequest, add uuid, and then add to the db
*/
