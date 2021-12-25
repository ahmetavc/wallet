package main

import (
	"github.com/ahmetavc/wallet/internal/application"
	"github.com/ahmetavc/wallet/pkg/infra"
)

func main() {
	repo := infra.NewCouchbaseRepository()

	service := application.NewService(repo)

	infra.SetupRouter(service)
}