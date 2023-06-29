package lg

type Level string

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = "DEBUG"
	// InfoLevel is the default logging priority.
	InfoLevel Level = "INFO"
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel Level = "WARN"
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel Level = "ERROR"
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel Level = "DPANIC"
	// PanicLevel logs a message, then panics.
	PanicLevel Level = "PANIC"
)
