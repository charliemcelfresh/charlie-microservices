package seed_database

import (
	"context"
	"fmt"
	"github.com/charliemcelfresh/charlie-microservices/internal/config"
	"math/rand"
	"syreclabs.com/go/faker"
)

var merchantNames, userEmails = map[string]struct{}{}, map[string]struct{}{}

func SeedDatabase() {
	merchantIds, err := makeMerchants(1000)
	if err != nil {
		panic(fmt.Sprintf("merchants failed with error %v\n", err))
	}

	userIds, err := makeUsers(10000)
	if err != nil {
		panic(fmt.Sprintf("users failed with error %v\n", err))
	}

	userAccounts, err := makeUserAccounts(userIds)
	if err != nil {
		panic(fmt.Sprintf("userAccounts failed with error %v\n", err))
	}

	err = makeTransactions(merchantIds, userAccounts)
	if err != nil {
		panic(fmt.Sprintf("transactions failed with error %v\n", err))
	}
}

func makeMerchants(ct int) ([][2]string, error) {
	merchantIds := make([][2]string, 0, ct)
	merchants := make([]string, ct)
	var i int
	for i < ct {
		e := faker.Company().Name()
		if _, ok := merchantNames[e]; ok {
			continue
		} else {
			merchantNames[e] = struct{}{}
			merchants[i] = e
			i++
		}
	}

	ctx := context.Background()
	txn, err := config.DB().BeginTx(ctx, nil)
	if err != nil {
		return [][2]string{}, err
	}

	for i = range merchants {
		statement := `INSERT INTO merchants (name) VALUES ($1) RETURNING id, name`
		var m [2]string
		row := txn.QueryRowContext(ctx, statement, merchants[i])
		err = row.Scan(&m[0], &m[1])
		if err != nil {
			txn.Rollback()
			return [][2]string{}, err
		}
		merchantIds = append(merchantIds, m)
	}
	err = txn.Commit()
	return merchantIds, err
}

func makeUsers(ct int) ([]string, error) {
	emails := make([]string, ct)
	var i int
	for i < ct {
		e := faker.Internet().Email()
		if _, ok := userEmails[e]; ok {
			continue
		} else {
			userEmails[e] = struct{}{}
			emails[i] = e
			i++
		}
	}

	ctx := context.Background()
	txn, err := config.DB().BeginTx(ctx, nil)
	if err != nil {
		return []string{}, err
	}

	userIds := make([]string, 0, ct)

	for i = range emails {
		statement := `INSERT INTO users (email) VALUES ($1) returning id`
		result := txn.QueryRowContext(ctx, statement, emails[i])
		var userID string
		err = result.Scan(&userID)
		if err != nil {
			txn.Rollback()
			return []string{}, err
		}
		userIds = append(userIds, userID)
	}
	err = txn.Commit()
	return userIds, err
}

func makeUserAccounts(userIds []string) ([][2]string, error) {
	usersAndAccounts := [][2]string{}
	ctx := context.Background()
	txn, err := config.DB().BeginTx(ctx, nil)
	if err != nil {
		return [][2]string{}, err
	}

	for _, uID := range userIds {
		for _, cn := range randomCardNetworks() {
			var r [2]string
			statement := `INSERT INTO user_accounts (user_id, card_network) VALUES ($1, $2) RETURNING id, card_network`
			result := txn.QueryRowContext(ctx, statement, uID, cn)
			if err != nil {
				txn.Rollback()
				return [][2]string{}, err
			}
			err = result.Scan(&r[0], &r[1])
			if err != nil {
				txn.Rollback()
				return [][2]string{}, err
			}
			usersAndAccounts = append(usersAndAccounts, r)
		}
	}
	err = txn.Commit()
	return usersAndAccounts, err
}

func makeTransactions(merchants, userAccounts [][2]string) error {
	n := len(merchants)
	ctx := context.Background()
	txn, err := config.DB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, sl := range userAccounts {
		for i := 0; i < 100; i++ {
			r := rand.Intn(n)
			m := merchants[r]
			total := 10.0 + float64(rand.Intn(3000))/10.0
			statement := `INSERT INTO transactions (user_account_id, card_network, merchant_id, merchant_name, total) VALUES ($1, $2, $3, $4, $5)`
			_, err = txn.ExecContext(ctx, statement, sl[0], sl[1], m[0], m[1], total)
			if err != nil {
				return err
			}
		}
	}
	err = txn.Commit()
	return nil
}

func randomCardNetworks() []string {
	sl := []string{"MasterCard", "Amex", "Visa"}
	n := len(sl)
	r := rand.Intn(n)
	rand.Shuffle(n, func(i, j int) { sl[i], sl[j] = sl[j], sl[i] })
	return sl[0 : r+1]
}
