package config_loader

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MysqlHost       string `envconfig:"MYSQL_HOST" required:"true"`
	MysqlPort       string `envconfig:"MYSQL_PORT" default:"3306"`
	MysqlUser       string `envconfig:"MYSQL_USER" required:"true"`
	MysqlPassword   string `envconfig:"MYSQL_PASSWORD" required:"true"`
	MysqlDbName     string `envconfig:"MYSQL_DB_NAME" required:"true"`
	SlackWebhookUrl string `envconfig:"SLACK_WEBHOOK_URL" required:"false"`
}

func LoadConfig() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}
