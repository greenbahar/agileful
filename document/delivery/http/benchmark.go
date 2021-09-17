package http

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	repo "testTask/repository/postgres"
)


type BenchMarkHandler struct {
	repo *repo.Repository
	db   *sql.DB
}

func NewBenchMarkHandler(r *fiber.App, repo *repo.Repository, db *sql.DB) (*BenchMarkHandler,error) {
	fmt.Println("START NewBenchMarkHandler")
	defer fmt.Println("END NewBenchMarkHandler")
	handler := &BenchMarkHandler{
		repo: repo,
		db: db,
	}

	//server.router.Use("/", handler.authCheck)
	r.Get("/benchmark", handler.benchmark)

	return handler,nil
}


func (b *BenchMarkHandler) benchmark(c *fiber.Ctx) error {
	fmt.Println("Start: return the slowest query ")
	defer fmt.Println("End: return the slowest query ")

	_,err:=repo.NewRepository(b.repo,b.db)
	if err!=nil{
		fmt.Println(err)
	}

	sortedBenchmark, err2 := b.repo.BenchMarkResult()
	fmt.Println(sortedBenchmark, err2)
	if err2!=nil{
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   "internal error!",
		})
	}else{
		return c.Status(400).JSON(&fiber.Map{
			"success": true,
			"sorted benchmark":   sortedBenchmark,
		})
	}

	//return c.Next()
}

func (b *BenchMarkHandler) authCheck(c *fiber.Ctx) error {
	//TODO
	fmt.Println("first handler")
	return c.Next()
}

