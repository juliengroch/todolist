package payloads

// Task payload for POST (create)
type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
}
