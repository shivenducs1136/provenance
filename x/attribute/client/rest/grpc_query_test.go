package rest_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/provenance-io/provenance/internal/pioconfig"
	"github.com/provenance-io/provenance/testutil"
	attributetypes "github.com/provenance-io/provenance/x/attribute/types"
	nametypes "github.com/provenance-io/provenance/x/name/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     testnet.Config
	testnet *testnet.Network

	accountAddr sdk.AccAddress
	accountKey  *secp256r1.PrivKey
	accountStr  string
}

func (s *IntegrationTestSuite) SetupSuite() {
	privKey, _ := secp256r1.GenPrivKey()
	s.accountKey = privKey
	addr, err := sdk.AccAddressFromHexUnsafe(s.accountKey.PubKey().Address().String())
	s.Require().NoError(err)
	s.accountAddr = addr
	s.accountStr = s.accountAddr.String()
	s.T().Log("setting up integration test suite")
	pioconfig.SetProvenanceConfig("atom", 0)

	cfg := testutil.DefaultTestNetworkConfig()

	genesisState := cfg.GenesisState
	cfg.NumValidators = 1

	var authData authtypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[authtypes.ModuleName], &authData))
	genAccount, err := codectypes.NewAnyWithValue(&authtypes.BaseAccount{
		Address:       s.accountAddr.String(),
		AccountNumber: 1,
		Sequence:      0,
	})
	s.Require().NoError(err)
	authData.Accounts = append(authData.Accounts, genAccount)

	// Configure Genesis data for name module
	var nameData nametypes.GenesisState
	nameData.Bindings = append(nameData.Bindings, nametypes.NewNameRecord("attribute", s.accountAddr, true))
	nameData.Bindings = append(nameData.Bindings, nametypes.NewNameRecord("example.attribute", s.accountAddr, false))
	nameData.Params.AllowUnrestrictedNames = false
	nameData.Params.MaxNameLevels = 3
	nameData.Params.MinSegmentLength = 3
	nameData.Params.MaxSegmentLength = 12
	nameDataBz, err := cfg.Codec.MarshalJSON(&nameData)
	s.Require().NoError(err)
	genesisState[nametypes.ModuleName] = nameDataBz

	// Configure Genesis data for account module
	var accountData attributetypes.GenesisState
	accountData.Attributes = append(accountData.Attributes,
		attributetypes.NewAttribute(
			"example.attribute",
			s.accountStr,
			attributetypes.AttributeType_String,
			[]byte("example attribute value string")))
	accountData.Params.MaxValueLength = 32
	accountDataBz, err := cfg.Codec.MarshalJSON(&accountData)
	s.Require().NoError(err)
	genesisState[attributetypes.ModuleName] = accountDataBz

	cfg.GenesisState = genesisState

	s.cfg = cfg
	//   TODO -- the following line needs to be patched because we must register our modules into this test node.
	s.testnet, err = testnet.New(s.T(), s.T().TempDir(), cfg)
	s.Require().NoError(err, "creating testnet")

	_, err = s.testnet.WaitForHeight(1)
	s.Require().NoError(err, "waiting for height 1")
}

func (s *IntegrationTestSuite) TearDownSuite() {
	testutil.CleanUp(s.testnet, s.T())
}

func (s *IntegrationTestSuite) TestGRPCQueries() {
	val := s.testnet.Validators[0]
	baseURL := val.APIAddress

	testCases := []struct {
		name     string
		url      string
		headers  map[string]string
		expErr   bool
		respType proto.Message
		expected proto.Message
	}{
		{
			name: "get attribute params",
			url: fmt.Sprintf("%s/provenance/attribute/v1/params", baseURL),
			headers: map[string]string{
				grpctypes.GRPCBlockHeightHeader: "1",
			},
			expErr: false,
			respType: &attributetypes.QueryParamsResponse{},
			expected: &attributetypes.QueryParamsResponse{Params: attributetypes.NewParams(32)},
		},
		{
			name: "get account attributes",
			url: fmt.Sprintf("%s/provenance/attribute/v1/attributes/%s", baseURL, s.accountAddr),
			headers: map[string]string{},
			expErr: false,
			respType: &attributetypes.QueryAttributesResponse{},
			expected: &attributetypes.QueryAttributesResponse{
				Account: s.accountAddr.String(),
				Attributes: []attributetypes.Attribute{
					attributetypes.NewAttribute("example.attribute",
						s.accountStr,
						attributetypes.AttributeType_String,
						[]byte("example attribute value string")),
				},
				Pagination: &query.PageResponse{NextKey: nil, Total: 1},
			},
		},
		{
			name: "get account attribute by name",
			url: fmt.Sprintf("%s/provenance/attribute/v1/attribute/%s/%s", baseURL, s.accountAddr, "example.attribute"),
			headers: map[string]string{},
			expErr: false,
			respType: &attributetypes.QueryAttributeResponse{},
			expected: &attributetypes.QueryAttributeResponse{
				Account: s.accountAddr.String(),
				Attributes: []attributetypes.Attribute{
					attributetypes.NewAttribute("example.attribute",
						s.accountStr,
						attributetypes.AttributeType_String,
						[]byte("example attribute value string")),
				},
				Pagination: &query.PageResponse{NextKey: nil, Total: 1},
			},
		},
		{
			name: "get non existint account attribute by name",
			url: fmt.Sprintf("%s/provenance/attribute/v1/attribute/%s/%s", baseURL, s.accountAddr, "im.not.here.attribute"),
			headers: map[string]string{},
			expErr: false,
			respType: &attributetypes.QueryAttributeResponse{},
			expected: &attributetypes.QueryAttributeResponse{
				Account:    s.accountAddr.String(),
				Attributes: nil,
				Pagination: &query.PageResponse{},
			},
		},
		{
			name: "scan account attribute by suffix",
			url: fmt.Sprintf("%s/provenance/attribute/v1/attribute/%s/scan/%s", baseURL, s.accountAddr, "attribute"),
			headers: map[string]string{},
			expErr: false,
			respType: &attributetypes.QueryScanResponse{},
			expected: &attributetypes.QueryScanResponse{
				Account: s.accountAddr.String(),
				Attributes: []attributetypes.Attribute{
					attributetypes.NewAttribute("example.attribute",
						s.accountStr,
						attributetypes.AttributeType_String,
						[]byte("example attribute value string")),
				},
				Pagination: &query.PageResponse{NextKey: nil, Total: 1},
			},
		},
		{
			name: "scan account attribute by suffix but send prefix",
			url: fmt.Sprintf("%s/provenance/attribute/v1/attribute/%s/scan/%s", baseURL, s.accountAddr, "example"),
			headers: map[string]string{},
			expErr: false,
			respType: &attributetypes.QueryScanResponse{},
			expected: &attributetypes.QueryScanResponse{
				Account:    s.accountAddr.String(),
				Attributes: nil,
				Pagination: &query.PageResponse{},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			resp, err := sdktestutil.GetRequestWithHeaders(tc.url, tc.headers)
			s.Require().NoError(err)
			err = val.ClientCtx.Codec.UnmarshalJSON(resp, tc.respType)

			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.expected.String(), tc.respType.String())
			}
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
