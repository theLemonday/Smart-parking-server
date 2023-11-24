package database

import "errors"

var (
	ErrAccountNotExists      = errors.New("account does not exist")
	ErrAccountMoneyNotEnough = errors.New("money in account is not enough")
)

type AccountsDatabase struct {
	accounts []Account
	CarParkStatusDAO
}

func NewAccountsDatabase(repo CarParkStatusDAO) *AccountsDatabase {
	return &AccountsDatabase{
		CarParkStatusDAO: repo,
		accounts: []Account{
			{"02990256126", "vsmzXHgJzQf4NPf", 3949},
			{"02132510843", "Oi3NvzMZWNZRe1M", 5244},
			{"02592867214", "JdeJuhzdBgjTXD0", 4567},
			{"02420410205", "69pU3r2rM3xzhyS", 1026},
			{"024470037538", "x2fXFAcKbDEF3__", 2421},
			{"02816016789", "UWmVP_e1AHcEN2E", 5113},
			{"02361591854", "LrjqGE29f4YvPh3", 3141},
			{"02566757027", "Qdfnh7gRwIaVNeW", 7718},
			{"02946328896", "5qPcPeHQMlsI1q1", 1463},
			{"020115150222", "4dHHwwVHrB9GdJq", 4090},
			{"02771204945", "Te3i9C6Kz7ndw_J", 5439},
			{"02875136150", "JBz9HVzhGwet9rX", 6281},
			{"023229915526", "ezpm_kyrpfYZWAe", 4822},
			{"02840549498", "SoxX4oo3yGkP2Py", 9709},
			{"02664506679", "2lSBQTaWSlVJkTR", 8051},
			{"028060016114", "7dhWGSyqW8PNIMw", 3017},
			{"02110882985", "G_2tEipy9ldDoqn", 1526},
			{"02454396250", "_L0wA7oazIqdWhe", 7532},
			{"027298281462", "RZ3qXHoUrm7KqxT", 1218},
			{"02757546110", "fmBSqesCxJ4LK3_", 4494},
		},
	}
}
func (a AccountsDatabase) AuthenticateUser(username, password string) (Account, error) {
	for _, account := range a.accounts {
		if username == account.Username && password == account.password {
			return account, nil
		}
	}

	return Account{}, ErrAccountNotExists
}

func (a AccountsDatabase) findAccount(username string) (*Account, error) {
	for _, account := range a.accounts {
		if username == account.Username {
			return &account, nil
		}
	}

	return nil, ErrAccountNotExists
}

func (a AccountsDatabase) GetAccountProfile(username string) (Account, error) {
	account, err := a.findAccount(username)
	return *account, err
}

func (a AccountsDatabase) UserCreditAccount(username string, creditMoney int) error {
	account, err := a.findAccount(username)
	if err != nil {
		return err
	}

	account.Balance += creditMoney

	return nil
}
