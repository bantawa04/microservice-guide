package service

import (
	"time"

	"github.com/bantawa04/go-mq/app/errors"
	"github.com/bantawa04/go-mq/app/model"
	"github.com/bantawa04/go-mq/app/repository"
	"github.com/bantawa04/go-mq/app/response"
	"github.com/bantawa04/go-mq/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(email, password string) (*response.TokenResponse, error)
	WithTrx(tx *gorm.DB) AuthService
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) WithTrx(tx *gorm.DB) AuthService {
	newService := &authService{
		userRepo: s.userRepo.WithTrx(tx),
	}
	return newService
}

func (s *authService) generateAccessToken(user *model.UserModel, jwt_secret []byte) (string, error) {
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     expirationTime.Unix(),
		"type":    "access",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwt_secret)
}

func (s *authService) generateRefreshToken(user *model.UserModel, jwt_secret []byte) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirationTime.Unix(),
		"type":    "refresh",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwt_secret)
}

func (s *authService) Login(email, password string) (*response.TokenResponse, error) {
	user, err := s.userRepo.GetUserByEmailForAuth(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.NewUnauthorizedError("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.NewUnauthorizedError("Invalid credentials")
	}
	env := config.NewEnv()
	// Generate access token
	accessToken, err := s.generateAccessToken(user, []byte(env.JWT_SECERET))
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	// Generate refresh token
	refreshToken, err := s.generateRefreshToken(user, []byte(env.JWT_SECERET))
	if err != nil {
		return nil, errors.NewInternalError(err)
	}

	return &response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    time.Now().Add(15 * time.Minute).Unix(),
	}, nil
}
