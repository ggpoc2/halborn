package halborn

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ggpoc2/halborn/x/halborn/keeper"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data GenesisState) {
	for _, advertisement := range data.Advertisements {
		k.SetAdvertisement(ctx, advertisement)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data GenesisState) {
	return NewGenesisState(k.GetAdvertisementsFromStore(ctx))
}

// ValidateGenesis validates the halborn genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	return nil
}
