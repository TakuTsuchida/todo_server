package infrastructures

import (
	"app/interfaces/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Init() {
	// echo instance
	e := echo.New()
	todoRouter(e)
	// start server
	e.Logger.Fatal(e.Start(":1323"))
}

// todo
func todoRouter(e *echo.Echo) {
	g := e.Group("/todos")
	todoController := controllers.NewTodoController(NewSqlHandler())
	g.GET("/", func(c echo.Context) error {
		todos := todoController.GetAll()
		return c.JSON(http.StatusOK, todos)
	})
}
