package keeper_test

import (
	"testing"

	testkeeper "beacon/testutil/keeper"
	"beacon/x/beacon/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BeaconKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
