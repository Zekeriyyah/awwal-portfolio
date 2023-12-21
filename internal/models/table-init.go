package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Initialize(dbDriver *sql.DB) error {
	stmt, driverErr := dbDriver.Prepare(contact)
	checkDriverErr(driverErr)

	_, err := stmt.Exec()
	checkExecErr(err, "contact")

	stmt, driverErr = dbDriver.Prepare(user)
	checkDriverErr(driverErr)

	_, err = stmt.Exec()
	checkExecErr(err, "user")

	stmt, driverErr = dbDriver.Prepare(home)
	checkDriverErr(driverErr)

	_, err = stmt.Exec()
	checkExecErr(err, "home")

	stmt, driverErr = dbDriver.Prepare(about)
	checkDriverErr(driverErr)

	_, err = stmt.Exec()
	checkExecErr(err, "about")

	stmt, driverErr = dbDriver.Prepare(experience)
	checkDriverErr(driverErr)

	_, err = stmt.Exec()
	checkExecErr(err, "exprience")

	stmt, driverErr = dbDriver.Prepare(review)
	checkDriverErr(driverErr)

	_, err = stmt.Exec()
	checkExecErr(err, "review")

	stmt, driverErr = dbDriver.Prepare(services)
	checkDriverErr(driverErr)

	_, err = stmt.Exec()
	checkExecErr(err, "services")

	stmt, driverErr = dbDriver.Prepare(project)
	checkDriverErr(driverErr)

	_, err = stmt.Exec()
	checkExecErr(err, "project")

	return nil
}

func checkExecErr(err error, table string) {
	if err != nil {
		log.Printf("%s table execution failed!\n", table)
		panic(err)
	} else {
		fmt.Printf("%s table executed successfully!\n", table)
	}
}

func checkDriverErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
