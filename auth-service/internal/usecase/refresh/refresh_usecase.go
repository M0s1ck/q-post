package refresh

import (
	"github.com/google/uuid"

	"auth-service/internal/dto"
	"auth-service/internal/usecase"
)

type RefreshUsecase struct {
	refreshTokenVerifier RefreshTokenVerifier
	userRepo             UserGetter
	accessTokenIssuer    usecase.AccessTokenIssuer
}

func NewRefreshUsecase(refreshTokenVerifier RefreshTokenVerifier, userRepo UserGetter, accessTokenIssuer usecase.AccessTokenIssuer) *RefreshUsecase {
	return &RefreshUsecase{
		refreshTokenVerifier: refreshTokenVerifier,
		accessTokenIssuer:    accessTokenIssuer,
		userRepo:             userRepo,
	}
}

func (u RefreshUsecase) Refresh(refreshToken uuid.UUID) (*dto.UserIdAndTokens, error) {
	userId, refErr := u.refreshTokenVerifier.Verify(refreshToken)

	if refErr != nil {
		return nil, refErr
	}

	us, usErr := u.userRepo.GetById(userId)

	if usErr != nil {
		return nil, usErr
	}

	jwt, err := u.accessTokenIssuer.IssueAccessToken(us.Id, us.Username, us.Role)

	if err != nil {
		return nil, err
	}

	response := dto.UserIdAndTokens{
		UserId:       userId,
		AccessToken:  jwt,
		RefreshToken: refreshToken,
	}

	return &response, nil
}
