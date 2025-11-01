package work

import "fmt"

// Anime represents an anime series with its title, description, genre, and a list of episodes.
type Anime struct {
	Title       string
	Description string
	Genre       string
	Episodes    []Episode
}

// Rating calculates the average rating of the anime based on the ratings of its episodes.
func (a *Anime) Rating() (float64, error) {
	totalRating := 0
	for _, episode := range a.Episodes {
		rating, err := episode.Rating()
		if err != nil {
			fmt.Sprintf("warning: could not get rating for episode %s: %v", episode.Title, err)
			continue
		}
		totalRating += rating
	}
	if len(a.Episodes) > 0 {
		averageRating := totalRating / len(a.Episodes)
		return float64(averageRating), nil
	} else {
		return 0, fmt.Errorf("no episodes found for anime: %s", a.Title)
	}
}
