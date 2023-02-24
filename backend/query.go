package main

import (
	"time"
)

type Sensor struct {
	ID int `json:"id"`
	Temp float64 `json:"temp" binding:"required"`
	Humid float64 `json:"humid" binding:"required"`
	Time time.Time `json:"time"`
}

func InsertData(data Sensor) (bool, error) {
	db, err := Connect()
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO dht11(timestamp, temp, humid) VALUES(?, ?, ?)", data.Time, data.Temp, data.Humid)
	if err != nil {
		return false, err
	}
	
	return true, nil
}

func Query() ([]Sensor, error) {

	var data []Sensor

	db, err := Connect()
	if err != nil {
		return data, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM dht11")
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		each := Sensor{}
		err := rows.Scan(&each.ID, &each.Time, &each.Temp, &each.Humid)

		if err != nil {
			return data, err
		}

		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		return data, err
	}

	return data, nil
}