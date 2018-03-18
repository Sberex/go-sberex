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

import (
	"hash"
)

const (
	BMTHash  = "BMT"
	SHA3Hash = "SHA3" // http://golang.org/pkg/hash/#Hash
)

type SwarmHash interface {
	hash.Hash
	ResetWithLength([]byte)
}

type HashWithLength struct {
	hash.Hash
}

func (self *HashWithLength) ResetWithLength(length []byte) {
	self.Reset()
	self.Write(length)
}
