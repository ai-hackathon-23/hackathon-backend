package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	hd "hackathon/handler"
	repository "hackathon/repository"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost     = "db"
	dbPort     = "3306"
	dbUser     = "myuser"
	dbPassword = "mypassword"
	dbName     = "hackathon_backend"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	clientRepository := repository.NewClientRepository(db)

	clientHandler := hd.NewClientHandler(clientRepository)
	
	stateRepository := repository.NewStateRepository(db)

     stateHandler := hd.NewStateHandler(stateRepository)
	http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "GET hello!\n")
		case "POST":
			clientHandler.HandleCreateClient(w,r)
		// ...省略
		default:
			fmt.Fprint(w, "Method not allowed.\n")
		}
	})

	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "GET hello!\n")
		case "POST":
			stateHandler.HandleCreateState(w,r)
		// ...省略
		default:
			fmt.Fprint(w, "Method not allowed.\n")
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func PatienceHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	dbUser, dbPassword, dbHost, dbPort, dbName))
if err != nil {
	log.Fatalf("failed to connect to database: %v", err)
}
    ins, err := db.Prepare("INSERT INTO Clients(name,age,living_info) VALUES(?,?,?,?)")
    if err != nil {
        log.Fatal(err)
    }
	jsonStr := `["apple", "orange", "banana"]`

    result, err := ins.Exec("太郎",20,"headache",jsonStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}
