package repository

import (
	"context"

	"github.com/ShasiChowdam/user-age-api/db/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{
		queries: q,
	}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	params sqlc.CreateUserParams,
) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, params)
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *UserRepository) ListUsers(
	ctx context.Context,
) ([]sqlc.User, error) {
	return r.queries.ListUsers(ctx)
}

func (r *UserRepository) UpdateUser(
	ctx context.Context,
	params sqlc.UpdateUserParams,
) (sqlc.User, error) {
	return r.queries.UpdateUser(ctx, params)
}

func (r *UserRepository) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) ListUsersPaginated(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]sqlc.User, error) {

	return r.queries.ListUsersPaginated(
		ctx,
		sqlc.ListUsersPaginatedParams{
			Limit:  limit,
			Offset: offset,
		},
	)
}

func (r *UserRepository) CountUsers(
	ctx context.Context,
) (int64, error) {

	return r.queries.CountUsers(ctx)
}