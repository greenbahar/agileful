package postgres

import (
	"database/sql"
	_ "github.com/google/uuid"
	"log"
	"testing"

	domain "testTask/domain"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var u = &domain.UserInfoModel{
	ID:    23,
	Name:  "Emma",
	Email: "emma@mail.com",
	PassWord: "123",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestfindByID(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, email, password FROM users WHERE id = \\$1"

	rows := sqlmock.NewRows([]string{"id", "namee", "email", "password"}).
		AddRow(u.ID, u.Name, u.Email, u.PassWord)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.findByID(u.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestfindByIDError(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, email, password FROM user WHERE id = \\$1"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"})

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.findByID(u.ID)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func Testfind(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, name, email, password FROM users"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(u.ID, u.Name, u.Email, u.PassWord)

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := repo.find()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func Testcreate(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO users \\(id, name, email, password\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID, u.Name, u.Email, u.PassWord).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.create(u)
	assert.NoError(t, err)
}

func TestCreateError(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO user \\(namee, email, password\\) VALUES \\(\\$1, \\$2, \\$3\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID, u.Name, u.Email, u.PassWord).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.create(u)
	assert.Error(t, err)
}

func Testupdate(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE users SET name = \\$1, email = \\$2, password = \\$3 WHERE id = \\$4"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.PassWord, u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.update(u)
	assert.NoError(t, err)
}

func TestUpdateErr(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "UPDATE user SET name = \\?, email = \\?, password = \\? WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Email, u.PassWord, u.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.update(u)
	assert.Error(t, err)
}

func Testdelete(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE FROM users WHERE id = \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.delete(u.ID)
	assert.NoError(t, err)
}

func TestdeleteError(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "DELETE FROM user WHERE id = \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.delete(u.ID)
	assert.Error(t, err)
}
