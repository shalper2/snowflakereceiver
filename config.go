package snowflakereceiver

import (
    "fmt"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type Config struct {
    scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
    confighttp.HTTPClientSettings           `mapstructure:",squash"`
    // Metrics metadata.MetricsSettings        `mapstructure:"metrics"`
    Username  string                        `mapstructure:"username"`
    Password  string                        `mapstructure:"password"`
    Account   string                        `mapstructure:"account"`
    Schema    string                        `mapstructure:"schema"`
    Warehouse string                        `mapstructure:"warehouse"`
    Database  string                        `mapstructure:"database"`
    Role      string                        `mapstructure:"role"`
}

func (cfg *Config) Validate() error {
    if (cfg.Username == "") {
        return fmt.Errorf("You must provide a valid snowflake username")
    }

    if (cfg.Password == "") {
        return fmt.Errorf("You must provide a password for the snowflake username")
    }

    if (cfg.Account == "") {
        return fmt.Errorf("You must provide a valid account name")
    }

    return nil
}
