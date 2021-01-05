package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrOrderInvalid                 = sdkerrors.Register(ModuleName, 1, "failed to validate order")
	ErrOrderNotFound                = sdkerrors.Register(ModuleName, 2, "no active order found for the hash")
	ErrPairSuspended                = sdkerrors.Register(ModuleName, 3, "trade pair suspended")
	ErrPairNotFound                 = sdkerrors.Register(ModuleName, 4, "trade pair not found")
	ErrPairExists                   = sdkerrors.Register(ModuleName, 5, "trade pair exists")
	ErrPairMismatch                 = sdkerrors.Register(ModuleName, 6, "trade pair mismatch")
	ErrBadField                     = sdkerrors.Register(ModuleName, 7, "struct field error")
	ErrMarketNotFound               = sdkerrors.Register(ModuleName, 8, "derivative market not found")
	ErrMarketInvalid                = sdkerrors.Register(ModuleName, 9, "failed to validate derivative market")
	ErrMarketExists                 = sdkerrors.Register(ModuleName, 10, "market exists")
	ErrMarketSuspended              = sdkerrors.Register(ModuleName, 11, "market suspended")
	ErrBadUpdateEvent               = sdkerrors.Register(ModuleName, 12, "order update event not confirmed")
	ErrUpdateSameValue              = sdkerrors.Register(ModuleName, 13, "cannot update the record's field with the same value")
	ErrOverLeveragedOrder           = sdkerrors.Register(ModuleName, 14, "cannot add overlevered order")
	ErrSubaccountNotFound           = sdkerrors.Register(ModuleName, 15, "subaccount not found")
	ErrOrderAlreadyExists           = sdkerrors.Register(ModuleName, 16, "order already exists")
	ErrInsufficientMargin           = sdkerrors.Register(ModuleName, 17, "subaccount has insufficient margin")
	ErrMarketExpired                = sdkerrors.Register(ModuleName, 18, "market has already expired")
	ErrOrderExpired                 = sdkerrors.Register(ModuleName, 19, "order has already expired")
	ErrInsufficientOrderQuantity    = sdkerrors.Register(ModuleName, 20, "order quantity invalid")
	ErrUnrecognizedOrderType        = sdkerrors.Register(ModuleName, 21, "unrecognized order type")
	ErrUnfundedPosition             = sdkerrors.Register(ModuleName, 22, "unfunded position for order type")
	ErrInsufficientPositionQuantity = sdkerrors.Register(ModuleName, 23, "position quantity insufficient for order type")
	ErrMarginNotBreached            = sdkerrors.Register(ModuleName, 24, "margin hold is not breached")
	ErrInsufficientTakerMargin      = sdkerrors.Register(ModuleName, 25, "taker has insufficient available margin")
)
