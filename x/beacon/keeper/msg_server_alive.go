package keeper

import (
	"context"
    "strconv"
	"beacon/x/beacon/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Alive(goCtx context.Context, msg *types.MsgAlive) (*types.MsgAliveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // Save block height as a decimal string
    blockHeight := strconv.FormatInt(ctx.BlockHeight(), 10)

    // Check if the sender is a real validator
    // That will possibly require PoS: 
    // cosmos/staking/v1beta1/query.proto
    // rpc Validators (QueryValidatorsRequest) returns (QueryValidatorsResponse)
    // see Cosmos SDK docs 

    newJitter := int64(0)

    whois, isFound := k.GetWhois(ctx, msg.Creator)
    if isFound {
        if blockHeight == whois.Latest {
            return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Beacon: repeated tx for the same block")
        }
        latest, err := strconv.ParseInt(whois.Latest, 10, 64)
        if err != nil {
            return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Beacon: latest is not an integer")
        }
        jitter, err := strconv.ParseInt(whois.Jitter, 10, 64)
        if err != nil {
            return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Beacon: jitter is not an integer")
        }
        // Accumulate jitter
        newJitter = jitter + (ctx.BlockHeight() - latest)
    }

    newWhois := types.Whois {
        Validator: msg.Creator,
        Latest: blockHeight, // Set to the current block height
        Jitter: strconv.FormatInt(newJitter, 10),
    }

    k.SetWhois(ctx, newWhois)
	return &types.MsgAliveResponse{}, nil
}
