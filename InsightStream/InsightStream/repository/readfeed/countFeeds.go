package readfeed

import (
	"database/sql"
	"insightstream/repository"
)

func CountFeeds() (int, error) {
	result, err := repository.CoreDatabase.Query("SELECT COUNT(*) FROM feeds")
	if err != nil {
		return -1, err
	}

	defer func(result *sql.Rows) {
		err := result.Close()
		if err != nil {
			return
		}

	}(result)

	var count int
	for result.Next() {
		err := result.Scan(&count)
		if err != nil {
			return -1, err
		}
	}

	return count, nil
}
