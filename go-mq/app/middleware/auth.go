package middleware

import (
	"github.com/bantawa04/go-mq/app/errors"
	"github.com/bantawa04/go-mq/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const UserContextKey = "user_context"

func Protected() fiber.Handler {
	env := config.NewEnv()
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(env.JWT_SECERET)},
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			// Extract user claims after successful JWT verification
			user := c.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// Store user context
			c.Locals(UserContextKey, map[string]interface{}{
				"user_id": claims["user_id"],
				"email":   claims["email"],
			})

			return c.Next()
		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	switch err.Error() {
	case "Missing or malformed JWT":
		return errors.JWTError("Authentication token is missing or malformed", err, fiber.StatusBadRequest)
	case "token has expired":
		return errors.JWTError("Authentication token has expired", err, fiber.StatusUnauthorized)
	default:
		return errors.JWTError("Invalid authentication token", err, fiber.StatusUnauthorized)
	}
}
