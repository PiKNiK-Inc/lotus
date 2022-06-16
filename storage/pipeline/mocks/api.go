// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/lotus/storage/pipeline (interfaces: SealingAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"

	address "github.com/filecoin-project/go-address"
	bitfield "github.com/filecoin-project/go-bitfield"
	abi "github.com/filecoin-project/go-state-types/abi"
	big "github.com/filecoin-project/go-state-types/big"
	market "github.com/filecoin-project/go-state-types/builtin/v8/market"
	miner "github.com/filecoin-project/go-state-types/builtin/v8/miner"
	crypto "github.com/filecoin-project/go-state-types/crypto"
	dline "github.com/filecoin-project/go-state-types/dline"
	network "github.com/filecoin-project/go-state-types/network"

	api "github.com/filecoin-project/lotus/api"
	types "github.com/filecoin-project/lotus/chain/types"
	sealing "github.com/filecoin-project/lotus/storage/pipeline"
)

// MockSealingAPI is a mock of SealingAPI interface.
type MockSealingAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSealingAPIMockRecorder
}

// MockSealingAPIMockRecorder is the mock recorder for MockSealingAPI.
type MockSealingAPIMockRecorder struct {
	mock *MockSealingAPI
}

// NewMockSealingAPI creates a new mock instance.
func NewMockSealingAPI(ctrl *gomock.Controller) *MockSealingAPI {
	mock := &MockSealingAPI{ctrl: ctrl}
	mock.recorder = &MockSealingAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSealingAPI) EXPECT() *MockSealingAPIMockRecorder {
	return m.recorder
}

// ChainBaseFee mocks base method.
func (m *MockSealingAPI) ChainBaseFee(arg0 context.Context, arg1 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainBaseFee", arg0, arg1)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainBaseFee indicates an expected call of ChainBaseFee.
func (mr *MockSealingAPIMockRecorder) ChainBaseFee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainBaseFee", reflect.TypeOf((*MockSealingAPI)(nil).ChainBaseFee), arg0, arg1)
}

// ChainGetMessage mocks base method.
func (m *MockSealingAPI) ChainGetMessage(arg0 context.Context, arg1 cid.Cid) (*types.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainGetMessage", arg0, arg1)
	ret0, _ := ret[0].(*types.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainGetMessage indicates an expected call of ChainGetMessage.
func (mr *MockSealingAPIMockRecorder) ChainGetMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainGetMessage", reflect.TypeOf((*MockSealingAPI)(nil).ChainGetMessage), arg0, arg1)
}

// ChainHead mocks base method.
func (m *MockSealingAPI) ChainHead(arg0 context.Context) (types.TipSetKey, abi.ChainEpoch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainHead", arg0)
	ret0, _ := ret[0].(types.TipSetKey)
	ret1, _ := ret[1].(abi.ChainEpoch)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChainHead indicates an expected call of ChainHead.
func (mr *MockSealingAPIMockRecorder) ChainHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainHead", reflect.TypeOf((*MockSealingAPI)(nil).ChainHead), arg0)
}

// ChainReadObj mocks base method.
func (m *MockSealingAPI) ChainReadObj(arg0 context.Context, arg1 cid.Cid) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainReadObj", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainReadObj indicates an expected call of ChainReadObj.
func (mr *MockSealingAPIMockRecorder) ChainReadObj(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainReadObj", reflect.TypeOf((*MockSealingAPI)(nil).ChainReadObj), arg0, arg1)
}

// SendMsg mocks base method.
func (m *MockSealingAPI) SendMsg(arg0 context.Context, arg1, arg2 address.Address, arg3 abi.MethodNum, arg4, arg5 big.Int, arg6 []byte) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSealingAPIMockRecorder) SendMsg(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSealingAPI)(nil).SendMsg), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// StateComputeDataCommitment mocks base method.
func (m *MockSealingAPI) StateComputeDataCommitment(arg0 context.Context, arg1 address.Address, arg2 abi.RegisteredSealProof, arg3 []abi.DealID, arg4 types.TipSetKey) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateComputeDataCommitment", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateComputeDataCommitment indicates an expected call of StateComputeDataCommitment.
func (mr *MockSealingAPIMockRecorder) StateComputeDataCommitment(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateComputeDataCommitment", reflect.TypeOf((*MockSealingAPI)(nil).StateComputeDataCommitment), arg0, arg1, arg2, arg3, arg4)
}

// StateGetRandomnessFromBeacon mocks base method.
func (m *MockSealingAPI) StateGetRandomnessFromBeacon(arg0 context.Context, arg1 crypto.DomainSeparationTag, arg2 abi.ChainEpoch, arg3 []byte, arg4 types.TipSetKey) (abi.Randomness, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateGetRandomnessFromBeacon", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(abi.Randomness)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateGetRandomnessFromBeacon indicates an expected call of StateGetRandomnessFromBeacon.
func (mr *MockSealingAPIMockRecorder) StateGetRandomnessFromBeacon(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateGetRandomnessFromBeacon", reflect.TypeOf((*MockSealingAPI)(nil).StateGetRandomnessFromBeacon), arg0, arg1, arg2, arg3, arg4)
}

// StateGetRandomnessFromTickets mocks base method.
func (m *MockSealingAPI) StateGetRandomnessFromTickets(arg0 context.Context, arg1 crypto.DomainSeparationTag, arg2 abi.ChainEpoch, arg3 []byte, arg4 types.TipSetKey) (abi.Randomness, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateGetRandomnessFromTickets", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(abi.Randomness)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateGetRandomnessFromTickets indicates an expected call of StateGetRandomnessFromTickets.
func (mr *MockSealingAPIMockRecorder) StateGetRandomnessFromTickets(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateGetRandomnessFromTickets", reflect.TypeOf((*MockSealingAPI)(nil).StateGetRandomnessFromTickets), arg0, arg1, arg2, arg3, arg4)
}

// StateLookupID mocks base method.
func (m *MockSealingAPI) StateLookupID(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateLookupID", arg0, arg1, arg2)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateLookupID indicates an expected call of StateLookupID.
func (mr *MockSealingAPIMockRecorder) StateLookupID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateLookupID", reflect.TypeOf((*MockSealingAPI)(nil).StateLookupID), arg0, arg1, arg2)
}

// StateMarketStorageDeal mocks base method.
func (m *MockSealingAPI) StateMarketStorageDeal(arg0 context.Context, arg1 abi.DealID, arg2 types.TipSetKey) (*api.MarketDeal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMarketStorageDeal", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.MarketDeal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMarketStorageDeal indicates an expected call of StateMarketStorageDeal.
func (mr *MockSealingAPIMockRecorder) StateMarketStorageDeal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMarketStorageDeal", reflect.TypeOf((*MockSealingAPI)(nil).StateMarketStorageDeal), arg0, arg1, arg2)
}

// StateMarketStorageDealProposal mocks base method.
func (m *MockSealingAPI) StateMarketStorageDealProposal(arg0 context.Context, arg1 abi.DealID, arg2 types.TipSetKey) (market.DealProposal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMarketStorageDealProposal", arg0, arg1, arg2)
	ret0, _ := ret[0].(market.DealProposal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMarketStorageDealProposal indicates an expected call of StateMarketStorageDealProposal.
func (mr *MockSealingAPIMockRecorder) StateMarketStorageDealProposal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMarketStorageDealProposal", reflect.TypeOf((*MockSealingAPI)(nil).StateMarketStorageDealProposal), arg0, arg1, arg2)
}

// StateMinerActiveSectors mocks base method.
func (m *MockSealingAPI) StateMinerActiveSectors(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (bitfield.BitField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerActiveSectors", arg0, arg1, arg2)
	ret0, _ := ret[0].(bitfield.BitField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerActiveSectors indicates an expected call of StateMinerActiveSectors.
func (mr *MockSealingAPIMockRecorder) StateMinerActiveSectors(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerActiveSectors", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerActiveSectors), arg0, arg1, arg2)
}

// StateMinerAvailableBalance mocks base method.
func (m *MockSealingAPI) StateMinerAvailableBalance(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerAvailableBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerAvailableBalance indicates an expected call of StateMinerAvailableBalance.
func (mr *MockSealingAPIMockRecorder) StateMinerAvailableBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerAvailableBalance", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerAvailableBalance), arg0, arg1, arg2)
}

// StateMinerInfo mocks base method.
func (m *MockSealingAPI) StateMinerInfo(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (api.MinerInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.MinerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerInfo indicates an expected call of StateMinerInfo.
func (mr *MockSealingAPIMockRecorder) StateMinerInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerInfo", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerInfo), arg0, arg1, arg2)
}

// StateMinerInitialPledgeCollateral mocks base method.
func (m *MockSealingAPI) StateMinerInitialPledgeCollateral(arg0 context.Context, arg1 address.Address, arg2 miner.SectorPreCommitInfo, arg3 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerInitialPledgeCollateral", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerInitialPledgeCollateral indicates an expected call of StateMinerInitialPledgeCollateral.
func (mr *MockSealingAPIMockRecorder) StateMinerInitialPledgeCollateral(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerInitialPledgeCollateral", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerInitialPledgeCollateral), arg0, arg1, arg2, arg3)
}

// StateMinerPartitions mocks base method.
func (m *MockSealingAPI) StateMinerPartitions(arg0 context.Context, arg1 address.Address, arg2 uint64, arg3 types.TipSetKey) ([]api.Partition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerPartitions", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]api.Partition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerPartitions indicates an expected call of StateMinerPartitions.
func (mr *MockSealingAPIMockRecorder) StateMinerPartitions(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerPartitions", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerPartitions), arg0, arg1, arg2, arg3)
}

// StateMinerPreCommitDepositForPower mocks base method.
func (m *MockSealingAPI) StateMinerPreCommitDepositForPower(arg0 context.Context, arg1 address.Address, arg2 miner.SectorPreCommitInfo, arg3 types.TipSetKey) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerPreCommitDepositForPower", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerPreCommitDepositForPower indicates an expected call of StateMinerPreCommitDepositForPower.
func (mr *MockSealingAPIMockRecorder) StateMinerPreCommitDepositForPower(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerPreCommitDepositForPower", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerPreCommitDepositForPower), arg0, arg1, arg2, arg3)
}

// StateMinerProvingDeadline mocks base method.
func (m *MockSealingAPI) StateMinerProvingDeadline(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (*dline.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerProvingDeadline", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dline.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerProvingDeadline indicates an expected call of StateMinerProvingDeadline.
func (mr *MockSealingAPIMockRecorder) StateMinerProvingDeadline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerProvingDeadline", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerProvingDeadline), arg0, arg1, arg2)
}

// StateMinerSectorAllocated mocks base method.
func (m *MockSealingAPI) StateMinerSectorAllocated(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerSectorAllocated", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerSectorAllocated indicates an expected call of StateMinerSectorAllocated.
func (mr *MockSealingAPIMockRecorder) StateMinerSectorAllocated(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerSectorAllocated", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerSectorAllocated), arg0, arg1, arg2, arg3)
}

// StateMinerSectorSize mocks base method.
func (m *MockSealingAPI) StateMinerSectorSize(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (abi.SectorSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerSectorSize", arg0, arg1, arg2)
	ret0, _ := ret[0].(abi.SectorSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerSectorSize indicates an expected call of StateMinerSectorSize.
func (mr *MockSealingAPIMockRecorder) StateMinerSectorSize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerSectorSize", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerSectorSize), arg0, arg1, arg2)
}

// StateMinerWorkerAddress mocks base method.
func (m *MockSealingAPI) StateMinerWorkerAddress(arg0 context.Context, arg1 address.Address, arg2 types.TipSetKey) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateMinerWorkerAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateMinerWorkerAddress indicates an expected call of StateMinerWorkerAddress.
func (mr *MockSealingAPIMockRecorder) StateMinerWorkerAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateMinerWorkerAddress", reflect.TypeOf((*MockSealingAPI)(nil).StateMinerWorkerAddress), arg0, arg1, arg2)
}

// StateNetworkVersion mocks base method.
func (m *MockSealingAPI) StateNetworkVersion(arg0 context.Context, arg1 types.TipSetKey) (network.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateNetworkVersion", arg0, arg1)
	ret0, _ := ret[0].(network.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateNetworkVersion indicates an expected call of StateNetworkVersion.
func (mr *MockSealingAPIMockRecorder) StateNetworkVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateNetworkVersion", reflect.TypeOf((*MockSealingAPI)(nil).StateNetworkVersion), arg0, arg1)
}

// StateSearchMsg mocks base method.
func (m *MockSealingAPI) StateSearchMsg(arg0 context.Context, arg1 cid.Cid) (*sealing.MsgLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSearchMsg", arg0, arg1)
	ret0, _ := ret[0].(*sealing.MsgLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSearchMsg indicates an expected call of StateSearchMsg.
func (mr *MockSealingAPIMockRecorder) StateSearchMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSearchMsg", reflect.TypeOf((*MockSealingAPI)(nil).StateSearchMsg), arg0, arg1)
}

// StateSectorGetInfo mocks base method.
func (m *MockSealingAPI) StateSectorGetInfo(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (*miner.SectorOnChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorGetInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*miner.SectorOnChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorGetInfo indicates an expected call of StateSectorGetInfo.
func (mr *MockSealingAPIMockRecorder) StateSectorGetInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorGetInfo", reflect.TypeOf((*MockSealingAPI)(nil).StateSectorGetInfo), arg0, arg1, arg2, arg3)
}

// StateSectorPartition mocks base method.
func (m *MockSealingAPI) StateSectorPartition(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (*sealing.SectorLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorPartition", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*sealing.SectorLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorPartition indicates an expected call of StateSectorPartition.
func (mr *MockSealingAPIMockRecorder) StateSectorPartition(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorPartition", reflect.TypeOf((*MockSealingAPI)(nil).StateSectorPartition), arg0, arg1, arg2, arg3)
}

// StateSectorPreCommitInfo mocks base method.
func (m *MockSealingAPI) StateSectorPreCommitInfo(arg0 context.Context, arg1 address.Address, arg2 abi.SectorNumber, arg3 types.TipSetKey) (*miner.SectorPreCommitOnChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSectorPreCommitInfo", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*miner.SectorPreCommitOnChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateSectorPreCommitInfo indicates an expected call of StateSectorPreCommitInfo.
func (mr *MockSealingAPIMockRecorder) StateSectorPreCommitInfo(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSectorPreCommitInfo", reflect.TypeOf((*MockSealingAPI)(nil).StateSectorPreCommitInfo), arg0, arg1, arg2, arg3)
}

// StateWaitMsg mocks base method.
func (m *MockSealingAPI) StateWaitMsg(arg0 context.Context, arg1 cid.Cid) (sealing.MsgLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateWaitMsg", arg0, arg1)
	ret0, _ := ret[0].(sealing.MsgLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateWaitMsg indicates an expected call of StateWaitMsg.
func (mr *MockSealingAPIMockRecorder) StateWaitMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateWaitMsg", reflect.TypeOf((*MockSealingAPI)(nil).StateWaitMsg), arg0, arg1)
}
