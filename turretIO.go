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
    "fmt"
    _ "log"
	"regexp"
	"errors"
	_ "runtime/debug"
)

// NewUser creates a new User instance.
// Must be provided a TurretInterface (TurretIO or AppEngineTurretIO)
func NewUser(inter TurretInterface) *User {
	u := &User{inter}
	return u
}

// NewTargetEmail creates a new TargetEmail instance.
// Must be provided a TurretInterface (TurretIO or AppEngineTurretIO)
func NewTargetEmail(inter TurretInterface) *TargetEmail {
	te := &TargetEmail{inter}
	return te
}

// NewTarget creates a new Target instance.
// Must be provided a TurretInterface (TurretIO or AppEngineTurretIO)
func NewTarget(inter TurretInterface) *Target {
	t := &Target{inter}
	return t
}

// NewAccount creates a new Account instance.
// Must be provided a TurretInterface (TurretIO or AppEngineTurretIO)
func NewAccount(inter TurretInterface) *Account {
	a := &Account{inter}
	return a
}

// Account provides API functionality for the account object
type Account struct {
	TH TurretInterface
}

// Get loads the account based on the owner of the authenticated API call
func (a *Account) Get() (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	url := fmt.Sprintf("%s", ACCOUNT_PATH)
	resp, err := a.TH.GetRequest(url, &payload, a.TH.GetHTTPClient())
	return resp, err
}

// Set updates the owner of the authenticated API call's account.
// Possible outgoing methods: turret.io, aws, smtp
// Options:
// turret.io: None
// aws: aws_access_key (required)
//      aws_secret_access_key (required)
// smtp: smtp_host (required)
// 		 smtp_username (required)
// 		 smtp_password (required)

func (a *Account) Set(outgoing_method string, options map[string]interface {}) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	// validate outgoing_method
	if ok, _ := regexp.MatchString(OUTGOING_METHOD_OPTIONS_REGEXP, outgoing_method); !ok {
		return nil, errors.New("Invalid outgoing method")
	}

	if outgoing_method == OUTGOING_METHOD_TURRET_IO_NAME && options != nil {
		return nil, errors.New("Turret.IO is outgoing method with non-nil options")
	}

	if outgoing_method != OUTGOING_METHOD_TURRET_IO_NAME && options == nil {
		return nil, errors.New("Non Turret.IO outgoing methods require options")
	}


	if outgoing_method == OUTGOING_METHOD_TURRET_IO_NAME {
		payload["type"] = OUTGOING_METHOD_TURRET_IO_NAME
	}

	if outgoing_method == OUTGOING_METHOD_AWS_NAME {
		// require aws_acces_key and aws_secret_access_key
		if _, ok := options[AWS_ACCESS_KEY_NAME]; !ok {
			return nil, errors.New("outgoing method AWS requires aws access key")
		}

		if _, ok := options[AWS_SECRET_ACCESS_KEY_NAME]; !ok {
			return nil, errors.New("outgoing method AWS requires aws secret access key")
		}
		payload["type"] = OUTGOING_METHOD_AWS_NAME
		payload["options"] = options

	}

	if outgoing_method == OUTGOING_METHOD_SMTP_NAME {
		// require aws_acces_key and aws_secret_access_key
		if _, ok := options[SMTP_HOST_NAME]; !ok {
			return nil, errors.New("outgoing method smtp requires options host name, username, password")
		}

		if _, ok := options[SMTP_USERNAME_NAME]; !ok {
			return nil, errors.New("outgoing method smtp requires options host name, username, password")
		}

		if _, ok := options[SMTP_PASSWORD_NAME]; !ok {
			return nil, errors.New("outgoing method smtp requires options host name, username, password")
		}

		payload["type"] = OUTGOING_METHOD_SMTP_NAME
		payload["options"] = options
	}

	url := fmt.Sprintf("%s/me", ACCOUNT_PATH)
	resp, err := a.TH.PostRequest(url, &payload, a.TH.GetHTTPClient())
	return resp, err
}

// Target provides API functionality for the target object
type Target struct {
	TH TurretInterface
}

// Get loads the target specified by target_name
func (t *Target) Get(target_name string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	url := fmt.Sprintf("%s/%s", TARGET_PATH, target_name)
	resp, err := t.TH.GetRequest(url, &payload, t.TH.GetHTTPClient())
	return resp, err
}

// Create adds a new target specified by the target_name with the attributes specified in attribute_map
// Example:
// Create("new_target", map[string]string{"location":"west coast", "logins":"10", "premium":"1"})
func (t *Target) Create(target_name string, attribute_list []map[string]interface {}) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	payload["attributes"] = attribute_list

	url := fmt.Sprintf("%s/%s", TARGET_PATH, target_name)
	resp, err := t.TH.PostRequest(url, &payload, t.TH.GetHTTPClient())
	return resp, err
}

// Update updates an existing target specified by the target_name with the attributes specified in the attribute_map.
// This works just like Create but updates an existing target.
func (t *Target) Update(target_name string, attribute_list []map[string]interface {}) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	payload["attributes"] = attribute_list

	url := fmt.Sprintf("%s/%s", TARGET_PATH, target_name)
	resp, err := t.TH.PostRequest(url, &payload, t.TH.GetHTTPClient())
	return resp, err
}

// TargetEmail provides API functionality for the TargetEmail object
type TargetEmail struct {
	TH TurretInterface
}

// Get loads the target email specified by the target_name and email_id
func (te *TargetEmail) Get(target_name string, email_id string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	url := fmt.Sprintf("%s/%s/email/%s", TARGET_EMAIL_PATH, target_name, email_id)
	resp, err := te.TH.GetRequest(url, &payload, te.TH.GetHTTPClient())
	return resp, err
}

// Create adds a new target email to the target specified by target_name with the subject, html_body, and plain_body provided
func (te *TargetEmail) Create(target_name string, subject string, html_body string, plain_body string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	payload["subject"] = subject
	payload["html"] = html_body
	payload["plain"] = plain_body

	url := fmt.Sprintf("%s/%s/email", TARGET_EMAIL_PATH, target_name)
	resp, err := te.TH.PostRequest(url, &payload, te.TH.GetHTTPClient())
	return resp, err
}

// Update updates an existing target email for the specified target_name based on the provided email_id and sets a new subject, html_body, and plain_body
func (te *TargetEmail) Update(target_name string, email_id string, subject string, html_body string, plain_body string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	payload["subject"] = subject
	payload["html"] = html_body
	payload["plain"] = plain_body

	url := fmt.Sprintf("%s/%s/email/%s", TARGET_EMAIL_PATH, target_name, email_id)
	resp, err := te.TH.PostRequest(url, &payload, te.TH.GetHTTPClient())
	return resp, err
}

// SendTest sends a test email to the target specified by target_name with email content from the email specified by email_id.
// from_email must match a verified sender on the account and the test email is sent to the address specified in recipient
func (te *TargetEmail) SendTest(target_name string, email_id string, from_email string, recipient string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	payload["email_from"] = from_email
	payload["recipient"] = recipient

	url := fmt.Sprintf("%s/%s/email/%s/sendTestEmail", TARGET_EMAIL_PATH, target_name, email_id)
	resp, err := te.TH.PostRequest(url, &payload, te.TH.GetHTTPClient())
	return resp, err
}

// Send sends an email to the target specified by target_name with email content from the email specified by email_id.
// from_email must match a verified sender on the account.
func (te *TargetEmail) Send(target_name string, email_id string, from_email string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	payload["email_from"] = from_email

	url := fmt.Sprintf("%s/%s/email/%s/send", TARGET_EMAIL_PATH, target_name, email_id)
	resp, err := te.TH.PostRequest(url, &payload, te.TH.GetHTTPClient())
	return resp, err
}

// User provides API functionality for the user object
type User struct {
	TH TurretInterface
}

// Get loads a user by email address
func (u *User) Get(email string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})

	url := fmt.Sprintf("%s/%s", USER_PATH, email)
	resp, err := u.TH.GetRequest(url, &payload, u.TH.GetHTTPClient())
	return resp, err
}

// Set updates an existing user or creates a new user with the email address specified.
// attribute_map is used to set all attributes for the user and classifies the user into matching targets
// property_map is used to set extra data for the user that's accessible when drafting emails, but not used to classify the user into targets
func (u *User) Set(email string, attribute_map map[string]string, property_map map[string]string) (*TurretIOResponse, error) {
	payload := make(map[string]interface{})
	for k, v := range attribute_map {
		payload[k] = v
	}
	if len(property_map) > 0 {
		properties := make(map[string]string)
		for k, v := range property_map {
			properties[k] = v
		}
		payload["properties"] = properties
	}
	url := fmt.Sprintf("%s/%s", USER_PATH, email)
	resp, err := u.TH.PostRequest(url, &payload, u.TH.GetHTTPClient())
	return resp, err
}

