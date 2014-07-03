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

package turretIO

import (
	"testing"
	"github.com/turretIO/turret-io-go"
	_ "fmt"
)

const API_KEY="YWJjMTIz"
const API_SECRET="ZGVmZ2hp"
const EMAIL_TEST="test@example.com"
const TARGET_NAME="eastwest"
const EMAIL_ID="aec71acea"
const TARGET_EMAIL_SUBJ="Subject"
const TARGET_EMAIL_HTML_BODY="<p>HTML Body</p>"
const TARGET_EMAIL_PLAIN_BODY="Plain body"
const TARGET_EMAIL_FROM="test@example.com"
const TARGET_EMAIL_RECIPIENT="test@example.com"
const OUTGOING_METHOD_TURRET="turret.io"
const OUTGOING_METHOD_AWS="aws"
const OUTGOING_METHOD_SMTP="smtp"
const AWS_ACCESS_KEY_NAME="aws_access_key"
const AWS_SECRET_ACCESS_KEY_NAME="aws_secret_access_key"

func TestNewInstance(t *testing.T) {
	inst := turretIO.NewTurretIO(API_KEY, API_SECRET)
	if inst.Apikey != API_KEY || inst.Apisecret != API_SECRET {
		t.Errorf("API Key or API Secret not setting")
	}
}


func TestNewUserInstance(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewUser(turret)
	if inst.TH.GetApikey() != API_KEY || inst.TH.GetApisecret() != API_SECRET {
		t.Errorf("NewUserInstance not setting API key or secret")
	}
}

func TestGetUser(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewUser(turret)
	resp, err := inst.Get(EMAIL_TEST)
	if err != nil {
		t.Errorf("GetUser error")
	}
	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestGetUser")
	}
}

func TestSetUser(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewUser(turret)
	attr_map := make(map[string]string)
	prop_map := make(map[string]string)

	attr_map["location"] = "midwest"
	prop_map["full_name"] = "john smith"
	resp, err := inst.Set(EMAIL_TEST, attr_map, prop_map)
	if err != nil {
		t.Errorf("SetUser error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestSetUser")
	}
}

func TestNewTargetEmailInstance(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTargetEmail(turret)
	if inst.TH.GetApikey() != API_KEY || inst.TH.GetApisecret() != API_SECRET {
		t.Errorf("NewTargetEmailInstance not setting API key or secret")
	}
}

func TestGetTargetEmail(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTargetEmail(turret)
	resp, err := inst.Get(TARGET_NAME, EMAIL_ID)
	if err != nil {
		t.Errorf("GetTargetEmail error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestGetTargetEmail")
	}
}

func TestCreateTargetEmail(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTargetEmail(turret)
	resp, err := inst.Create(TARGET_NAME, TARGET_EMAIL_SUBJ, TARGET_EMAIL_HTML_BODY, TARGET_EMAIL_PLAIN_BODY)
	if err != nil {
		t.Errorf("CreateTargetEmail error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestCreateTargetEmail")
	}
}

func TestUpdateTargetEmail(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTargetEmail(turret)
	resp, err := inst.Update(TARGET_NAME, EMAIL_ID, TARGET_EMAIL_SUBJ, TARGET_EMAIL_HTML_BODY, TARGET_EMAIL_PLAIN_BODY)
	if err != nil {
		t.Errorf("UpdateTargetEmail error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestUpdateTargetEmail")
	}
}

func TestSendTestTargetEmail(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTargetEmail(turret)
	resp, err := inst.SendTest(TARGET_NAME, EMAIL_ID, TARGET_EMAIL_FROM, TARGET_EMAIL_RECIPIENT)
	if err != nil {
		t.Errorf("SendTestTargetEmail error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestSendTestTargetEmail")
	}
}

func TestSendTargetEmail(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTargetEmail(turret)
	resp, err := inst.Send(TARGET_NAME, EMAIL_ID, TARGET_EMAIL_FROM)
	if err != nil {
		t.Errorf("SendTargetEmail error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestSendTargetEmail")
	}
}

func TestNewTargetInstance(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTarget(turret)
	if inst.TH.GetApikey() != API_KEY || inst.TH.GetApisecret() != API_SECRET {
		t.Errorf("NewTarget not setting API key or secret")
	}
}

func TestGetTarget(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTarget(turret)
	resp, err := inst.Get(TARGET_NAME)
	if err != nil {
		t.Errorf("GetTarget error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestGetTarget")
	}
}

func TestCreateTarget(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTarget(turret)
	attr_map := make(map[string]interface {})
	attr_map["location"] = "midwest"
	resp, err := inst.Create(TARGET_NAME, attr_map)

	if err != nil {
		t.Errorf("CreateTarget error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestGetTarget")
	}
}

func TestUpdateTarget(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewTarget(turret)
	attr_map := make(map[string]interface {})
	attr_map["location"] = "midwest"
	resp, err := inst.Update(TARGET_NAME, attr_map)

	if err != nil {
		t.Errorf("UpdateTarget error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestUpdateTarget")
	}
}

func TestNewAccountInstance(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewAccount(turret)
	if inst.TH.GetApikey() != API_KEY || inst.TH.GetApisecret() != API_SECRET {
		t.Errorf("NewTarget not setting API key or secret")
	}
}

func TestGetAccount(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewAccount(turret)
	resp, err := inst.Get()
	if err != nil {
		t.Errorf("GetAccount error")
	}

	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestGetAccount")
	}
}

func TestSetAccount(t *testing.T) {
	turret := turretIO.NewTurretIO(API_KEY, API_SECRET)
	inst := turretIO.NewAccount(turret)
	resp, err := inst.Set(OUTGOING_METHOD_TURRET, nil)
	if err != nil {
		t.Errorf("SetAccount error")
	}
	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestSetAccount")
	}

	resp, err = inst.Set("invalid_method", nil)
	if err == nil { // This is an error *IF NO ERROR* if produced
		t.Errorf("Set should return error if an invalid method is provided")
	}

	resp, err = inst.Set(OUTGOING_METHOD_AWS, nil)
	if err == nil { // This is an error *IF NO ERROR* if produced
		t.Errorf("Set should return error if AWS is used and no options present")
	}

	options := make(map[string]interface {})
	options[AWS_ACCESS_KEY_NAME] = "abc123"
	resp, err = inst.Set(OUTGOING_METHOD_AWS, options)
	if err == nil { // This is an error *IF NO ERROR* if produced
		t.Errorf("Set should return error if AWS is used and only %s OR %s are present", AWS_ACCESS_KEY_NAME, AWS_SECRET_ACCESS_KEY_NAME)
	}

	options = make(map[string]interface {})
	options[AWS_SECRET_ACCESS_KEY_NAME] = "defghi"
	resp, err = inst.Set(OUTGOING_METHOD_AWS, options)
	if err == nil { // This is an error *IF NO ERROR* if produced
		t.Errorf("Set should return error if AWS is used and only %s OR %s are present", AWS_ACCESS_KEY_NAME, AWS_SECRET_ACCESS_KEY_NAME)
	}

	options = make(map[string]interface {})
	options[AWS_ACCESS_KEY_NAME] = "abc123"
	options[AWS_SECRET_ACCESS_KEY_NAME] = "defghi"

	resp, err = inst.Set(OUTGOING_METHOD_AWS, options)
	// We should see a 401 here because we're using a bad API KEY / SECRET
	if resp.Status != "401 Unauthorized" {
		t.Errorf("%s Should throw 401 Unauthorized, unless you changed the API key and secret", "TestSetAccount")
	}
}

