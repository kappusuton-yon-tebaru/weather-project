package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port                       int    `mapstructure:"PORT"`
	Dev                        bool   `mapstructure:"DEV"`
	GeolocationServiceEndpoint string `mapstructure:"GEOLOCATION_SERVICE_ENDPOINT"`
	WeatherServiceEndpoint  string `mapstructure:"WEATHER_SERVICE_ENDPOINT"`
}

func Load() (*Config, error) {
	cfg := &Config{}

	err := bindStruct(Config{})
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func bindStruct(s interface{}) error {
	ct := reflect.TypeOf(s)

	if ct.Kind() != reflect.Struct {
		return fmt.Errorf("listStructKeys: %v is not a struct", ct)
	}

	for i := range ct.NumField() {
		field := ct.Field(i)
		tag := field.Tag.Get("mapstructure")

		if field.Type.Kind() == reflect.Struct {
			err := bindStruct(reflect.New(field.Type).Elem().Interface())
			if err != nil {
				return err
			}
		} else {
			viper.BindEnv(strings.Split(tag, ",")[0])
		}
	}

	return nil
}
