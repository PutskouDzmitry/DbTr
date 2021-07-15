package main

import (
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/PutskouDzmitry/DbTr/pkg/api"
	"github.com/PutskouDzmitry/DbTr/pkg/const_db"
	"github.com/PutskouDzmitry/DbTr/pkg/data"
	"github.com/PutskouDzmitry/DbTr/pkg/db"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = const_db.Host
	}
	if port == "" {
		port = const_db.Port
	}
	if user == "" {
		user = const_db.User
	}
	if dbname == "" {
		dbname = const_db.DbName
	}
	if password == "" {
		password = const_db.Password
	}
	if sslmode == "" {
		sslmode = const_db.Sslmode
	}
}

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}
	// 2. create router that allows to set routes
	r := mux.NewRouter()
	// 3. connect to data layer
	userData := data.NewBookData(conn)
	// 4. send data layer to api layer
	api.ServeUserResource(r, *userData)
	// 5. cors for making requests from any domain
	r.Use(mux.CORSMethodMiddleware(r))
	// 6. start server
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Server Listen port...", err)
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}
