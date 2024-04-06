package config

type (

	Config struct {
	}

	Server struct {
		port int `mapstructure:"port" validate:"required"`
	}

	OAuth2 struct {

	}

	State struct {

	}

	Database struct {

	}

)