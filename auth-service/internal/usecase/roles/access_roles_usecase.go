package roles

import (
	"fmt"

	"github.com/google/uuid"

	"auth-service/internal/domain"
	"auth-service/internal/domain/user"
	"auth-service/internal/usecase"
)

type UserRoleUpdater interface {
	UpdateRole(userId uuid.UUID, newRoleStr user.UserRole) error
}

type AccessRolesUsecase struct {
	repo           UserRoleUpdater
	tokenValidator usecase.AccessTokenValidator
}

func NewAccessRolesUsecase(rep UserRoleUpdater, tokenValidator usecase.AccessTokenValidator) *AccessRolesUsecase {
	return &AccessRolesUsecase{
		repo:           rep,
		tokenValidator: tokenValidator,
	}
}

func (uc *AccessRolesUsecase) UpdateUserRole(userId uuid.UUID, newRoleStr string, jwt string) error {
	claims, err := uc.tokenValidator.ValidateAccessToken(jwt)

	if err != nil {
		return fmt.Errorf("%w: %v", domain.ErrInvalidToken, err)
	}

	var newRole user.UserRole = user.RoleIdsByNames[newRoleStr]
	var claimedRole user.UserRole = claims.Role

	if claimedRole == user.RoleUser || claimedRole == user.RoleModer && newRole == user.RoleAdmin {
		return fmt.Errorf("%w: can't update role %v having role %v", domain.ErrWeakRole, newRoleStr, user.RoleNamesById[claimedRole])
	}

	repoErr := uc.repo.UpdateRole(userId, newRole)
	return repoErr
}
