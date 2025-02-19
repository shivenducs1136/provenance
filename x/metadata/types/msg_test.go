package types

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ownerPartyList(addresses ...string) []Party {
	retval := make([]Party, len(addresses))
	for i, addr := range addresses {
		retval[i] = Party{Address: addr, Role: PartyType_PARTY_TYPE_OWNER}
	}
	return retval
}

func TestWriteScopeRoute(t *testing.T) {
	var scope = NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		ownerPartyList("data_owner"),
		[]string{"data_accessor"},
		"value_owner",
	)
	var msg = NewMsgWriteScopeRequest(*scope, []string{})

	require.Equal(t, sdk.MsgTypeURL(msg), "/provenance.metadata.v1.MsgWriteScopeRequest")
	yaml := `scope:
  scope_id: scope1qzxcpvj6czy5g354dews3nlruxjsahhnsp
  specification_id: scopespec1qs30c9axgrw5669ft0kffe6h9gysfe58v3
  owners:
  - address: data_owner
    role: 5
  data_access:
  - data_accessor
  value_owner_address: value_owner
signers: []
scope_uuid: ""
spec_uuid: ""
`
	require.Equal(t, yaml, msg.String())

	bz, _ := ModuleCdc.MarshalJSON(msg)
	require.Equal(t, "{\"scope\":{\"scope_id\":\"scope1qzxcpvj6czy5g354dews3nlruxjsahhnsp\",\"specification_id\":\"scopespec1qs30c9axgrw5669ft0kffe6h9gysfe58v3\",\"owners\":[{\"address\":\"data_owner\",\"role\":\"PARTY_TYPE_OWNER\"}],\"data_access\":[\"data_accessor\"],\"value_owner_address\":\"value_owner\"},\"signers\":[],\"scope_uuid\":\"\",\"spec_uuid\":\"\"}", string(bz))
}

func TestWriteScopeValidation(t *testing.T) {
	var scope = NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		ownerPartyList("data_owner"),
		[]string{"data_accessor"},
		"value_owner",
	)
	var msg = NewMsgWriteScopeRequest(*scope, []string{"invalid"})
	err := msg.ValidateBasic()
	require.EqualError(t, err, "invalid scope owners: invalid party address [data_owner]: decoding bech32 failed: invalid separator index -1")
	require.Panics(t, func() { msg.GetSigners() }, "panics due to invalid addresses")

	err = msg.Scope.ValidateBasic()
	require.Error(t, err, "invalid addresses")
	require.Equal(t, "invalid scope owners: invalid party address [data_owner]: decoding bech32 failed: invalid separator index -1", err.Error())

	msg.Scope = *NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		[]Party{},
		[]string{},
		"",
	)
	err = msg.Scope.ValidateBasic()
	require.Error(t, err, "no owners")
	require.Equal(t, "invalid scope owners: at least one party is required", err.Error())

	msg.Scope = *NewScope(
		ScopeMetadataAddress(uuid.MustParse("8d80b25a-c089-4446-956e-5d08cfe3e1a5")),
		ScopeSpecMetadataAddress(uuid.MustParse("22fc17a6-40dd-4d68-a95b-ec94e7572a09")),
		ownerPartyList("cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"),
		[]string{},
		"",
	)
	msg.Signers = []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}
	err = msg.Scope.ValidateBasic()
	require.NoError(t, err, "valid add scope request")
	requiredSigners := msg.GetSigners()
	require.Equal(t, 1, len(requiredSigners))
	x, err := hex.DecodeString("85EA54E8598B27EC37EAEEEEA44F1E78A9B5E671")
	require.NoError(t, err)
	require.Equal(t, sdk.AccAddress(x), requiredSigners[0])
}

func TestAddScopeDataAccessValidateBasic(t *testing.T) {
	notAScopeId := RecordMetadataAddress(uuid.New(), "recordname")
	actualScopeId := ScopeMetadataAddress(uuid.New())

	cases := map[string]struct {
		msg      *MsgAddScopeDataAccessRequest
		wantErr  bool
		errorMsg string
	}{
		"should fail to validate basic, incorrect scope id type": {
			msg: NewMsgAddScopeDataAccessRequest(notAScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: fmt.Sprintf("address is not a scope id: %v", notAScopeId.String()),
		},
		"should fail to validate basic, requires at least one data access address": {
			msg: NewMsgAddScopeDataAccessRequest(actualScopeId, []string{}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: "at least one data access address is required",
		},
		"should fail to validate basic, incorrect data access address format": {
			msg: NewMsgAddScopeDataAccessRequest(actualScopeId, []string{"notabech32address"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: "data access address is invalid: notabech32address",
		},
		"should fail to validate basic, requires at least one signer": {
			msg: NewMsgAddScopeDataAccessRequest(actualScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{}),
			wantErr: true,
			errorMsg: "at least one signer is required",
		},
		"should successfully validate basic": {
			msg: NewMsgAddScopeDataAccessRequest(actualScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: false,
			errorMsg: "",
		},
	}

	for n, tc := range cases {
		tc := tc

		t.Run(n, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.wantErr {
				require.Error(t, err)
				require.Equal(t, tc.errorMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDeleteScopeDataAccessValidateBasic(t *testing.T) {
	notAScopeId := RecordMetadataAddress(uuid.New(), "recordname")
	actualScopeId := ScopeMetadataAddress(uuid.New())

	cases := map[string]struct {
		msg      *MsgDeleteScopeDataAccessRequest
		wantErr  bool
		errorMsg string
	}{
		"should fail to validate basic, incorrect scope id type": {
			msg: NewMsgDeleteScopeDataAccessRequest(notAScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: fmt.Sprintf("address is not a scope id: %v", notAScopeId.String()),
		},
		"should fail to validate basic, requires at least one data access address": {
			msg: NewMsgDeleteScopeDataAccessRequest(actualScopeId, []string{}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: "at least one data access address is required",
		},
		"should fail to validate basic, incorrect data access address format": {
			msg: NewMsgDeleteScopeDataAccessRequest(actualScopeId, []string{"notabech32address"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: "data access address is invalid: notabech32address",
		},
		"should fail to validate basic, requires at least one signer": {
			msg: NewMsgDeleteScopeDataAccessRequest(actualScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{}),
			wantErr: true,
			errorMsg: "at least one signer is required",
		},
		"should successfully validate basic": {
			msg: NewMsgDeleteScopeDataAccessRequest(actualScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: false,
			errorMsg: "",
		},
	}

	for n, tc := range cases {
		tc := tc

		t.Run(n, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.wantErr {
				require.Error(t, err)
				require.Equal(t, tc.errorMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddScopeOwnersValidateBasic(t *testing.T) {
	notAScopeId := RecordMetadataAddress(uuid.New(), "recordname")
	actualScopeId := ScopeMetadataAddress(uuid.New())

	cases := []struct {
		name     string
		msg      *MsgAddScopeOwnerRequest
		errorMsg string
	}{
		{
			name: "should fail to validate basic, incorrect scope id type",
			msg: NewMsgAddScopeOwnerRequest(
				notAScopeId,
				[]Party{{Address: "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", Role: PartyType_PARTY_TYPE_OWNER}},
				[]string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			errorMsg: fmt.Sprintf("address is not a scope id: %v", notAScopeId.String()),
		},
		{
			name: "should fail to validate basic, requires at least one owner address",
			msg: NewMsgAddScopeOwnerRequest(
				actualScopeId,
				[]Party{},
				[]string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"},
			),
			errorMsg: "invalid owners: at least one party is required",
		},
		{
			name: "should fail to validate basic, incorrect owner address format",
			msg: NewMsgAddScopeOwnerRequest(
				actualScopeId,
				[]Party{{Address: "notabech32address", Role: PartyType_PARTY_TYPE_OWNER}},
				[]string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"},
			),
			errorMsg: "invalid owners: invalid party address [notabech32address]: decoding bech32 failed: invalid separator index -1",
		},
		{
			name: "should fail to validate basic, incorrect party type",
			msg: NewMsgAddScopeOwnerRequest(
				actualScopeId,
				[]Party{{Address: "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", Role: PartyType_PARTY_TYPE_UNSPECIFIED}},
				[]string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"},
			),
			errorMsg: "invalid owners: invalid party type for party cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck",
		},
		{
			name: "should fail to validate basic, requires at least one signer",
			msg: NewMsgAddScopeOwnerRequest(
				actualScopeId,
				[]Party{{Address: "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", Role: PartyType_PARTY_TYPE_OWNER}},
				[]string{},
			),
			errorMsg: "at least one signer is required",
		},
		{
			name: "should successfully validate basic",
			msg: NewMsgAddScopeOwnerRequest(
				actualScopeId,
				[]Party{{Address: "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", Role: PartyType_PARTY_TYPE_OWNER}},
				[]string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"},
			),
			errorMsg: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if len(tc.errorMsg) > 0 {
				require.EqualError(t, err, tc.errorMsg, "MsgAddScopeOwnerRequest.ValidateBasic expected error")
			} else {
				require.NoError(t, err, "MsgAddScopeOwnerRequest.ValidateBasic unexpected error")
			}
		})
	}
}

func TestDeleteScopeOwnerValidateBasic(t *testing.T) {
	notAScopeId := RecordMetadataAddress(uuid.New(), "recordname")
	actualScopeId := ScopeMetadataAddress(uuid.New())

	cases := map[string]struct {
		msg      *MsgDeleteScopeOwnerRequest
		wantErr  bool
		errorMsg string
	}{
		"should fail to validate basic, incorrect scope id type": {
			msg: NewMsgDeleteScopeOwnerRequest(notAScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: fmt.Sprintf("address is not a scope id: %v", notAScopeId.String()),
		},
		"should fail to validate basic, requires at least one owner address": {
			msg: NewMsgDeleteScopeOwnerRequest(actualScopeId, []string{}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: "at least one owner address is required",
		},
		"should fail to validate basic, incorrect data access address format": {
			msg: NewMsgDeleteScopeOwnerRequest(actualScopeId, []string{"notabech32address"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: "owner address is invalid: notabech32address",
		},
		"should fail to validate basic, requires at least one signer": {
			msg: NewMsgDeleteScopeOwnerRequest(actualScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{}),
			wantErr: true,
			errorMsg: "at least one signer is required",
		},
		"should successfully validate basic": {
			msg: NewMsgDeleteScopeOwnerRequest(actualScopeId, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: false,
			errorMsg: "",
		},
	}

	for n, tc := range cases {
		tc := tc

		t.Run(n, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.wantErr {
				require.Error(t, err)
				require.Equal(t, tc.errorMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgAddContractSpecToScopeSpecRequestValidateBasic(t *testing.T) {
	contractSpecID := ContractSpecMetadataAddress(uuid.New())
	scopeSpecID := ScopeSpecMetadataAddress(uuid.New())

	cases := map[string]struct {
		msg      *MsgAddContractSpecToScopeSpecRequest
		wantErr  bool
		errorMsg string
	}{
		"should fail to validate basic, incorrect contract spec id type": {
			msg: NewMsgAddContractSpecToScopeSpecRequest(scopeSpecID, scopeSpecID, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: fmt.Sprintf("address is not a contract specification id: %v", scopeSpecID.String()),
		},
		"should fail to validate basic, incorrect scope spec id type": {
			msg: NewMsgAddContractSpecToScopeSpecRequest(contractSpecID, contractSpecID, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: fmt.Sprintf("address is not a scope specification id: %v", contractSpecID.String()),
		},
		"should fail to validate basic, requires at least one signer": {
			msg: NewMsgAddContractSpecToScopeSpecRequest(contractSpecID, scopeSpecID, []string{}),
			wantErr: true,
			errorMsg: "at least one signer is required",
		},
		"should successfully validate basic": {
			msg: NewMsgAddContractSpecToScopeSpecRequest(contractSpecID, scopeSpecID, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: false,
			errorMsg: "",
		},
	}

	for n, tc := range cases {
		tc := tc

		t.Run(n, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.wantErr {
				require.Error(t, err)
				require.Equal(t, tc.errorMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgDeleteContractSpecFromScopeSpecRequestValidateBasic(t *testing.T) {
	contractSpecID := ContractSpecMetadataAddress(uuid.New())
	scopeSpecID := ScopeSpecMetadataAddress(uuid.New())

	cases := map[string]struct {
		msg      *MsgDeleteContractSpecFromScopeSpecRequest
		wantErr  bool
		errorMsg string
	}{
		"should fail to validate basic, incorrect contract spec id type": {
			msg: NewMsgDeleteContractSpecFromScopeSpecRequest(scopeSpecID, scopeSpecID, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: fmt.Sprintf("address is not a contract specification id: %v", scopeSpecID.String()),
		},
		"should fail to validate basic, incorrect scope spec id type": {
			msg: NewMsgDeleteContractSpecFromScopeSpecRequest(contractSpecID, contractSpecID, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: true,
			errorMsg: fmt.Sprintf("address is not a scope specification id: %v", contractSpecID.String()),
		},
		"should fail to validate basic, requires at least one signer": {
			msg: NewMsgDeleteContractSpecFromScopeSpecRequest(contractSpecID, scopeSpecID, []string{}),
			wantErr: true,
			errorMsg: "at least one signer is required",
		},
		"should successfully validate basic": {
			msg: NewMsgDeleteContractSpecFromScopeSpecRequest(contractSpecID, scopeSpecID, []string{"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck"}),
			wantErr: false,
			errorMsg: "",
		},
	}

	for n, tc := range cases {
		tc := tc

		t.Run(n, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.wantErr {
				require.Error(t, err)
				require.Equal(t, tc.errorMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBindOSLocator(t *testing.T) {
	var bindRequestMsg = NewMsgBindOSLocatorRequest(ObjectStoreLocator{Owner: "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", LocatorUri: "http://foo.com"})

	err := bindRequestMsg.ValidateBasic()
	require.NoError(t, err)
	signers := bindRequestMsg.GetSigners()
	require.Equal(t, "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", signers[0].String())
	require.Equal(t, "/provenance.metadata.v1.MsgBindOSLocatorRequest", sdk.MsgTypeURL(bindRequestMsg))

	bz, _ := ModuleCdc.MarshalJSON(bindRequestMsg)
	require.Equal(t, "{\"locator\":{\"owner\":\"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck\",\"locator_uri\":\"http://foo.com\",\"encryption_key\":\"\"}}", string(bz))
}

func TestModifyOSLocator(t *testing.T) {
	var modifyRequest = NewMsgModifyOSLocatorRequest(ObjectStoreLocator{Owner: "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", LocatorUri: "http://foo.com"})

	err := modifyRequest.ValidateBasic()
	require.NoError(t, err)
	signers := modifyRequest.GetSigners()
	require.Equal(t, "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", signers[0].String())
	require.Equal(t, "/provenance.metadata.v1.MsgModifyOSLocatorRequest", sdk.MsgTypeURL(modifyRequest))

	bz, _ := ModuleCdc.MarshalJSON(modifyRequest)
	require.Equal(t, "{\"locator\":{\"owner\":\"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck\",\"locator_uri\":\"http://foo.com\",\"encryption_key\":\"\"}}", string(bz))
}

func TestDeleteOSLocator(t *testing.T) {
	var deleteRequest = NewMsgDeleteOSLocatorRequest(ObjectStoreLocator{Owner: "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", LocatorUri: "http://foo.com"})

	err := deleteRequest.ValidateBasic()
	require.NoError(t, err)

	signers := deleteRequest.GetSigners()
	require.Equal(t, "cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", signers[0].String())
	require.Equal(t, "/provenance.metadata.v1.MsgDeleteOSLocatorRequest", sdk.MsgTypeURL(deleteRequest))

	bz, _ := ModuleCdc.MarshalJSON(deleteRequest)
	require.Equal(t, "{\"locator\":{\"owner\":\"cosmos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck\",\"locator_uri\":\"http://foo.com\",\"encryption_key\":\"\"}}", string(bz))
}

func TestBindOSLocatorInvalid(t *testing.T) {
	var bindRequestMsg = NewMsgBindOSLocatorRequest(ObjectStoreLocator{Owner: "vamonos1sh49f6ze3vn7cdl2amh2gnc70z5mten3y08xck", LocatorUri: "http://foo.com"})

	err := bindRequestMsg.ValidateBasic()
	require.Error(t, err)
}

func TestBindOSLocatorInvalidAddr(t *testing.T) {
	var bindRequestMsg = NewMsgBindOSLocatorRequest(ObjectStoreLocator{Owner: "", LocatorUri: "http://foo.com"})

	err := bindRequestMsg.ValidateBasic()
	require.Error(t, err)
}

func TestBindOSLocatorInvalidURI(t *testing.T) {
	var bindRequestMsg = NewMsgBindOSLocatorRequest(ObjectStoreLocator{Owner: "", LocatorUri: "foo://foo.com"})

	err := bindRequestMsg.ValidateBasic()
	require.Error(t, err)
}

type MsgTypeURL interface {
	MsgTypeURL() string
}

func TestPrintMessageTypeStrings(t *testing.T) {
	messageTypes := []sdk.Msg{
		&MsgWriteScopeRequest{},
		&MsgDeleteScopeRequest{},
		&MsgAddScopeDataAccessRequest{},
		&MsgDeleteScopeDataAccessRequest{},
		&MsgAddScopeOwnerRequest{},
		&MsgDeleteScopeOwnerRequest{},
		&MsgWriteSessionRequest{},
		&MsgWriteRecordRequest{},
		&MsgDeleteRecordRequest{},
		&MsgWriteScopeSpecificationRequest{},
		&MsgDeleteScopeSpecificationRequest{},
		&MsgWriteContractSpecificationRequest{},
		&MsgDeleteContractSpecificationRequest{},
		&MsgAddContractSpecToScopeSpecRequest{},
		&MsgDeleteContractSpecFromScopeSpecRequest{},
		&MsgWriteRecordSpecificationRequest{},
		&MsgDeleteRecordSpecificationRequest{},
		&MsgBindOSLocatorRequest{},
		&MsgDeleteOSLocatorRequest{},
		&MsgModifyOSLocatorRequest{},
		// add  any new messages here
	}

	for _, msg := range messageTypes {
		expected := sdk.MsgTypeURL(msg)
		// compare to what we currently have in msg.go
		mtu, ok := msg.(MsgTypeURL)
		if assert.True(t, ok, "MsgTypeURL function for %s is not defined.", expected) {
			actual := mtu.MsgTypeURL()
			assert.Equal(t, expected, actual)
		}
		fmt.Println(expected)
	}
}
