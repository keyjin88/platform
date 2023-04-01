package services

import (
	"context"
	"platform/config"
	"platform/logging"
	"reflect"
)

const ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

func newServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(serviceMap))
	} else {
		return c
	}
}

func ConfigurationFactory() config.Configuration {
	//TODO create struct that implements Configuration interface
	return nil
}

func LoggerFactory(cfg config.Configuration) logging.Logger {
	//TODO create struct that implements Logger interface
	return nil
}
