package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"user-service/internal/domain"
	"user-service/internal/domain/user"
)

const duplicateErrCode = "23505"

type UserRepo struct {
	*BaseRepo
}

func NewUserRepo(baseRepo *BaseRepo) *UserRepo {
	return &UserRepo{BaseRepo: baseRepo}
}

func (repo *UserRepo) GetById(id uuid.UUID) (*user.User, error) {
	ctx := context.Background()
	us, err := gorm.G[user.User](repo.db).Where("id = ?", id).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("%w: get by id: %v", domain.ErrNotFound, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%w: get by id: %v", domain.UnhandledDbError, err)
	}

	return &us, nil
}

func (repo *UserRepo) Create(us *user.User) error {
	ctx := context.Background()
	err := gorm.G[user.User](repo.db).Create(ctx, us)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == duplicateErrCode {
		return fmt.Errorf("%w: create user: %v", domain.ErrDuplicate, err)
	}

	if err != nil {
		return fmt.Errorf("%w: create user: %v", domain.UnhandledDbError, err)
	}

	return nil
}

// TODO: maybe refactor for better architecture, so that user is ensured to have updated details

func (repo *UserRepo) UpdateDetails(id uuid.UUID, details *user.UserDetails) error {
	ctx := context.Background()

	affected, err := gorm.G[user.User](repo.db).Where("id = ?", id).
		Updates(ctx, user.User{Name: details.Name, Description: details.Description, Birthday: details.Birthday})

	if affected == 0 {
		return fmt.Errorf("%w: update details", domain.ErrNotFound)
	}

	if err != nil {
		return fmt.Errorf("%w: update details: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (repo *UserRepo) SaveFollowCounts(us *user.User, ctx context.Context) error {
	tx := repo.getTx(ctx)
	affected, err := gorm.G[user.User](tx).Where("id = ?", us.Id).
		Select("friends_count", "followees_count", "followers_count").
		Updates(ctx, user.User{
			FriendsCount:   us.FriendsCount,
			FolloweesCount: us.FolloweesCount,
			FollowersCount: us.FollowersCount})

	if affected == 0 {
		return fmt.Errorf("%w: update follow counts", domain.ErrNotFound)
	}

	if err != nil {
		return fmt.Errorf("%w: update follow counts: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (repo *UserRepo) Delete(id uuid.UUID) error {
	ctx := context.Background()
	affected, err := gorm.G[user.User](repo.db).Where("id = ?", id).Delete(ctx)

	if affected == 0 {
		return fmt.Errorf("%w: delete", domain.ErrNotFound)
	}

	if err != nil {
		return fmt.Errorf("%w: delete: %v", domain.UnhandledDbError, err)
	}

	return nil
}

func (repo *UserRepo) GetUsers(ids []uuid.UUID) ([]user.User, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	ctx := context.Background()
	users, err := gorm.G[user.User](repo.db).Where("id IN ?", ids).Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("%w: get users: %v", domain.UnhandledDbError, err)
	}

	return users, nil
}

func (repo *UserRepo) ExistsBYId(id uuid.UUID) (bool, error) {
	ctx := context.Background()
	_, err := gorm.G[user.User](repo.db).Where("id = ?", id).First(ctx)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("%w: user exists by id: %v", domain.UnhandledDbError, err)
	}

	return true, nil
}
