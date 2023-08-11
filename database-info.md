# Music Metrics Database Information

---

## General Information

* Name: `mm`
* Type: MySQL
* Host: `pratt@musicmetrics.app:3306`

The database contains 4 tables that store user info, auth tokens, listening history, and track metadata.
All inputs are sanitized before being stored in the database.

## 👨‍🦰 User Table

Table name: `user`

| Column Name | Data Type     | Null? | Key | Description                                                                           |
|-------------|---------------|-------|-----|---------------------------------------------------------------------------------------|
| username    | varchar(255)  | NO    | PRI | The user's username (obtained from Spotify)                                           |
| displayName | varchar(255)  | NO    |     | The user's display name (obtained from Spotify)                                       |
| email       | varchar(255)  | NO    |     | The user's email (obtained from Spotify)                                              |
| refresh     | varchar(1023) | YES   |     | The refresh token used to obtain new Spotify access tokens (shouldn't change)         |
| timestamp   | bigint        | NO    |     | When MusicMetrics began tracking the user's listening, formatted as Unix milliseconds |

## 🔐 Auth Token Table

Table name: `authtoken`

| Column Name | Data Type    | Null?  | Key | Description                                                             |
|-------------|--------------|--------|-----|-------------------------------------------------------------------------|
| token       | varchar(255) | NO     | PRI | The auth token that is used to identify the user                        |
| username    | varchar(255) | NO     |     | The username of the user that the auth token belongs to                 |
| expiration  | varchar(255) | YES    |     | The expiration time of the auth token, in Unix milliseconds as a string |

## 🎧 Listening History Table

Table name: `listen`

| Column Name  | Data Type    | Null? | Key      | Description                                                   |
|--------------|--------------|-------|----------|---------------------------------------------------------------|
| username     | varchar(255) | NO    | PK1      | The username of the user that this listen belongs to          |
| timestamp    | bigint       | NO    | PK2      | The timestamp of the listen, in Unix milliseconds as a string |
| trackID      | varchar(255) | NO    | PK3, FOR | The ID of the track that was listened to                      |

## 🎵 Track Metadata Table

Table name: `track`

| Column Name | Data Type     | Null? | Key | Description                                                         |
|-------------|---------------|-------|-----|---------------------------------------------------------------------|
| id          | varchar(255)  | NO    | PRI | The track's Spotify ID                                              |
| name        | varchar(255)  | NO    |     | The track's name                                                    |
| artist      | varchar(1023) | NO    |     | The track's artist(s) as a string separated by `;;`                 |
| album       | varchar(255)  | NO    | FOR | The name of the album containing the track                          |
| explicit    | tinyint(1)    | NO    |     | Boolean representing whether or not the track is marked as explicit |
| popularity  | int           | NO    |     | The track's popularity on a scale of 0 to 100 as rated by Spotify   |
| duration    | int           | NO    |     | The track's duration in milliseconds                                |
| artistID    | varchar(1023) | NO    |     | The Spotify ID of the track's artist(s) separated by `;;`           |

## 💽 Album Metadata Table

Table name: `album`  
*Note: `genre` and `popularity` are currently not supported by the Spotify API*

| Column Name | Data Type     | Null? | Key | Description                                                  |
|-------------|---------------|-------|-----|--------------------------------------------------------------|
| id          | varchar(255)  | NO    | PRI | The album's Spotify ID                                       |
| name        | varchar(255)  | NO    |     | The album's name                                             |
| artist      | varchar(1023) | NO    |     | The album's artist(s) as a string separated by `;;`          |
| genre       | varchar(1023) | YES   |     | The album's genre(s) as a string separated by `;;`           |
| totalTracks | int           | NO    |     | The number of tracks on the album                            |
| year        | int           | NO    |     | The year the album was released                              |
| image       | varchar(1023) | NO    |     | The URL of the album's highest resolution cover art          |
| popularity  | int           | YES   |     | The album's popularity on a scale of 0 to 100                |
| artistID    | varchar(1023) | NO    |     | The Spotify ID(s) of the album's artist(s) separated by `;;` |