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

const (
	// These are the multipliers for sbr denominations.
	// Example: To get the leto value of an amount in 'sbr', use
	//
	//    new(big.Int).Mul(value, big.NewInt(params.Sbr))
	//
	Leto      = 1
	Bit       = 1e16
	Sbr       = 1e18
)
