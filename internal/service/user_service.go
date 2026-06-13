package service

import (
	"context"
	"time"

	"github.com/ShasiChowdam/user-age-api/db/sqlc"
	"github.com/ShasiChowdam/user-age-api/internal/repository"
	"github.com/ShasiChowdam/user-age-api/internal/models"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob string,
) (sqlc.User, error) {

	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return sqlc.User{}, err
	}

	return s.repo.CreateUser(
		ctx,
		sqlc.CreateUserParams{
		Name: name,
		Dob: pgtype.Date{
		Time:  parsedDOB,
		Valid: true,
	},
},
	)
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int32,
) (sqlc.User, int, error) {

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return sqlc.User{}, 0, err
	}

	age := CalculateAge(user.Dob.Time)

	return user, age, nil
}

func (s *UserService) ListUsers(
	ctx context.Context,
) ([]models.UserWithAgeResponse, error) {

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var response []models.UserWithAgeResponse

	for _, user := range users {
		response = append(response, models.UserWithAgeResponse{
			ID:   user.ID,
			Name: user.Name,
			DOB:  user.Dob.Time.Format("2006-01-02"),
			Age:  CalculateAge(user.Dob.Time),
		})
	}

	return response, nil
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int32,
	name string,
	dob string,
) (sqlc.User, error) {

	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return sqlc.User{}, err
	}

	return s.repo.UpdateUser(
		ctx,
		sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob: pgtype.Date{
		Time:  parsedDOB,
		Valid: true,
	},
},
	)
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int32,
) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsersPaginated(
	ctx context.Context,
	page int,
	limit int,
) (models.PaginatedUsersResponse, error) {

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, err := s.repo.ListUsersPaginated(
		ctx,
		int32(limit),
		int32(offset),
	)
	if err != nil {
		return models.PaginatedUsersResponse{}, err
	}

	total, err := s.repo.CountUsers(ctx)
	if err != nil {
		return models.PaginatedUsersResponse{}, err
	}

	var response []models.UserWithAgeResponse

	for _, user := range users {

		response = append(response,
			models.UserWithAgeResponse{
				ID:   user.ID,
				Name: user.Name,
				DOB:  user.Dob.Time.Format("2006-01-02"),
				Age:  CalculateAge(user.Dob.Time),
			},
		)
	}

	return models.PaginatedUsersResponse{
		Page:  page,
		Limit: limit,
		Total: total,
		Users: response,
	}, nil
}