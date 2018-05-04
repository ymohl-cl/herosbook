



import (
	"fmt"

	"github.com/labstack/echo"
)

func Skip(c echo.Context) bool {
	fmt.Println("Skip ? ", c.Path())
	if c.Path() == "/" {
		fmt.Println("equal")
		return true
	}
	fmt.Println("not equal")
	return false
}

func Test(s string, c echo.Context) (bool, error) {
	if s == "api-key" {
		return true, nil
	}
	return false, nil
}
