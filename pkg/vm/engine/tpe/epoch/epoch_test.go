// Copyright 2021 Matrix Origin
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

package epoch

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEpochHandler_SetEpoch(t *testing.T) {
	convey.Convey("set/get",t, func() {
		eh := NewEpochHandler()

		for i := 0; i < 100; i++ {
			eh.SetEpoch(uint64(i))
			convey.So(uint64(i),convey.ShouldEqual, eh.GetEpoch())
		}
	})
}