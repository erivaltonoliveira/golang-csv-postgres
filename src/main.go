package main

import (
	"database/sql"
	"log"
	"strings"

	"./control"
	"./db"
	"./lib/util"
	"./model"

	_ "github.com/lib/pq"
)

var connection *sql.DB

func main() {
	var content []string
	var row []string
	var err error
	content = nil

	control.Init()

	filename := "../tmp/tmp_file.txt"
	content, err = util.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file:", err)
	}

	connection, err = db.Open()
	if err != nil {
		panic(err)
	}

	err = model.CreateTableCustomer(connection)
	if err != nil {
		panic(err)
	}

	for indice, linha := range content {

		if indice > 0 {
			row = strings.Fields(linha)
			customer := model.Customer{}
			customer.Cpf = row[0]
			customer.Private = util.StrToInt(row[1])
			customer.Incompleto = util.StrToInt(row[2])
			customer.DataUltimaCompra = util.StrToDate(row[3])
			customer.TicketMedio = util.StrToFloat(row[4])
			customer.TicketUltimaCompra = util.StrToFloat(row[5])
			customer.CnpjMaisFrequente = row[6]
			customer.CnpjUltimaCompra = row[7]
			model.InsertRowCustomer(connection, customer)
		}
	}

	model.ValidateCustomerDocs(connection)
}
