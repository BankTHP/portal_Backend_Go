package middleware

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"strings"
	"pccth/portal-blog/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	PublicKey *rsa.PublicKey
}

func NewAuthMiddleware(publicKey *rsa.PublicKey) *AuthMiddleware {
	return &AuthMiddleware{PublicKey: publicKey}
}

func (am *AuthMiddleware) HasRole(roles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(model.NewErrorResponse(
				"UNAUTHORIZED",
				"ไม่พบ token การยืนยันตัวตน",
			))
		}

		for _, role := range roles {
			hasRole, err := am.CheckUserRole(token, role)
			if err != nil {
				return ctx.Status(fiber.StatusUnauthorized).JSON(model.NewErrorResponse(
					"TOKEN_VALIDATION_ERROR",
					"การตรวจสอบ token ล้มเหลว: " + err.Error(),
				))
			}
			if hasRole {
				return ctx.Next()
			}
		}

		return ctx.Status(fiber.StatusForbidden).JSON(model.NewErrorResponse(
			"FORBIDDEN",
			"ไม่มีสิทธิ์เข้าถึง",
		))
	}
}

func (am *AuthMiddleware) CheckUserRole(tokenString, role string) (bool, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("วิธีการเซ็นไม่ถูกต้อง: %v", token.Header["alg"])
		}
		return am.PublicKey, nil
	})

	if err != nil {
		return false, fmt.Errorf("ไม่สามารถแยก token ได้: %w", err)
	}

	if !token.Valid {
		return false, errors.New("token ไม่ถูกต้อง")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("ไม่สามารถอ่านข้อมูล claims จาก token ได้")
	}

	if role != "client_user" && role != "client_admin" {
		return false, fmt.Errorf("บทบาทไม่ได้รับอนุญาต: %s", role)
	}

	resourceAccess, ok := claims["resource_access"].(map[string]interface{})
	if !ok {
		return false, errors.New("ไม่พบข้อมูล resource_access ใน token")
	}

	clientAccess, ok := resourceAccess["sso-client-api"].(map[string]interface{})
	if !ok {
		return false, errors.New("ไม่พบข้อมูล sso-client-api ใน resource_access")
	}

	roles, ok := clientAccess["roles"].([]interface{})
	if !ok {
		return false, errors.New("ไม่พบข้อมูลบทบาทใน sso-client-api")
	}

	for _, r := range roles {
		if r.(string) == role {
			return true, nil
		}
	}

	return false, nil
}
