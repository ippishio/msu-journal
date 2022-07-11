package connect

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Users struct {
	Login    string `json:"name"`
	Password string `json:"age"`
	Email    string `json:"email"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "store"
	first    = "Login"
	second   = "Password"
	third    = "Email"
)

func Connect(Log string, Pas string, Email string) {
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

	ins := fmt.Sprintf("INSERT INTO public.example (%s, %s,%s) VALUES ('%s','%s', '%s')", first, second, third, Log, Pas, Email)
	_, err = connf.Query(context.Background(), ins)
	if err != nil {
		panic(err)
	}

	// choose := "SELECT * FROM example"
	// _, err = connf.Query(context.Background(), choose)
	// if err != nil {
	// 	panic(err)
	// }
	// _, errs := connf.Exec(context.Background(), "DELETE FROM example WHERE ")
	// if errs != nil {
	// 	panic(err)
	// }

	// choose := "SELECT Login, Password FROM example"
	// res, err := connf.Query(context.Background(), choose)
	// if err != nil {
	// 	panic(err)
	// }

	// for res.Next() {
	// 	var u Users
	// 	err = res.Scan(&u.Login, &u.Password)
	// 	if err != nil {
	// 		panic("no date")
	// 	}
	// 	fmt.Println(fmt.Sprintf("Your log is %s and Your pass is %s", u.Login, u.Password))
	// }

}
