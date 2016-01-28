package config

type Config struct {
	Executor struct {
		URL string `json:"url"`
	} `json:"executor,omitempty"`
	HTTP struct {
		BindAddress string `json:"bind_address"`
		BindPort    int    `json:"bind_port"`
		Certificate string `json:"certificate,omitempty"`
		Key         string `json:"key,omitempty"`
		Scheme      string `json:"scheme,omitempty"`
	} `json:"http"`
	Log Log `json:"log"`
}

type Log struct {
	Hook []struct {
		URL      string `json:"url,omitempty"`
		Level    string `json:"level"`
		Protocol string `json:"protocol"`
		Type     string `json:"type"`
	} `json:"hook,omitempty"`
	Output struct {
		Level string `json:"level"`
		Path  string `json:"path"`
	} `json:"output,omitempty"`
}
