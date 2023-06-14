package groupDataPool

import (
	"errors"
	"insightstream/driver/harmony"
	"net/http"
)

func UpdateGroupDataPool() error {
	const (
		host   = "host.docker.internal:5100"
		path   = "/api/v1/group_and_store_by_url"
		method = http.MethodPost
	)

	cl, err := harmony.Driver(host, path, method, nil)
	if err != nil {
		return errors.New("initializing harmony driver failed")
	}

	if cl.StatusCode != 200 {
		return errors.New("group and store by url failed")
	}

	return nil
}
