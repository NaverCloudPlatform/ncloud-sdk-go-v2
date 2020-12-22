package credentials

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NaverCloudPlatform/ncloud-sdk-go-v2/ncloud/metadata"
	"strings"
	"time"
)

type ServerRoleProvider struct {
	ApiClient *metadata.ApiClient
}

func (p *ServerRoleProvider) Name() string {
	return "ServerRoleProvider"
}

func (p *ServerRoleProvider) Retrieve() (Value, error) {
	credentials, err := reqCredentialsList(p.ApiClient)
	if err != nil {
		return Value{}, err
	}

	if len(credentials) == 0 {
		return Value{}, errors.New("empty role list")
	}
	roleId := credentials[0]

	roleCreds, err := reqCredentials(p.ApiClient, roleId)
	if err != nil {
		return Value{}, err
	}

	return Value{
		AccessKey:  roleCreds.AccessKeyID,
		SecretKey:  roleCreds.SecretAccessKey,
		Expiration: roleCreds.Expiration,
	}, nil
}

func reqCredentialsList(client *metadata.ApiClient) ([]string, error) {
	resp, err := client.GetMetadata(iamSecurityCredsPath)
	if err != nil {
		return nil, errors.New("no server role found")
	}
	var credentials []string
	s := bufio.NewScanner(strings.NewReader(resp))
	for s.Scan() {
		credentials = append(credentials, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, errors.New("failed to read server role from metadata api")
	}
	return credentials, nil
}

type roleCredRespBody struct {
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	Expiration      time.Time `json:"Expiration"`
	Code            string    `json:"Code"`
	Message         string    `json:"Message"`
}

const iamSecurityCredsPath = "iam/security-credentials/"

func reqCredentials(client *metadata.ApiClient, roleId string) (roleCredRespBody, error) {
	resp, err := client.GetMetadata(iamSecurityCredsPath + roleId)
	if err != nil {
		return roleCredRespBody{},
			errors.New(fmt.Sprintf("failed to get %s server role credentials", roleId))
	}
	respCreds := roleCredRespBody{}
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&respCreds); err != nil {
		return roleCredRespBody{},
			errors.New(fmt.Sprintf("failed to decode %s server role credentials", roleId))
	}
	return respCreds, nil
}
