package payloads

// Task payload for POST (create) and  PATCH (update)
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
}
