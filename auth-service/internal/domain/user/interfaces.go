package user

type PassHasher interface {
	Hash(str string) (string, error)
}

type AuthUserCreator interface {
	Create(authUser *AuthUser) error
}
