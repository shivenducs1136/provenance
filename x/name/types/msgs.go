package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// name message types
const (
	TypeMsgBindNameRequest       = "bind_name"
	TypeMsgDeleteNameRequest     = "delete_name"
	TypeMsgCreateRootNameRequest = "create_root_name"
)

// Compile time interface checks.
var _, _, _ sdk.Msg = &MsgBindNameRequest{}, &MsgDeleteNameRequest{}, &MsgCreateRootNameRequest{}

// NewMsgBindNameRequest creates a new bind name request
func NewMsgBindNameRequest(record, parent NameRecord) *MsgBindNameRequest {
	return &MsgBindNameRequest{
		Parent: parent,
		Record: record,
	}
}

// Route implements Msg
func (msg MsgBindNameRequest) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgBindNameRequest) Type() string { return TypeMsgBindNameRequest }

// ValidateBasic runs stateless validation checks on the message.
func (msg MsgBindNameRequest) ValidateBasic() error {
	if strings.TrimSpace(msg.Parent.Name) == "" {
		return fmt.Errorf("parent name cannot be empty")
	}
	if strings.TrimSpace(msg.Parent.Address) == "" {
		return fmt.Errorf("parent address cannot be empty")
	}
	if strings.TrimSpace(msg.Record.Name) == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if strings.Contains(msg.Record.Name, ".") {
		return fmt.Errorf("invalid name: \".\" is reserved")
	}
	if strings.TrimSpace(msg.Record.Address) == "" {
		return fmt.Errorf("address cannot be empty")
	}
	return nil
}

// GetSigners indicates that the message must have been signed by the parent.
func (msg MsgBindNameRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Parent.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// NewMsgDeleteNameRequest creates a new Delete Name Request
func NewMsgDeleteNameRequest(record NameRecord) *MsgDeleteNameRequest {
	return &MsgDeleteNameRequest{
		Record: record,
	}
}

// Route implements Msg
func (msg MsgDeleteNameRequest) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgDeleteNameRequest) Type() string { return TypeMsgDeleteNameRequest }

// ValidateBasic runs stateless validation checks on the message.
func (msg MsgDeleteNameRequest) ValidateBasic() error {
	if strings.TrimSpace(msg.Record.Name) == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if strings.TrimSpace(msg.Record.Address) == "" {
		return fmt.Errorf("address cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteNameRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners indicates that the message must have been signed by the record owner.
func (msg MsgDeleteNameRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Record.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// NewMsgCreateRootNameRequest creates a new Create Root Name Request
func NewMsgCreateRootNameRequest(title string, description string, metadata *Metadata, authority string) *MsgCreateRootNameRequest {
	return &MsgCreateRootNameRequest{
		Title:       title,
		Description: description,
		Authority:   authority,
		Metadata:    metadata,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRootNameRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic runs stateless validation checks on the message.
func (msg MsgCreateRootNameRequest) ValidateBasic() error {
	if strings.TrimSpace(msg.Metadata.Owner) != "" {
		if _, err := sdk.AccAddressFromBech32(msg.Metadata.Owner); err != nil {
			return ErrInvalidAddress
		}
	}

	if strings.TrimSpace(msg.Metadata.Name) == "" {
		return ErrInvalidLengthName
	}

	if strings.Contains(msg.Metadata.Name, ".") {
		return ErrNameContainsSegments
	}

	return nil
}

// GetSigners Implements Msg.
func (msg MsgCreateRootNameRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{fromAddress}
}
