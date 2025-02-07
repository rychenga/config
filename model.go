package main

type GinConfig1 struct {
	Port        string `yaml:"port"`
	DebugMode   string `yaml:"debug_mode"`
	SwaggerMode string `yaml:"swagger_mode"`
}

type Config1 struct {
	Config struct {
		GinConfig GinConfig1 `yaml:"http_config"`
	} `yaml:"config"`
}

type GinConfig2 struct {
	Port        string `mapstructure:"port"`
	DebugMode   string `mapstructure:"debug_mode"`
	SwaggerMode string `mapstructure:"swagger_mode"`
}

type Config2 struct {
	GinConfig GinConfig2 `mapstructure:"http_config"`
}
