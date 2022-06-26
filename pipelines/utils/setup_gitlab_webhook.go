package utils

import (
	"github.com/go-playground/webhooks/v6/gitlab"
	"os"
)

var gitlabHook *gitlab.Webhook

func SetupGitlabWebhook() {
	gitlabHook, _ = gitlab.New(gitlab.Options.Secret(getGitlabWebhookSecret()))
}

func GetGitlabHook() *gitlab.Webhook{
	return gitlabHook
}

func getGitlabWebhookSecret() string {
	gitlabWebhookSecret := os.Getenv("GITLAB_WEBHOOK_SECRET")
	if gitlabWebhookSecret == "" {
		gitlabWebhookSecret = "gitlabSecret"
	}
	return gitlabWebhookSecret
}