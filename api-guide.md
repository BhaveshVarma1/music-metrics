# Music Metrics API Guide

The root url for this API is `https://dev.musicmetrics.app/api/v1`.  
The WebSocket is accessed via `wss://dev.musicmetrics.app/ws`.

# Currently Available Endpoints
* All APIs are require authentication via HTTP header `Authorization`.  
* All `GET` endpoints can respond with only `200 OK`, `401 Bad Token`, and `500 Internal Server Error`.  
* Unless otherwise specified, all `GET` endpoints return JSON formatted data. For lists of things, the data is stored 
as a single object, which is an array of objects called `items`. In the case of single values, the data is stored as a
single object, which is an integer called `value`.  
* Last Updated: April 11, 2023

## `POST /updateCode`
Called only when a user logs in with Spotify from the frontend. The `update-code` service uses the access token
obtained from Spotify (see authorization-info.md) to get user information from Spotify. If the user is already in the
database, their authtoken is updated both in the database and in local storage. If the user is not in the database,
they are added to the database and their authtoken is stored in local storage, meaning they will be added to the
tracking script as well, which runs every 2 hours. This endpoint can respond with only 200, 400, and 500.

## `GET /averageLength/{username}`
Returns the average duration, in seconds, of all songs ever listened to by the user.

## `GET /averagePopularity/{username}`
Returns 3 objects each containing the name of the song, the artist, and the popularity of that song. All three songs
will be of the same popularity, which is the average. If no songs of the average popularity have been listened to,
the array is null but will still return `200`.

## `GET /averageYear/{username}`
Returns the average release year of all songs ever listened to by the user.

## `GET /decadeBreakdown/{username}`
Returns an array of objects, each containing the decade (formatted as '2000s') and the number of songs listened to
in that decade, ordered by most listened to descending.

## `GET /hourBreakdown/{username}`
Returns an array of integers, whose indices correspond to the hour of the day (0-23) and whose values correspond to the
number of songs listened to in that hour by the user. Currently doesn't take time zones into account.

## `GET /medianYear/{username}`
Similar to `averageYear`, but returns the median year instead.

## `GET /modeYear/{username}`
Returns the 3 objects, each containing the year as an int and the percent of total songs listened to that that year
represents.

## `GET /percentExplicit/{username}`
Returns a single integer, which is the percent of songs ever listened to by the user that are marked as explicit
by Spotify.

## `GET /topAlbums/{username}`
Returns an array of TopAlbum objects, which contain the album's name, artist, URL to its cover art, and a count
representing how many times a song on that album has been listened to by the user.

## `GET /topAlbumsTime/{username}`
Same as `topAlbums`, but returns the top albums listened to by time, where the count is the total time in seconds
instead of the number of songs.

## `GET /topArtists/{username}`
Returns an array of TopArtist objects, which contain the artist's name and a count representing how many times a
song by that artist has been listened to by the user.

## `GET /topArtistsTime/{username}`
Same as `topArtists`, but returns the top artists listened to by time, where the count is the total time in seconds.

## `GET /topSongs/{username}`
Returns an array of TopSong objects, which contain the song's name, artist, and a count representing how many times
that song has been listened to by the user.

## `GET /topSongsTime/{username}`
Same as `topSongs`, but returns the top songs listened to by time, where the count is the total time in seconds.

## `GET /totalSongs/{username}`
Returns the total number of songs the user has listened to, including duplicates.

## `GET /uniqueAlbums/{username}`
Returns the number of distinct albums the user has listened to.

## `GET /uniqueArtists/{username}`
Returns the number of distinct artists the user has listened to.

## `GET /uniqueSongs/{username}`
Returns the number of distinct songs the user has listened to.

## `GET /weekDayBreakdown/{username}`
Returns an array of integers, whose indices correspond to the day of the week (0-6) and whose values correspond to the
number of songs listened to on that day by the user. Currently doesn't take time zones into account.
