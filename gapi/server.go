package gapi

import (
	"fmt"

	db "github.com/thien-nhat/simplebank/db/sqlc"
	"github.com/thien-nhat/simplebank/pb"
	"github.com/thien-nhat/simplebank/token"
	"github.com/thien-nhat/simplebank/util"
	"github.com/thien-nhat/simplebank/worker"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store db.Store
	tokenMaker token.TokenMaker
	taskDistributor worker.TaskDistributor
}



func NewServer(config util.Config,store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("error creating tokenMaker: %v", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker, taskDistributor:taskDistributor}

	
	return server, nil
}

