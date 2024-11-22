package cache

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

type FileCache struct {
	db *badger.DB
}

func NewFileCache() *FileCache {
	db, err := badger.Open(badger.DefaultOptions("./tmp/cache"))
	if err != nil {
		log.Fatal(err)
	}

	return &FileCache{
		db: db,
	}
}

func (c *FileCache) Get(key string) (string, error) {
	var value []byte

	err := c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil
		})

		return err
	})

	return string(value), err
}

func (c *FileCache) Set(key string, value string) error {
	return c.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
}

func (c *FileCache) Exists(key string) (bool, error) {
	var exists bool

	err := c.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		exists = err == nil
		return nil
	})

	return exists, err
}

func (c *FileCache) Close() error {
	return c.db.Close()
}
