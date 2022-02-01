package store

import (
	"bytes"
	"encoding/gob"

	"github.com/peterbourgon/diskv/v3"
)

const (
	DefaultCacheSizeMax = 1024 * 1024 // 1MB cache.
	DefaultBasePath     = ".data"
)

// Store represents a persistent key-value store using
// the filesystem to store arbitrary data on fhe disk.
// Under the hood, it uses a cache for increased perfomance.
type Store struct {
	disk *diskv.Diskv
}

// New returns a Store that stores data on the disk under baseDir
// and with an internal cache of cacheSizeMax bytes.
func New(baseDir string, cacheSizeMax uint64) Store {
	// Simplest transform function: put all the data files into the base dir.
	flatTransform := func(s string) []string { return []string{} }

	d := diskv.New(diskv.Options{
		BasePath:     baseDir,
		Transform:    flatTransform,
		CacheSizeMax: cacheSizeMax,
	})
	return Store{disk: d}
}

// NewDefault returns a Store configured with default options.
func NewDefault() Store {
	return New(DefaultBasePath, DefaultCacheSizeMax)
}

// Set writes the key-value pair to the disk, encoding the value to bytes.
// Returns a non-nil error if any is encountered during the process.
func (s Store) Set(key string, value interface{}) error {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(value)
	if err != nil {
		return err
	}
	return s.disk.Write(key, buf.Bytes())
}

// Get reads the key and returns the decoded value.
// Returns a non-nil error if any is encountered during the process.
func (s Store) Get(key string) (value interface{}, err error) {
	bslice, err := s.disk.Read(key)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	_, err = buf.Write(bslice)
	if err != nil {
		return nil, err
	}

	err = gob.NewDecoder(&buf).Decode(value)
	if err != nil {
		return nil, err
	}
	return value, nil
}
