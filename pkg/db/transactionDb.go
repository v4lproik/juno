package db

import (
	"github.com/NethermindEth/juno/internal/log"
	"github.com/torquem-ch/mdbx-go/mdbx"
	"math/bits"
	"runtime"
)

// Transactioner describes methods relating to an abstract key-value
// database oriented to transactions.
type Transactioner interface {
	Databaser
	// Begin starts a new transaction.
	Begin()
	// Commit saves all the information included in the current transaction
	Commit() error
	// Rollback rollsback the database to a previous state.
	Rollback()
}

type TransactionDb struct {
	txn  *mdbx.Txn
	env  *mdbx.Env
	path string
}

// NewTransactionDb creates a new key-value database based on transactions.
func NewTransactionDb(path string, flags uint) *TransactionDb {
	env, err := mdbx.NewEnv()
	if err != nil {
		// notest
		return nil
	}

	// Set flags.
	// Based on https://github.com/torquem-ch/mdbx-go/blob/96f31f483af593377e52358a079e834256d5af55/mdbx/env_test.go#L495
	err = env.SetOption(mdbx.OptMaxDB, 1024)
	if err != nil {
		// notest
		return nil
	}
	const pageSize = 4096
	err = env.SetGeometry(-1, -1, 64*1024*pageSize, -1, -1, pageSize)
	if err != nil {
		// notest
		return nil
	}
	err = env.Open(path, bits.Reverse(mdbx.Exclusive), 0664)
	if err != nil {
		// notest
		return nil
	}
	return &TransactionDb{env: env, path: path}
}

func (d *TransactionDb) GetEnv() *mdbx.Env {
	return d.env
}
func (d *TransactionDb) initTransaction() {
	txn, err := d.env.BeginTxn(nil, 0)
	if err != nil {
		return
	}
	d.txn = txn
}

// Has returns true if the value at the provided key is in the
// database.
func (d *TransactionDb) Has(key []byte) (has bool, err error) {
	val, err := d.getOne(key)
	if err != nil {
		return false, err
	}
	return val != nil, nil
}

// getOne returns the value associated with the provided key in the
// database or returns an error otherwise.
func (d *TransactionDb) getOne(key []byte) (val []byte, err error) {
	var dbi mdbx.DBI
	dbi, err = d.txn.OpenRoot(mdbx.Create)
	if err != nil {
		return nil, err
	}
	runtime.LockOSThread()
	val, err = d.txn.Get(dbi, key)
	if err != nil {
		if mdbx.IsNotFound(err) {
			err = nil
			return nil, nil
		}
		return nil, err
	}
	runtime.UnlockOSThread()
	return val, nil
}

// Get returns the value associated with the provided key in the
// database or returns an error otherwise.
func (d *TransactionDb) Get(key []byte) ([]byte, error) {
	return d.getOne(key)
}

// Put inserts a key-value pair into the database.
func (d *TransactionDb) Put(key, value []byte) error {
	dbi, err := d.txn.OpenRoot(mdbx.Create)
	if err != nil {
		return err
	}
	runtime.LockOSThread()
	err = d.txn.Put(dbi, key, value, 0)
	runtime.UnlockOSThread()

	return err
}

// Delete removes a previous inserted key or returns an error otherwise.
func (d *TransactionDb) Delete(key []byte) error {
	db, err := d.txn.OpenRoot(mdbx.Create)
	if err != nil {
		return err
	}
	err = d.txn.Del(db, key, nil)
	if mdbx.IsNotFound(err) {
		return nil
	}
	return err
}

// NumberOfItems returns the number of items in the database.
func (d *TransactionDb) NumberOfItems() (uint64, error) {
	stats, err := d.env.Stat()
	if err != nil {
		log.Default.With("Error", err).Info("Unable to get stats from env.")
		return 0, err
	}
	return stats.Entries, err
}

// Close closes the environment.
func (d *TransactionDb) Close() {
	d.Rollback()
	d.env.Close()
}

// Commit saves all the information included in the current transaction
func (d *TransactionDb) Commit() error {
	if d.txn != nil {
		_, err := d.txn.Commit()
		if err != nil {
			d.txn = nil
			return err
		}
	}
	return nil
}

// Begin starts a new transaction.
func (d *TransactionDb) Begin() {
	d.initTransaction()
}

// Rollback rolls back the database to a previous state.
func (d *TransactionDb) Rollback() {
	if d.txn != nil {
		d.txn.Abort()
	}
}