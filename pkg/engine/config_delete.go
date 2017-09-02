package engine

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-fabric/pkg/tokens"
)

func (eng *Engine) DeleteConfig(envName string, key string) error {
	info, err := eng.initEnvCmdName(tokens.QName(envName), "")
	if err != nil {
		return err
	}
	config := info.Target.Config

	if config != nil {
		delete(config, tokens.Token(key))

		if err = eng.Environment.SaveEnvironment(info.Target, info.Snapshot); err != nil {
			return errors.Wrap(err, "could not save configuration value")
		}
	}

	return nil
}