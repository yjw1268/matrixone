// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unary

import (
	"testing"

	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/testutil"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
	"github.com/smartystreets/goconvey/convey"
)

func TestHex(t *testing.T) {
	procs := testutil.NewProc()
	cases := []struct {
		name     string
		proc     *process.Process
		inputstr []string
		expected []string
		isScalar bool
	}{
		{
			name:     "Normal test",
			proc:     procs,
			inputstr: []string{"a"},
			expected: []string{"61"},
			isScalar: false,
		},
		{
			name:     "Scalar empty string",
			proc:     procs,
			inputstr: []string{""},
			expected: []string{""},
			isScalar: true,
		},
		{
			name:     "Number test",
			proc:     procs,
			inputstr: []string{"255"},
			expected: []string{"323535"},
			isScalar: true,
		},
		{
			name:     "multi row test",
			proc:     procs,
			inputstr: []string{"Hello", "Gopher!"},
			expected: []string{"48656c6c6f", "476f7068657221"},
			isScalar: false,
		},
		{
			name:     "Null",
			proc:     procs,
			expected: []string(nil),
			isScalar: true,
		},
	}
	for _, c := range cases {
		convey.Convey(c.name, t, func() {
			inVector := testutil.MakeVarcharVector(c.inputstr, nil)
			result, err := HexString([]*vector.Vector{inVector}, c.proc)
			convey.So(err, convey.ShouldBeNil)
			col := result.Col.(*types.Bytes)
			s := BytesToString(col)
			convey.So(s, convey.ShouldResemble, c.expected)
			convey.So(result.IsScalar(), convey.ShouldEqual, c.isScalar)
		})
	}

	procs = testutil.NewProc()
	cases2 := []struct {
		name     string
		proc     *process.Process
		inputnum []int64
		expected []string
		isScalar bool
	}{
		{
			name:     "Number test",
			proc:     procs,
			inputnum: []int64{255},
			expected: []string{"FF"},
			isScalar: false,
		}, {
			name:     "Number test",
			proc:     procs,
			inputnum: []int64{231323423423421},
			expected: []string{"D2632E7B3BBD"},
			isScalar: false,
		}, {
			name:     "Number test",
			proc:     procs,
			inputnum: []int64{123, 234, 345},
			expected: []string{"7B", "EA", "159"},
			isScalar: false,
		}, {
			name:     "Null",
			proc:     procs,
			expected: []string(nil),
			isScalar: true,
		},
	}
	for _, c := range cases2 {
		convey.Convey(c.name, t, func() {
			inVector := testutil.MakeInt64Vector(c.inputnum, nil)
			result, err := HexInt64([]*vector.Vector{inVector}, c.proc)
			convey.So(err, convey.ShouldBeNil)
			col := result.Col.(*types.Bytes)
			s := BytesToString(col)
			convey.So(s, convey.ShouldResemble, c.expected)
			convey.So(result.IsScalar(), convey.ShouldEqual, c.isScalar)
		})
	}

}

func BytesToString(src *types.Bytes) []string {
	var s []string
	for i := 0; i < int(len(src.Lengths)); i++ {
		s = append(s, string(src.Data[src.Offsets[i]:src.Offsets[i]+src.Lengths[i]]))
	}
	return s
}
