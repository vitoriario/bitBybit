package user

//Reader interface
type Reader interface {
	Find(id int) (*User, error)
}

//Writer bookmark writer
type Writer interface {
	Store(b *User) (int, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
