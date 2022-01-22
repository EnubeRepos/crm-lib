package context

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config contains the configurations for application.
type Config struct {
	// Application version.
	Version string

	HOST_CRM string

	// HOST CRM
	TOKEN_CRM string
}

// NewConfig load the configuration for application.
func NewConfig(host string) (*Context, error) {
	config := &Config{
		Version:  viper.GetString("VERSION"),
		HOST_CRM: host,
	}

	c := &Context{
		Config: *config,
		Logger: NewLogger(),
	}

	return c, nil
}

// Context represent a context for application.
type Context struct {
	Config Config
	Logger logrus.FieldLogger
}

var (
	logrusinstanceonce sync.Once
	logrusinstance     *logrus.Logger
)

// GlobalLogger return a global logger (for logs outside some context).
func GlobalLogger() *logrus.Logger {
	logrusinstanceonce.Do(func() {
		logrusinstance = NewLogger()
	})
	return logrusinstance
}

// NewLogger create a new instace of logrus logger.
func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
	})
	return l
}
