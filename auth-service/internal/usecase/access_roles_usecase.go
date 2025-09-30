package usecase

import "github.com/google/uuid"

type AccessRolesUsecase struct {
}

func NewAccessRolesUsecase() *AccessRolesUsecase {
	return &AccessRolesUsecase{}
}

func (uc *AccessRolesUsecase) updateUserRole(userId uuid.UUID, newRole string, jwt string) {

}
