package sqlbase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Andrew-UA/product-list/app/models"
	"github.com/Masterminds/squirrel"
)

type AuthRepository struct {
	DB        *sql.DB
	Format    squirrel.PlaceholderFormat
	TableName string
}

func (a AuthRepository) CreateToken(ctx context.Context, token models.JwtAuth) (*models.JwtAuth, error) {
	query, args, err := squirrel.
		Insert("jwt_auth").
		Columns("user_id", "token_id", "expires_at").
		Values(token.UserID, token.TokenID, token.ExpiresAt).
		ToSql()
	if err != nil {
		return nil, err
	}

	result, err := a.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	token.ID = uint64(id)

	return &token, nil
}

func (a AuthRepository) GetTokenByUserId(ctx context.Context, userId uint64) (*models.JwtAuth, error) {
	query, args, err := squirrel.
		Select("id", "user_id", "token_id", "created_at", "expired_at").
		From(a.TableName).
		Where(squirrel.Eq{"user_id": userId}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	var jwtAuth models.JwtAuth
	row := a.DB.QueryRowContext(ctx, query, args...)

	if err := row.Scan(&jwtAuth.ID, &jwtAuth.UserID, &jwtAuth.TokenID, &jwtAuth.CreatedAt, &jwtAuth.ExpiresAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &jwtAuth, nil

}

func (a AuthRepository) GetTokenByTokenId(ctx context.Context, tokenId string) (*models.JwtAuth, error) {
	query, args, err := squirrel.
		Select("id", "user_id", "token_id", "created_at", "expired_at").
		From(a.TableName).
		Where(squirrel.Eq{"tokenId": tokenId}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	var jwtAuth models.JwtAuth
	row := a.DB.QueryRowContext(ctx, query, args...)

	if err := row.Scan(&jwtAuth.ID, &jwtAuth.UserID, &jwtAuth.TokenID, &jwtAuth.CreatedAt, &jwtAuth.ExpiresAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &jwtAuth, nil
}

func (a AuthRepository) DeleteToken(ctx context.Context, token models.JwtAuth) error {
	query, args, err := squirrel.
		Delete(a.TableName).
		Where(squirrel.Eq{"id": token.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = a.DB.ExecContext(ctx, query, args...)
	return err
}
