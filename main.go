package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Width    int `mapstructure:"width"`
	Height   int `mapstructure:"height"`
	MaxIter  int `mapstructure:"maxIter"`
	Contrast int `mapstructure:"contrast"`
}

var config Config

func init() {
	// Set the configuration file name (config.yaml in this case)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read the configuration from the YAML file using viper
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}

	// Unmarshal the configuration into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling configuration: %v", err)
	}
}

func mandelbrot(c complex64) int {
	z := complex64(0)
	for i := 0; i < config.MaxIter; i++ {
		z = z*z + c
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			return i
		}
	}
	return config.MaxIter
}

func drawMandelbrot() {
	for y := 0; y < config.Height; y++ {
		for x := 0; x < config.Width; x++ {
			realPart := float32(x)/float32(config.Width)*3.5 - 2.5
			imagPart := float32(y)/float32(config.Height)*2 - 1

			c := complex(realPart, imagPart)
			value := mandelbrot(c)
			if value == config.MaxIter {
				fmt.Print(" ")
			} else {
				fmt.Print(string(48 + value%config.Contrast))
			}
		}
		fmt.Println()
	}
}

func main() {
	drawMandelbrot()
}
