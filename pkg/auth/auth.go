package auth

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

func GetToken(username, password string) (string, error) {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://keystone.cloud.syseleven.net:5000/v3",
		Username:         username,
		Password:         password,
		DomainName:       "Default",
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return "", err
	}

	return provider.Token(), nil
}
