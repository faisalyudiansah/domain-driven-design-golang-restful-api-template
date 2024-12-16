package controller

import (
	"server/internal/auth/dto"
	"server/internal/auth/usecase"
	"server/internal/auth/utils"
	"server/pkg/utils/ginutils"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	refreshTokenUseCase usecase.RefreshTokenUseCase
}

func NewRefreshTokenController(refreshTokenUseCase usecase.RefreshTokenUseCase) *RefreshTokenController {
	return &RefreshTokenController{
		refreshTokenUseCase: refreshTokenUseCase,
	}
}

func (c *RefreshTokenController) Logout(ctx *gin.Context) {
	req := &dto.DeleteRefreshTokenRequest{JTI: utils.GetJTIFromToken(ctx)}
	if err := c.refreshTokenUseCase.Delete(ctx, req); err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOKPlain(ctx)
}
