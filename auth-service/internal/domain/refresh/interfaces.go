package refresh

type Hasher interface {
	Hash(str string) (string, error)
}

type TokenSaver interface {
	Create(refreshModel *RefreshToken) error
}
