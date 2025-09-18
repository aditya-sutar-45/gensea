package main

import (
	"context"
	"log"

	"github.com/aditya-sutar-45/gensea/ingestion/config"
	"github.com/aditya-sutar-45/gensea/ingestion/db"
	"github.com/aditya-sutar-45/gensea/ingestion/models"
	"github.com/aditya-sutar-45/gensea/ingestion/parsers"
)

func main() {
	ctx := context.Background()

	DB, err := config.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer DB.Close()

	if err := db.EnsureTables(ctx, DB); err != nil {
		log.Fatalln("failed to create table!", err)
		return
	}
	log.Println("tables exists")

	raw_ocean_data, err := parsers.LoadCSV[models.RawOceanData]("../data_lake/csv/oceanographic_data.csv")
	if err != nil {
		log.Fatalln("failed to load csv!", err)
		return
	}
	log.Println("loaded ocean data from csv")

	if err := db.ImportRecords(ctx, DB, raw_ocean_data); err != nil {
		log.Fatalln("failed to insert data into table", err)
		return
	}
	log.Println("imported records of len: ", len(raw_ocean_data))
}
