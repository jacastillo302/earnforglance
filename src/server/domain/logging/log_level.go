package domain

// LogLevel represents a log level
type LogLevel int

const (
	// Debug represents debug level
	Debug LogLevel = 10

	// Information represents information level
	Information LogLevel = 20

	// Warning represents warning level
	Warning LogLevel = 30

	// Error represents error level
	Error LogLevel = 40

	// Fatal represents fatal level
	Fatal LogLevel = 50
)
