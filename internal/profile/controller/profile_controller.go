package controller

import (
	"strconv"

	"server/internal/auth/utils"
	appErrorProfile "server/internal/profile/apperror"
	dtoProfile "server/internal/profile/dto"
	"server/internal/profile/usecase"
	"server/pkg/utils/ginutils"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUseCase usecase.ProfileUseCase
}

func NewProfileController(
	profileUseCase usecase.ProfileUseCase,
) *ProfileController {
	return &ProfileController{
		profileUseCase: profileUseCase,
	}
}

func (pc *ProfileController) GetMyProfile(ctx *gin.Context) {
	myProfile, err := pc.profileUseCase.GetProfile(ctx, utils.GetValueUserIdFromToken(ctx), "all")
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOK(ctx, myProfile)
}

func (pc *ProfileController) PutMyProfile(ctx *gin.Context) {
	req := new(dtoProfile.RequestPutMyProfile)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(err)
		return
	}
	myProfile, err := pc.profileUseCase.PutMyProfile(ctx, req, utils.GetValueUserIdFromToken(ctx), int64(utils.GetValueRoleUserFromToken(ctx)))
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseOK(ctx, myProfile)
}

func (pc *ProfileController) GetProfileById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil || userId <= 0 {
		ctx.Error(appErrorProfile.NewInvalidIdUserProfileError())
		return
	}
	myProfile, err := pc.profileUseCase.GetProfile(ctx, int64(userId), "byId")
	if err != nil {
		ctx.Error(err)
		return
	}
	ginutils.ResponseOK(ctx, myProfile)
}
