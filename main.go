package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os/exec"

	"go.uber.org/zap"

	config_loader "github.com/chiguhagu/amenokaguyama-mysql/internal"
)

var config *config_loader.Config
var logger *zap.Logger

type payload struct {
	Text string `json:"text"`
}

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic("zap akan")
	}
	config, err = config_loader.LoadConfig()
	if err != nil {
		panic("config_loader akan")
	}
}

func main() {
	cmd := exec.Command(
		"mysqldef",
		"--dry-run",
		"-h", config.MysqlHost,
		"-P", config.MysqlPort,
		"-u", config.MysqlUser,
		"-p", config.MysqlPassword,
		"--file=./schema/init.sql",
		config.MysqlDbName)
	b, err := cmd.Output()
	if err != nil {
		logger.Fatal("failed to execute mysqldef command: " + string(err.(*exec.ExitError).Stderr))
		return
	}
	p, err := json.Marshal(payload{Text: "```" + string(b) + "```"})
	if err != nil {
		logger.Fatal("failed to marshal slack payload: " + err.Error())
	}
	resp, err := http.PostForm(config.SlackWebhookUrl, url.Values{"payload": {string(p)}})
	if err != nil {
		logger.Fatal("failed to post message to slack: " + err.Error())
		return
	}
	defer resp.Body.Close()
}
