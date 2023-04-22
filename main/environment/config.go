package environment

type (
	Config struct {
		DB     DB     `cue:"db" json:"db,omitempty"`
		Sentry Sentry `cue:"sentry" json:"sentry,omitempty"`
	}

	// DB
	DB struct {
		DBConfigs []DBConfig `cue:"dbconfigs" json:"dbconfigs,omitempty"`
		DebugMode bool       `cue:"debug_mode" json:"debug_mode,omitempty"`
	}

	DBConfig struct {
		Name        string `cue:"name" json:"name,omitempty"`
		DBMS        string `cue:"dbms" json:"dbms,omitempty"`
		User        string `cue:"user" json:"user,omitempty"`
		Password    string `cue:"password" json:"password,omitempty"`
		Protocol    string `cue:"protocol" json:"protocol,omitempty"`
		Host        string `cue:"host" json:"host,omitempty"`
		Port        int    `cue:"port" json:"port,omitempty"`
		Schema      string `cue:"schema" json:"schema,omitempty"`
		Option      string `cue:"option" json:"option,omitempty"`
		MaxOpen     int    `cue:"max_open" json:"max_open,omitempty"`
		MaxIdle     int    `cue:"max_idle" json:"max_idle,omitempty"`
		LifetimeSec int    `cue:"conn_lifetime_sec" json:"conn_lifetime_sec,omitempty"`
	}

	// Sentry
	Sentry struct {
		SentryConfig SentryConfig `cue:"config" json:"config,omitempty"`
		DebugMode    bool         `cue:"debug_mode" json:"debug_mode,omitempty"`
	}

	SentryConfig struct {
		Dsn string `cue:"dsn" json:"dsn,omitempty"`
	}
)

var configData Config

func GetConfig() Config {
	return configData
}
