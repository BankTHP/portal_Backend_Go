package model

type CreateReleaseNoteRequest struct {
	ReleaseNoteBody   string `json:"releaseNoteBody" validate:"required"`
	ReleaseNoteHeader string `json:"releaseNoteHeader" validate:"required"`
}

type UpdateReleaseNoteRequest struct {
	ReleaseNoteHeader string `json:"releaseNoteHeader" validate:"required"`
	ReleaseNoteBody   string `json:"releaseNoteBody" validate:"required"`
}
