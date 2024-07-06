package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/thien-nhat/simplebank/api"
	db "github.com/thien-nhat/simplebank/db/sqlc"
	"github.com/thien-nhat/simplebank/gapi"
	"github.com/thien-nhat/simplebank/pb"
	"github.com/thien-nhat/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	runGrpcServer(config, store)
}
func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create grpc server:", err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer,server)
	reflection.Register(grpcServer) //like documentation

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	log.Printf("start grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}