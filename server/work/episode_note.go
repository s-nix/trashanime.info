package work

// EpisodeNote represents a note associated with an episode, including its timestamp, content, related images, and type.
type EpisodeNote struct {
	// Timestamp indicates the time in the episode when the note is relevant.
	Timestamp int

	// Note contains the textual content of the episode note.
	Note string

	// Images holds a list of image URLs related to the episode note.
	Images []string

	// Type represents the type of the episode note, which affects the rating of the episode.
	Type EpisodeNoteType
}
