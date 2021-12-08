package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ggpoc2/halborn/x/halborn/types"
)

// NewQuerier creates a new querier for halborn clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListAdvertisement:
			return listAdvertisement(ctx, k)
		case types.QueryGetAdvertisement:
			return getAdvertisement(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown halborn query endpoint")
		}
	}
}
