package models

type Athlete struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Category string `json:"category"`
	Age      int    `json:"age"`
	Country  string `json:"country"`
	Races    int    `json:"races"`
	Wins     int    `json:"wins"`
	Rating   int    `json:"rating"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	ImageKey string `json:"image_key"`
}

// Для обратной совместимости с существующими шаблонами
type Event struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Date        string  `json:"date"`
	Location    string  `json:"location"`
	Price       float64 `json:"price"`
	ImageKey    string  `json:"image_key"`
	VideoKey    string  `json:"video_key"`
	Description string  `json:"description"`
}

type TeamApplication struct {
	ID           string       `json:"id"`
	TeamName     string       `json:"team_name"`
	CaptainID    string       `json:"captain_id"`
	Members      []TeamMember `json:"members"`
	TotalMembers int          `json:"total_members"`
	MaxMembers   int          `json:"max_members"`
	Status       string       `json:"status"`
	CreatedAt    string       `json:"created_at"`
}

type TeamMember struct {
	AthleteID string `json:"athlete_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	EventID   string `json:"event_id"`
	EventType string `json:"event_type"`
	ImageKey  string `json:"image_key"`
}

type AthleteEvent struct {
	Event   Event
	Athlete Athlete
	PB      string `json:"pb"`
	Stats   string `json:"stats"`
}
