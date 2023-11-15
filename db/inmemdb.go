package db

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	accounts = []account{
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
	}

	uidTags = []string{
		"B3AD9715",
		"1365B515",
		"B3726B0F",
		"B3F1CO15",
		"536FBF15",
		"73C4B215",
	}
)

type cacheDb struct {
	currentUsers map[string]User
}

func SetupCacheDb() *cacheDb {
	return &cacheDb{
		currentUsers: map[string]User{},
	}
}

func (d *cacheDb) GetAllUsers() []User {
	var users = make([]User, 0, len(d.currentUsers))
	for _, v := range d.currentUsers {
		users = append(users, v)
	}

	return users
}

func (d *cacheDb) IsRFIDTagValid(uid string) bool {
	for _, v := range uidTags {
		if v == uid {
			return true
		}
	}

	return false
}

func (d *cacheDb) InsertNewCurrentUser(id string, identifiedBy string) {
	log.Info().Msgf("Insert new user into database")
	d.currentUsers[id] = User{id: id, identifiedBy: identifiedBy, goInTimestamp: time.Now()}
}

func (d *cacheDb) GetGoInTimestampOfUser(id string) (time.Time, error) {
	if v, ok := d.currentUsers[id]; ok {
		return v.goInTimestamp, nil
	}

	return time.Time{}, errors.New("no user with given Id")
}

func (d *cacheDb) DeleteUser(id string) {
	delete(d.currentUsers, id)
}
