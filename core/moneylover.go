package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func GetTransactions(token string, since *time.Time, until *time.Time)([]*Transaction, error){
	input := TransactionListRequest{
		WalletID: "all",
		Force: true,
		StartDate: since,
		EndDate: until,
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(input); err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", moneyLoverBaseUrl + "/transaction/list", b)
	req.Header.Set("authorization", "AuthJWT " + token)
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		return nil, errors.New("Error! status is " + res.Status)
	}
	var transactionListResponse TransactionListResponse
	if err := json.NewDecoder(res.Body).Decode(&transactionListResponse); err != nil {
		return nil, err
	}
	return transactionListResponse.Data.Transactions, nil
}