package inputUtil

import (
	"fmt"
	"os"

	"github.com/jsmith312/soundcloud/track"
)

//ChooseTrackName chooses a track name from the given list
func ChooseTrackName(tracks []track.Track) track.Track {
	var tmap = make(map[string]track.Track)
	fmt.Println("Choose track name:")
	var count int
	for _, track := range tracks {
		count++
		fmt.Printf("%d: %s\n", count, track.Title)
		tmap[track.Title] = track
	}
	tname := ""
	fmt.Scanf("%s", &tname)
	track, ok := tmap[tname]
	// doesn't exist
	if !ok {
		fmt.Println("\033[0;31m[ERROR]\tPlease provide a valid track name")
		os.Exit(1)
	}
	return track
}

//ChooseNewGroups option to choose new groups to upload track to
func ChooseNewGroups() bool {
	fmt.Println("Upload to new groups?")
	opt := ""
	fmt.Scanf("%s", &opt)
	if opt == "yes" {
		return true
	}
	return false
}

//InputGenre has user specify what genre to query for groups
func InputGenre(numGroups int) string {
	// TODO
	return ""
}
