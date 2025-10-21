package qpost_api_clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"auth-service/internal/domain"
	"auth-service/internal/dto"
)

type JwtApiServiceIssuer interface {
	IssueJwtForApiService() (string, error)
}

type UserServiceApiClient struct {
	baseURL             string
	jwtApiServiceIssuer JwtApiServiceIssuer
	httpClient          *http.Client
}

func NewUserServiceClient(domain string, port string, jwtApiServiceIssuer JwtApiServiceIssuer) *UserServiceApiClient {
	baseUrl := "http://" + domain + ":" + port

	return &UserServiceApiClient{
		baseURL:             baseUrl,
		jwtApiServiceIssuer: jwtApiServiceIssuer,
		httpClient:          &http.Client{Timeout: 10 * time.Second},
	}
}

func (client *UserServiceApiClient) CreateUserRequest(us *dto.UserToCreate) error {
	url := client.baseURL + "/users/create"

	request, err := client.buildUserCreateRequest(url, us)

	if err != nil {
		return err
	}

	log.Printf("sending create user request to %v \n", url)

	resp, err := client.httpClient.Do(request)

	if err != nil {
		log.Printf("user create request error: %s \n", err)
		return err
	}

	// if 200-204
	if http.StatusOK <= resp.StatusCode && resp.StatusCode <= http.StatusNoContent {
		log.Printf("create user response: success")
		return nil
	}

	var errResp dto.ErrorResponse
	decErr := json.NewDecoder(resp.Body).Decode(&errResp)

	if decErr != nil {
		log.Printf("user create request error: %s \n", decErr)
		return decErr
	}

	log.Printf("create user error response: %v \n", errResp.Message)
	_ = resp.Body.Close()

	err = fmt.Errorf("%w: %v", domain.ErrUnexpectedApiResponse, errResp.Message)
	return err
}

func (client *UserServiceApiClient) buildUserCreateRequest(url string, us *dto.UserToCreate) (*http.Request, error) {
	reqBody, err := json.Marshal(us)

	if err != nil {
		log.Printf("user create request error: %s \n", err.Error())
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

	if err != nil {
		log.Printf("user create request error: %s \n", err.Error())
		return nil, err
	}

	jwt, jwtErr := client.jwtApiServiceIssuer.IssueJwtForApiService()

	if jwtErr != nil {
		log.Printf("user create request error: %s \n", jwtErr.Error())
		return nil, jwtErr
	}

	request.Header.Add("Authorization", jwt)
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
