package infra

import (
	"github.com/ahmetavc/wallet/internal/domain/wallet"
	"github.com/couchbase/gocb/v2"
	"time"
)

type Repository interface {
	Upsert(key string, value *wallet.Wallet) error
	Get(key string) (*wallet.Wallet, error)
	Remove(key string) error
	Close() error
}

type CouchbaseRepository struct {
	collection *gocb.Collection
	cluster    *gocb.Cluster
}

func NewCouchbaseRepository() *CouchbaseRepository {
	//couchbaseHost := os.Getenv("COUCHBASE_HOST")
	couchbaseHost := "localhost:8091"
	cluster, err := gocb.Connect(couchbaseHost, gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	})

	if err != nil {
		panic(err)
	}

	bucket := cluster.Bucket("wallet")

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	collection := bucket.DefaultCollection()

	return &CouchbaseRepository{collection: collection, cluster: cluster}
}

func (repository *CouchbaseRepository) Upsert(key string, value *wallet.Wallet) error {
	_, err := repository.collection.Upsert(key, value, &gocb.UpsertOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (repository *CouchbaseRepository) Get(key string) (*wallet.Wallet, error) {
	data, err := repository.collection.Get(key, &gocb.GetOptions{})

	if err != nil {
		return nil, err
	}

	var content *wallet.Wallet
	if err := data.Content(content); err != nil {
		panic(err)
	}

	return content, nil
}

func (repository *CouchbaseRepository) Remove(key string) error {
	_, err := repository.collection.Remove(key, &gocb.RemoveOptions{})

	if err != nil {
		return err
	}

	return nil
}

func (repository *CouchbaseRepository) Close() error {
	if err := repository.cluster.Close(&gocb.ClusterCloseOptions{}); err != nil {
		return err
	}

	return nil
}
