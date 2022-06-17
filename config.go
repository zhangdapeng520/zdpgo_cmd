package zdpgo_cmd

type Config struct {
	PidName     string `json:"pid_name" yaml:"pid_name" env:"pid_name"`
	EnvFileName string `json:"env_file_name" yaml:"env_file_name" env:"env_file_name"`
}
