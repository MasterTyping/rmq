package config

type Config struct {
	Version  string   `yaml:"version"`
	Services Services `yaml:"services"`
}

type Services struct {
	RabbitMQ RabbitMQ `yaml:"rabbitmq"`
}

type RabbitMQ struct {
	Image       string   `yaml:"image"`
	Ports       []string `yaml:"ports"`
	Volumes     []string `yaml:"volumes"`
	Environment []string `yaml:"environment"`
}

func NewConfig() *Config {
	return &Config{
		Version: "3",
		Services: Services{
			RabbitMQ: RabbitMQ{
				Image: "rabbitmq:3.8.17-management",
				Ports: []string{"5672:5672", "15672:15672"},
				Volumes: []string{
					"./rabbitmq:/var/lib/rabbitmq",
				},
				Environment: []string{
					"RABBITMQ_DEFAULT_USER=mastertyping",
					"RABBITMQ_DEFAULT_PASS=mastertyping",
				},
			},
		},
	}
}
