package payloads

// Task payload for POST (create) and  PATCH (update)
type Task struct {
	Title       string `json:"title" valid:"required,stringlength(4,20)"`
	Description string `json:"description" valid:"optional"`
	Priority    int8   `json:"priority" valid:"optional"`
}
