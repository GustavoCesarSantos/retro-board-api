package identity

import (
	"go.uber.org/fx"

	identityApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	userDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/nativeSql"
	identityProvider "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/integration/provider"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation/handlers"
)



var Module = fx.Module(
	"identity",
	fx.Provide(
		// Repositories
		userDb.NewUserRepository,

		// Providers
		identityProvider.NewUserPublicApiProvider,

		// Applications
		identityApplication.NewCreateAuthToken,
		identityApplication.NewDecodeAuthToken,
		identityApplication.NewFindUserByEmail,
		identityApplication.NewFindUserBySigninToken,
		identityApplication.NewIncrementVersion,
		identityApplication.NewSaveUser,
		identityApplication.NewUpdateSigninToken,

		// Handlers
		identity.NewRefreshAuthToken,
		identity.NewSigninUser,
		identity.NewSigninWithGoogle,
		identity.NewSigninWithGoogleCallback,
		identity.NewSignoutUser,

		// Handlers
		identity.NewHandlers,
	),
)