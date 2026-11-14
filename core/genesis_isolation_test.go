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

package core

import (
	"testing"

	"github.com/Sberex/go-sberex/common"
	"github.com/Sberex/go-sberex/params"
)

func TestDeveloperGenesisConfigIsolation(t *testing.T) {
	original := params.AllCliqueProtocolChanges.Clique.Period
	defer func() { params.AllCliqueProtocolChanges.Clique.Period = original }()
	first := DeveloperGenesisBlock(2, common.Address{1})
	firstHash := first.ToBlock(nil).Hash()
	second := DeveloperGenesisBlock(7, common.Address{2})
	if first.Config.Clique.Period != 2 {
		t.Errorf("first period changed to %d", first.Config.Clique.Period)
	}
	if second.Config.Clique.Period != 7 {
		t.Errorf("second period is %d", second.Config.Clique.Period)
	}
	if params.AllCliqueProtocolChanges.Clique.Period != original {
		t.Error("global Clique configuration changed")
	}
	if first.Config.Clique == second.Config.Clique || first.Config.Clique == params.AllCliqueProtocolChanges.Clique {
		t.Error("genesis configurations share a Clique pointer")
	}
	if first.ToBlock(nil).Hash() != firstHash {
		t.Error("first genesis hash changed")
	}
}
