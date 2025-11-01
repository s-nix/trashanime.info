package work

// EpisodeNoteType represents different types of episode notes that can affect the rating based on a weight system.
type EpisodeNoteType int

const (
    // FanService indicates content that is intended to sexually stimulate the audience.
    FanService EpisodeNoteType = iota

    // SexualContent indicates scenes or themes that directly involve sexual activity.
    SexualContent

    // Nudity indicates the presence of unclothed characters.
    Nudity

    // Violence indicates scenes involving physical force intended to hurt or damage.
    Violence

    // Gore indicates graphic depictions of injury or death.
    Gore

    // Profanity indicates the use of offensive language.
    Profanity

    // Blasphemy indicates irreverence towards religious beliefs.
    Blasphemy

    // PromotedDrugUse indicates scenes that encourage drug use.
    PromotedDrugUse

    // MoralAmbiguity indicates situations where the distinction between right and wrong is unclear.
    MoralAmbiguity
)

// Weight returns the weight associated with the EpisodeNoteType for rating calculations.
func (ent EpisodeNoteType) Weight() float64 {
    switch ent {
    case FanService:
        return 1.0
    case SexualContent:
        return 1.5
    case Nudity:
        return 2.0
    case Violence:
        return 0.2
    case Gore:
        return 0.8
    case Profanity:
        return 0.5
    case Blasphemy:
        return 1.0
    case PromotedDrugUse:
        return 1.0
    case MoralAmbiguity:
        return 0.5
    default:
        return 0.0
    }
}
