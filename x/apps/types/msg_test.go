package types

import (
	"fmt"
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	"math/rand"
	"reflect"
	"testing"
)

var msgAppStake MsgAppStake
var msgBeginAppUnstake MsgBeginAppUnstake
var msgAppUnjail MsgAppUnjail

func init() {
	var pub crypto.Ed25519PublicKey
	_, err := rand.Read(pub[:])
	if err != nil {
		_ = err
	}

	moduleCdc = codec.New()
	RegisterCodec(moduleCdc)
	codec.RegisterCrypto(moduleCdc)
	moduleCdc.Seal()

	msgAppStake = MsgAppStake{
		PubKey: pub,
		Chains: []string{"0001"},
		Value:  sdk.NewInt(10),
	}
	msgAppUnjail = MsgAppUnjail{sdk.Address(pub.Address())}
	msgBeginAppUnstake = MsgBeginAppUnstake{sdk.Address(pub.Address())}
}

func TestMsgApp_GetSigners(t *testing.T) {
	type args struct {
		msgAppStake MsgAppStake
	}
	tests := []struct {
		name string
		args
		want sdk.Address
	}{
		{
			name: "return signers",
			args: args{msgAppStake},
			want: sdk.Address(msgAppStake.PubKey.Address()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppStake.GetSigner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigner() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgApp_GetSignBytes(t *testing.T) {
	type args struct {
		msgAppStake MsgAppStake
	}
	tests := []struct {
		name string
		args
		want []byte
	}{
		{
			name: "return signers",
			args: args{msgAppStake},
			want: sdk.MustSortJSON(moduleCdc.MustMarshalJSON(msgAppStake)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppStake.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgApp_Route(t *testing.T) {
	type args struct {
		msgAppStake MsgAppStake
	}
	tests := []struct {
		name string
		args
		want string
	}{
		{
			name: "return signers",
			args: args{msgAppStake},
			want: RouterKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppStake.Route(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgApp_Type(t *testing.T) {
	type args struct {
		msgAppStake MsgAppStake
	}
	tests := []struct {
		name string
		args
		want string
	}{
		{
			name: "return signers",
			args: args{msgAppStake},
			want: MsgAppStakeName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppStake.Type(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgApp_ValidateBasic(t *testing.T) {
	type args struct {
		msgAppStake MsgAppStake
	}
	tests := []struct {
		name string
		args
		want sdk.Error
		msg  string
	}{
		{
			name: "errs if no Address",
			args: args{MsgAppStake{}},
			want: ErrNilApplicationAddr(DefaultCodespace),
		},
		{
			name: "errs if no stake lower than zero",
			args: args{MsgAppStake{PubKey: msgAppStake.PubKey, Value: sdk.NewInt(-1)}},
			want: ErrBadStakeAmount(DefaultCodespace),
		},
		{
			name: "errs if no native chains supported",
			args: args{MsgAppStake{PubKey: msgAppStake.PubKey, Value: sdk.NewInt(1), Chains: []string{}}},
			want: ErrNoChains(DefaultCodespace),
		},
		{
			name: "returns err",
			args: args{MsgAppStake{PubKey: msgAppStake.PubKey, Value: msgAppStake.Value, Chains: []string{"aaaaaa"}}},
			want: ErrInvalidNetworkIdentifier("application", fmt.Errorf("net id length is > 2")),
		},
		{
			name: "returns nil if valid address",
			args: args{msgAppStake},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppStake.ValidateBasic(); got != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("ValidatorBasic() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMsgBeginAppUnstake_GetSigners(t *testing.T) {
	type args struct {
		msgBeginAppUnstake MsgBeginAppUnstake
	}
	tests := []struct {
		name string
		args
		want sdk.Address
	}{
		{
			name: "return signers",
			args: args{msgBeginAppUnstake},
			want: sdk.Address(msgAppStake.PubKey.Address()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgBeginAppUnstake.GetSigner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgBeginAppUnstake_GetSignBytes(t *testing.T) {
	type args struct {
		msgBeginAppUnstake MsgBeginAppUnstake
	}
	tests := []struct {
		name string
		args
		want []byte
	}{
		{
			name: "return signers",
			args: args{msgBeginAppUnstake},
			want: sdk.MustSortJSON(moduleCdc.MustMarshalJSON(msgBeginAppUnstake)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgBeginAppUnstake.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgBeginAppUnstake_Route(t *testing.T) {
	type args struct {
		msgBeginAppUnstake MsgBeginAppUnstake
	}
	tests := []struct {
		name string
		args
		want string
	}{
		{
			name: "return signers",
			args: args{msgBeginAppUnstake},
			want: RouterKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgBeginAppUnstake.Route(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgBeginAppUnstake_Type(t *testing.T) {
	type args struct {
		msgBeginAppUnstake MsgBeginAppUnstake
	}
	tests := []struct {
		name string
		args
		want string
	}{
		{
			name: "return signers",
			args: args{msgBeginAppUnstake},
			want: MsgAppUnstakeName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgBeginAppUnstake.Type(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgBeginAppUnstake_ValidateBasic(t *testing.T) {
	type args struct {
		msgBeginAppUnstake MsgBeginAppUnstake
	}
	tests := []struct {
		name string
		args
		want sdk.Error
		msg  string
	}{
		{
			name: "errs if no Address",
			args: args{MsgBeginAppUnstake{}},
			want: ErrNilApplicationAddr(DefaultCodespace),
		},
		{
			name: "returns nil if valid address",
			args: args{msgBeginAppUnstake},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgBeginAppUnstake.ValidateBasic(); got != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("ValidatorBasic() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMsgAppUnjail_Route(t *testing.T) {
	type args struct {
		msgAppUnjail MsgAppUnjail
	}
	tests := []struct {
		name string
		args
		want string
	}{
		{
			name: "return signers",
			args: args{msgAppUnjail},
			want: RouterKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppUnjail.Route(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgAppUnjail_Type(t *testing.T) {
	type args struct {
		msgAppUnjail MsgAppUnjail
	}
	tests := []struct {
		name string
		args
		want string
	}{
		{
			name: "return signers",
			args: args{msgAppUnjail},
			want: MsgAppUnjailName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppUnjail.Type(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgAppUnjail_GetSigners(t *testing.T) {
	type args struct {
		msgAppUnjail MsgAppUnjail
	}
	tests := []struct {
		name string
		args
		want sdk.Address
	}{
		{
			name: "return signers",
			args: args{msgAppUnjail},
			want: sdk.Address(msgAppUnjail.AppAddr),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppUnjail.GetSigner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgAppUnjail_GetSignBytes(t *testing.T) {
	type args struct {
		msgAppUnjail MsgAppUnjail
	}
	tests := []struct {
		name string
		args
		want []byte
	}{
		{
			name: "return signers",
			args: args{msgAppUnjail},
			want: sdk.MustSortJSON(moduleCdc.MustMarshalJSON(msgAppUnjail)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppUnjail.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMsgAppUnjail_ValidateBasic(t *testing.T) {
	type args struct {
		msgAppUnjail MsgAppUnjail
	}
	tests := []struct {
		name string
		args
		want sdk.Error
	}{
		{
			name: "errs if no Address",
			args: args{MsgAppUnjail{}},
			want: ErrBadApplicationAddr(DefaultCodespace),
		},
		{
			name: "returns nil if valid address",
			args: args{msgAppUnjail},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.msgAppUnjail.ValidateBasic(); got != nil {
				if !reflect.DeepEqual(got.Error(), tt.want.Error()) {
					t.Errorf("GetSigners() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
