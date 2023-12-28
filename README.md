Groupie Tracker
Groupie Tracker is a web application that provides information about artists, their concerts, and related details. The application fetches data from the Groupie Tracker API and presents it in a user-friendly manner.

Table of Contents
Introduction
Table of Contents
Features
Prerequisites
Installation
Usage
Endpoints
Models
Routes


Features
#Fetch a list of artists from the Groupie Tracker API.
#Retrieve detailed information about a specific artist by ID.
#Display artist details, including images, members, creation date, and first album.
#Fetch additional data such as artist locations, concert dates, and relations.

Prerequisites
Before you begin, ensure you have the following installed:

#GoLang
#Git

Usage
Start the Golang backend server:

bash
Copy code
go run main.go
The server will be running at http://localhost:3089.

Start the React frontend:

cd frontend
npm start
The React app will be accessible at http://localhost:3000.

Endpoints
GET /artists: Fetch a list of artists.
GET /artists/:id: Fetch detailed information about an artist by ID.
Models
Artist

type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
    Locations    string   `json:"locations"`
    ConcertDates string   `json:"concertDates"`
    Relations    string   `json:"relations"`
}
Date

type Date struct {
    ID       int      `json:"id"`
    Dates    []string `json:"dates"`
    ArtistID int      `json:"artistId"`
}
ArtistDetails

type ArtistDetails struct {
    Artist    Artist        `json:"artist"`
    Locations []Location    `json:"locations"`
    Dates     []interface{} `json:"dates"`
    Relations []interface{} `json:"relations"`
}
Location

type Location struct {
    ID        int      `json:"id"`
    Locations []string `json:"locations"`
    Dates     string   `json:"dates"`
    ArtistID  int      `json:"artistId"`
}

Routes
GET /artists: Fetch a list of artists.
GET /artists/:id: Fetch detailed information about an artist by ID.
Contributing
If you would like to contribute to Groupie Tracker, follow these steps:

Fork the repository.
Create a new branch for your feature or bug fix.
Make your changes and submit a pull request.
