// Copyright 2014 Loop Science
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build appengine

package turretIO

import (
	"testing"
	"appengine/aetest"
	"github.com/turretIO/turret-io-go"
	_ "fmt"
)

func TestGAEInstance(t *testing.T) {
	// create GAE Context
	ctx, err := aetest.NewContext(nil)
	inst := turretIO.NewAppEngineTurretIO(API_KEY, API_SECRET, ctx)
	if err != nil {
		t.Errorf("Can't create GAE test context")
	}

	if inst.GAEContext != ctx {
		t.Errorf("Can't set GAEContext")
	}

	u := turretIO.NewUser(inst)
	if u.TH.GetApikey() != API_KEY || u.TH.GetApisecret() != API_SECRET {
		t.Errorf("GAE instance not setting API key/secret")
	}
}

