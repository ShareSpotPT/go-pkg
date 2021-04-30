package env

const (
	Main    = "main"
	Staging = "staging"
	CICD    = "CICD"
	Dev     = "dev"
)

type Mode string
func (m Mode) ConfigFile() string {
	return "mode/" + string(m) + ".yaml"
}

type ErrModeNotFound string
func (e ErrModeNotFound) Error() string {
	return "mode not found: " + string(e)
}

type ErrSecuredEnvironment string
func (e ErrSecuredEnvironment) Error() string {
	return "cannot drop secured environment: " + string(e)
}

func GetMode(s string) (Mode, error) {
	switch s {
	case Main:
		return Main, nil
	case Staging:
		return Staging, nil
	case CICD:
		return CICD, nil
	case Dev:
		return Dev, nil
	}

	return "",  ErrModeNotFound(s)
}

func (m Mode) String() string {
	return string(m)
}
func GetModeConfigFile(s string) (string, error) {
	m, err := GetMode(s)
	if err != nil {
		return "", err
	}
	return m.ConfigFile(), nil
}