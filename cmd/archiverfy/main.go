package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

var (
	playlistId        = ""
	archivePlaylistId = ""
	client            = &spotify.Client{}
	ctx               = context.Background()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file it will be ignored.")
	}

	playlistId = os.Getenv("PLAYLIST_ID")
	if len(playlistId) <= 0 {
		log.Fatalln("PLAYLIST_ID is not set.")
	}
	archivePlaylistId = os.Getenv("ARCHIVE_PLAYLIST_ID")
	if len(archivePlaylistId) <= 0 {
		log.Fatalln("ARCHIVE_PLAYLIST_ID is not set.")
	}

	auth := spotifyauth.New()
	token := &oauth2.Token{
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
	}
	client = spotify.New(auth.Client(ctx, token))
	user, err := client.CurrentUser(ctx)
	fmt.Println("Logged in as user: ", user.DisplayName)

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Panicln("Failed to initalize scheduler: ", err.Error())
	}

	_, err = s.NewJob(
		gocron.CronJob(os.Getenv("CRON"), true),
		gocron.NewTask(archive),
	)
	if err != nil {
		log.Panicln("Failed to add job to scheduler: ", err.Error())
	}

	s.Start()

	sigInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(sigInterruptChannel, os.Interrupt)
	<-sigInterruptChannel
}

func archive() {
	log.Println("Archiving Playlist...")
	trackPage, err := client.GetPlaylistItems(ctx, spotify.ID(playlistId))
	if err != nil {
		log.Println("Failed to archive playlist: ", err.Error())
		return
	}

	var ids = []spotify.ID{}
	for _, item := range trackPage.Items {
		ids = append(ids, item.Track.Track.ID)
	}
	_, err = client.AddTracksToPlaylist(ctx, spotify.ID(archivePlaylistId), ids...)
	if err != nil {
		log.Println("Failed to add items to playlist: ", err.Error())
		return
	}
	log.Println("Done archiving.")
}
