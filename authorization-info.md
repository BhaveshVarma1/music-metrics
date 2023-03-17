# How To Use Spotify Web API
Similar documentation can be found [here](https://developer.spotify.com/documentation/general/guides/authorization/code-flow/).  
Note: the user must be signed in to Music Metrics.


## Step 1: User Authorization

---

At the click of a button on the web page, a JavaScript function is executed that sends a `GET` request to the
Spotify API. The request contains the following parameters in the URI:

  * `client_id` - The client ID of the application - this can be public
  * `response_type` - The type of response that is expected from the API ("code" in this case)
  * `redirect_uri` - The URI to which the user is redirected after authorization, success or failure - must be defined in the [application settings](https://developer.spotify.com/dashboard/applications)
  * `scope` - The [permissions](https://developer.spotify.com/documentation/general/guides/authorization/scopes/) that the application requests (in a space-separated list)
  * `state` - A randomly generated token to prevent [CSRF](https://en.wikipedia.org/wiki/Cross-site_request_forgery) attacks

As this request is processed by Spotify, the user will be redirected to a Spotify pop-up which shows the user the scopes
in detail that they will be allowing MusicMetrics access to. Whether or not the user authorizes the application, they will
be redirected to the `redirect_uri` with the following parameters:

  * `code` - The authorization code that can be exchanged for an access token
  * `state` - The randomly generated token that was sent in the original request

The front end then validates the `state` and if it is valid, stores the `code` in the database
using the `/updateCode` endpoint with their Music Metrics auth token. The code is now ready to be used to fetch
access tokens from Spotify.

## Step 2: Request Access Tokens

---

Now that we have the code and client secret stored on the back end, we can make a `POST` request to the Spotify Web API
to exchange the code for an access token. The request contains the following parameters in the body:

  * `grant_type` - The type of grant that is being requested ("authorization_code" in this case)
  * `code` - The authorization code that was received in Step 1
  * `redirect_uri` - Not used for actual redirection in this case, but for validation (must match the URI used in Step 1)

In addition to these parameters, the request must contain the following headers:

  * `Authorization` - Base 64 encoded string that contains the client ID and client secret key. The field must have the format: `Authorization: Basic [base64 encoded client_id:client_secret]`
  * `Content-Type` - The type of data being sent ("application/x-www-form-urlencoded" in this case)

If the request is successful, the response will contain the following parameters:

  * `access_token` - The access token that can be used to make requests to the Spotify Web API
  * `token_type` - The type of token ("Bearer" in this case)
  * `scope` - A space-separated list of scopes
  * `expires_in` - The number of seconds until the token expires
  * `refresh_token` - The refresh token that can be used to request a new access token

This `access_token` is now ready to be used to make requests to the Spotify Web API.

## Step 3: Refreshing the Token

---

The access token that was received in Step 2 will expire after a certain amount of time (usually 1 hour). To prevent this, we can use the
refresh token that was received in Step 2 to request a new access token. The request contains the following parameters in the body:

  * `grant_type` - The type of grant that is being requested ("refresh_token" in this case)
  * `refresh_token` - The refresh token that was received in Step 2

In addition to these parameters, the request must contain the same headers as in Step 2. The response will contain the same
parameters as in Step 2, except that the `refresh_token` will not be included.