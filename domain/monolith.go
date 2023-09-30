package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/skyshi-todolist/domain/activity"
	"github.com/jadahbakar/skyshi-todolist/domain/health"
	"github.com/jadahbakar/skyshi-todolist/domain/todos"
	"github.com/jmoiron/sqlx"
)

func MonolithIOC(f *fiber.App, db *sqlx.DB) {
	r := f.Group("/")
	health.AddRoutes(r)

	routerActivity := r.Group("/activity-groups")
	activityRepo := activity.NewRepository(db)
	activitySrv := activity.NewService(activityRepo)
	activity.NewHandler(routerActivity, activitySrv)

	routerTodo := r.Group("/todo-items")
	todoRepo := todos.NewRepository(db)
	todoSrv := todos.NewService(todoRepo)
	todos.NewHandler(routerTodo, todoSrv)

}
