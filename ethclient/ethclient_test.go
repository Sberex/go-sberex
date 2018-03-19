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

package ethclient

import "github.com/Sberex/go-sberex"

// Verify that Client implements the sberex interfaces.
var (
	_ = sberex.ChainReader(&Client{})
	_ = sberex.TransactionReader(&Client{})
	_ = sberex.ChainStateReader(&Client{})
	_ = sberex.ChainSyncReader(&Client{})
	_ = sberex.ContractCaller(&Client{})
	_ = sberex.GasEstimator(&Client{})
	_ = sberex.GasPricer(&Client{})
	_ = sberex.LogFilterer(&Client{})
	_ = sberex.PendingStateReader(&Client{})
	// _ = sberex.PendingStateEventer(&Client{})
	_ = sberex.PendingContractCaller(&Client{})
)
