package routes

import (
	"github.com/alexanderlesser/golf-tracker/internal/controllers"
	"github.com/alexanderlesser/golf-tracker/internal/db"
	"github.com/alexanderlesser/golf-tracker/internal/repositories"
	"github.com/alexanderlesser/golf-tracker/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterClubRoutes(r *gin.Engine, dbConn *db.MysqlStorage) {
	clubRepo := repositories.NewClubRepo(dbConn.DB)
	clubService := services.NewClubService(clubRepo)
	clubController := controllers.NewClubController(clubService)

	clubs := r.Group("/clubs")

	clubs.GET("/", clubController.GetClubs)
	clubs.PUT("/update", clubController.UpdateClub)
	clubs.POST("/create", clubController.CreateNewClub)
	clubs.DELETE("/:id", clubController.RemoveClub)
}
