package repositories

type UserRepository struct {
	*BaseRepository // inheritance
}

func NewUserRepository(base *BaseRepository) *UserRepository {
	return &UserRepository{base}
}
