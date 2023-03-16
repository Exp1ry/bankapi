package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	DeleteCompany() error
	CreateCompany(*FurnitureStore) error
	GetAccountByID(id int) (*FurnitureStore, error)
	GetAllAccounts() ([]*FurnitureStore, error)
}

type PostgressStore struct {
	db *sql.DB
}

func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil
}

func (s *PostgressStore) CreateCompany(company *FurnitureStore) error {
	query := `INSERT INTO companies
	(type, name, location, city, phone, product_and_services, created_at)
	values ($1, $2, $3, $4, $5, $6, $7)
	`
	resp, err := s.db.Exec(query, company.Type, company.Name, company.Location, company.City, company.Phone, company.ProductsAndServices, company.CreatedAt)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}

// func (s *PostgressStore) DeleteAccount(id int) error {
// 	return nil
// }
// func (s *PostgressStore) UpdateAccount(*Account) error {
// 	return nil
// }
func (s *PostgressStore) GetAllAccounts() ([]*FurnitureStore, error) {
	query := `SELECT * FROM companies`

	resp, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	companies := []*FurnitureStore{}
	for resp.Next() {
		company := new(FurnitureStore)
		err := resp.Scan(&company.ID, &company.Name, &company.Location, &company.City, &company.Phone, &company.ProductsAndServices, &company.CreatedAt)

		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}
func (s *PostgressStore) GetAccountByID(id int) (*FurnitureStore, error) {

	resp, err := s.db.Query(`SELECT id FROM companies WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	for resp.Next() {
		return scanIntoAccount(resp)
	}

	return nil, err
}

func (s *PostgressStore) DeleteCompany() error {
	query := `DROP TABLE companies`

	if _, err := s.db.Query(query); err != nil {
		return err
	}
	return nil
}

func (s *PostgressStore) Init() error {
	return s.createCompanyTable()
}
func (s *PostgressStore) createCompanyTable() error {
	err := s.DeleteCompany()
	if err != nil {
		return err
	}

	query := `CREATE table if not exists companies (
		id serial primary key,
		type varchar(500),
		name varchar(500),
		location varchar(50000),
		city varchar(500),
		phone varchar(500),
		product_and_services varchar(50000),
		created_at timestamp
		)`

	_, err2 := s.db.Exec(query)
	return err2
}

func scanIntoAccount(resp *sql.Rows) (*FurnitureStore, error) {
	company := new(FurnitureStore)
	err := resp.Scan(&company.Type, &company.ID, &company.Name, &company.Location, &company.City, &company.Phone, &company.ProductsAndServices, &company.CreatedAt)

	if err != nil {
		return nil, err
	}

	return company, err
}
