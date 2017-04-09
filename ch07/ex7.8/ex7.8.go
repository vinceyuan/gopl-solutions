// ex7.8: Many GUIs provide a table widget with a stateful multi-tier sort:
// the primary sort key is he most recently clicked column head, the
// secondary sort key is the scond most recently clicked column head, and
// so on. Define an implementation of sort.Interface for use by such table.
// Compare that approach with repeated sorting using sort.Stable
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type clickedSort struct {
	tracks []*Track
	less   func(track1, track2 *Track) bool
}

func (x clickedSort) Len() int           { return len(x.tracks) }
func (x clickedSort) Less(i, j int) bool { return x.less(x.tracks[i], x.tracks[j]) }
func (x clickedSort) Swap(i, j int)      { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main() {
	recentClicks := make([]string, 4) // Recent clicks. First one is the most recent and so on
	recentClicks[0] = "Title"
	recentClicks[1] = "Year"
	recentClicks[2] = "Artist"

	sort.Sort(clickedSort{tracks, func(track1, track2 *Track) bool {
		if len(recentClicks) == 0 {
			return track1.Title < track2.Title
		}

		for _, recentClick := range recentClicks {
			switch recentClick {
			case "Title":
				if track1.Title != track2.Title {
					return track1.Title < track2.Title
				}
			case "Year":
				if track1.Year != track2.Year {
					return track1.Year < track2.Year
				}
			case "Length":
				if track1.Length != track2.Length {
					return track1.Length < track2.Length
				}
			case "Artist":
				if track1.Artist != track2.Artist {
					return track1.Artist < track2.Artist
				}
			}
		}
		return false
	}})
	printTracks(tracks)
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tablewriter := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tablewriter, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tablewriter, format, "-----", "-----", "-----", "-----", "-----")
	for _, track := range tracks {
		fmt.Fprintf(tablewriter, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}
	tablewriter.Flush()
}

func length(stringDuration string) time.Duration {
	duration, err := time.ParseDuration(stringDuration) // e.g. "4m32s"
	if err != nil {
		panic(stringDuration)
	}
	return duration
}

/*
// recentClicks = ["Title", "Artist", "Year"]

Title       Artist          Album              Year   Length
-----       -----           -----              -----  -----
Go          Delilah         From the Roots Up  2012   3m38s
Go          Moby            Moby               1992   3m37s
Go Ahead    Alicia Keys     As I Am            2007   4m36s
Ready 2 Go  Martin Solveig  Smash              2011   4m24s

// recentClicks = []

Title       Artist          Album              Year   Length
-----       -----           -----              -----  -----
Go          Delilah         From the Roots Up  2012   3m38s
Go          Moby            Moby               1992   3m37s
Go Ahead    Alicia Keys     As I Am            2007   4m36s
Ready 2 Go  Martin Solveig  Smash              2011   4m24s

// recentClicks = ["Title", "Year", "Artist"]

Title       Artist          Album              Year   Length
-----       -----           -----              -----  -----
Go          Moby            Moby               1992   3m37s
Go          Delilah         From the Roots Up  2012   3m38s
Go Ahead    Alicia Keys     As I Am            2007   4m36s
Ready 2 Go  Martin Solveig  Smash              2011   4m24s
*/
