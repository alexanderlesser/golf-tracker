package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alexanderlesser/golf-tracker/internal/requests"
	"github.com/alexanderlesser/golf-tracker/internal/response"
	"github.com/alexanderlesser/golf-tracker/internal/services"
	"github.com/alexanderlesser/golf-tracker/internal/utils"
	"github.com/alexanderlesser/golf-tracker/pkg/httpstatus"
	"github.com/gin-gonic/gin"
)

// * wraps the club service
type ClubController struct {
	ClubService services.ClubService
}

// * returns a new controller instance
func NewClubController(cs services.ClubService) *ClubController {
	return &ClubController{ClubService: cs}
}

// GET /clubs
func (cc *ClubController) GetClubs(c *gin.Context) {
	clubs, err := cc.ClubService.FetchAllClubs()
	if err != nil {
		response.Error(err, "Failed to fetch clubs", httpstatus.InternalServerError, c)
		return
	}
	response.Success(clubs, "Fetched clubs successfully", httpstatus.OK, c)
}

// PUT /clubs/rename
func (cc *ClubController) UpdateClub(c *gin.Context) {
	var req requests.RenameClubRequest

	utils.BindJSONRequest(c, &req)

	if req.Name == "" {
		response.Error(errors.New("club name cannot be empty"), "Invalid club name", httpstatus.BadRequest, c)
		return
	}

	if req.Loft == "" {
		response.Error(errors.New("loft cannot be empty"), "Invalid club Loft", httpstatus.BadRequest, c)
		return
	}

	err := cc.ClubService.UpdateClubValues(req.ID, req.Name, req.Loft)
	if err != nil {
		fmt.Println("Error in updating club name:", err)
		response.Error(err, "Could not update name", httpstatus.InternalServerError, c)
		return
	}

	response.Success(nil, "Club updated successfully", httpstatus.OK, c)

}

// DELETE /clubs/:id
func (cc *ClubController) RemoveClub(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(err, "Invalid club id", httpstatus.BadRequest, c)
		return
	}

	err = cc.ClubService.RemoveClub(id)
	if err != nil {
		response.Error(err, "Failed to remove club", httpstatus.InternalServerError, c)
		return
	}

	response.Success(nil, "Fetched clubs successfully", httpstatus.Removed, c)
}

// POST /clubs/create
func (cc *ClubController) CreateNewClub(c *gin.Context) {
	var req requests.CreateClubRequest

	utils.BindJSONRequest(c, &req)

	if req.Name == "" {
		response.Error(errors.New("club name cannot be empty"), "Invalid club name", httpstatus.BadRequest, c)
		return
	}
	if req.Loft == "" {
		response.Error(errors.New("club loft cannot be empty"), "Invalid club loft", httpstatus.BadRequest, c)
		return
	}

	newClub, err := cc.ClubService.CreateClub(req.Name, req.Loft)
	if err != nil {
		response.Error(err, "Failed to create club", httpstatus.InternalServerError, c)
		return
	}

	response.Success(newClub, "New club created", httpstatus.Created, c)

}
