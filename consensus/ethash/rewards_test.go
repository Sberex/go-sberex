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

package ethash

import (
	"math/big"
	"sync"
	"testing"

	"github.com/Sberex/go-sberex/common"
	"github.com/Sberex/go-sberex/core/state"
	"github.com/Sberex/go-sberex/core/types"
	"github.com/Sberex/go-sberex/ethdb"
	"github.com/Sberex/go-sberex/params"
)

// Finalization by independent miners and importers must not share mutable rewards.
func TestAccumulateRewardsConcurrent(t *testing.T) {
	original := new(big.Int).Set(SberexBlockReward)
	// Preserve the global even when running this regression against the old code.
	defer SberexBlockReward.Set(original)
	tests := []struct{ number, reward int64 }{
		{1, 2428872331211146752},
		{1000000, 2417906998582689792},
		{10000000, 1544881456109227264},
		{20000000, 397526951224410752},
	}
	var wg sync.WaitGroup
	start := make(chan struct{})
	for _, test := range tests {
		wg.Add(1)
		go func(number, reward int64) {
			defer wg.Done()
			db, err := ethdb.NewMemDatabase()
			if err != nil {
				t.Error(err)
				return
			}
			defer db.Close()
			statedb, err := state.New(common.Hash{}, state.NewDatabase(db))
			if err != nil {
				t.Error(err)
				return
			}
			header := &types.Header{Number: big.NewInt(number), Coinbase: common.Address{1}}
			uncle := &types.Header{Coinbase: common.Address{2}}
			<-start
			for i := 0; i < 100; i++ {
				accumulateRewards(params.AllEthashProtocolChanges, statedb, header, []*types.Header{uncle})
				want := new(big.Int).Mul(big.NewInt(reward), big.NewInt(int64(i+1)))
				if got := statedb.GetBalance(header.Coinbase); got.Cmp(want) != 0 {
					t.Errorf("block %d iteration %d: balance %v, want %v", number, i, got, want)
					return
				}
			}
			if statedb.GetBalance(uncle.Coinbase).Sign() != 0 {
				t.Error("unexpected uncle reward")
			}
		}(test.number, test.reward)
	}
	close(start)
	wg.Wait()
	if SberexBlockReward.Cmp(original) != 0 {
		t.Error("finalization changed the global block reward")
	}
}
