package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ApiError struct {
	Error string
}

type APIServer struct {
	ListenAddr string
	Store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
		Store:      store,
	}
}

func (s *APIServer) Run() {
	app := fiber.New()

	app.Get("/account", s.handleAccount)
	app.Get("/account/:id", s.handleGetAccount)
	app.Post("/account/new", s.handleCreateAccount)
	app.Delete("/account/delete", s.handleDeleteAccount)
	if err := app.Listen(s.ListenAddr); err != nil {

		log.Println("FAILED TO START")
	}
	log.Println("Server on port", s.ListenAddr)
}

func (s *APIServer) handleAccount(c *fiber.Ctx) error {

	acc, err := s.Store.GetAllAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(c, fiber.StatusOK, acc)
}
func (s *APIServer) handleGetAccount(c *fiber.Ctx) error {
	queries := c.Params("id")
	return WriteJSON(c, fiber.StatusOK, queries)
}

func (s *APIServer) handleCreateAccount(c *fiber.Ctx) error {
	accReq := new(CreateAccountRequest)
	if err := json.Unmarshal(c.Body(), accReq); err != nil {
		return err
	}
	account := NewAccount(accReq.FirstName, accReq.LastName)
	if err := s.Store.CreateAccount(account); err != nil {
		return err
	}
	return WriteJSON(c, fiber.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(c *fiber.Ctx) error {
	return nil
}

// func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

func WriteJSON(c *fiber.Ctx, status int, v interface{}) error {
	c.Status(status).Set("Content-Type", "application/json")

	return json.NewEncoder(c).Encode(v)
}
