package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Users struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Email    string `json:"Email,string"`
	Products []string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "store"
	first    = "Login"
	second   = "Password"
	insert   = "UPDATE public.example SET  products = array_append(Products,'hello') WHERE Login = 'Login';"
)

func GetProfile(Log string) *Users {
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
	choose := "SELECT Login, Password, Email  FROM example"
	res, err := connf.Query(context.Background(), choose)
	if err != nil {
		panic(err)
	}
	//fmt.Println(Log + " " + Pas)
	for res.Next() {
		var u Users
		err = res.Scan(&u.Login, &u.Password, &u.Email)
		if err != nil {
			panic("no date")
		}
		if Log == u.Login {
			return &Users{
				Login:    u.Login,
				Password: u.Password,
				Email:    u.Email,
			}

		}
		//fmt.Println(fmt.Sprintf("Your log is %s and Your pass is %s", u.Login, u.Password))
	}
	//return false
	//return u;
	return &Users{
		Login:    "",
		Password: "",
		Email:    "",
	}
}
