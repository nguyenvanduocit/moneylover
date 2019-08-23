package core

import "time"

type TransactionListRequest struct {
	WalletID  string    `json:"walletId"`
	StartDate *time.Time `json:"startDate"`
	Force     bool      `json:"force"`
	Type      string    `json:"type"`
	EndDate   *time.Time `json:"endDate"`
}

type TransactionListResponse struct {
	Error  int    `json:"error"`
	Msg    string `json:"msg"`
	Action string `json:"action"`
	Data   *Data   `json:"data"`
}

type Account struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	CurrencyID  int    `json:"currency_id"`
	AccountType int    `json:"account_type"`
	Icon        string `json:"icon"`
}

type Category struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Account  string `json:"account"`
	Type     int    `json:"type"`
	Metadata string `json:"metadata"`
	Parent   *Category `json:"parent"`
}

type Transaction struct {
	ID            string        `json:"_id"`
	Note          string        `json:"note"`
	Account       *Account       `json:"account"`
	Category      *Category      `json:"category,omitempty"`
	Amount        int           `json:"amount"`
	Longtitude    int           `json:"longtitude"`
	Latitude      int           `json:"latitude"`
	DisplayDate   time.Time     `json:"displayDate"`
	Remind        int           `json:"remind"`
	Metadata      string        `json:"metadata"`
	Images        []interface{} `json:"images"`
	ExcludeReport bool          `json:"exclude_report"`
	Campaign      []interface{} `json:"campaign"`
	With          []interface{} `json:"with"`
}

type Data struct {
	Transactions []*Transaction `json:"transactions"`
}

var moneyLoverBaseUrl = "https://web.moneylover.me/api"