package dto

import (
	"server/internal/auth/entity"
)

type UserOrAdminResponse struct {
	ID             int64   `json:"id"`
	Role           int64   `json:"role"`
	Email          string  `json:"email"`
	Name           *string `json:"name"`
	WhatsappNumber *string `json:"whatsapp_number"`
	IsVerified     bool    `json:"is_verified"`
}

type SearchUserRequest struct {
	Role  []int64 `form:"role" binding:"required,dive,numeric,role"`
	Name  string  `form:"name"`
	Last  string  `form:"last" binding:"omitempty,numeric,gte=0"`
	Limit int64   `form:"limit" binding:"numeric,gte=1,lte=25"`
}

func ConvertToUserOrAdminResponses(users []*entity.UserOrAdmin) []*UserOrAdminResponse {
	responses := []*UserOrAdminResponse{}
	for _, user := range users {
		responses = append(responses, ConvertToUserOrAdminResponse(user))
	}
	return responses
}

func ConvertToUserOrAdminResponse(user *entity.UserOrAdmin) *UserOrAdminResponse {
	return &UserOrAdminResponse{
		ID:             user.ID,
		Role:           user.Role,
		Email:          user.Email,
		Name:           user.Name,
		WhatsappNumber: user.WhatsappNumber,
		IsVerified:     user.IsVerified,
	}
}
