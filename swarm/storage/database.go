// This file is part of the go-sberex library. The go-sberex library is 
// free software: you can redistribute it and/or modify it under the terms 
// of the GNU Lesser General Public License as published by the Free 
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// The go-sberex library is distributed in the hope that it will be useful, 
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser 
// General Public License <http://www.gnu.org/licenses/> for more details.

package storage

// this is a clone of an earlier state of the sberex ethdb/database
// no need for queueing/caching

import (
	"fmt"

	"github.com/Sberex/go-sberex/compression/rle"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const openFileLimit = 128

type LDBDatabase struct {
	db   *leveldb.DB
	comp bool
}

func NewLDBDatabase(file string) (*LDBDatabase, error) {
	// Open the db
	db, err := leveldb.OpenFile(file, &opt.Options{OpenFilesCacheCapacity: openFileLimit})
	if err != nil {
		return nil, err
	}

	database := &LDBDatabase{db: db, comp: false}

	return database, nil
}

func (self *LDBDatabase) Put(key []byte, value []byte) {
	if self.comp {
		value = rle.Compress(value)
	}

	err := self.db.Put(key, value, nil)
	if err != nil {
		fmt.Println("Error put", err)
	}
}

func (self *LDBDatabase) Get(key []byte) ([]byte, error) {
	dat, err := self.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	if self.comp {
		return rle.Decompress(dat)
	}

	return dat, nil
}

func (self *LDBDatabase) Delete(key []byte) error {
	return self.db.Delete(key, nil)
}

func (self *LDBDatabase) LastKnownTD() []byte {
	data, _ := self.Get([]byte("LTD"))

	if len(data) == 0 {
		data = []byte{0x0}
	}

	return data
}

func (self *LDBDatabase) NewIterator() iterator.Iterator {
	return self.db.NewIterator(nil, nil)
}

func (self *LDBDatabase) Write(batch *leveldb.Batch) error {
	return self.db.Write(batch, nil)
}

func (self *LDBDatabase) Close() {
	// Close the leveldb database
	self.db.Close()
}
