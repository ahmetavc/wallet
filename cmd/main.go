package main

import "github.com/ahmetavc/wallet/pkg/infra"

func main() {
	_ = infra.NewCouchbaseRepository()
	infra.SetupRouter()
}