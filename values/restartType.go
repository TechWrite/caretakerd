package values

import (
    "strings"
    "github.com/echocat/caretakerd/errors"
    "encoding/json"
)

type RestartType struct {
    onSuccess  bool
    onFailures bool
}

var Never RestartType = RestartType{
    onSuccess: false,
    onFailures: false,
}
var OnFailures RestartType = RestartType{
    onSuccess: false,
    onFailures: true,
}
var Always RestartType = RestartType{
    onSuccess: true,
    onFailures: true,
}
var AllRestartTypes []RestartType = []RestartType{
    Never,
    OnFailures,
    Always,
}

func (i RestartType) String() string {
    result, err := i.CheckedString()
    if err != nil {
        panic(err)
    }
    return result
}

func (i RestartType) CheckedString() (string, error) {
    switch i {
    case Never:
        return "never", nil
    case OnFailures:
        return "onFailures", nil
    case Always:
        return "always", nil
    }
    return "", errors.New("Illegal restart type: %v", i)
}

func (i *RestartType) Set(value string) error {
    lowerValue := strings.ToLower(value)
    for _, candidate := range AllRestartTypes {
        if strings.ToLower(candidate.String()) == lowerValue {
            (*i) = candidate
            return nil
        }
    }
    return errors.New("Illegal restart type: %s", value)
}

func (i RestartType) MarshalYAML() (interface{}, error) {
    return i.CheckedString()
}

func (i *RestartType) UnmarshalYAML(unmarshal func(interface{}) error) error {
    var value string
    if err := unmarshal(&value); err != nil {
        return err
    }
    return i.Set(value)
}

func (i RestartType) MarshalJSON() ([]byte, error) {
    s, err := i.CheckedString()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(s)
}

func (i *RestartType) UnmarshalJSON(b []byte) error {
    var value string
    if err := json.Unmarshal(b, &value); err != nil {
        return err
    }
    return i.Set(value)
}

func (i RestartType) OnSuccess() bool {
    return i.onSuccess
}

func (i RestartType) OnFailures() bool {
    return i.onFailures
}

func (i RestartType) Validate() error {
    _, err := i.CheckedString()
    return err
}