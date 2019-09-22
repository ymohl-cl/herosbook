package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ymohl-cl/herosbook/pkg/app/auth"
	"github.com/ymohl-cl/herosbook/pkg/config"
	"golang.org/x/xerrors"
)

// package name
const pkgName = "_auth"

// App component
type App struct {
	driver *echo.Echo
	ssl    SSL
	port   string
	jwtKey string
}

// New app http
func New(appName string) (App, error) {
	var c Config
	var err error

	if err = config.ParseEnv(appName, &c); err != nil {
		return App{}, err
	}
	return NewWithConfig(appName, c)
}

// NewWithConfig server http
func NewWithConfig(appName string, c Config) (App, error) {
	a := App{
		driver: echo.New(),
		ssl:    c.Ssl,
		port:   c.Port,
		jwtKey: c.Auth.JwtKey,
	}
	a.driver.Use(middleware.Logger())
	a.driver.GET("/ping", Ping)
	if c.Auth.Enable {
		if a.jwtKey == "" {
			return App{}, xerrors.New("You can't enable the jwt authentication with a jwtkey empty value")
		}
		driverAuth, err := auth.New(appName)
		if err != nil {
			return App{}, err
		}
		a.driver.POST("/login", driverAuth.Login)
		a.driver.POST("/register", driverAuth.Register)
	}
	return a, nil
}

// SubRouter return with the prefix path specified on parameter
func (a App) SubRouter(prefix string, jwtSecure bool) (*echo.Group, error) {
	router := a.driver.Group(prefix)
	if jwtSecure {
		if a.jwtKey == "" {
			return nil, xerrors.New("Can't use jwtSecure without enable this service")
		}
		router.Use(middleware.JWT([]byte(a.jwtKey)))
	}
	return router, nil
}

// Start run the server
func (a App) Start() error {
	var err error

	defer a.driver.Close()
	if a.ssl.Enable {
		if err = a.driver.StartTLS(":"+a.port, a.ssl.Cert, a.ssl.Key); err != nil {
			return err
		}
	} else {
		if err = a.driver.Start(":" + a.port); err != nil {
			return err
		}
	}
	return nil
}
