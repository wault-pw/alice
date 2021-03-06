package storage

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/wault-pw/alice/pkg/validator"
)

type Storage struct {
	db        IDb
	validator validator.IValidator
	sseKey    []byte // Server-side encryption key
}

func Connect(dsn string) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", dsn)
}

func NewStorage(db *sqlx.DB, validator validator.IValidator, sseKey []byte) *Storage {
	return &Storage{db: NewSqlxDB(db), validator: validator, sseKey: sseKey}
}

func NewSavepointStorage(db *sqlx.DB, validator validator.IValidator, sseKey []byte) (*Storage, *SavepointDB) {
	savepoint := NewSavepointDB(db)
	return &Storage{db: savepoint, validator: validator, sseKey: sseKey}, savepoint
}

func (s *Storage) SetValidator(validator validator.IValidator) *Storage {
	s.validator = validator
	return s
}

// SqlDB returns raw SQL connection, required for
// goose migrations
func (s *Storage) SqlDB() *sql.DB {
	return s.db.SqlDB()
}

// Builder returns SQL Builder object
func Builder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
