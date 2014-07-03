// Turret.IO Go Client package
//
// This package provides client API support for the following Turret.IO types:
// Account
// Target
// TargetEmail
// User
//
// Google AppEngine support is also provided via the AppEngineTurretIO type



package turretIO

import (
    "bytes"
    "crypto/hmac"
    "crypto/sha512"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    _ "log"
    "net/http"
    "strconv"
    "strings"
    "time"
)

const ENDPOINT = "https://api.turret.io"
const USER_PATH = "/latest/user"
const TARGET_EMAIL_PATH = "/latest/target"
const TARGET_PATH = "/latest/target"
const ACCOUNT_PATH = "/latest/account"

const OUTGOING_METHOD_OPTIONS_REGEXP="(turret\\.io|aws|smtp)"
const OUTGOING_METHOD_TURRET_IO_NAME="turret.io"
const OUTGOING_METHOD_AWS_NAME="aws"
const OUTGOING_METHOD_SMTP_NAME="smtp"

const AWS_ACCESS_KEY_NAME="aws_access_key"
const AWS_SECRET_ACCESS_KEY_NAME="aws_secret_access_key"
const SMTP_HOST_NAME="smtp_host"
const SMTP_USERNAME_NAME="smtp_username"
const SMTP_PASSWORD_NAME="smtp_password"

// NewTurretIO is used to create a new TurretIO base instance to be provided
// to other types during instantiation
func NewTurretIO(api_key string, api_secret string) *TurretIO {
    t := new(TurretIO)
    t.Apikey = api_key
    t.Apisecret = api_secret
    return t
}

type TurretIOResponse struct {
	JSONBody map[string]interface {}
	Status	string
}

type TurretInterface interface {
	GetHTTPClient() (*http.Client)
	GetRequest(url string, payload *map[string]interface {}, client *http.Client) (*TurretIOResponse, error)
	PostRequest(url string, payload *map[string]interface {}, client *http.Client) (*TurretIOResponse, error)
	GetApikey() (string)
	GetApisecret() (string)
}

type TurretIO struct {
    Apikey     string
    Apisecret  string
}

func (t *TurretIO) GetApikey() (string) {
	return t.Apikey
}

func (t *TurretIO) GetApisecret() (string) {
	return t.Apisecret
}

func (t *TurretIO) GetHTTPClient() (*http.Client) {
	return &http.Client{}
}

func (t *TurretIO) buildPath(resourceType string, resource string) string {
    return fmt.Sprintf("%s/latest/%s/%s", ENDPOINT, resourceType, resource)
}

func (t *TurretIO) makeSignature(url string, payload *map[string]interface {}, timestamp int64) (string, error) {
    // cut ENDPOINT from url
    u := strings.Replace(url, ENDPOINT, "", -1)
    j, err := json.Marshal(payload)
    if err != nil {
        return "", err
    }

    stringToSign := fmt.Sprintf("%s%s%s", u, string(j), strconv.FormatInt(timestamp, 10))
    k, err := base64.StdEncoding.DecodeString(t.Apisecret)
    if err != nil {
        return "", err
    }
    h := hmac.New(sha512.New, []byte(k))
    h.Write([]byte(stringToSign))
    return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func (t *TurretIO) request(url string, method string, payload *map[string]interface {}, client *http.Client) (*TurretIOResponse, error) {
    // make timestamp
    timestamp := int64(time.Now().Unix())
    // sign request
	full_url := fmt.Sprintf("%s%s", ENDPOINT, url)
    sig, err := t.makeSignature(full_url, payload, timestamp)
    if err != nil {
        return nil, err
    }

    j, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    var b bytes.Buffer

    b.Write([]byte(base64.StdEncoding.EncodeToString(j)))

    req, err := http.NewRequest(method, full_url, &b)
    if err != nil {
        return nil, err
    }

    req.Header.Set("X-Ls-Auth", sig)
    req.Header.Set("X-Ls-Time", strconv.FormatInt(timestamp, 10))
    req.Header.Set("X-Ls-Key", t.Apikey)

    response, err := client.Do(req)

    if err != nil {
        return nil, err
    }

    //log.Print(response)

    defer response.Body.Close()

    r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var jresponse map[string]interface{}
	err = json.Unmarshal(r, &jresponse)
	if err != nil {
		if response.Status == "401 Unauthorized" {
			// Here, err is actually "Invalid character 'U'" because of
			// "Unauthorized" being returned in the body, so let's not send
			// that as an actual error
			return &TurretIOResponse{nil, response.Status}, nil
		}

		return nil, err
	}
    return &TurretIOResponse{jresponse, response.Status}, err
}

func (t *TurretIO) GetRequest(url string, payload *map[string]interface {}, client *http.Client) (*TurretIOResponse, error) {
	resp, err := t.request(url, "GET", payload, client)
	return resp, err
}

func (t *TurretIO) PostRequest(url string, payload *map[string]interface {}, client *http.Client) (*TurretIOResponse, error) {
	resp, err := t.request(url, "POST", payload, client)
	return resp, err
}
