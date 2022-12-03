package plasticcredit

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gopkg.in/yaml.v2"
)

// NewParams creates a new Params instance
func NewParams(issuerCreator string) Params {
	return Params{
		IssuerCreator: issuerCreator,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("")
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.IssuerCreator != "" {
		_, err := sdk.AccAddressFromBech32(p.IssuerCreator)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid issuer creator address (%s)", err)
		}
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func (idc IDCounters) Validate() error {
	if idc.NextIssuerId == 0 {
		return errors.Wrap(ErrInvalidValue, "next_issuer_id is zero")
	}

	if idc.NextProjectId == 0 {
		return errors.Wrap(ErrInvalidValue, "next_project_id is zero")
	}

	if idc.NextApplicantId == 0 {
		return errors.Wrap(ErrInvalidValue, "next_collector_id is zero")
	}

	return nil
}

func (is Issuer) Validate() error {
	if is.Id == 0 {
		return errors.Wrap(ErrInvalidValue, "id is zero")
	}

	if is.Name == "" {
		return errors.Wrap(ErrInvalidValue, "name is empty")
	}

	if is.Admin == "" {
		return errors.Wrap(ErrInvalidValue, "admin is empty")
	}

	if _, err := sdk.AccAddressFromBech32(is.Admin); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	return nil
}

func (is Issuer) AddressHasAuthorization(addr sdk.AccAddress) bool {
	return is.Admin == addr.String()
}

func (a Applicant) Validate() error {
	if a.Id == 0 {
		return errors.Wrap(ErrInvalidValue, "id is zero")
	}

	if a.Name == "" {
		return errors.Wrap(ErrInvalidValue, "name is empty")
	}

	if a.Admin == "" {
		return errors.Wrap(ErrInvalidValue, "admin is empty")
	}

	if _, err := sdk.AccAddressFromBech32(a.Admin); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	return nil
}

func (cc CreditClass) Validate() error {
	if cc.Abbreviation == "" {
		return errors.Wrap(ErrInvalidValue, "abbreviation is empty")
	}

	if cc.IssuerId == 0 {
		return errors.Wrap(ErrInvalidValue, "issuer_id is zero")
	}

	if cc.Name == "" {
		return errors.Wrap(ErrInvalidValue, "name is empty")
	}

	return nil
}

func (cc CreditCollection) Validate() error {
	if cc.Denom == "" {
		return errors.Wrap(ErrInvalidValue, "denom is empty")
	}
	if cc.ProjectId == 0 {
		return errors.Wrap(ErrInvalidValue, "project id is empty or zero")
	}
	if cc.TotalAmount.Active == 0 && cc.TotalAmount.Retired == 0 {
		return errors.Wrap(ErrInvalidValue, "cannot issue 0 credits")
	}
	for _, data := range cc.CreditData {
		if data.Uri == "" {
			return errors.Wrap(ErrInvalidValue, "empty credit data uri")
		}
		if data.Hash == "" {
			return errors.Wrap(ErrInvalidValue, "empty credit data hash")
		}
	}
	return nil
}

func (icb CreditBalance) Validate() error {
	if _, err := sdk.AccAddressFromBech32(icb.Owner); err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid credit owner address (%s)", err)
	}
	if icb.Denom == "" {
		return errors.Wrap(ErrInvalidValue, "denom is empty")
	}
	return nil
}
