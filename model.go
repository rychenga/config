package main

type GinConfig struct {
	Port        string `yaml:"port"`
	DebugMode   string `yaml:"debug_mode"`
	SwaggerMode string `yaml:"swagger_mode"`
}

type Config struct {
	GinConfig GinConfig `yaml:"http_config"`
}
