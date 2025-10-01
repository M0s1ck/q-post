package usecase

import (
	"fmt"

	"github.com/google/uuid"

	"auth-service/internal/domain"
)

type UserRoleUpdater interface {
	UpdateRole(userId uuid.UUID, newRoleStr domain.UserRole) error
}

type AccessRolesUsecase struct {
	repo           UserRoleUpdater
	tokenValidator TokenValidator
}

func NewAccessRolesUsecase(rep UserRoleUpdater, tokenValidator TokenValidator) *AccessRolesUsecase {
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

	var newRole domain.UserRole = domain.RoleIdsByNames[newRoleStr]
	var claimedRole domain.UserRole = claims.Role

	if claimedRole == domain.RoleUser || claimedRole == domain.RoleModer && newRole == domain.RoleAdmin {
		return fmt.Errorf("%w: can't update role %v having role %v", domain.ErrWeakRole, newRoleStr, domain.RoleNamesById[claimedRole]) // TODO: test
	}

	repoErr := uc.repo.UpdateRole(userId, newRole)
	return repoErr
}
