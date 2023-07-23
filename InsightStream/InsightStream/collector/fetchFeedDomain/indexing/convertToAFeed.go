package indexing

import (
	"errors"
	"insightstream/driver/harmony"
	"net/http"
)

func PingToSync() error {
	driver, err := harmony.Driver("host.docker.internal:5100", "/api/v1/parse_and_store_feeds", http.MethodPost, nil)
	if err != nil {
		return errors.New("baseFeeds sync failed")
	}

	if driver.StatusCode != 200 {
		return errors.New("baseFeeds sync failed")
	}

	return nil
}

func PingToSyncOnlyLatestFeeds() error {
	driver, err := harmony.Driver("host.docker.internal:5100", "/api/v1/parse_and_store_latest_feeds", http.MethodPost, nil)
	if err != nil {
		return errors.New("baseFeeds sync failed")
	}

	if driver.StatusCode != 200 {
		return errors.New("baseFeeds sync failed")
	}

	return nil
}
