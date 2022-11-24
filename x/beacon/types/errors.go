package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/beacon module sentinel errors
var (
    ErrBeaconRepeated = sdkerrors.Register(ModuleName, 1200, "Beacon: repeated tx for the same block")
    ErrBeaconJitterNotAnInt = sdkerrors.Register(ModuleName, 1300, "Beacon: jitter is not an integer")
    ErrBeaconLatestNotAnInt = sdkerrors.Register(ModuleName, 1400, "Beacon: latest is not an integer")
)
