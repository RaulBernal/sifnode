package ibctransfer_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/stretchr/testify/require"

	sifapp "github.com/Sifchain/sifnode/app"
	test2 "github.com/Sifchain/sifnode/x/ethbridge/test"
	"github.com/Sifchain/sifnode/x/ibctransfer"
	"github.com/Sifchain/sifnode/x/ibctransfer/keeper"
	"github.com/Sifchain/sifnode/x/ibctransfer/keeper/testhelpers"
	"github.com/Sifchain/sifnode/x/tokenregistry/test"
	tokenregistrytypes "github.com/Sifchain/sifnode/x/tokenregistry/types"
)

func TestOnAcknowledgementMaybeConvert_Source(t *testing.T) {
	sifapp.SetConfig(false)
	addrs, _ := test2.CreateTestAddrs(2)

	rowanToken := tokenregistrytypes.RegistryEntry{
		IsWhitelisted:        true,
		Denom:                "rowan",
		IbcCounterpartyDenom: "xrowan",
		Decimals:             18,
	}

	xrowanToken := tokenregistrytypes.RegistryEntry{
		IsWhitelisted: true,
		Denom:         "xrowan",
		UnitDenom:     "rowan",
		Decimals:      10,
	}

	// successAck := channeltypes.NewResultAcknowledgement([]byte{byte(1)})
	errorAck := channeltypes.NewErrorAcknowledgement("failed packet transfer")

	msgSourceTransfer := types.NewMsgTransfer(
		"transfer",
		"channel-0",
		sdk.NewCoin("rowan", sdk.NewInt(123456789123456789)),
		addrs[0],
		addrs[1].String(),
		clienttypes.NewHeight(0, 0),
		0,
	)

	type args struct {
		goCtx           context.Context
		msg             *types.MsgTransfer
		destChannel     string
		transferToken   tokenregistrytypes.RegistryEntry
		packetToken     tokenregistrytypes.RegistryEntry
		acknowledgement channeltypes.Acknowledgement
	}

	tests := []struct {
		name   string
		args   args
		err    error
		events sdk.Events
	}{
		{
			name: "Ack err sender is source, causes refund - success",
			args: args{
				context.Background(),
				msgSourceTransfer,
				"channel-1",
				rowanToken,
				xrowanToken,
				errorAck,
			},
			err:    nil,
			events: []sdk.Event{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app, ctx, _ := test.CreateTestApp(false)
			app.TokenRegistryKeeper.SetToken(ctx, &tt.args.transferToken)
			app.TokenRegistryKeeper.SetToken(ctx, &tt.args.packetToken)
			// Setup the send conversion before testing ACK.
			tokenDeduction, tokensConverted := keeper.ConvertCoinsForTransfer(sdk.WrapSDKContext(ctx), tt.args.msg, tt.args.transferToken, tt.args.packetToken)

			initCoins := sdk.NewCoins(tt.args.msg.Token)
			sender, err := sdk.AccAddressFromBech32(tt.args.msg.Sender)
			require.NoError(t, err)

			err = app.BankKeeper.AddCoins(ctx, sender, initCoins)
			require.NoError(t, err)

			err = keeper.PrepareToSendConvertedCoins(sdk.WrapSDKContext(ctx), tt.args.msg, tokenDeduction, tokensConverted, app.BankKeeper)
			require.NoError(t, err)
			require.Equal(t, tokensConverted.String(), app.BankKeeper.GetBalance(ctx, sender, tokensConverted.Denom).String())
			require.Equal(t, tt.args.msg.Token.Sub(tokenDeduction).String(), app.BankKeeper.GetBalance(ctx, sender, tt.args.msg.Token.Denom).String())

			// Simulate send with SDK stub.
			sdkSentDenom, err := testhelpers.SendStub(ctx, app.TransferKeeper, app.BankKeeper, tokensConverted, sender, tt.args.msg.SourcePort, tt.args.msg.SourceChannel)
			require.Equal(t, tt.args.msg.Token.Sub(tokenDeduction).String(), app.BankKeeper.GetBalance(ctx, sender, tt.args.msg.Token.Denom).String())
			require.Equal(t, "0"+tokensConverted.Denom, app.BankKeeper.GetBalance(ctx, sender, tokensConverted.Denom).String())

			// Test Ack.
			packet := channeltypes.Packet{
				SourceChannel:      "channel-0",
				SourcePort:         "transfer",
				DestinationChannel: "channel-1",
				DestinationPort:    "transfer",
				Data: app.AppCodec().MustMarshalJSON(&types.FungibleTokenPacketData{
					Denom:    sdkSentDenom,
					Amount:   tokensConverted.Amount.Uint64(),
					Sender:   tt.args.msg.Sender,
					Receiver: tt.args.msg.Receiver,
				}),
			}

			_, err = ibctransfer.OnAcknowledgementMaybeConvert(ctx, app.TransferKeeper, app.TokenRegistryKeeper, app.BankKeeper, packet, app.AppCodec().MustMarshalJSON(&tt.args.acknowledgement))
			require.ErrorIs(t, err, tt.err)
			require.Equal(t, tt.args.msg.Token.String(), app.BankKeeper.GetBalance(ctx, sender, tt.args.msg.Token.Denom).String())
		})
	}
}

func TestOnAcknowledgementMaybeConvert_Sink(t *testing.T) {
	sifapp.SetConfig(false)

	addrs, _ := test2.CreateTestAddrs(2)

	denomTrace := types.DenomTrace{
		// A token coming from source will have this chain's source channel prepended when this chain generates hash.
		Path:      "transfer/channel-0",
		BaseDenom: "uatom",
	}

	atomToken := tokenregistrytypes.RegistryEntry{
		IsWhitelisted: true,
		Denom:         denomTrace.IBCDenom(),
		BaseDenom:     "uatom",
		Decimals:      6,
	}

	/*croGweiDenomTrace := types.DenomTrace{
		// A token coming from source will have this chain's source channel prepended when this chain generates hash.
		Path:      "transfer/channel-0",
		BaseDenom: "gwei",
	}

	croKweiDenomTrace := types.DenomTrace{
		// A token coming from source will have this chain's source channel prepended when this chain generates hash.
		Path:      "transfer/channel-0",
		BaseDenom: "kwei",
	}

	croGweiToken := tokenregistrytypes.RegistryEntry{
		IsWhitelisted: true,
		Denom:     croGweiDenomTrace.IBCDenom(),
		BaseDenom: "gwei",
		Decimals:  18,
		IbcCounterpartyDenom: croKweiDenomTrace.IBCDenom(),
	}

	croKweiToken := tokenregistrytypes.RegistryEntry{
		IsWhitelisted: true,
		Denom:     croKweiDenomTrace.IBCDenom(),
		BaseDenom: "kwei",
		Decimals:  10,
		UnitDenom: croGweiToken.Denom,
	}*/

	errorAck := channeltypes.NewErrorAcknowledgement("failed packet transfer")

	msgSinkTransfer := types.NewMsgTransfer(
		"transfer",
		"channel-0", // Sent from this chain back to source
		sdk.NewCoin(atomToken.Denom, sdk.NewIntFromUint64(123456789123456789)),
		addrs[0],
		addrs[1].String(),
		clienttypes.NewHeight(0, 0),
		0,
	)

	/*msgSinkTransferWithConv := types.NewMsgTransfer(
		"transfer",
		"channel-0", // Sent from this chain back to source
		sdk.NewCoin(croGweiToken.Denom, sdk.NewIntFromUint64(123456789123456789)),
		addrs[0],
		addrs[1].String(),
		clienttypes.NewHeight(0, 0),
		0,
	)*/

	type args struct {
		goCtx           context.Context
		msg             *types.MsgTransfer
		transferToken   tokenregistrytypes.RegistryEntry
		transferAsToken tokenregistrytypes.RegistryEntry
		acknowledgement channeltypes.Acknowledgement
	}

	tests := []struct {
		name   string
		args   args
		err    error
		events sdk.Events
	}{
		{
			name: "Ack err sender is sink, causes refund without conversion - success",
			args: args{
				context.Background(),
				msgSinkTransfer,
				atomToken,
				atomToken,
				errorAck,
			},
			err:    nil,
			events: []sdk.Event{},
		},
		/*{
			name: "Ack err sender is sink, causes refund with conversion - not supported",
			args: args{
				context.Background(),
				msgSinkTransferWithConv,
				croGweiToken,
				croKweiToken,
				errorAck,
			},
			err:    nil,
			events: []sdk.Event{},
		},*/
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app, ctx, _ := test.CreateTestApp(false)
			app.TokenRegistryKeeper.SetToken(ctx, &tt.args.transferToken)
			app.TokenRegistryKeeper.SetToken(ctx, &tt.args.transferAsToken)

			recvTokenPacket := types.FungibleTokenPacketData{
				Denom:  tt.args.transferToken.BaseDenom,
				Amount: tt.args.msg.Token.Amount.Uint64(),
				Sender: tt.args.msg.Receiver,
				// Fund the addr that will do a send later.
				Receiver: tt.args.msg.Sender,
			}
			recvPacket := channeltypes.Packet{
				SourceChannel:      "channel-1",
				SourcePort:         "transfer",
				DestinationChannel: "channel-0",
				DestinationPort:    "transfer",
				Data:               app.AppCodec().MustMarshalJSON(&recvTokenPacket),
			}

			sender, err := sdk.AccAddressFromBech32(tt.args.msg.Sender)
			require.NoError(t, err)

			// Simulate OnRecv so that IBC hash is stored in transfer keeper and can be,
			// converted to denom trace during processing ack.
			err = app.TransferKeeper.OnRecvPacket(ctx, recvPacket, recvTokenPacket)
			require.NoError(t, err)
			require.Equal(t, tt.args.msg.Token.Amount.String(), app.BankKeeper.GetBalance(ctx, sender, tt.args.transferToken.Denom).Amount.String())

			// Simulate send from this chain, with SDK stub.
			sdkSentDenom, err := testhelpers.SendStub(ctx, app.TransferKeeper, app.BankKeeper, tt.args.msg.Token, sender, tt.args.msg.SourcePort, tt.args.msg.SourceChannel)
			require.Equal(t, "0", app.BankKeeper.GetBalance(ctx, sender, tt.args.transferToken.Denom).Amount.String())

			// Test Ack.
			ackPacket := channeltypes.Packet{
				SourceChannel:      "channel-0",
				SourcePort:         "transfer",
				DestinationChannel: "channel-1",
				DestinationPort:    "transfer",
				Data: app.AppCodec().MustMarshalJSON(&types.FungibleTokenPacketData{
					Denom:    sdkSentDenom,
					Amount:   tt.args.msg.Token.Amount.Uint64(),
					Sender:   tt.args.msg.Sender,
					Receiver: tt.args.msg.Receiver,
				}),
			}

			_, err = ibctransfer.OnAcknowledgementMaybeConvert(ctx, app.TransferKeeper, app.TokenRegistryKeeper, app.BankKeeper, ackPacket, app.AppCodec().MustMarshalJSON(&tt.args.acknowledgement))
			require.ErrorIs(t, err, tt.err)
			if tt.err != nil {
				return
			}
			require.Equal(t, tt.args.msg.Token.String(), app.BankKeeper.GetBalance(ctx, sender, tt.args.msg.Token.Denom).String())
		})
	}
}
