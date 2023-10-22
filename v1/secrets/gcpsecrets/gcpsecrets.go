// Package gcpsecrets implements unwrapping of secrets via GCP Secret Manager.
package gcpsecrets

import (
	"context"
	"fmt"
	"path"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

const latestVersion = "latest"

type Provider struct {
	client    *secretmanager.Client
	projectId string
}

func New(client *secretmanager.Client, projectId string) *Provider {
	return &Provider{client, projectId}
}

func (p *Provider) Unwrap(name, value string) (string, error) {
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName(p.projectId, value, latestVersion),
	}

	sec, err := p.client.AccessSecretVersion(cxt, req)
	if err != nil {
		return "", fmt.Errorf("Could not access secret for %s=%s: %w", name, value, err)
	} else if sec.Payload == nil {
		return "", fmt.Errorf("Secret payload is empty for %s=%s", name, value)
	}

	return string(sec.Payload.Data), nil
}

func secretName(projectId, secretId, versionId string) string {
	return path.Join("projects", projectId, "secrets", secretId, "versions", versionId)
}
