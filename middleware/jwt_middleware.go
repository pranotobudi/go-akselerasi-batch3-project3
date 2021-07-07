package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func JwtMiddleWare() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(key),
	})
}

func JwtMiddleWareWithRedirect() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &JwtCustomClaims{},
		SigningKey:              []byte(key),
		ErrorHandlerWithContext: RedirectToLogin,
	})
}
func RedirectToLogin(err error, c echo.Context) error {
	// response := helper.ResponseFormatter(http.StatusBadRequest, "error", "please login or register first", err.Error())

	return c.Redirect(http.StatusSeeOther, "/api/login")

	// return c.JSON(http.StatusBadRequest, response)

}
func GetJwtRole(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	role := claims.Role
	// id := claims.ID
	return role
}
func GetJwtID(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	fmt.Println("==================CLAIMS: %s", claims)
	// role := claims.Role
	id := claims.ID
	return id
}

func AdminJwtMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		id := claims.ID
		role := claims.Role
		c.String(http.StatusOK, "Welcome, id: "+string(id)+", role: "+role+"!")

		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func RoleAccessMiddleware(rolesAllowed ...string) echo.MiddlewareFunc {
	var roles []string
	for _, role := range rolesAllowed {
		roles = append(roles, role)
	}
	return RoleAccessMiddlewareImpl(roles)
}

func RoleAccessMiddlewareImpl(roleAllowedList []string) echo.MiddlewareFunc {
	fmt.Println("ROLEACCESSMIDDLEWAREIMPL CALLED ....")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenRole := GetJwtRole(c)
			fmt.Println("role:", tokenRole, "roleAllowed:", roleAllowedList)
			if !Contains(tokenRole, roleAllowedList) {
				c.JSON(http.StatusUnauthorized, fmt.Sprintf("c.JSON Limited Access, only allowed for: %v", roleAllowedList))
				c.Error(fmt.Errorf("c.Error limited access, only allowed for: %v", roleAllowedList))
				return fmt.Errorf("return limited access, only allowed for: %v", roleAllowedList)
			}
			fmt.Println("jwt :", tokenRole)
			fmt.Println("PATH :", c.Path())

			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

func Contains(element string, slices []string) bool {
	for _, x := range slices {
		fmt.Println("contains: x:", x, " element: ", element)
		if x == element {
			return true
		}
	}
	return false
}
