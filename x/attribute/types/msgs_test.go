package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	priv1 = ed25519.GenPrivKey()
	priv2 = ed25519.GenPrivKey()
	addrs = []sdk.AccAddress{
		sdk.AccAddress(priv1.PubKey().Address()),
		sdk.AccAddress(priv2.PubKey().Address()),
	}
)

// test ValidateBasic for TestMsgAddAttribute
func TestMsgAddAttribute(t *testing.T) {
	tests := []struct {
		account            string
		owner              sdk.AccAddress
		name, proposalType string
		proposalValue      string
		expectPass         bool
	}{
		{
		    account: "",
		    owner: addrs[1],
		    name: "test",
		    proposalType: "string",
		    proposalValue: "nil owner",
		    expectPass: false,
		},
		{
		    account: addrs[0].String(),
		    owner: nil,
		    name: "test",
		    proposalType: "string",
		    proposalValue: "nil account",
		    expectPass: false,
		},
		{
		    account: "",
		    owner: nil,
		    name: "test",
		    proposalType: "string",
		    proposalValue: "nil owner and account",
		    expectPass: false,

		},
		{
            account: addrs[0].String(),
            owner: addrs[1],
            name: "test",
            proposalType: "string",
            proposalValue: "valid attribute",
            expectPass:  true,
        },
	}

	for i, tc := range tests {
		at, err := AttributeTypeFromString(tc.proposalType)
		require.NoError(t, err)
		msg := NewMsgAddAttributeRequest(
			tc.account,
			tc.owner,
			tc.name,
			at,
			[]byte(tc.proposalValue),
		)

		if tc.expectPass {
			require.NoError(t, msg.ValidateBasic(), "test: %v", i)
		} else {
			require.Error(t, msg.ValidateBasic(), "test: %v", i)
		}
	}
}

// test ValidateBasic for TestMsgUpdateAttribute
func TestMsgUpdateAttribute(t *testing.T) {
	tests := []struct {
		account       string
		owner         sdk.AccAddress
		name          string
		originalValue []byte
		originalType  AttributeType
		updateValue   []byte
		updateType    AttributeType
		expectPass    bool
		expectedError string
	}{
		{
        account: addrs[0].String(),
        owner:  addrs[1],
        name:   "example",
        originalValue:    []byte("original"),
        originalType:     AttributeType_String,
        updateValue:      []byte("update"),
        updateType:       AttributeType_Bytes,
        expectPass:        true,
        expectedError:         "",
        },
		{
		account: "",
		owner:  addrs[1],
		name:   "example",
		originalValue:    []byte("original"),
		originalType:     AttributeType_String,
		updateValue:      []byte("update"),
		updateType:       AttributeType_Bytes,
		expectPass:        false,
		expectedError:         "",
		},
		{
		account: addrs[0].String(),
		owner:  nil,
		name:   "example",
		originalValue:    []byte(""),
		originalType:     AttributeType_String,
		updateValue:      []byte("update"),
		updateType:       AttributeType_Bytes,
		expectPass:        false,
		expectedError:         "",
		},
	}


	for _, tc := range tests {
		msg := NewMsgUpdateAttributeRequest(tc.account, tc.owner, tc.name, tc.originalValue, tc.updateValue, tc.originalType, tc.updateType)

		if tc.expectPass {
			require.NoError(t, msg.ValidateBasic(), "test: %v", tc)
		} else {
			require.Error(t, msg.ValidateBasic(), "test: %v", tc)
		}
	}
}

// test ValidateBasic for TestMsgDeleteDistinctAttribute
func TestMsgDeleteDistinctAttribute(t *testing.T) {
	tests := []struct {
		account    string
		owner      sdk.AccAddress
		name       string
		value      []byte
		attrType   AttributeType
		expectPass bool
	}{
		{
		account: addrs[0].String(),
		owner:  addrs[1],
		name:   "example",
		value:    []byte("original"),
		attrType:     AttributeType_String,
		expectPass:      true,
		},
		{
		account: "",
		owner:  addrs[1],
		name:   "example",
		value:    []byte("original"),
		attrType:     AttributeType_String,
		expectPass:      false,
		},
		{
		account: addrs[0].String(),
		owner:  nil,
		name:   "example",
		value:    []byte(""),
		attrType:     AttributeType_String,
		expectPass:      false,
		},
	}

	for _, tc := range tests {
		msg := NewMsgDeleteDistinctAttributeRequest(tc.account, tc.owner, tc.name, tc.value)

		if tc.expectPass {
			require.NoError(t, msg.ValidateBasic(), "test: %v", tc)
		} else {
			require.Error(t, msg.ValidateBasic(), "test: %v", tc)
		}
	}
}
