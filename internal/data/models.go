package data

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Books interface {
		Insert(movie *Book) error
		Get(id int64) (*Book, error)
		GetAll(title string, genres []string, filters Filters) ([]*Book, Metadata, error)
		Update(book *Book) error
		Delete(id int64) error
	}
	Users interface {
		Insert(user *User) error
		GetByEmail(email string) (*User, error)
		Update(user *User) error
		GetForToken(tokenScope, tokenPlaintext string) (*User, error)
	}
	Tokens interface {
		New(userID int64, ttl time.Duration, scope string) (*Token, error)
		Insert(token *Token) error
		DeleteAllForUser(scope string, userID int64) error
	}
}

func NewModels(db *pgxpool.Pool) Models {
	return Models{
		Books:  BookModel{Pool: db},
		Users:  UserModel{Pool: db},
		Tokens: TokenModel{Pool: db},
	}
}

func NewMockModels() Models {
	return Models{
		Books:  MockBookModel{},
		Users:  MockUserModel{},
		Tokens: MockTokenModel{},
	}
}
