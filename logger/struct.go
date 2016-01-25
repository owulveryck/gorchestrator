package logger

type Log struct {
	Hook []struct {
		URL      string `json:"address,omitempty"`
		Level    string `json:"level"`
		Protocol string `json:"protocol"`
		Type     string `json:"type"`
	} `json:"hook,omitempty"`
	Output struct {
		Level string `json:"level"`
		Path  string `json:"path"`
	} `json:"output,omitempty"`
}
