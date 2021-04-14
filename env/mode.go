package env

const (
	Main Mode = iota + 1
	Staging
	CICD
	Dev
)

const (
	main    = "main"
	staging = "staging"
	cicd    = "cicd"
	dev     = "dev"
)

type Mode uint

func (m Mode) String() string {
	return []string{"", main, staging, cicd, dev}[m]
}

func (m Mode) ConfigFile() string {
	return "config/" + m.String() + ".yaml"
}

func GetMode(s string) (Mode, error) {
	m := map[string]Mode{main: Main, staging: Staging, cicd: CICD, dev: Dev}
	res, ok := m[s]
	if !ok {
		return 0, ErrBadMode(s + " mode doesn't exist")
	}
	return res, nil
}

func GetModeConfigFile(s string) (string, error) {
	m, err := GetMode(s)
	if err != nil {
		return "", err
	}
	return m.ConfigFile(), nil
}

