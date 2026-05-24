type atom struct {
	ID          string `json:"id"`
	SpaceId     string `json:"space_id"`
	FilePath    string `json:"file_path"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
	CreatedAt   string `json:"created_at"`
}
