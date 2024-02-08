package user

type IUserRepository interface {
}

type newUserRepoFunc func() *IUserRepository

func NewUserRepository(repoFunc newUserRepoFunc) *IUserRepository {
	if repoFunc != nil {
		return repoFunc()
	}
	return nil
}
