package delete_data

import "github.com/charliemcelfresh/charlie-microservices/internal/config"

func DeleteAll() {
	cmds := []string{
		"DELETE FROM transactions",
		"DELETE FROM user_accounts",
		"DELETE FROM users",
		"DELETE FROM merchants",
	}

	for _, cmd := range cmds {
		_, err := config.DB().Exec(cmd)
		if err != nil {
			panic(err)
		}
	}
}
