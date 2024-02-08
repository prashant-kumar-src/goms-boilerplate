package user

type pgUserRepository struct {
	// database *Database
}

func NewPgUserRepository() *pgUserRepository {
	return &pgUserRepository{}
}
