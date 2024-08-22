// Copyright © 2024 JR team
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package luascript_test

import (
	"context"
	"testing"

	"github.com/jrnd-io/jr/pkg/producers/luascript"
)

func TestProducer(t *testing.T) {

	testCases := []struct {
		name   string
		config luascript.Config
	}{
		{
			name: "testprint",
			config: luascript.Config{
				Script: ` json = require "json"
				          print(k) 
						  json.decode(v)`,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		someJson := `{"key": "value"}`

		t.Run(tc.name, func(t *testing.T) {
			p := &luascript.Producer{}
			p.InitializeFromConfig(tc.config)
			p.Produce(context.TODO(), []byte("somekey"), []byte(someJson), nil)
		})

	}
}