package database

import (
	"database/sql"
	"log"

	"github.com/goccy/go-json"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite", "./incidents.db")
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS incidents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT, age TEXT, location TEXT, date TEXT, year TEXT, reason TEXT, 
		by TEXT, protection TEXT, method TEXT, status TEXT, source TEXT, image TEXT, url TEXT
	);
	`
	if _, err = db.Exec(createTable); err != nil {
		log.Fatalf("Table creation error: %v", err)
	}
}

func InsertIncident(incident Incident) error {
	query := `
	INSERT INTO incidents 
	(name, age, location, date, year, reason, by, protection, method, status, source, image, url) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	source, _ := json.Marshal(incident.Source)
	_, err := db.Exec(query, incident.Name, incident.Age, incident.Location, incident.Date, incident.Year, incident.Reason, incident.By, incident.Protection, incident.Method, incident.Status, source, incident.Image, incident.Url)
	return err
}

func fetchIncidents(query string, args ...interface{}) ([]Incident, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []Incident
	for rows.Next() {
		var incident Incident
		var source string
		if err := rows.Scan(&incident.Id, &incident.Name, &incident.Age, &incident.Location, &incident.Date, &incident.Year, &incident.Reason, &incident.By, &incident.Protection, &incident.Method, &incident.Status, &source, &incident.Image, &incident.Url); err != nil {
			return nil, err
		}
		incidents = append(incidents, incident)
	}

	return incidents, rows.Err()
}

func fetchIncident(query string, args ...interface{}) (Incident, error) {
	row := db.QueryRow(query, args...)

	var incident Incident
	var source string
	err := row.Scan(&incident.Id, &incident.Name, &incident.Age, &incident.Location, &incident.Date, &incident.Year, &incident.Reason, &incident.By, &incident.Protection, &incident.Method, &incident.Status, &source, &incident.Image, &incident.Url)
	if err != nil {
		return Incident{}, err
	}

	return incident, nil
}

func ListIncidents() ([]Incident, error) {
	return fetchIncidents("SELECT id, name, age, location, date, year, reason, by, protection, method, status, source, image, url FROM incidents")
}

func ListIncidentsByYear(year string) ([]Incident, error) {
	return fetchIncidents("SELECT id, name, age, location, date, year, reason, by, protection, method, status, source, image, url FROM incidents WHERE year = ?", year)
}

func SearchIncidentsByName(name string) ([]Incident, error) {
	return fetchIncidents("SELECT id, name, age, location, date, year, reason, by, protection, method, status, source, image, url FROM incidents WHERE name LIKE ?", "%"+name+"%")
}

func GetIncidentById(id int) (Incident, error) {
	return fetchIncident("SELECT id, name, age, location, date, year, reason, by, protection, method, status, source, image, url FROM incidents WHERE id = ?", id)
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatalf("DB close error: %v", err)
	}
}
