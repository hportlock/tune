package tune

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
)

type envVarLoader struct {
}

func (e envVarLoader) Getenv(key string) string {
	return os.Getenv(key)
}

// LoadConfig reads toml files from a specified directory and parses them into
// a config struct.
//	configDir: The directory containing the toml config files.
//	env: The environment, this should be the prefix to the toml file. For
//		example if env is "development" the config files should be named
//		"development.toml"
//	config: The struct to write the config values to. It's structure should
//		match the config file.
func LoadConfig(configDir string, env string, config interface{}) error {
	settingsFile := "settings.toml"
	envFile := fmt.Sprintf("%s.toml", strings.ToLower(env))
	localFile := fmt.Sprintf("%s.local.toml", strings.ToLower(env))

	err := readSettings(configDir, settingsFile, config)
	if err != nil {
		return fmt.Errorf("Error reading config file: %s", err)
	}

	err = readSettings(configDir, envFile, config)
	if err != nil {
		return fmt.Errorf("Error reading config file: %s", err)
	}

	err = readSettings(configDir, localFile, config)
	if err != nil {
		return fmt.Errorf("Error reading config file: %s", err)
	}

	return nil
}

func readSettings(configDir string, configFile string, config interface{}) error {
	configBytes, err := ioutil.ReadFile(path.Join(configDir, configFile))
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	configTemplate, err := template.New("config").Parse(string(configBytes))
	if err != nil {
		return fmt.Errorf("Can't parse template: %s", err)
	}

	configContent := new(bytes.Buffer)
	configTemplate.Execute(configContent, envVarLoader{})
	if err != nil {
		return fmt.Errorf("Can't process template: %s", err)
	}

	_, err = toml.DecodeReader(configContent, config)
	if err != nil {
		return fmt.Errorf("Can't parse toml: %s", err)
	}

	return nil
}

func parseConfig(configBytes []byte, config interface{}) error {
	configTemplate, err := template.New("config").Parse(string(configBytes))
	if err != nil {
		return fmt.Errorf("Can't parse template: %s", err)
	}

	configContent := new(bytes.Buffer)
	configTemplate.Execute(configContent, envVarLoader{})
	if err != nil {
		return fmt.Errorf("Can't process template: %s", err)
	}

	_, err = toml.DecodeReader(configContent, config)
	if err != nil {
		return fmt.Errorf("Can't parse toml: %s", err)
	}

	return nil
}
