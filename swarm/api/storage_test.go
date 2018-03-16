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

package api

import (
	"testing"
)

func testStorage(t *testing.T, f func(*Storage)) {
	testApi(t, func(api *Api) {
		f(NewStorage(api))
	})
}

func TestStoragePutGet(t *testing.T) {
	testStorage(t, func(api *Storage) {
		content := "hello"
		exp := expResponse(content, "text/plain", 0)
		// exp := expResponse([]byte(content), "text/plain", 0)
		bzzhash, err := api.Put(content, exp.MimeType)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// to check put against the Api#Get
		resp0 := testGet(t, api.api, bzzhash, "")
		checkResponse(t, resp0, exp)

		// check storage#Get
		resp, err := api.Get(bzzhash)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		checkResponse(t, &testResponse{nil, resp}, exp)
	})
}
