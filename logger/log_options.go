package logger

type LogOptions struct {
	Level         string `json:"level,omitempty" yaml:"level,omitempty" mapstructure:"level"`
	Format        string `json:"format,omitempty" yaml:"format,omitempty" mapstructure:"format"`
	DisableQuote  bool   `json:"disable_quote,omitempty" yaml:"disable_quote,omitempty" mapstructure:"disable_quote"`
	RotationSize  int    `json:"rotation_size,omitempty" yaml:"rotation_size,omitempty" mapstructure:"rotation_size"`
	RotationCount int    `json:"rotation_count,omitempty" yaml:"rotation_count,omitempty" mapstructure:"rotation_count"`
	MaxAge        int    `json:"max_age,omitempty" yaml:"max_age,omitempty" mapstructure:"max_age"`
	Compress      bool   `json:"compress,omitempty" yaml:"compress,omitempty" mapstructure:"compress"`
	Path          string `json:"path,omitempty" yaml:"path,omitempty" mapstructure:"path"`
}

func NewLogOptions() *LogOptions {
	return &LogOptions{
		Level:         "debug",
		Format:        "text",
		DisableQuote:  false,
		RotationSize:  4,
		RotationCount: 10,
		MaxAge:        7,
		Compress:      false,
	}
}
