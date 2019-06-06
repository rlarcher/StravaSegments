package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/strava/go.strava"
)

func main() {
	var segmentId int64
	var accessToken string

	flag.Int64Var(&segmentId, "id", 229781, "Segment Id")
	flag.StringVar(&accessToken, "token", "", "Access Token")

	flag.Parse()
	fmt.Printf("%d and %s\n", segmentId, accessToken)
	if accessToken == "" {
		return
	}
	fmt.Println("Having trouble")
	client := strava.NewClient(accessToken)
	// athlete, _ := strava.NewAthletesService(client).Get(11841412).Do()
	segments, _ := strava.NewAthletesService(client).ListStarredSegments(11841412).Do()
	fmt.Printf("len is %d", len(segments))
	for _, segment := range segments {
		parseSegmentById(segment.SegmentSummary.Id, client)
	}

	// parseSegmentById(segmentId, client)
}

func parseSegmentById(segmentId int64, client *strava.Client) {
	segment, err := strava.NewSegmentsService(client).Get(segmentId).Do()
	if err != nil {
		return
	}

	distance := segment.SegmentSummary.Distance
	count := segment.AthleteCount

	fmt.Printf("Distance in meters is %v and the number of runners who have done it is %d\n", distance, count)

	leaderboard, err := strava.NewSegmentsService(client).GetLeaderboard(segmentId).Do()
	for _, entry := range leaderboard.Entries {
		name := entry.AthleteName
		elapsedTime := entry.ElapsedTime
		paceMPS := distance / float64(elapsedTime)
		paceMPH := paceMPS * 2.2369
		pace := float64(60) / paceMPH
		minutes := math.Floor(pace)
		seconds := (pace - minutes) * 60
		fmt.Printf("%s ran %.2f meters in %d at %d:%02.0f\n", name, distance, elapsedTime, int(minutes), seconds)
		calc_func := linearRegressionLSE()
		record := calc_func(distance)
		if elapsedTime < int(record) {
			fmt.Printf("%s did not set a WR, ran %.2f meters in %d at %d:%02.0f\n", name, distance, elapsedTime, int(minutes), seconds)
		} else {
			fmt.Printf("%s set a WR!!, ran %.2f meters in %d at %d:%02.0f\n", name, distance, elapsedTime, int(minutes), seconds)
		}
	}
}
