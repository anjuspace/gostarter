package log

// Config holds details necessary for logging.
type Config struct {
	// Format specifies the output log format.
	// Accepted values are: json, logfmt
	Format string

	// Level is the minimum log level that should appear on the output.
	Level string

	// NoColor makes sure that no log output gets colorized.
	NoColor bool

	// EnableFile saves logs to a file
	EnableFile bool

	// EnableConsole  actives console to print out logs
	EnableConsole bool

	// FileLocation is the path of log file
	FileLocation string
}
