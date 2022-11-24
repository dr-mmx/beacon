package keeper_test

import (
	"context"
	"testing"

	keepertest "beacon/testutil/keeper"
	"beacon/x/beacon/keeper"
	"beacon/x/beacon/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BeaconKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
