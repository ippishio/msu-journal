package enter

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Users struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "store"
	first    = "Login"
	second   = "Password"
)

func Entry(Log string, Pas string) bool {
	sqlcon := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	conn, err := pgx.ParseConfig(sqlcon)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config: %v\n", err)
		os.Exit(1)
	}

	connf, err := pgx.ConnectConfig(context.Background(), conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer connf.Close(context.Background())
	choose := "SELECT Login, Password FROM example"
	res, err := connf.Query(context.Background(), choose)
	if err != nil {
		panic(err)
	}
	fmt.Println(Log + " " + Pas)
	for res.Next() {
		var u Users
		err = res.Scan(&u.Login, &u.Password)
		fmt.Println(u.Login + " " + u.Password)
		if err != nil {
			panic("no date")
		}
		if Log == u.Login {
			if Pas == u.Password {

				fmt.Println("succes entry")
				return true

			}
		}
		//fmt.Println(fmt.Sprintf("Your log is %s and Your pass is %s", u.Login, u.Password))
	}
	return false
}
