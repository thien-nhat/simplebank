package gapi

import (
	"fmt"

	db "github.com/thien-nhat/simplebank/db/sqlc"
	"github.com/thien-nhat/simplebank/pb"
	"github.com/thien-nhat/simplebank/token"
	"github.com/thien-nhat/simplebank/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store db.Store
	tokenMaker token.TokenMaker
}



func NewServer(config util.Config,store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("error creating tokenMaker: %v", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	
	return server, nil
}

