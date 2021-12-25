package infra

import (
	"github.com/couchbase/gocb/v2"
	"time"
)

type CouchbaseRepository struct {
	collection *gocb.Collection
	cluster    *gocb.Cluster
}

type WalletDTO struct {
	Id      string  `json:"id"`
	Balance float64 `json:"balance"`
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

	//couchbaseBucketName := os.Getenv("COUCHBASE_BUCKET")
	couchbaseBucketName := "wallet"

	bucket := cluster.Bucket(couchbaseBucketName)

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	collection := bucket.DefaultCollection()

	return &CouchbaseRepository{collection: collection, cluster: cluster}
}

func (repository *CouchbaseRepository) Upsert(key string, value *WalletDTO) error {
	_, err := repository.collection.Upsert(key, value, &gocb.UpsertOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (repository *CouchbaseRepository) Get(key string) (WalletDTO, error) {
	data, err := repository.collection.Get(key, &gocb.GetOptions{})

	if err != nil {
		return WalletDTO{}, err
	}

	var content WalletDTO
	if err := data.Content(&content); err != nil {
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
