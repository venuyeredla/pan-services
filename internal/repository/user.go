package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/venuyeredla/pan-services/internal/models"

	_ "github.com/go-sql-driver/mysql" // _ refers to for side effects.
)

/*
1. User, with roles.
2. Password need to be encrypted.
3. How do you implment password reset.
4. Password generate
*/
const (
	AUTH_QUERY = `SELECT U.user_id, U.first_name, U.last_name, U.email, GROUP_CONCAT(R.ROLE_NAME ORDER BY R.ROLE_NAME ASC SEPARATOR ', ') as ROLES from USER U INNER JOIN USER_ROLE UR on U.user_id=UR.USER_ID INNER JOIN ROLE R on UR.ROLE_ID=R.ROLE_ID where U.email ='%s' AND U.pwd='%s' GROUP BY U.user_id,
U.first_name, U.last_name, U.email`
)

func Create(custoemr *models.User) {
	ctx := context.Background()
	sqlConn, con_err := db_con_pool.Conn(ctx)
	if con_err != nil {
		log.Fatal("Exceptin happned in getting connection")
	}
	defer sqlConn.Close()
	tx, tx_error := sqlConn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if tx_error != nil {
		log.Fatal("Error in initialzing transaction")
	}
	insert_query := "insert into ecom_user(first_name, last_name, email,pwd) values ('venu','gopal','venugopal@ecom.com','ecom#24'"
	_, exec_erro := sqlConn.ExecContext(ctx, insert_query)

	if exec_erro != nil {
		log.Fatal(exec_erro.Error())
	}
	tx.Commit()
}

func Authenticate(authReq models.AuthRequest) (*models.User, error) {
	ctx := context.Background()
	sqlConn, con_err := db_con_pool.Conn(ctx)
	if con_err != nil {
		log.Fatal("Exceptin happned in getting connection")
	}
	defer sqlConn.Close()
	tx, tx_error := sqlConn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	if tx_error != nil {
		log.Fatal("Error in initialzing transaction")
	}

	auth_query := fmt.Sprintf(AUTH_QUERY, authReq.UserName, authReq.Password)

	row := sqlConn.QueryRowContext(ctx, auth_query)
	var euser models.User
	var rolesString string
	scanError := row.Scan(&euser.Id, &euser.Firstname, &euser.Lastname, &euser.Email, &rolesString)
	if scanError != nil {
		log.Default().Println(scanError.Error())
		tx.Rollback()
		return nil, scanError
	}
	if len(rolesString) > 0 {
		euser.Roles = strings.Split(rolesString, ", ")
	}
	tx.Commit()
	return &euser, nil
}

func GetUserInfos() []*models.User {
	//db, err := sql.Open("sqlite3", ":memory:")
	ctx := context.Background()
	sqlConn, con_err := db_con_pool.Conn(ctx)
	if con_err != nil {
		log.Fatal("Exceptin happned in getting connection")
	}
	defer sqlConn.Close()
	rows, err := sqlConn.QueryContext(ctx, "select * from ecom_user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	custmers := make([]*models.User, 0)

	for rows.Next() {
		var customer models.User
		err = rows.Scan(&customer.Firstname, &customer.Lastname, &customer.Firstname)
		if err != nil {
			log.Fatal(err)
		}
		custmers = append(custmers, &customer)
	}
	return custmers
}

func Update(custoemr *models.User) {
	//db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open(DATA_BASE_NAME, DB_CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
