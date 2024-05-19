# Archiverfy

A tooling to archive my Spotify Discover Weekly's.


## TL;DR - I want this thing running

1. First you need to create a new Spotify app [here](https://developer.spotify.com/dashboard/applications).
2. Copy your Client ID and Client Secret.
3. Get your refresh token. This can be done with the `cmd/archiverfy-token-fetcher` app (Don't forget to create an .env file with `SPOTIFY_ID` & `SPOTIFY_SECRET`).
4. You can copy PlaylistID's within the Spotify Web Player (look in the URL there you can find it).
5. Lastly you can set, when the archiverfy should run and copy your Discover Weekly.

You can run it with Docker:
```shell
docker run \
 -e SPOTIFY_ID=<YOUR CLIENT_ID> \
 -e SPOTIFY_SECRET=<YOUR CLIENT_SECRET> \
 -e REFRESH_TOKEN=<YOUR REFRESH_TOKEN> \
 -e PLAYLIST_ID=<DISCOVER_WEEKLY_PLAYLIST_ID> \
 -e ARCHIVE_PLAYLIST_ID=<ARCHIVE_PLAYLIST_ID> \
 -e CRON=<THE CRON SCHEDULE> \
 ghcr.io/alaust-dev/archiverfy:1.0.0
```

### Cron Format:
```
--------- Seconds: 0-59
| --------- Minutes: 0-59
| | --------- Hours: 0-23
| | | --------- Day of Month: 1-31
| | | | --------- Months: 0-11 (Jan-Dec)
| | | | | --------- Day of Week: 0-6 (Sun-Sat)
| | | | | |
0 0 3 * * 2
```
This example shows a scheduled time at **03:00:00 on Tuesday**.

## Kubernetes Deployment
I deploy this Software on Kubernetes, you can find the helm chart in the charts directory.
