package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewAuthenticator,
	NewAuthorizer,
	NewDiscovery,
	NewUserServiceClient,
	NewUserTokenServiceClient,
)
