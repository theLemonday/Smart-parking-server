package presenter

import "github.com/thelemonday/smart-parking-iot-server/db"

type AccountAuthenticationSuccessResponse struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

func NewAccountAuthenticationSuccessResponse(account *db.Account) *AccountAuthenticationSuccessResponse {
	return &AccountAuthenticationSuccessResponse{
		Type:     "authentication",
		Username: account.Username,
		Balance:  account.Balance,
	}
}
