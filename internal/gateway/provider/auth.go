package provider

import (
	controllerAuth "server/internal/auth/controller"
	repositoryAuth "server/internal/auth/repository"
	routeAuth "server/internal/auth/route"
	usecaseAuth "server/internal/auth/usecase"
	controllerProfile "server/internal/profile/controller"
	repositoryProfile "server/internal/profile/repository"
	routeProfile "server/internal/profile/route"
	usecaseProfile "server/internal/profile/usecase"

	"github.com/gin-gonic/gin"
)

var (
	authUserRepository              repositoryAuth.UserRepository
	authResetTokenRepository        repositoryAuth.ResetTokenRepository
	authVerificationTokenRepository repositoryAuth.VerificationTokenRepository
	addressRepository               repositoryProfile.AddressRepository
	clusterRepository               repositoryProfile.ClusterRepository
	profileRepository               repositoryProfile.ProfileRepository
)

var (
	authUserUseCase usecaseAuth.UserUseCase
	oauthUseCase    usecaseAuth.OauthUseCase
	clusterUseCase  usecaseProfile.ClusterUseCase
	addressUseCase  usecaseProfile.AddressUseCase
	profileUseCase  usecaseProfile.ProfileUseCase
)

var (
	authUserController     *controllerAuth.UserController
	oauthController        *controllerAuth.OauthController
	refreshTokenController *controllerAuth.RefreshTokenController
	addressController      *controllerProfile.AddressController
	profileController      *controllerProfile.ProfileController
	clusterController      *controllerProfile.ClusterController
)

func ProvideAuthModule(router *gin.Engine) {
	injectAuthModuleRepository()
	injectAuthModuleUseCase()
	injectAuthModuleController()

	routeAuth.UserControllerRoute(authUserController, router)
	routeAuth.OauthControllerRoute(oauthController, router)
	routeAuth.RefreshTokenControllRoute(refreshTokenController, router, authMiddleware)
	routeProfile.AddressControllerRoute(addressController, router, authMiddleware)
	routeProfile.ProfileControllerRoute(profileController, router, authMiddleware)
	routeProfile.ClusterControllerRoute(clusterController, router)
}

func injectAuthModuleRepository() {
	authUserRepository = repositoryAuth.NewUserRepository(db)
	authResetTokenRepository = repositoryAuth.NewResetTokenRepository(db)
	authVerificationTokenRepository = repositoryAuth.NewVerificationTokenRepository(db)
	addressRepository = repositoryProfile.NewAddressRepository(db)
	clusterRepository = repositoryProfile.NewClusterRepository(db)
	profileRepository = repositoryProfile.NewProfileRepository(db)
}

func injectAuthModuleUseCase() {
	authUserUseCase = usecaseAuth.NewUserUseCase(
		redisUtil,
		jwtUtil,
		passwordEncryptor,
		base64Encryptor,
		emailTask,
		authUserRepository,
		authResetTokenRepository,
		refreshTokenRepository,
		authVerificationTokenRepository,
		store,
	)
	oauthUseCase = usecaseAuth.NewOauthUseCase(jwtUtil, authUserRepository, refreshTokenRepository, store)
	clusterUseCase = usecaseProfile.NewClusterUseCase(clusterRepository)
	addressUseCase = usecaseProfile.NewAddressUseCase(addressRepository, authUserRepository, store)
	profileUseCase = usecaseProfile.NewProfileUseCase(profileRepository, addressRepository, authUserRepository, store, cloudinaryUtil)
}

func injectAuthModuleController() {
	authUserController = controllerAuth.NewUserController(authUserUseCase)
	oauthController = controllerAuth.NewOauthController(oauthUseCase, refreshTokenUseCase)
	refreshTokenController = controllerAuth.NewRefreshTokenController(refreshTokenUseCase)
	clusterController = controllerProfile.NewClusterController(clusterUseCase)
	addressController = controllerProfile.NewAddressController(addressUseCase)
	profileController = controllerProfile.NewProfileController(profileUseCase)
}
