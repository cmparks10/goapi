package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234",
		Username:  "alex",
	},
	"bob": {
		AuthToken: "5678",
		Username:  "bob",
	},
	"corey": {
		AuthToken: "9123",
		Username:  "corey",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"bob": {
		Coins:    200,
		Username: "bob",
	},
	"corey": {
		Coins:    600,
		Username: "corey",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//simulate db call
	time.Sleep(time.Second * 1)

	//get the value associated with the username
	//if it is found, assign the data to clientData, and true boolean to the 'ok' variable
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//simulate db call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatebase() error {
	return nil
}
