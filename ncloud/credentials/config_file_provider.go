package credentials

import (
	"bufio"
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// file example: $HOME/.ncloud/configure
type ConfigFileProvider struct {
	// Path to the ncloud configure file.
	// The default is the current user's home directory.
	// 	Linux/OSX: "$HOME/.ncloud/configure"
	// 	Windows:   "%USERPROFILE%\.ncloud\configure"
	Filename string

	// retrieved states if the credentials have been successfully retrieved.
	retrieved bool
}

func (p *ConfigFileProvider) Name() string {
	return "ConfigFileProvider"
}

func (p *ConfigFileProvider) Retrieve() (Value, error) {
	p.retrieved = false
	credentials, err := p.loadCredentials()
	if err != nil {
		return Value{}, err
	}
	p.retrieved = true
	return credentials, nil
}

func (p *ConfigFileProvider) loadCredentials() (Value, error) {
	usr, err := user.Current()
	if err != nil {
		return Value{}, err
	}

	if usr.HomeDir == "" {
		return Value{}, errors.New("user home directory not found")
	}
	configureFile, err := p.filename()
	if err != nil {
		return Value{}, err
	}

	file, err := os.Open(configureFile)
	if err != nil {
		return Value{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	credValue := Value{}
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "=")
		switch strings.TrimSpace(s[0]) {
		case "ncloud_access_key_id":
			credValue.AccessKey = strings.TrimSpace(s[1])
		case "ncloud_secret_access_key":
			credValue.SecretKey = strings.TrimSpace(s[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return Value{}, err
	}

	return credValue, nil
}

func (p *ConfigFileProvider) filename() (string, error) {
	if len(p.Filename) != 0 {
		return p.Filename, nil
	}

	if home := UserHomeDir(); len(home) == 0 {
		return "", errors.New("user home directory not found")
	}
	p.Filename = filepath.Join(UserHomeDir(), ".ncloud", "configure")
	return p.Filename, nil
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" { // Windows
		return os.Getenv("USERPROFILE")
	}

	// *nix
	return os.Getenv("HOME")
}
