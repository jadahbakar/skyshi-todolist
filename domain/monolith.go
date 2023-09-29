package domain

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/skyshi-todolist/domain/activity"
	"github.com/jadahbakar/skyshi-todolist/domain/health"
)

func MonolithIOC(f *fiber.App, db *sql.DB) {
	r := f.Group("/")
	health.AddRoutes(r)

	routerActivity := r.Group("/activity-groups")
	activityRepo := activity.NewRepository(db)
	activitySrv := activity.NewService(activityRepo)
	activity.NewHandler(routerActivity, activitySrv)

}
