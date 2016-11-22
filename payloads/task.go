package payloads

// Task payload for POST (create)
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
}
