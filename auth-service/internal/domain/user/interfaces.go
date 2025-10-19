package user

type AuthUserCreatorGetter interface {
	Create(authUser *AuthUser) error
	GetByUsername(username string) (*AuthUser, error)
}

type HasherVerifier interface {
	Hash(str string) (string, error)
	Verify(secret string, hash string) (bool, error)
}
