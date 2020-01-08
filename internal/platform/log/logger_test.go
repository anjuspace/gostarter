package log

import "testing"

import "github.com/stretchr/testify/require"

func TestLogger(t *testing.T) {
	config := Config{
		EnableConsole: true,
		Level:         Debug,
		Format:        "json",
		EnableFile:    true,
		NoColor:       false,
		FileLocation:  "log.log",
	}
	zapLogger, _ := NewLogger(config, InstanceZapLogger)
	require.NotNil(t, zapLogger)

	contextLogger := zapLogger.WithFields(Fields{"key1": "value1"})
	contextLogger.Debugf("Starting with zap")
	contextLogger.Infof("Zap is awesome")

	rusLogger, _ := NewLogger(config, InstanceLogrusLogger)
	contextLogger = rusLogger.WithFields(Fields{"key1": "value1"})
	contextLogger.Debugf("Starting with logrus")

	contextLogger.Infof("Logrus is awesome")
}
