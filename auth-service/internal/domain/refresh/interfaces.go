package refresh

type TokenRepo interface {
	Create(refreshModel *RefreshToken) error
	GetByTokenHash(tokenHash string) (*RefreshToken, error)
	RemoveByTokenHash(tokenHash string) error
}

type Hasher interface {
	Hash(str string) (string, error)
}
