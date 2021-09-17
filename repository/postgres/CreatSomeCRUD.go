package postgres

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	_ "sync"
	"testTask/domain"
	"time"
)

//var w sync.WaitGroup
var working = 6
var user = &domain.UserInfoModel{
	Name:     "Emma",
	Email:    "emma@mial.com",
	PassWord: "123",
}

func (r *Repository) BenchMarkResult() ([]*domain.ExecutionTimeModel, error) {
	fmt.Println("START BenchMarkResult")
	defer fmt.Println("END BenchMarkResult")

	err := r.crud()
	if err != nil {
		return nil, err
	}
	records, err := r.sortByTimeSpent()
	if err != nil {
		return nil, err
	}
	return records, err
}

// createSampleData create sample data
func (r *Repository) createSampleData() error {
	fmt.Println("START createSampleData")
	defer fmt.Println("END createSampleData")

	for i := 1; i < 50; i++ {
		user.Name = "ema" + strconv.Itoa(i)
		err := r.create(user)
		if err != nil {
			return err
		}
	}
	return nil
}

// updateSampleData update the created sample data
func (r *Repository) updateSampleData() error {
	fmt.Println("START updateSampleData")
	defer fmt.Println("END updateSampleData")

	for i := 1; i < 30; i++ {
		user.Name = "Rose"
		randomID := rand.Intn(100)
		user.ID = randomID
		err := r.update(user)
		if err != nil {
			return err
		}
	}
	return nil
}

// findByIDSampleData find by ID among the created sample data
func (r *Repository) findByIDSampleData() error {
	fmt.Println("START findByIDSampleData")
	defer fmt.Println("END findByIDSampleData")

	for i := 1; i < 30; i++ {
		randomID := rand.Intn(49)
		_, err := r.findByID(randomID)
		if err != nil {
			return err
		}
	}
	return nil
}

// findSampleData select limit=10 rows of the created sample data
func (r *Repository) findSampleData() error {
	fmt.Println("START findSampleData")
	defer fmt.Println("END findSampleData")

	for i := 1; i < 3; i++ {
		_, err := r.find()
		if err != nil {
			panic(err)
			return err
		}
	}
	return nil
}

// paginateSampleData return pages of the created sample data
func (r *Repository) paginateSampleData() error {
	fmt.Println("START paginateSampleData")
	defer fmt.Println("END paginateSampleData")

	for i := 1; i < 10; i++ {
		_, err := r.pagination(i)
		if err != nil {
			return err
		}
	}
	return nil
}

// deleteSampleData delete by random ID from created sample data
func (r *Repository) deleteSampleData() error {
	fmt.Println("START deleteSampleData")
	defer fmt.Println("END deleteSampleData")

	for i := 1; i < 10; i++ {
		randomID := rand.Intn(100)
		err := r.delete(randomID)
		if err != nil {
			return err
		}
	}
	return nil
}

// crud perform CRUD operations in the database
func (r *Repository) crud() error {
	err := r.createSampleData()
	if err != nil {
		return err
	}
	err = r.updateSampleData()
	if err != nil {
		return err
	}
	err = r.findByIDSampleData()
	if err != nil {
		return err
	}
	err = r.findSampleData()
	if err != nil {
		return err
	}
	err = r.paginateSampleData()
	if err != nil {
		return err
	}
	err = r.deleteSampleData()
	if err != nil {
		return err
	}
	return nil
}

// sortByTimeSpent
func (r *Repository) sortByTimeSpent() ([]*domain.ExecutionTimeModel, error) {
	fmt.Println("START sortByTimeSpent")
	defer fmt.Println("END sortByTimeSpent")
	records := make([]*domain.ExecutionTimeModel, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT id, query, time_spent FROM executionTime ORDER BY time_spent DESC")
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		record := new(domain.ExecutionTimeModel)
		err = rows.Scan(&record.ID, &record.Query, &record.TimeSpent)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
