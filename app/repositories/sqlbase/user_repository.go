package sqlbase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/app/models"
	rErrors "github.com/Andrew-UA/product-list/app/repositories/errors"
	"github.com/Masterminds/squirrel"
	"time"
)

type UserRepository struct {
	DB     *sql.DB
	Format squirrel.PlaceholderFormat
	Table  string
}

func (u *UserRepository) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.
		Select(
			"id",
			"first_name",
			"second_name",
			"email",
			"nickname",
			"role",
			"password_hash",
			"created_at",
			"updated_at",
			"deleted_at",
		).
		From(u.Table)
}

func (u *UserRepository) ScanUser(scanner squirrel.RowScanner, user *models.User) error {
	err := scanner.Scan(
		&user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Nickname,
		&user.Role,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	return err
}

func (u *UserRepository) GetList(ctx context.Context) ([]models.User, error) {
	query, args, err := u.SelectBuilder().
		Where(squirrel.Eq{"deleted_at": nil}).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := u.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = u.ScanUser(rows, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetById(ctx context.Context, userId uint64) (*models.User, error) {
	query, args, err := u.SelectBuilder().
		Where(squirrel.Eq{
			"id":         userId,
			"deleted_at": nil, // ігноруємо soft-deleted
		}).
		PlaceholderFormat(squirrel.Question). // для MySQL
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := u.DB.QueryRowContext(ctx, query, args...)

	var user models.User
	err = u.ScanUser(row, &user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, rErrors.NotFoundError
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query, args, err := u.SelectBuilder().
		Where(squirrel.Eq{
			"email":      email,
			"deleted_at": nil,
		}).
		PlaceholderFormat(squirrel.Question).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := u.DB.QueryRowContext(ctx, query, args...)

	var user models.User
	err = u.ScanUser(row, &user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, rErrors.NotFoundError
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Create(ctx context.Context, data dto.UserDTO) (*models.User, error) {
	now := time.Now()
	insertData := make(map[string]any)

	if data.FirstName != nil {
		insertData["first_name"] = *data.FirstName
	}
	if data.SecondName != nil {
		insertData["second_name"] = *data.SecondName
	}
	if data.Email != nil {
		insertData["email"] = *data.Email
	}
	if data.Role != nil {
		insertData["role"] = *data.Role
	}
	if nickname, ok := data.Nickname.Value(); ok {
		insertData["nickname"] = nickname
	}
	if password, ok := data.Password.Value(); ok {
		insertData["password_hash"] = password
	}

	insertData["created_at"] = now
	insertData["updated_at"] = now

	query, args, err := squirrel.
		Insert(u.Table).
		SetMap(insertData).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return nil, err
	}

	result, err := u.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return u.GetById(ctx, uint64(lastInsertID))
}

func (u *UserRepository) Update(ctx context.Context, data dto.UserDTO, user *models.User) (*models.User, error) {
	updateData := make(map[string]any)

	if data.FirstName != nil {
		updateData["first_name"] = *data.FirstName
	}
	if data.SecondName != nil {
		updateData["second_name"] = *data.SecondName
	}
	if data.Email != nil {
		updateData["email"] = *data.Email
	}
	if data.Role != nil {
		updateData["role"] = *data.Role
	}
	if nickname, ok := data.Nickname.Value(); ok {
		updateData["nickname"] = nickname
	}
	if password, ok := data.Password.Value(); ok {
		updateData["password_hash"] = password
	}

	if len(updateData) == 0 {
		return user, nil
	}

	updateData["updated_at"] = time.Now()

	query, args, err := squirrel.
		Update(u.Table).
		SetMap(updateData).
		Where(squirrel.Eq{"id": user.ID}).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = u.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return u.GetById(ctx, user.ID)
}

func (u *UserRepository) Delete(ctx context.Context, user *models.User) error {
	now := time.Now()

	query, args, err := squirrel.
		Update(u.Table).
		Set("deleted_at", now).
		Where(squirrel.Eq{"id": user.ID}).
		PlaceholderFormat(squirrel.Question).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.DB.ExecContext(ctx, query, args...)
	return err
}
