package persistence

import (
	"context"
	"fmt"
	"os"

	"github.com/adolsalamanca/go-clean-boilerplate/internal/domain/entities"
	"github.com/adolsalamanca/go-clean-boilerplate/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PsqlRepository struct {
	pool *pgxpool.Pool
}

func NewPsqlRepository(config config.Provider) PsqlRepository {
	dbUser := config.GetString("DB_USER")
	dbHost := config.GetString("DB_HOST")
	dbPort := config.GetInt("DB_PORT")
	dbName := config.GetString("DB_NAME")

	psqlConnectString := fmt.Sprintf("postgres://%s:@%s:%d/%s", dbUser, dbHost, dbPort, dbName)

	pool, err := pgxpool.Connect(context.Background(), psqlConnectString)
	if err != nil {
		fmt.Printf("could not connect to DB, %v \n", err)
		os.Exit(1)
	}

	return PsqlRepository{
		pool: pool,
	}
}

func (p PsqlRepository) FindAllItems() ([]entities.Item, error) {
	conn, err := p.pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error acquiring connection: %v", err)
	}
	defer conn.Release()

	rows, err := conn.Query(context.Background(), "SELECT id, name, price, created_at, updated_at from ITEMS")
	var outputRows []entities.Item
	for rows.Next() {
		row := entities.Item{}
		err := rows.Scan(&row.Id, &row.Name, &row.Price, &row.CreatedAt, &row.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("unexpected error for rows.Values(): %v", err)
		}
		outputRows = append(outputRows, row)
	}

	return outputRows, nil
}

func (p PsqlRepository) StoreItem(i entities.Item) error {
	conn, err := p.pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("error acquiring connection: %v", err)
	}
	defer conn.Release()

	if _, err = conn.Exec(context.Background(), "INSERT INTO ITEMS(NAME, PRICE) VALUES($1, $2)", i.Name, i.Price); err != nil {
		return fmt.Errorf("unable to insert due to: %v", err)
	}

	fmt.Println("Inserted row successfully")
	return nil
}
