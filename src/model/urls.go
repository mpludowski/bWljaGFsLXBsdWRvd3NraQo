package model

import (
	"context"
	"fmt"
	"log"
)

type Url struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	Interval int    `json:"interval"`
}

func SaveUrl(url string, interval int) string {
	query := fmt.Sprintf("INSERT INTO urls(url, interval) VALUES ('%s', %d) RETURNING id", url, interval)
	var id string
	Db.QueryRow(context.Background(), query).Scan(&id)

	return id
}

func DeleteUrl(id string) {
	query := fmt.Sprintf("DELETE FROM urls WHERE id = '%s'", id)
	Db.Exec(context.Background(), query)
}

func GetUrl(id string) Url {
	var url Url
	query := fmt.Sprintf("SELECT id, url, interval FROM urls WHERE id = '%s'", id)
	Db.QueryRow(context.Background(), query).Scan(&url.Id, &url.Url, &url.Interval)

	return url
}

func ListUrls() []Url {
	var urls []Url

	rows, err := Db.Query(context.Background(), "SELECT id, url, interval FROM urls")
	if err != nil {
		log.Panic(err)
	}

	var url Url

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&url.Id, &url.Url, &url.Interval); err != nil {
			log.Panic(err)
		}
		urls = append(urls, url)
	}

	return urls
}
