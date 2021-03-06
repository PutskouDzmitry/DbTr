package main

import (
	"fmt"
	"github.com/PutskouDzmitry/DbTr/pkg/api"
	"github.com/PutskouDzmitry/DbTr/pkg/const_db"
	"github.com/PutskouDzmitry/DbTr/pkg/data"
	"github.com/PutskouDzmitry/DbTr/pkg/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/cenkalti/backoff"
)

var (
	host       = os.Getenv("POSTGRES_HOST_SERVER")
	port       = os.Getenv("POSTGRES_PORT_SERVER")
	user       = os.Getenv("POSTGRES_USER_SERVER")
	dbname     = os.Getenv("POSTGRES_DB_NAME_SERVER")
	password   = os.Getenv("POSTGRES_PASSWORD_SERVER")
	sslmode    = os.Getenv("POSTGRES_SSLMODE_SERVER")
	portServer = os.Getenv("SERVER_OUT_PORT")
)

func initialization() {
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
	if portServer == "" {
		portServer = "8081"
	}
}

func main() {
	initialization()
	var conn *gorm.DB
	back := config()
	var err error
	for {
		timeWait := back.NextBackOff()
		time.Sleep(timeWait)
		conn, err = db.GetConnection(host, port, user, dbname, password, sslmode)
		if err != nil {
			logrus.Error("we wait connect to redis, time: ", timeWait)
		} else {
			break
		}
	}
	conn, err = db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!: %v", err)
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
	listener, err := net.Listen("tcp", fmt.Sprint(":"+portServer))
	if err != nil {
		log.Fatal("Server Listen port...", err)
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}

func config() *backoff.ExponentialBackOff {
	back := backoff.NewExponentialBackOff()
	back.MaxInterval = 20 * time.Second
	back.Multiplier = 2
	return back
}
