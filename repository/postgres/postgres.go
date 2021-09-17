package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	domain "testTask/domain"
)

// Repository represent the repository model
type Repository struct {
	db *sql.DB
	//mu *sync.Mutex
}

// NewRepository will create a variable that represent the Repository struct
func NewRepository(repo *Repository, db *sql.DB) (*Repository, error) {
	fmt.Println("START NewRepository")
	defer fmt.Println("END START NewRepository")

	repo.db=db

	err := db.Ping()
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}

	db.SetMaxIdleConns(500)
	db.SetMaxOpenConns(500)

	return repo, nil
}

// Close attaches the provider and close the connection
func (r *Repository) Close() {
	r.db.Close()
}

// findByID attaches the user repository and find data based on id
func (r *Repository) findByID(id int) (*domain.UserInfoModel, error) {

	user := new(domain.UserInfoModel)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	queryStart := time.Now()
	err := r.db.QueryRowContext(ctx, "SELECT id, namee, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.PassWord)
	if err != nil {
		return nil, err
	}
	queryEnd := time.Now()
	executionTime := queryEnd.Sub(queryStart).String()
	r.insertTimeSpent("FindByID",executionTime)
	return user, nil
}

// find attaches the user repository and find all data
func (r *Repository) find() ([]*domain.UserInfoModel, error) {

	users := make([]*domain.UserInfoModel, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	limit:=10

	queryStart := time.Now()
	rows, err := r.db.QueryContext(ctx, "SELECT id, namee, email, password FROM users LIMIT=$1",limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	queryEnd := time.Now()
	executionTime := queryEnd.Sub(queryStart).String()
	r.insertTimeSpent("Find",executionTime,)

	for rows.Next() {
		user := new(domain.UserInfoModel)
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.PassWord,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// create attaches the user repository and creating the data
func (r *Repository) create(user *domain.UserInfoModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO users (namee, email, password) VALUES ($1, $2, $3)"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	queryStart := time.Now()
	result, err := stmt.ExecContext(ctx, user.Name, user.Email, user.PassWord)
	fmt.Println("result: ", &result)
	if err!=nil{
		return err
	}
	queryEnd := time.Now()
	executionTime := queryEnd.Sub(queryStart).String()
	err=r.insertTimeSpent("Create",executionTime)
	if err!=nil{
		panic(err)
	}
	return nil
}

// update attaches the user repository and update data based on id
func (r *Repository) update(user *domain.UserInfoModel) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "UPDATE users SET namee = $1, email = $2, password = $3 WHERE id = $4"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	queryStart := time.Now()
	_, err = stmt.ExecContext(ctx, user.Name, user.Email, user.PassWord, user.ID)
	if err!=nil{
		return err
	}
	queryEnd := time.Now()
	executionTime := queryEnd.Sub(queryStart).String()
	r.insertTimeSpent("Update",executionTime)
	return nil
}

// delete attaches the user repository and delete data based on id
func (r *Repository) delete(id int) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	queryStart := time.Now()
	_, err = stmt.ExecContext(ctx, id)
	if err!=nil{
		return err
	}
	queryEnd := time.Now()
	executionTime := queryEnd.Sub(queryStart).String()
	r.insertTimeSpent("Delete",executionTime)
	return nil
}

// pagination attach to user repository and returns a page
func (r *Repository) pagination(page int) ([]*domain.UserInfoModel, error) {

	users := make([]*domain.UserInfoModel, 0)

	limit := 10
	offset := limit * (page - 1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id,namee,email,password FROM users ORDER BY id LIMIT=$2 OFFSET=$1`

	queryStart := time.Now()
	rows, err := r.db.QueryContext(ctx,query, offset, limit)
	if err != nil {
		return nil,err
	}
	queryEnd := time.Now()
	executionTime := queryEnd.Sub(queryStart).String()
	r.insertTimeSpent("Pagination"+strconv.Itoa(page),executionTime)

	defer rows.Close()
	for rows.Next() {
		p := &domain.UserInfoModel{}
		err = rows.Scan(&p)
		if err != nil {
			log.Println(err)
			return nil,err
		}
		users = append(users,p)
	}

	return nil,nil
}

// insertTimeSpent insert the log of the time spent in query execution
func (r *Repository) insertTimeSpent(queryName,duration string) error{

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	SQL := "INSERT INTO executionTime (query, time_spent) VALUES ($1, $2)"
	stmt, err := r.db.PrepareContext(ctx, SQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, queryName, duration)
	if err!=nil{
		return err
	}

	return nil
}

