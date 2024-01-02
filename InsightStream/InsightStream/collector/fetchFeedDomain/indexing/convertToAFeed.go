package indexing

import (
	"fmt"
	"insightstream/driver/harmony"
	"net/http"
)

func PingToSync() error {
	driver, err := harmony.Driver("host.docker.internal:5100", "/api/v1/parse_and_store_feeds", http.MethodPost, nil)
	if err != nil {
		return fmt.Errorf("baseFeeds sync failed: %w", err)
	}

	if driver.StatusCode != 200 {
		return fmt.Errorf("baseFeeds sync failed: %w", err)
	}

	return nil
}

func PingToSyncOnlyLatestFeeds() error {
	driver, err := harmony.Driver("host.docker.internal:5100", "/api/v1/parse_and_store_latest_feeds", http.MethodPost, nil)
	if err != nil {
		return fmt.Errorf("baseFeeds sync failed: %w", err)
	}

	if driver.StatusCode != 200 {
		return fmt.Errorf("baseFeeds sync failed: %w", err)
	}

	return nil
}
