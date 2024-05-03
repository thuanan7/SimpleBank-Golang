package gapi

import (
	"fmt"

	db "github.com/thuanan7/SimpleBank-Golang/db/sqlc"
	"github.com/thuanan7/SimpleBank-Golang/pb"
	"github.com/thuanan7/SimpleBank-Golang/token"
	"github.com/thuanan7/SimpleBank-Golang/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
