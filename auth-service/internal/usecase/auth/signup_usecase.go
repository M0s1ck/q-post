package auth

import (
	"sync"

	"github.com/google/uuid"

	"auth-service/internal/domain/user"
	"auth-service/internal/dto"
	"auth-service/internal/usecase"
)

type UserCreator interface {
	Create(userId uuid.UUID, username string, pass string, role user.UserRole) error
}

type CreateUserClient interface {
	CreateUserRequest(*dto.UserToCreate) error
}

type SignUpUsecase struct {
	userCreator           UserCreator
	createUserClient      CreateUserClient
	refreshTokenGenerator RefreshTokenGenerator
	accessTokenIssuer     usecase.AccessTokenIssuer
}

func NewSignUpUsecase(userCreator UserCreator, createUserClient CreateUserClient,
	refreshTokenGenerator RefreshTokenGenerator, accessTokenIssuer usecase.AccessTokenIssuer) *SignUpUsecase {
	return &SignUpUsecase{
		userCreator:           userCreator,
		createUserClient:      createUserClient,
		refreshTokenGenerator: refreshTokenGenerator,
		accessTokenIssuer:     accessTokenIssuer,
	}
}

func (uc *SignUpUsecase) SignUpWithUsername(usPass *dto.UsernamePass) (*dto.UserIdAndTokens, error) {
	userId := uuid.New()
	username := usPass.Username
	pass := usPass.Password
	userRole := user.RoleUser

	createErr := uc.userCreator.Create(userId, username, pass, userRole)

	if createErr != nil {
		return nil, createErr
	}

	// Call here to user-service
	// TODO: update user create in user-service
	reqDto := &dto.UserToCreate{UserId: userId, Username: username}
	errChan := make(chan error, 1)
	waitGroup := &sync.WaitGroup{} // here it's also possible without wait group
	waitGroup.Add(1)
	go uc.callToCreateUserInUserService(reqDto, waitGroup, errChan)

	refresh, refreshErr := uc.refreshTokenGenerator.GenerateNewAndSave(userId)

	if refreshErr != nil {
		return nil, refreshErr
	}

	accessToken, tokenErr := uc.accessTokenIssuer.IssueAccessToken(userId, username, userRole)

	if tokenErr != nil {
		return nil, tokenErr
	}

	waitGroup.Wait()
	apiErr := <-errChan
	if apiErr != nil {
		return nil, apiErr
	}

	response := dto.UserIdAndTokens{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refresh,
	}

	return &response, nil
}

func (uc *SignUpUsecase) callToCreateUserInUserService(reqDto *dto.UserToCreate, wg *sync.WaitGroup, errChan chan error) {
	err := uc.createUserClient.CreateUserRequest(reqDto)
	errChan <- err
	wg.Done()
}
