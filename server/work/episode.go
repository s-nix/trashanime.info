package work

import "fmt"

// Episode represents a single episode of an anime, including its title, duration, episode number, and associated notes.
type Episode struct {
	// Title is the title of the episode.
	Title string

	// Duration is the duration of the episode in seconds.
	Duration int // Duration in seconds

	// EpisodeNumber is the sequential number of the episode within the anime series.
	EpisodeNumber int

	// EpisodeNotes is a list of notes associated with the episode.
	EpisodeNotes []EpisodeNote
}

// Rating returns the rating of the episode based on its notes.
func (e *Episode) Rating() (int, error) {
	totalWeight := 0.0
	if len(e.EpisodeNotes) == 0 {
		return 0, fmt.Errorf("no episode notes found for episode: %s", e.Title)
	}
	for _, note := range e.EpisodeNotes {
		totalWeight += note.Type.Weight()
	}
	// Convert total weight to an integer rating (for simplicity, we can just cast it)
	return int(totalWeight), nil
}
