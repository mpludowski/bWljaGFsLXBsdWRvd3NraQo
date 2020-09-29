package model

import (
	"context"
	"fmt"
	"log"
	"time"
)

type HistoryRecord struct {
	Id string `json:"-"`
	UrlId string `json:"-"`
	Response *string `json:"response"`
	Duration float32 `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
}

func SaveToHistory(urlId, response string, duration float64) {
	query := fmt.Sprintf(
		"INSERT INTO history(url_id, response, duration) VALUES ('%s', '%s', %f)",
		urlId,
		response,
		duration,
	)

	_, err := Db.Exec(context.Background(), query)
	if err != nil {
		log.Panic(err)
	}
}

func GetHistory(id string) []HistoryRecord {
	var history []HistoryRecord

	rows, err := Db.Query(
		context.Background(),
		fmt.Sprintf(
			"SELECT id, url_id, response, duration, created_at FROM history WHERE url_id = '%s'", id,
		),
	)
	if err != nil {
		log.Panic(err)
	}

	var row HistoryRecord
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&row.Id, &row.UrlId, &row.Response, &row.Duration , &row.CreatedAt); err != nil {
			log.Panic(err)
		}
		if respValue := *row.Response; respValue == "" {
			row.Response = nil
		}
		history = append(history, row)
	}

	return history
}
