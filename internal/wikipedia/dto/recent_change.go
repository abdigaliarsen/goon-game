package dto

type RecentChange struct {
	Event string
	Id    []struct {
		Topic     string  `json:"topic"`
		Partition *int    `json:"partition,omitempty"`
		Offset    *int    `json:"offset,omitempty"`
		Timestamp *uint64 `json:"timestamp,omitempty"`
	}
	Data struct {
		Schema string `json:"$schema"`
		Meta   struct {
			Uri       string `json:"uri"`
			RequestId string `json:"request_id"`
			Id        string `json:"id"`
			Dt        string `json:"dt"`
			Domain    string `json:"domain"`
			Stream    string `json:"stream"`
			Topic     string `json:"topic"`
			Partition int    `json:"partition"`
			Offset    int    `json:"offset"`
		} `json:"meta"`
		Id        int64  `json:"id"`
		Type      string `json:"type"`
		Namespace int    `json:"namespace"`
		Title     string `json:"title"`
		TitleURL  string `json:"title_url"`
		Comment   string `json:"comment"`
		Timestamp int64  `json:"timestamp"`
		User      string `json:"user"`
		Bot       bool   `json:"bot"`
		NotifyURL string `json:"notify_url"`
		Minor     bool   `json:"minor"`
		Patrolled bool   `json:"patrolled"`
		Length    struct {
			Old uint64 `json:"old"`
			New uint64 `json:"new"`
		} `json:"length"`
		ServerURL        string `json:"server_url"`
		ServerName       string `json:"server_name"`
		ServerScriptPath string `json:"server_script_path"`
		Wiki             string `json:"wiki"`
		ParsedComment    string `json:"parsed_comment"`
	}
}
