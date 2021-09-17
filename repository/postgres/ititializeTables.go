package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (r *Repository) CreateTables(db *sql.DB) error{
	fmt.Println("START createTable")
	defer fmt.Println("END createTable")

	r.db=db
	err := r.deleteTables()
	if err!=nil{
		return err
	}

	err = r.createUsersTable()
	if err!=nil{
		return err
	}
	err = r.createExecutionTimeTable()
	if err!=nil{
		return err
	}

	return nil
}

// DeleteTables delete all tables in the data base
func (r *Repository) deleteTables() error{
	fmt.Println("START DeleteTables")
	defer fmt.Println("END DeleteTables")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sqlStatement := "DROP TABLE users, executionTime;"

	_, err :=r.db.QueryContext(ctx,sqlStatement)
	if err!=nil{
		fmt.Println("delete err: ",err)
		return err
	}
	return nil
}

// createUsersTable create the user table in app initialization
func (r *Repository) createUsersTable() error {
	fmt.Println("START createUsersTable")
	defer fmt.Println("END createUsersTable-user table created")

	sqlStatement := "CREATE TABLE users (id serial NOT NULL," +
		"namee character varying(255) NOT NULL," +
		"email character varying(255) NOT NULL," +
		"password character varying(255) NOT NULL," +
		"PRIMARY KEY (id));"

	_, err :=r.db.Exec(sqlStatement)
	if err!=nil{
		return err
	}
	return nil
}

// createExecutionTimeTable create a log table for time spend to execute the queries
func (r *Repository) createExecutionTimeTable() error {
	fmt.Println("START createUsersTable")
	defer fmt.Println("END createUsersTable-user table created")

	sqlStatement := "CREATE TABLE executionTime (id serial NOT NULL," +
		"query character varying(255) NOT NULL," +
		"time_spent character varying(255) NOT NULL," +
		"PRIMARY KEY (id));"

	_, err :=r.db.Exec(sqlStatement)
	if err!=nil{
		return err
	}
	return nil
}

