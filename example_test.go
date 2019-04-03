package config_test

import (
	"fmt"
	"os"

	"github.com/JeremyLoy/config"
)

type MySubConfig struct {
	IPWhitelist []string
}

type MyConfig struct {
	DatabaseURL string
	Port        int
	FeatureFlag bool
	SubConfig   MySubConfig
	// etc.
}

func init() {
	os.Setenv("DATABASEURL", "db://")
	os.Setenv("PORT", "1234")
	os.Setenv("FEATUREFLAG", "true")
	os.Setenv("SUBCONFIG__IPWHITELIST", "0.0.0.0,1.1.1.1,2.2.2.2")
}

func Example() {

	var c MyConfig
	config.FromEnv().To(&c)

	fmt.Println(c.DatabaseURL)
	fmt.Println(c.Port)
	fmt.Println(c.FeatureFlag)
	fmt.Println(c.SubConfig.IPWhitelist)

	// Output:
	// db://
	// 1234
	// true
	// [0.0.0.0,1.1.1.1,2.2.2.2]
}
