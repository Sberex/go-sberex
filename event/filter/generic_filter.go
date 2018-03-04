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

package filter

type Generic struct {
	Str1, Str2, Str3 string
	Data             map[string]struct{}

	Fn func(data interface{})
}

// self = registered, f = incoming
func (self Generic) Compare(f Filter) bool {
	var strMatch, dataMatch = true, true

	filter := f.(Generic)
	if (len(self.Str1) > 0 && filter.Str1 != self.Str1) ||
		(len(self.Str2) > 0 && filter.Str2 != self.Str2) ||
		(len(self.Str3) > 0 && filter.Str3 != self.Str3) {
		strMatch = false
	}

	for k := range self.Data {
		if _, ok := filter.Data[k]; !ok {
			return false
		}
	}

	return strMatch && dataMatch
}

func (self Generic) Trigger(data interface{}) {
	self.Fn(data)
}
