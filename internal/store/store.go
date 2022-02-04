package store

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

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

// Store writes the key-value pair to the disk, encoding value via gob.Encoder.
// Returns a non-nil error if any is encountered encoding or writing processes.
func (s Store) Store(key string, value interface{}) error {
	buf := &bytes.Buffer{}
	err := gob.NewEncoder(buf).Encode(value)
	if err != nil {
		return err
	}
	return s.disk.Write(key, buf.Bytes())
}

// Load reads the key and loads the value into dst. The value is decoded
// via gob.Decoder. Returns a non-nil error if any is encountered during
// reading or decoding processes.
func (s Store) Load(key string, dst interface{}) error {
	bslice, err := s.disk.Read(key)
	if err != nil {
		return fmt.Errorf("error reading disk: %#v", err)
	}

	buf := bytes.NewBuffer(bslice)
	err = gob.NewDecoder(buf).Decode(dst)
	if err != nil && err != io.EOF {
		println(err)
		return fmt.Errorf("error decoding gob: %#v", err)
	}
	return nil
}
