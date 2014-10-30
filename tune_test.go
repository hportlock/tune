package tune

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testConfig struct {
	DatabaseURL string
}

func TestLoadConfig(t *testing.T) {
	configRoot := "fixtures"

	Convey("LoadConfig", t, func() {
		Convey("Correctly loads the config", func() {
			config := testConfig{}
			err := LoadConfig(configRoot, "test", &config)
			if err != nil {
				t.Fatal(err)
			}
			So(config.DatabaseURL, ShouldEqual, "foo")
		})

		Convey("Correctly overrides with a local config", func() {
			config := testConfig{}
			err := LoadConfig(configRoot, "test2", &config)
			if err != nil {
				t.Fatal(err)
			}
			So(config.DatabaseURL, ShouldEqual, "bar")
		})

		Convey("Parses in environment variables", func() {
			foo := os.Getenv("FOO")
			defer os.Setenv("FOO", foo)
			os.Setenv("FOO", "baz")
			config := testConfig{}
			err := LoadConfig(configRoot, "test3", &config)
			if err != nil {
				t.Fatal(err)
			}
			So(config.DatabaseURL, ShouldEqual, "baz")
		})

		Convey("Parses in environment variables in the local config", func() {
			bar := os.Getenv("BAR")
			defer os.Setenv("BAR", bar)
			os.Setenv("BAR", "snat")
			config := testConfig{}
			err := LoadConfig(configRoot, "test4", &config)
			if err != nil {
				t.Fatal(err)
			}
			So(config.DatabaseURL, ShouldEqual, "snat")
		})
	})
}
