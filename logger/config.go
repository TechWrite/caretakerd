package logger

import (
    . "github.com/echocat/caretakerd/values"
)

var defaults = map[string]interface{} {
    "Level": Info,
    "StdoutLevel": Info,
    "StderrLevel": Error,
    "Filename": String("console"),
    "MaxSizeInMb": NonNegativeInteger(500),
    "MaxBackups": NonNegativeInteger(5),
    "MaxAgeInDays": NonNegativeInteger(1),
    "Pattern": Pattern("%d{YYYY-MM-DD HH:mm:ss} [%-5.5p] [%c] %m%n%P{%m}"),
}

// @id Logger
// @type struct
//
// ## Description
//
// A logger handles every output generated by the daemon itself, the process or other parts controlled by the daemon.
type Config struct {
    // @id      level
    // @default info
    //
    // Minimal log level the logger will output its messages. All below will be ignored.
    Level        Level `json:"level" yaml:"level"`

    // @id      stdoutLevel
    // @default info
    //
    // If the service prints something to ``stdout`` this will logged with this level.
    StdoutLevel  Level `json:"stdoutLevel" yaml:"stdoutLevel"`

    // @id      stderrLevel
    // @default error
    //
    // If the service prints something to ``stderr`` this will logged with this level.
    StderrLevel  Level `json:"stderrLevel" yaml:"stderrLevel"`

    // @id      filename
    // @default "console"
    //
    // Target file of the logger. The file will be created if not exist - but not the parent directory.
    //
    // If this value is set to ``console`` the whole output will go to ``stdout`` or to ``stderr`` on every log level
    // above or equal to {@ref Level#warn}.
    Filename     String `json:"filename" yaml:"filename"`

    // @id      maxSizeInMb
    // @default 500
    //
    // Maximum size in megabytes of the log file before it gets rotated.
    //
    // This is ignored if {@ref #filename} os set to ``console``.
    MaxSizeInMb  NonNegativeInteger `json:"maxSizeInMb" yaml:"maxSizeInMb"`

    // @id      maxBackups
    // @default 500
    //
    // Maximum number of old log files to retain.
    //
    // This is ignored if {@ref #filename} os set to ``console``.
    MaxBackups   NonNegativeInteger `json:"maxBackups" yaml:"maxBackups"`

    // @id      maxAgeInDays
    // @default 1
    //
    // Maximum number of days to retain old log files based on the
    // timestamp encoded in their filename.  Note that a day is defined as 24
    // hours and may not exactly correspond to calendar days due to daylight
    // savings, leap seconds, etc.
    //
    // This is ignored if {@ref #filename} os set to ``console``.
    MaxAgeInDays NonNegativeInteger `json:"maxAgeInDays" yaml:"maxAgeInDays"`

    // @id      pattern
    // @default "%d{YYYY-MM-DD HH:mm:ss} [%-5.5p] [%c] %m%n%P{%m}"
    //
    // Pattern how to format the log messages to output with.
    Pattern      Pattern `json:"pattern" yaml:"pattern"`
}

func NewConfig() Config {
    result := Config{}
    result.init()
    return result
}

func (this Config) Validate() error {
    err := this.StdoutLevel.Validate()
    if err == nil {
        err = this.StderrLevel.Validate()
    }
    return err
}

func (this *Config) init() {
    SetDefaultsTo(defaults, this)
}

func (this *Config) BeforeUnmarshalYAML() error {
    this.init()
    return nil
}
