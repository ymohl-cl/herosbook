package hasher

const (
	defaultScryptN   = 32768
	defaultScryptR   = 8
	defaultScryptP   = 1
	defaultScryptLen = 64
)

// Config to hash the byte key
type Config struct {
	ScryptN   int
	ScryptP   int
	ScryptR   int
	ScryptLen int
}

// DefaultConf statement
func DefaultConf() Config {
	return Config{
		ScryptN:   defaultScryptN,
		ScryptP:   defaultScryptP,
		ScryptR:   defaultScryptR,
		ScryptLen: defaultScryptLen,
	}
}
