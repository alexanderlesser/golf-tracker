package requests

type RenameClubRequest struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Loft string `json:"loft" binding:"required"`
}

type CreateClubRequest struct {
	Name string `json:"name" binding:"required"`
	Loft string `json:"loft" binding:"required"`
}
