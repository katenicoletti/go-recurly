package api

import (
	"encoding/xml"
	"fmt"

	recurly "github.com/blacklightcms/go-recurly"
)

var _ recurly.TransactionsService = &TransactionsService{}

// TransactionsService handles communication with the transactions related methods
// of the recurly API.
type TransactionsService struct {
	client *Client
}

// List returns a list of transactions
// https://dev.recurly.com/docs/list-transactions
func (s *TransactionsService) List(params recurly.Params) (*recurly.Response, []recurly.Transaction, error) {
	req, err := s.client.newRequest("GET", "transactions", params, nil)
	if err != nil {
		return nil, nil, err
	}

	var v struct {
		XMLName      xml.Name              `xml:"transactions"`
		Transactions []recurly.Transaction `xml:"transaction"`
	}
	resp, err := s.client.do(req, &v)

	return resp, v.Transactions, err
}

// ListAccount returns a list of transactions for an account
// https://dev.recurly.com/docs/list-accounts-transactions
func (s *TransactionsService) ListAccount(accountCode string, params recurly.Params) (*recurly.Response, []recurly.Transaction, error) {
	action := fmt.Sprintf("accounts/%s/transactions", accountCode)
	req, err := s.client.newRequest("GET", action, params, nil)
	if err != nil {
		return nil, nil, err
	}

	var v struct {
		XMLName      xml.Name              `xml:"transactions"`
		Transactions []recurly.Transaction `xml:"transaction"`
	}
	resp, err := s.client.do(req, &v)

	return resp, v.Transactions, err
}

// Get returns account and billing information at the time the transaction was
// submitted. It may not reflect the latest account information. A
// transaction_error section may be included if the transaction failed.
// Please see transaction error codes for more details.
// https://dev.recurly.com/docs/lookup-transaction
func (s *TransactionsService) Get(uuid string) (*recurly.Response, *recurly.Transaction, error) {
	action := fmt.Sprintf("transactions/%s", uuid)
	req, err := s.client.newRequest("GET", action, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var dst recurly.Transaction
	resp, err := s.client.do(req, &dst)

	return resp, &dst, err
}

// Create creates a new transaction. The Recurly API provides a shortcut for
// creating an invoice, charge, and optionally account, and processing the
// payment immediately. When creating an account all of the required account
// attributes must be supplied. When charging an existing account only the
// account_code must be supplied.
//
// See the documentation and Transaction.MarshalXML function for a detailed field list.
// https://dev.recurly.com/docs/create-transaction
func (s *TransactionsService) Create(t recurly.Transaction) (*recurly.Response, *recurly.Transaction, error) {
	req, err := s.client.newRequest("POST", "transactions", nil, t)
	if err != nil {
		return nil, nil, err
	}

	var dst recurly.Transaction
	resp, err := s.client.do(req, &dst)

	return resp, &dst, err
}
