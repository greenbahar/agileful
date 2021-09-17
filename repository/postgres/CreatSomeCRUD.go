package postgres

import (
	"context"
	"math/rand"
	"strconv"
	"testTask/domain"
	"time"
)

var user = &domain.UserInfoModel{
	Name: "Emma",
	Email: "emma@mial.com",
	PassWord: "123",
}

func (r *Repository) BenchMarkResult() ( []*domain.ExecutionTimeModel, error) {
	err := r.createSampleData()
	if err!=nil{
		return nil, err
	}
	err = r.updateSampleData()
	if err!=nil{
		return nil, err
	}
	err = r.findByIDSampleData()
	if err!=nil{
		return nil, err
	}
	err = r.findSampleData()
	if err!=nil{
		return nil, err
	}
	err = r.paginateSampleData()
	if err!=nil{
		return nil, err
	}
	err = r.deleteSampleData()
	if err!=nil{
		return nil, err
	}

	records,err := r.sortByTimeSpent()
	if err!=nil{
		return nil, err
	}

	return records,nil
}

// createSampleData create sample data
func (r *Repository) createSampleData () error{
	for i:=1;i<50;i++{
		user.Name="ema"+strconv.Itoa(i)
		err:=r.create(user)
		if err!=nil{
			return err
		}
	}
	return nil
}

// updateSampleData update the created sample data
func (r *Repository) updateSampleData () error{
	for i:=1;i<30;i++{
		user.Name = "Rose"
		randomID := rand.Intn(100)
		user.ID = randomID
		err:=r.update(user)
		if err!=nil{
			return err
		}
	}
	return nil
}

// findByIDSampleData find by ID among the created sample data
func (r *Repository) findByIDSampleData () error{
	for i:=1;i<30;i++{
		randomID := rand.Intn(100)
		_,err:=r.findByID(randomID)
		if err!=nil{
			return err
		}
	}
	return nil
}

// findSampleData select limit=10 rows of the created sample data
func (r *Repository) findSampleData () error{
	for i:=1;i<3;i++{
		_,err:=r.find()
		if err!=nil{
			return err
		}
	}
	return nil
}

// paginateSampleData return pages of the created sample data
func (r *Repository) paginateSampleData () error{
	for i:=1;i<10;i++{
		_,err:=r.pagination(i)
		if err!=nil{
			return err
		}
	}
	return nil
}

// deleteSampleData delete by random ID from created sample data
func (r *Repository) deleteSampleData () error{
	for i:=1;i<10;i++{
		randomID := rand.Intn(100)
		err:=r.delete(randomID)
		if err!=nil{
			return err
		}
	}
	return nil
}

// sortByTimeSpent
func (r *Repository) sortByTimeSpent() ([]*domain.ExecutionTimeModel, error){
	records := make([]*domain.ExecutionTimeModel, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, "SELECT id, query, time_spent FROM executionTime ORDER BY time_spent ASC")
	if err != nil {
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
