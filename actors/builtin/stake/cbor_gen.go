// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package stake

import (
	"fmt"
	"io"

	abi "github.com/filecoin-project/go-state-types/abi"
	stake "github.com/filecoin-project/specs-actors/v2/actors/builtin/stake"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

var lengthBufState = []byte{144}

func (t *State) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufState); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.RootKey (address.Address) (struct)
	if err := t.RootKey.MarshalCBOR(w); err != nil {
		return err
	}

	// t.TotalStakePower (big.Int) (struct)
	if err := t.TotalStakePower.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MaturePeriod (abi.ChainEpoch) (int64)
	if t.MaturePeriod >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.MaturePeriod)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.MaturePeriod-1)); err != nil {
			return err
		}
	}

	// t.RoundPeriod (abi.ChainEpoch) (int64)
	if t.RoundPeriod >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.RoundPeriod)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.RoundPeriod-1)); err != nil {
			return err
		}
	}

	// t.PrincipalLockDuration (abi.ChainEpoch) (int64)
	if t.PrincipalLockDuration >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.PrincipalLockDuration)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.PrincipalLockDuration-1)); err != nil {
			return err
		}
	}

	// t.MinDepositAmount (big.Int) (struct)
	if err := t.MinDepositAmount.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MaxRewardPerRound (big.Int) (struct)
	if err := t.MaxRewardPerRound.MarshalCBOR(w); err != nil {
		return err
	}

	// t.InflationFactor (big.Int) (struct)
	if err := t.InflationFactor.MarshalCBOR(w); err != nil {
		return err
	}

	// t.LastRoundReward (big.Int) (struct)
	if err := t.LastRoundReward.MarshalCBOR(w); err != nil {
		return err
	}

	// t.StakePeriodStart (abi.ChainEpoch) (int64)
	if t.StakePeriodStart >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.StakePeriodStart)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.StakePeriodStart-1)); err != nil {
			return err
		}
	}

	// t.NextRoundEpoch (abi.ChainEpoch) (int64)
	if t.NextRoundEpoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.NextRoundEpoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.NextRoundEpoch-1)); err != nil {
			return err
		}
	}

	// t.LockedPrincipalMap (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.LockedPrincipalMap); err != nil {
		return xerrors.Errorf("failed to write cid field t.LockedPrincipalMap: %w", err)
	}

	// t.AvailablePrincipalMap (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.AvailablePrincipalMap); err != nil {
		return xerrors.Errorf("failed to write cid field t.AvailablePrincipalMap: %w", err)
	}

	// t.StakePowerMap (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.StakePowerMap); err != nil {
		return xerrors.Errorf("failed to write cid field t.StakePowerMap: %w", err)
	}

	// t.VestingRewardMap (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.VestingRewardMap); err != nil {
		return xerrors.Errorf("failed to write cid field t.VestingRewardMap: %w", err)
	}

	// t.AvailableRewardMap (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.AvailableRewardMap); err != nil {
		return xerrors.Errorf("failed to write cid field t.AvailableRewardMap: %w", err)
	}

	return nil
}

func (t *State) UnmarshalCBOR(r io.Reader) error {
	*t = State{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 16 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.RootKey (address.Address) (struct)

	{

		if err := t.RootKey.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.RootKey: %w", err)
		}

	}
	// t.TotalStakePower (big.Int) (struct)

	{

		if err := t.TotalStakePower.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalStakePower: %w", err)
		}

	}
	// t.MaturePeriod (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.MaturePeriod = abi.ChainEpoch(extraI)
	}
	// t.RoundPeriod (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.RoundPeriod = abi.ChainEpoch(extraI)
	}
	// t.PrincipalLockDuration (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.PrincipalLockDuration = abi.ChainEpoch(extraI)
	}
	// t.MinDepositAmount (big.Int) (struct)

	{

		if err := t.MinDepositAmount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinDepositAmount: %w", err)
		}

	}
	// t.MaxRewardPerRound (big.Int) (struct)

	{

		if err := t.MaxRewardPerRound.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MaxRewardPerRound: %w", err)
		}

	}
	// t.InflationFactor (big.Int) (struct)

	{

		if err := t.InflationFactor.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.InflationFactor: %w", err)
		}

	}
	// t.LastRoundReward (big.Int) (struct)

	{

		if err := t.LastRoundReward.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.LastRoundReward: %w", err)
		}

	}
	// t.StakePeriodStart (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.StakePeriodStart = abi.ChainEpoch(extraI)
	}
	// t.NextRoundEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.NextRoundEpoch = abi.ChainEpoch(extraI)
	}
	// t.LockedPrincipalMap (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.LockedPrincipalMap: %w", err)
		}

		t.LockedPrincipalMap = c

	}
	// t.AvailablePrincipalMap (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.AvailablePrincipalMap: %w", err)
		}

		t.AvailablePrincipalMap = c

	}
	// t.StakePowerMap (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.StakePowerMap: %w", err)
		}

		t.StakePowerMap = c

	}
	// t.VestingRewardMap (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.VestingRewardMap: %w", err)
		}

		t.VestingRewardMap = c

	}
	// t.AvailableRewardMap (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.AvailableRewardMap: %w", err)
		}

		t.AvailableRewardMap = c

	}
	return nil
}

var lengthBufLockedPrincipals = []byte{129}

func (t *LockedPrincipals) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufLockedPrincipals); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Data ([]stake.LockedPrincipal) (slice)
	if len(t.Data) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Data was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Data))); err != nil {
		return err
	}
	for _, v := range t.Data {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *LockedPrincipals) UnmarshalCBOR(r io.Reader) error {
	*t = LockedPrincipals{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Data ([]stake.LockedPrincipal) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Data: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Data = make([]LockedPrincipal, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v LockedPrincipal
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Data[i] = v
	}

	return nil
}

var lengthBufLockedPrincipal = []byte{130}

func (t *LockedPrincipal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufLockedPrincipal); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Epoch (abi.ChainEpoch) (int64)
	if t.Epoch >= 0 {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Epoch)); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajNegativeInt, uint64(-t.Epoch-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *LockedPrincipal) UnmarshalCBOR(r io.Reader) error {
	*t = LockedPrincipal{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Amount: %w", err)
		}

	}
	// t.Epoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Epoch = abi.ChainEpoch(extraI)
	}
	return nil
}

var lengthBufVestingFunds = []byte{129}

func (t *VestingFunds) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufVestingFunds); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Funds ([]stake.VestingFund) (slice)
	if len(t.Funds) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Funds was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajArray, uint64(len(t.Funds))); err != nil {
		return err
	}
	for _, v := range t.Funds {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *VestingFunds) UnmarshalCBOR(r io.Reader) error {
	*t = VestingFunds{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Funds ([]stake.VestingFund) (slice)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Funds: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Funds = make([]stake.VestingFund, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v stake.VestingFund
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Funds[i] = v
	}

	return nil
}
