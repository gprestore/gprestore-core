package flip

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type FlipClient struct {
	baseUrl             *url.URL
	authorizationHeader string
	httpClient          *http.Client
}

func NewFlipClient() *FlipClient {
	secretKey := viper.GetString("flip.secret_key")
	encodedSecretKey := base64.StdEncoding.EncodeToString([]byte(secretKey + ":"))
	authorizationHeader := "Basic " + encodedSecretKey

	baseUrl, err := url.Parse(viper.GetString("flip.base_url"))
	if err != nil {
		log.Fatal(err)
	}

	return &FlipClient{
		baseUrl:             baseUrl,
		authorizationHeader: authorizationHeader,
		httpClient:          &http.Client{},
	}
}

func (s *FlipClient) CreatePayment(flipBillRequest *FlipBillRequest) (*FlipBill, error) {
	endpoint := s.baseUrl.JoinPath("/pwf/bill")

	requestJson, err := json.Marshal(flipBillRequest)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, endpoint.String(), bytes.NewBuffer(requestJson))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", s.authorizationHeader)

	resp, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var flipError *FlipError
		err = json.NewDecoder(resp.Body).Decode(&flipError)
		if err != nil {
			return nil, err
		}
		errJson, _ := json.Marshal(flipError)
		return nil, errors.New(string(errJson))
	}

	var flipBill *FlipBill
	err = json.NewDecoder(resp.Body).Decode(&flipBill)
	if err != nil {
		return nil, err
	}

	return flipBill, err
}

func (s *FlipClient) GetBills() ([]*FlipBill, error) {
	endpoint := s.baseUrl.JoinPath("/pwf/bill")

	request, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", s.authorizationHeader)

	resp, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var flipError *FlipError
		err = json.NewDecoder(resp.Body).Decode(&flipError)
		if err != nil {
			return nil, err
		}
		errJson, _ := json.Marshal(flipError)
		return nil, errors.New(string(errJson))
	}

	var flipBill []*FlipBill
	err = json.NewDecoder(resp.Body).Decode(&flipBill)
	if err != nil {
		return nil, err
	}

	return flipBill, err
}

func (s *FlipClient) GetBill(id int) (*FlipBill, error) {
	endpoint := s.baseUrl.JoinPath(fmt.Sprintf("/pwf/%d/bill", id))

	request, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", s.authorizationHeader)

	resp, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var flipError *FlipError
		err = json.NewDecoder(resp.Body).Decode(&flipError)
		if err != nil {
			return nil, err
		}
		errJson, _ := json.Marshal(flipError)
		return nil, errors.New(string(errJson))
	}

	var flipBill *FlipBill
	err = json.NewDecoder(resp.Body).Decode(&flipBill)
	if err != nil {
		return nil, err
	}

	return flipBill, err
}

func (s *FlipClient) GetPayment(id int) (*FlipPayment, error) {
	endpoint := s.baseUrl.JoinPath(fmt.Sprintf("/pwf/%d/payment", id))

	request, err := http.NewRequest(http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", s.authorizationHeader)

	resp, err := s.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var flipError *FlipError
		err = json.NewDecoder(resp.Body).Decode(&flipError)
		if err != nil {
			return nil, err
		}
		errJson, _ := json.Marshal(flipError)
		return nil, errors.New(string(errJson))
	}

	var flipPayment *FlipPayment
	err = json.NewDecoder(resp.Body).Decode(&flipPayment)
	if err != nil {
		return nil, err
	}

	return flipPayment, err
}
