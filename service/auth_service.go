package service

import (
	"context"
	"errors"
	"shiny-journey/ent"
	"shiny-journey/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	adminRepo repository.AdminRepository
	jwtSecret []byte
}

func NewAuthService(adminRepo repository.AdminRepository, jwtSecret string) *AuthService {
	return &AuthService{
		adminRepo: adminRepo,
		jwtSecret: []byte(jwtSecret),
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*ent.Admin, error) {

	admin, err := s.adminRepo.GetAdminByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if admin == nil || !verifyPassword(admin.Password, password) {
		return nil, errors.New("invalid password")
	}

	return admin, nil
}

func verifyPassword(storedHash, providedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedPassword)) == nil
}

func (s *AuthService) LoginWithJWT(ctx context.Context, email, password string) (string, error) {

	admin, err := s.Login(ctx, email, password)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin_id"] = admin.ID
	claims["email"] = admin.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) Register(ctx context.Context, admin *ent.Admin) (*ent.Admin, error) {
	existingAdmin, err := s.adminRepo.GetAdminByEmail(ctx, admin.Email)
	if existingAdmin != nil {
		return nil, errors.New("email not exists")
	}

	hashedPassword, err := hashPassword(admin.Password)
	if err != nil {
		return nil, err
	}

	admin.Password = hashedPassword

	createdAdmin, err := s.adminRepo.CreateAdmin(ctx, admin)
	if err != nil {
		return nil, err
	}

	return createdAdmin, nil
}

func hashPassword(providedPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(providedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
