package app

import (
	"github.com/PriceChain/rd_net/app/types"
	authtypes "github.com/PriceChain/rd_net/app/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/authz"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcante "github.com/cosmos/ibc-go/v3/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
)

const (
	RD_NET_COIN = "uprc"
)

// HandlerOptions extends the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions

	StakingKeeper authtypes.StakingKeeper
	BankKeeper    authtypes.BankKeeper
	IBCKeeper     *ibckeeper.Keeper
	Cdc           codec.BinaryCodec
}

type MinCommissionDecorator struct {
	options HandlerOptions
	cdc     codec.BinaryCodec
}

func NewMinCommissionDecorator(options HandlerOptions) MinCommissionDecorator {
	return MinCommissionDecorator{options, options.Cdc}
}

func (min MinCommissionDecorator) CheckEmissionRule(totalSupply int64, amount int64, suggestEmission int64) bool {
	emission := types.STAKE_AMT_EMISSION_5
	if amount > types.STAKE_AMT_EMISSION_20 {
		emission = types.EMISSION_20
	} else if amount > types.STAKE_AMT_EMISSION_19 {
		emission = types.EMISSION_19
	} else if amount > types.STAKE_AMT_EMISSION_18 {
		emission = types.EMISSION_18
	} else if amount > types.STAKE_AMT_EMISSION_17 {
		emission = types.EMISSION_17
	} else if amount > types.STAKE_AMT_EMISSION_16 {
		emission = types.EMISSION_16
	} else if amount > types.STAKE_AMT_EMISSION_15 {
		emission = types.EMISSION_15
	} else if amount > types.STAKE_AMT_EMISSION_14 {
		emission = types.EMISSION_14
	} else if amount > types.STAKE_AMT_EMISSION_13 {
		emission = types.EMISSION_13
	} else if amount > types.STAKE_AMT_EMISSION_12 {
		emission = types.EMISSION_12
	} else if amount > types.STAKE_AMT_EMISSION_11 {
		emission = types.EMISSION_11
	} else if amount > types.STAKE_AMT_EMISSION_10 {
		emission = types.EMISSION_10
	} else if amount > types.STAKE_AMT_EMISSION_9 {
		emission = types.EMISSION_9
	} else if amount > types.STAKE_AMT_EMISSION_8 {
		emission = types.EMISSION_8
	} else if amount > types.STAKE_AMT_EMISSION_7 {
		emission = types.EMISSION_7
	} else if amount > types.STAKE_AMT_EMISSION_6 {
		emission = types.EMISSION_6
	} else {
		emission = types.EMISSION_5
	}

	if suggestEmission > int64(emission) {
		return false
	}

	return true
}

func (min MinCommissionDecorator) AnteHandle(
	ctx sdk.Context, tx sdk.Tx,
	simulate bool, next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	minCommissionRate := sdk.NewDecWithPrec(5, 2)
	maxCommissionRate := sdk.NewDecWithPrec(20, 2)

	// Calc total supply
	totalSupply := min.options.BankKeeper.GetSupply(ctx, RD_NET_COIN).Amount.Int64()

	validMsg := func(m sdk.Msg, totalSupply int64) error {
		switch msg := m.(type) {
		case *stakingtypes.MsgCreateValidator:
			// prevent new validators joining the set with
			// commission set below 5%
			c := msg.Commission
			if c.Rate.LT(minCommissionRate) {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Commission can't be lower than 5%")
			}

			if c.Rate.GT(maxCommissionRate) {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Commission can't be greater than 20%")
			}

			if !min.CheckEmissionRule(totalSupply, msg.Value.Amount.Int64(), msg.Commission.Rate.RoundInt64()) {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Invalid commission rate.")
			}
		case *stakingtypes.MsgEditValidator:
			// if commission rate is nil, it means only
			// other fields are affected - skip
			if msg.CommissionRate == nil {
				break
			}

			if msg.CommissionRate.LT(minCommissionRate) {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "commission can't be lower than 5%")
			}

			if msg.CommissionRate.GT(maxCommissionRate) {
				return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "commission can't be greater than 20%")
			}
		}

		return nil
	}

	validAuthz := func(execMsg *authz.MsgExec, totalSupply int64) error {
		for _, v := range execMsg.Msgs {
			var innerMsg sdk.Msg
			err := min.cdc.UnpackAny(v, &innerMsg)
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot unmarshal authz exec msgs")
			}

			err = validMsg(innerMsg, totalSupply)
			if err != nil {
				return err
			}
		}

		return nil
	}

	for _, m := range msgs {
		if msg, ok := m.(*authz.MsgExec); ok {
			if err := validAuthz(msg, totalSupply); err != nil {
				return ctx, err
			}
			continue
		}

		// validate normal msgs
		err = validMsg(m, totalSupply)
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}

	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		NewMinCommissionDecorator(options),
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.HandlerOptions.BankKeeper, options.FeegrantKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
