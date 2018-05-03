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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Sberex network.
var MainnetBootnodes = []string{
	// Sberex Foundation Go Bootnodes
	"enode://0e1634b3c7f97ff32a87fe18e4d70d6ddc44a0473ed375a2c0effb304d4439b521779cefe320e192d09f2861334b860d733156da6e880cfffd82e65421168b66@192.168.88.31:30303",
	"enode://0e1634b3c7f97ff32a87fe18e4d70d6ddc44a0473ed375a2c0effb304d4439b521779cefe320e192d09f2861334b860d733156da6e880cfffd82e65421168b66@176.104.8.150:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
