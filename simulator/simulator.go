package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/wizlif/dfcu_bank/api"
	db "github.com/wizlif/dfcu_bank/db/sqlc"
	"github.com/wizlif/dfcu_bank/util"
)

type SimulationRequest struct {
	User      *api.CreateUserRequest `json:"account"`
	AccountNo string                 `json:"account_no"`
	Loans     []*db.Loan             `json:"loans"`
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	simulations := []*SimulationRequest{}

	dummyUsers := []*api.CreateUserRequest{
		{
			Username: "acc1",
			Email:    "acc1@example.com",
			Password: "acc1@example.com",
			FullName: "Account One",
		}, {
			Username: "acc2",
			Email:    "acc2@example.com",
			Password: "acc2@example.com",
			FullName: "Account Two",
		},
	}

	// accounts := []*api.UserResponse{}

	log.Info().Msg(">>> Creating accounts")

	for _, account := range dummyUsers {
		log.Info().Msg(">> Creating " + account.Email)
		acc, err := createAccount(config, account)

		if err != nil {
			log.Fatal().Msg("> Failed to create account " + account.Email)
		}

		log.Info().Msg(">> Created " + account.Email)

		log.Info().Msg(">> Creating Loans for " + account.Email)

		loans := createLoansForAccount(config, account)

		simulations = append(simulations, &SimulationRequest{
			AccountNo: acc.AccountNo,
			User:      account,
			Loans:     loans,
		})
	}

	body, err := json.Marshal(simulations)

	if err != nil {
		log.Fatal().Err(err).Msg("Error parsing body")
	}

	// Open a new file for writing. If the file already exists, it will be truncated.
	file, err := os.Create("simulation.json")
	if err != nil {
		// Handle error
		return
	}
	defer file.Close()

	// Write a string to the file
	_, err = file.Write(body)
	if err != nil {
		// Handle error
		return
	}

	path, _ := filepath.Abs(file.Name())
	log.Info().Msg("Simulation witten to " + path)
}

func createAccount(config util.Config, request *api.CreateUserRequest) (*api.UserResponse, error) {
	// Convert request to json
	jsonReq, err := json.Marshal(request)

	if err != nil {
		log.Err(err).Msg("Error parsing request body:")
		return nil, err
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("POST", "http://"+config.HTTPServerAddress+"/users", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Err(err).Msg("Error creating HTTP request:")
		return nil, err
	}

	// Send the request and get the response
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("Error sending HTTP request:")
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("Error reading HTTP response body:")
		return nil, err
	}

	// Print the response body
	response := api.UserResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Err(err).Msg("Error parsing response body:")
		return nil, err
	}

	return &response, nil
}

func login(config util.Config, request *api.LoginUserRequest) (*api.LoginUserResponse, error) {
	// Convert request to json
	jsonReq, err := json.Marshal(request)

	if err != nil {
		log.Err(err).Msg("Error parsing request body:")
		return nil, err
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("POST", "http://"+config.HTTPServerAddress+"/users/login", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Err(err).Msg("Error creating HTTP request:")
		return nil, err
	}

	// Send the request and get the response
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("Error sending HTTP request:")
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("Error reading HTTP response body:")
		return nil, err
	}

	// Print the response body
	response := api.LoginUserResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Err(err).Msg("Error parsing response body:")
		return nil, err
	}

	return &response, nil
}

func createLoan(config util.Config, auth *api.LoginUserResponse) (*db.Loan, error) {
	// Convert request to json
	jsonReq, err := json.Marshal(api.CreateLoanRequest{
		Amount: util.RandomAmount(),
	})

	if err != nil {
		log.Err(err).Msg("Error parsing request body:")
		return nil, err
	}

	// Create a new HTTP GET request
	req, err := http.NewRequest("POST", "http://"+config.HTTPServerAddress+"/loans", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Err(err).Msg("Error creating HTTP request:")
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+auth.AccessToken)

	// Send the request and get the response
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("Error sending HTTP request:")
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("Error reading HTTP response body:")
		return nil, err
	}

	// Print the response body
	response := db.Loan{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Err(err).Msg("Error parsing response body:")
		return nil, err
	}

	return &response, nil
}

func createLoansForAccount(config util.Config, acc *api.CreateUserRequest) (loans []*db.Loan) {

	auth, err := login(config, &api.LoginUserRequest{
		Username: acc.Username,
		Password: acc.Password,
	})

	if err != nil {
		log.Fatal().Err(err)
	}

	for i := 0; i < int(util.RandomInt(1, 8)); i++ {
		loan, err := createLoan(config, auth)

		if err != nil {
			log.Fatal().Err(err).Msg("error creating loan")
		}

		loans = append(loans, loan)
	}

	return loans
}
