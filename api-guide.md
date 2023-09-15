# Music Metrics API Guide

The root url for this API is `https://musicmetrics.app`.  
The WebSocket is accessed via `wss://musicmetrics.app/ws`.  
All APIs are require authentication via HTTP header `Authorization`.  
Last Updated: August 10, 2023

# `PUT /code`
Called only when a user logs in with Spotify from the frontend. The `update-code` service uses the access token
obtained from Spotify (see authorization-info.md) to get user information from Spotify. If the user is already in the
database, their authtoken is updated both in the database and in local storage. If the user is not in the database,
they are added to the database and their authtoken is stored in local storage, meaning they will be added to the
tracking script as well, which runs every 2 hours. This endpoint will respond with `200`, `400`, or `500`.

# `GET /stats/[username]/[range]`
Executes the primary function of this software, computing and returning all statistics. The first parameter is the
username for whom to retrieve stats. The second is the time range, formatted like "x-y" where x is the start time and y
is the end time. Both times are inclusive and must be in unix milliseconds. All metrics described below are computed
with this time range considered.

This endpoint will respond with `200`, `401`, or `500`. The successful response is one JSON object
containing one child object per metric, all of which are either value objects (an object with one attribute which is an
integer called `value`) or list objects (an object with one attribute which is an array called `items`). Here are the
current attributes of the successful response object:

### `averageLength`
A value object containing the average duration, in seconds, of the tracks listened to by the user.

### `averagePopularity`
A list object containing 3 objects which are formatted like so:

| Attribute    | Type   | Description                                                       |
|--------------|--------|-------------------------------------------------------------------|
| `artist`     | string | The artist's name                                                 |
| `artistId`   | string | The artist's Spotify ID                                           |
| `popularity` | int    | An integer 0-100 from Spotify representing the track's popularity |
| `track`      | string | The track's name                                                  |
| `trackId`    | string | The track's Spotify ID                                            |
If no tracks of the average popularity have been listened to, the array is null.

### `averageYear`
A value object containing the average release year of the tracks listened to by the user.

### `decadeBreakdown`
A list object containing 1 object for every decade containing year(s) in which a track was released that the user
listened to. The objects are formatted like so and are sorted by descending `count`:

| Attribute | Type   | Description                                                     |
|-----------|--------|-----------------------------------------------------------------|
| `count`   | int    | The number of tracks from this decade that the user listened to |
| `decade`  | string | The decade in question, formatted like "1990s"                  |

### `hourBreakdown`
A list object containing an array of 24 integers, whose indices correspond to the hour of the day and whose values
correspond to the number of tracks listened to in that hour by the user. Time zone is not taken into account.

### `medianYear`
Similar to `averageYear`, but returns the median year instead.

### `modeYear`
A list object containing 3 objects which are formatted like so and are sorted by descending `count`:

| Attribute | Type | Description                                                                           |
|-----------|------|---------------------------------------------------------------------------------------|
| `count`   | int  | An integer representing what percent of the user's listens were released in this year |
| `year`    | int  | A year                                                                                |

### `percentExplicit`
A value object containing an integer representing what percentage of the user's listens are marked as explicit by
Spotify.

### `topAlbums`
A list object containing 1000 objects which are formatted like so and are sorted by descending `count`:

| Attribute  | Type   | Description                                                          |
|------------|--------|----------------------------------------------------------------------|
| `album`    | string | The album's name                                                     |
| `albumId`  | string | The album's Spotify ID                                               |
| `artist`   | string | The album's artist(s) as a string separated by `;;`                  |
| `artistId` | string | The Spotify ID(s) of the album's artist(s) separated by `;;`         |
| `count`    | int    | The number of times the user has listened to a track from this album |
| `image`    | string | The URL of the album's highest resolution cover art                  |

### `topAlbumsTime`
Similar to `topAlbums`, but the `count` attribute is the total number of seconds the user has spent listening to that
album.

### `topArtists`
A list object containing 1000 objects which are formatted like so and are sorted by descending `count`:

| Attribute  | Type   | Description                                                         |
|------------|--------|---------------------------------------------------------------------|
| `artist`   | string | The artist's name                                                   |
| `artistId` | string | The artist's Spotify ID                                             |
| `count`    | int    | The number of times the user has listened to a track by this artist |

### `topArtistsTime`
Similar to `topArtists`, but the `count` attribute is the total number of seconds the user has spent listening to that
artist.

### `topTracks`
A list object containing 1000 objects which are formatted like so and are sorted by descending `count`:

| Attribute  | Type   | Description                                                  |
|------------|--------|--------------------------------------------------------------|
| `artist`   | string | The track's artist(s) as a string separated by `;;`          |
| `artistId` | string | The Spotify ID(s) of the album's artist(s) separated by `;;` |
| `count`    | int    | The number of times the user has listened to this track      |
| `image`    | string | The URL of the album's highest resolution cover art          |
| `track`    | string | The track's name                                             |
| `trackId`  | string | The track's Spotify ID                                       |

### `topTracksTime`
Similar to `topTracks`, but the `count` attribute is the total number of seconds the user has spent listening to that
track.

### `totalMinutes`
A value object containing the total number of minutes the user has spent listening to Spotify.

### `totalTracks`
A value object containing the total number of tracks the user has listened to, including duplicates.

### `uniqueAlbums`
A value object containing the number of distinct albums the user has listened to.

### `uniqueArtists`
A value object containing the number of distinct artists the user has listened to.

### `uniqueTracks`
A value object containing the number of distinct tracks the user has listened to.

### `weekDayBreakdown`
A list object containing an array of integers, whose indices correspond to the day of the week (0-6) and whose values
correspond to the number of tracks listened to on that day by the user. Time zone is not taken into account.

# `POST /data/[username]`
Asynchronously loads the provided data into the database for the given username. Request body must be formatted as a
JSON array of Spotify's Extended Streaming History objects. This endpoint will respond with `200`, `400`, `401`, or
`500`.

# `DELETE /data/[username]`
Deletes account information and any auth tokens for this user (deleting listening data not yet supported). This endpoint
will respond with `200`, `401`, or `500`.
