package usecase

type AccessTokenApiValidator interface {
	ValidateTokenIssuedAt(token string, issuer string) error
}
