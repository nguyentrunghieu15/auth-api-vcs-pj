package main

import (
	"log"
	"net"

	"github.com/nguyentrunghieu15/auth-api-vcs-pj/pkg/db"
	"github.com/nguyentrunghieu15/auth-api-vcs-pj/pkg/repository"
	"github.com/nguyentrunghieu15/auth-api-vcs-pj/pkg/server"
	"github.com/nguyentrunghieu15/common-vcs-prj/apu/auth"
	"google.golang.org/grpc"
)

type ConfigServer struct {
	Network string
	Address string
}

func main() {

	var configDb = db.DBConfig{
		Host:     "postgre",
		Port:     "5432",
		Database: "on_demand_services_db",
		Username: "hiro",
		Password: "1",
		TimeOut:  10,
		SslModel: "prefer",
	}

	var configServer = ConfigServer{
		Network: "tcp",
		Address: "localhost:3456",
	}

	// Connect to Postgre DB
	db, err := db.FactoryConnection(configDb, db.POSTGRE)
	if err != nil {
		log.Fatalln("Can't connect DB")
	}
	log.Println("Connected Database")

	// create a listener on config server
	lis, err := net.Listen(configServer.Network, configServer.Address)

	if err != nil {
		log.Fatalf("Cant not listen on address %v\n,Error:%v\n", configServer.Address, err)
	}

	s := grpc.NewServer()

	auth.RegisterAuthServiceServer(s,
		&server.AuthServer{UserRepo: &repository.UserRepository{Db: db}})

	log.Println("Creatting GRPC server")
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Cant not create gRPC server", err)
	} else {
	}
	defer lis.Close()
}
