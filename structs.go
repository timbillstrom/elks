package elks

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// Elks46 holds the client for working with the 46Elks API
type Elks46 struct {
	Dry      bool
	Client   *http.Client
	username string
	password string
}

// NewClient returns an instance of Elks46
func NewClient(u, p string, dryRun bool) *Elks46 {
	return &Elks46{
		Client:   &http.Client{},
		Dry:      dryRun,
		username: u,
		password: p,
	}
}

// SendMessage sends a text message
func (e *Elks46) SendMessage(sms *SMS) (map[string]interface{}, error) {
	return call(e, request(http.MethodPost, sendURL, e, convertStruct(sms, e.Dry)))
}

// SMS structs the model for outgoing TextMessages
type SMS struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
	DryRun  string `json:"dryrun"`
}

func convertStruct(in interface{}, testing bool) url.Values {
	var (
		vals = url.Values{}
		data map[string]string
	)
	data, err := toMap(in)
	if err != nil {
		panic(err)
	}
	for k, v := range data {
		if v != "" {
			vals.Add(k, v)
		}
	}
	if testing {
		vals.Add("dryrun", "yes")
	}
	return vals
}

// toMap converts a struct to a map using the struct's tags.
func toMap(in interface{}) (map[string]string, error) {
	out := make(map[string]string)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		out[strings.ToLower(fi.Name)] = v.Field(i).String()
	}
	return out, nil
}
