package user

type PassHasher interface {
	Hash(str string) (string, error)
	Verify(str string, hashedPassword string) (bool, error)
}

type AuthUserCreatorGetter interface {
	Create(authUser *AuthUser) error
	GetByUsername(username string) (*AuthUser, error)
}
