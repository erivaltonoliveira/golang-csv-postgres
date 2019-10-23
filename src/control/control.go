package control

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"database/sql"
	"log"
	"strings"

	"../db"
	"../lib/util"
	"../model"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	err = ioutil.WriteFile("/tmp/tmp_file.txt", fileBytes, 0644)
	if err != nil {
		//return err
	}
	
	processFile()

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Imported File\n")

}

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./view")))

	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func Init() {
	fmt.Println("Server running: listening")
	setupRoutes()
}

func processFile() {
	var connection *sql.DB	
	var content []string
	var col []string
	var err error
	content = nil
	filename := "/tmp/tmp_file.txt"
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
			col = strings.Fields(linha)
			customer := model.Customer{}
			customer.Cpf = col[0]
			customer.Private = util.StrToInt(col[1])
			customer.Incompleto = util.StrToInt(col[2])
			customer.DataUltimaCompra = util.StrToDate(col[3])
			customer.TicketMedio = util.StrToFloat(col[4])
			customer.TicketUltimaCompra = util.StrToFloat(col[5])
			customer.CnpjMaisFrequente = col[6]
			customer.CnpjUltimaCompra = col[7]
			model.InsertRowCustomer(connection, customer)
			
			fmt.Printf("Row inserted: %+v\n", indice)
		}
	}

	model.ValidateCustomerDocs(connection)
}