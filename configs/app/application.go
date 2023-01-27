package app

type Configs struct {
	Application ApplicationConfigs `mapstructure:"application"`
}

type ApplicationConfigs struct {
	Clients   ClientsConfig  `mapstructure:"clients"`
	NewRelic  NewRelicConfig `mapstructure:"newrelic"`
	Postgres  PGSettings     `mapstructure:"postgres"`
	JWTConfig JWTConfig      `mapstructure:"jwt-config"`
}

type ClientsConfig struct {
	Workspace ClientConfig `mapstructure:"workspace"`
}

type ClientConfig struct {
	BaseURL string `mapstructure:"base-url"`
	Timeout int    `mapstructure:"timeout"`
}

type NewRelicConfig struct {
	ApplicationKey  string `mapstructure:"application-key"`
	ApplicationName string `mapstructure:"application-name"`
}

type AuthConfig struct {
	Google GeneralOAuthSettings `mapstructure:"google"`
	Github GeneralOAuthSettings `mapstructure:"github"`
}

type GeneralOAuthSettings struct {
	Scopes       []string `mapstructure:"scopes"`
	RedirectURL  string   `mapstructure:"redirect-url"`
	ClientID     string   `mapstructure:"client-id"`
	ClientSecret string   `mapstructure:"client-secret"`
}

type PGSettings struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Debug    bool   `mapstructure:"debug"`
}

type JWTConfig struct {
	TTL     int    `mapstructure:"ttl"`
	SignKey string `mapstructure:"sk"`
}
