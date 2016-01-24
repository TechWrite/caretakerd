package main

import (
    . "github.com/echocat/caretakerd/values"
    "github.com/echocat/caretakerd/config"
    "github.com/echocat/caretakerd/errors"
    "github.com/echocat/caretakerd/defaults"
)

type ConfigWrapper struct {
    instance    *config.Config
    filename    String
    explicitSet bool
}

func NewConfigWrapper() *ConfigWrapper {
    instance := config.NewConfig()
    return &ConfigWrapper{
        instance: &instance,
        filename: defaults.ConfigFilename(),
        explicitSet: false,
    }
}

func (this ConfigWrapper) String() string {
    return this.filename.String()
}

func (this *ConfigWrapper) Set(value string) error {
    if len(value) == 0 {
        return errors.New("There is an empty filename for configuration provided.")
    }
    filename := String(value)
    conf, err := config.LoadFromYamlFile(filename)
    if err != nil {
        return err
    }
    conf = conf.EnrichFromEnvironment()
    this.filename = filename
    this.instance = &conf
    this.explicitSet = true
    return nil
}

func (this ConfigWrapper) IsExplicitSet() bool {
    return this.explicitSet
}

func (this *ConfigWrapper) ConfigureAndValidate(listenAddress *FlagWrapper, pemFile *FlagWrapper, validateAlsoMaster bool) error {
    listenAddress.AssignIfExplicitSet(&this.instance.Rpc.Listen)
    pemFile.AssignIfExplicitSet(&this.instance.Control.Access.PemFile)
    err := this.instance.Validate()
    if err != nil {
        return err
    }
    if validateAlsoMaster {
        return this.instance.ValidateMaster()
    }
    return nil
}

func (this ConfigWrapper) Instance() *config.Config {
    return this.instance
}
