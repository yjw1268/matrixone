// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unary

import (
	"encoding/hex"

	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func Hex(vectors []*vector.Vector, proc *process.Process) (*vector.Vector, error) {
	inputVector := vectors[0]
	resultType := types.Type{Oid: types.T_varchar, Size: 24}
	resultElementSize := int(resultType.Size)
	inputValues := vector.MustBytesCols(inputVector)
	if inputVector.IsScalar() {
		if inputVector.ConstVectorIsNull() {
			return proc.AllocScalarNullVector(resultType), nil
		}
		resultVector := vector.New(resultType)
		resultValues := make([]string, 1)
		vector.SetCol(resultVector, HexEncode(inputValues, resultValues))
		return resultVector, nil
	} else {
		resultVector, err := proc.AllocVector(resultType, int64(resultElementSize*len(inputValues.Lengths)))
		if err != nil {
			return nil, err
		}
		resultValues := types.DecodeStringSlice(resultVector.Data)
		resultValues = resultValues[:len(inputValues.Lengths)]
		nulls.Set(resultVector.Nsp, inputVector.Nsp)
		vector.SetCol(resultVector, HexEncode(inputValues, resultValues))
		return resultVector, nil
	}
}

func HexEncode(xs *types.Bytes, rs []string) []string {
	s := xs.GetString(0)
	dst := hex.EncodeToString([]byte(s))
	rs[0] = dst
	return rs
}
