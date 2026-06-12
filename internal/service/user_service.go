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