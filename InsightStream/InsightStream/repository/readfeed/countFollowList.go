package readfeed

import (
	"database/sql"
	"insightstream/repository"
	"log"
)

func CountFollowList() (int, error) {
	result, err := repository.CoreDatabae.Query("SELECT TABLE_ROWS FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME = 'follow_list'")
	if err != nil {
		log.Fatalln(err)
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
