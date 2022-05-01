package zdpgo_cmd

type Config struct {
	Debug       bool   `json:"debug" yaml:"debug" env:"debug"`
	LogFileName string `json:"log_file_name" yaml:"log_file_name" env:"log_file_name"`
	PidName     string `json:"pid_name" yaml:"pid_name" env:"pid_name"`
	EnvFileName string `json:"env_file_name" yaml:"env_file_name" env:"env_file_name"`
}
