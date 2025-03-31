package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Andrew-UA/product-list/app/dto"
	"github.com/Andrew-UA/product-list/app/models"
	"github.com/Andrew-UA/product-list/app/repositories/mongodb"
	"github.com/Andrew-UA/product-list/app/repositories/mysql"
	"github.com/Andrew-UA/product-list/app/repositories/postgres"
	"github.com/Andrew-UA/product-list/app/repositories/sqlite"
	"github.com/Andrew-UA/product-list/internal/db"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	CreateToken(ctx context.Context, token models.JwtAuth) (*models.JwtAuth, error)
	GetTokenByUserId(ctx context.Context, userId uint64) (*models.JwtAuth, error)
	GetTokenByTokenId(ctx context.Context, tokenId string) (*models.JwtAuth, error)
	DeleteToken(ctx context.Context, token models.JwtAuth) error
}

type UserRepository interface {
	GetList(ctx context.Context) ([]models.User, error)
	GetById(ctx context.Context, userId uint64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, data dto.UserDTO) (*models.User, error)
	Update(ctx context.Context, data dto.UserDTO, user *models.User) (*models.User, error)
	Delete(ctx context.Context, user *models.User) error
}

type RepositoryFactory interface {
	UserRepository() UserRepository
	AuthRepository() AuthRepository
}

func GetRepositoryFactory(dbConnector db.DatabaseConnector) (RepositoryFactory, error) {
	conn, err := dbConnector.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	switch dbConnector.Type() {
	case db.MySQL:
		dbConn, _ := conn.(*sql.DB)
		return &MySqlRepository{
			db: dbConn,
		}, nil
	case db.Postgres:
		dbConn, _ := conn.(*sql.DB)
		return &PostgresRepository{
			db: dbConn,
		}, nil
	case db.SQLite:
		dbConn, _ := conn.(*sql.DB)
		return &SqliteRepository{
			db: dbConn,
		}, nil
	case db.Mongo:
		dbConn, _ := conn.(*mongo.Client)
		return &MongoRepository{
			client: dbConn,
		}, nil
	default:
		return nil, errors.New("invalid db type")
	}
}

type MySqlRepository struct {
	db       *sql.DB
	authRepo AuthRepository
	userRepo UserRepository
}

func (r *MySqlRepository) AuthRepository() AuthRepository {
	if r.authRepo == nil {
		r.authRepo = mysql.NewAuthRepository(r.db)
	}

	return r.authRepo
}

func (r *MySqlRepository) UserRepository() UserRepository {
	if r.userRepo == nil {
		r.userRepo = mysql.NewUserRepository(r.db)
	}

	return r.userRepo
}

type SqliteRepository struct {
	db       *sql.DB
	authRepo AuthRepository
	userRepo UserRepository
}

func (r *SqliteRepository) AuthRepository() AuthRepository {
	if r.authRepo == nil {
		r.authRepo = sqlite.NewAuthRepository(r.db)
	}

	return r.authRepo
}

func (r *SqliteRepository) UserRepository() UserRepository {
	if r.userRepo == nil {
		r.userRepo = sqlite.NewUserRepository(r.db)
	}

	return r.userRepo
}

type PostgresRepository struct {
	db       *sql.DB
	authRepo AuthRepository
	userRepo UserRepository
}

func (r *PostgresRepository) AuthRepository() AuthRepository {
	if r.authRepo == nil {
		r.authRepo = postgres.NewAuthRepository(r.db)
	}

	return r.authRepo
}

func (r *PostgresRepository) UserRepository() UserRepository {
	if r.userRepo == nil {
		r.userRepo = postgres.NewUserRepository(r.db)
	}

	return r.userRepo
}

type MongoRepository struct {
	client   *mongo.Client
	authRepo AuthRepository
	userRepo UserRepository
}

func (r *MongoRepository) AuthRepository() AuthRepository {
	if r.authRepo == nil {
		r.authRepo = mongodb.NewAuthRepository(r.client)
	}

	return r.authRepo
}

func (r *MongoRepository) UserRepository() UserRepository {
	if r.userRepo == nil {
		r.userRepo = mongodb.NewUserRepository(r.client)
	}

	return r.userRepo
}
