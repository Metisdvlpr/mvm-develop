// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages-common.proto

package trezor

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Failure_FailureType int32

const (
	Failure_Failure_UnexpectedMessage Failure_FailureType = 1
	Failure_Failure_ButtonExpected    Failure_FailureType = 2
	Failure_Failure_DataError         Failure_FailureType = 3
	Failure_Failure_ActionCancelled   Failure_FailureType = 4
	Failure_Failure_PinExpected       Failure_FailureType = 5
	Failure_Failure_PinCancelled      Failure_FailureType = 6
	Failure_Failure_PinInvalid        Failure_FailureType = 7
	Failure_Failure_InvalidSignature  Failure_FailureType = 8
	Failure_Failure_ProcessError      Failure_FailureType = 9
	Failure_Failure_NotEnoughFunds    Failure_FailureType = 10
	Failure_Failure_NotInitialized    Failure_FailureType = 11
	Failure_Failure_PinMismatch       Failure_FailureType = 12
	Failure_Failure_FirmwareError     Failure_FailureType = 99
)

var Failure_FailureType_name = map[int32]string{
	1:  "Failure_UnexpectedMessage",
	2:  "Failure_ButtonExpected",
	3:  "Failure_DataError",
	4:  "Failure_ActionCancelled",
	5:  "Failure_PinExpected",
	6:  "Failure_PinCancelled",
	7:  "Failure_PinInvalid",
	8:  "Failure_InvalidSignature",
	9:  "Failure_ProcessError",
	10: "Failure_NotEnoughFunds",
	11: "Failure_NotInitialized",
	12: "Failure_PinMismatch",
	99: "Failure_FirmwareError",
}

var Failure_FailureType_value = map[string]int32{
	"Failure_UnexpectedMessage": 1,
	"Failure_ButtonExpected":    2,
	"Failure_DataError":         3,
	"Failure_ActionCancelled":   4,
	"Failure_PinExpected":       5,
	"Failure_PinCancelled":      6,
	"Failure_PinInvalid":        7,
	"Failure_InvalidSignature":  8,
	"Failure_ProcessError":      9,
	"Failure_NotEnoughFunds":    10,
	"Failure_NotInitialized":    11,
	"Failure_PinMismatch":       12,
	"Failure_FirmwareError":     99,
}

func (x Failure_FailureType) Enum() *Failure_FailureType {
	p := new(Failure_FailureType)
	*p = x
	return p
}

func (x Failure_FailureType) String() string {
	return proto.EnumName(Failure_FailureType_name, int32(x))
}

func (x *Failure_FailureType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Failure_FailureType_value, data, "Failure_FailureType")
	if err != nil {
		return err
	}
	*x = Failure_FailureType(value)
	return nil
}

func (Failure_FailureType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{1, 0}
}

//*
// Type of button request
type ButtonRequest_ButtonRequestType int32

const (
	ButtonRequest_ButtonRequest_Other                 ButtonRequest_ButtonRequestType = 1
	ButtonRequest_ButtonRequest_FeeOverThreshold      ButtonRequest_ButtonRequestType = 2
	ButtonRequest_ButtonRequest_ConfirmOutput         ButtonRequest_ButtonRequestType = 3
	ButtonRequest_ButtonRequest_ResetDevice           ButtonRequest_ButtonRequestType = 4
	ButtonRequest_ButtonRequest_ConfirmWord           ButtonRequest_ButtonRequestType = 5
	ButtonRequest_ButtonRequest_WipeDevice            ButtonRequest_ButtonRequestType = 6
	ButtonRequest_ButtonRequest_ProtectCall           ButtonRequest_ButtonRequestType = 7
	ButtonRequest_ButtonRequest_SignTx                ButtonRequest_ButtonRequestType = 8
	ButtonRequest_ButtonRequest_FirmwareCheck         ButtonRequest_ButtonRequestType = 9
	ButtonRequest_ButtonRequest_Address               ButtonRequest_ButtonRequestType = 10
	ButtonRequest_ButtonRequest_PublicKey             ButtonRequest_ButtonRequestType = 11
	ButtonRequest_ButtonRequest_MnemonicWordCount     ButtonRequest_ButtonRequestType = 12
	ButtonRequest_ButtonRequest_MnemonicInput         ButtonRequest_ButtonRequestType = 13
	ButtonRequest_ButtonRequest_PassphraseType        ButtonRequest_ButtonRequestType = 14
	ButtonRequest_ButtonRequest_UnknownDerivationPath ButtonRequest_ButtonRequestType = 15
)

var ButtonRequest_ButtonRequestType_name = map[int32]string{
	1:  "ButtonRequest_Other",
	2:  "ButtonRequest_FeeOverThreshold",
	3:  "ButtonRequest_ConfirmOutput",
	4:  "ButtonRequest_ResetDevice",
	5:  "ButtonRequest_ConfirmWord",
	6:  "ButtonRequest_WipeDevice",
	7:  "ButtonRequest_ProtectCall",
	8:  "ButtonRequest_SignTx",
	9:  "ButtonRequest_FirmwareCheck",
	10: "ButtonRequest_Address",
	11: "ButtonRequest_PublicKey",
	12: "ButtonRequest_MnemonicWordCount",
	13: "ButtonRequest_MnemonicInput",
	14: "ButtonRequest_PassphraseType",
	15: "ButtonRequest_UnknownDerivationPath",
}

var ButtonRequest_ButtonRequestType_value = map[string]int32{
	"ButtonRequest_Other":                 1,
	"ButtonRequest_FeeOverThreshold":      2,
	"ButtonRequest_ConfirmOutput":         3,
	"ButtonRequest_ResetDevice":           4,
	"ButtonRequest_ConfirmWord":           5,
	"ButtonRequest_WipeDevice":            6,
	"ButtonRequest_ProtectCall":           7,
	"ButtonRequest_SignTx":                8,
	"ButtonRequest_FirmwareCheck":         9,
	"ButtonRequest_Address":               10,
	"ButtonRequest_PublicKey":             11,
	"ButtonRequest_MnemonicWordCount":     12,
	"ButtonRequest_MnemonicInput":         13,
	"ButtonRequest_PassphraseType":        14,
	"ButtonRequest_UnknownDerivationPath": 15,
}

func (x ButtonRequest_ButtonRequestType) Enum() *ButtonRequest_ButtonRequestType {
	p := new(ButtonRequest_ButtonRequestType)
	*p = x
	return p
}

func (x ButtonRequest_ButtonRequestType) String() string {
	return proto.EnumName(ButtonRequest_ButtonRequestType_name, int32(x))
}

func (x *ButtonRequest_ButtonRequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ButtonRequest_ButtonRequestType_value, data, "ButtonRequest_ButtonRequestType")
	if err != nil {
		return err
	}
	*x = ButtonRequest_ButtonRequestType(value)
	return nil
}

func (ButtonRequest_ButtonRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{2, 0}
}

//*
// Type of PIN request
type PinMatrixRequest_PinMatrixRequestType int32

const (
	PinMatrixRequest_PinMatrixRequestType_Current   PinMatrixRequest_PinMatrixRequestType = 1
	PinMatrixRequest_PinMatrixRequestType_NewFirst  PinMatrixRequest_PinMatrixRequestType = 2
	PinMatrixRequest_PinMatrixRequestType_NewSecond PinMatrixRequest_PinMatrixRequestType = 3
)

var PinMatrixRequest_PinMatrixRequestType_name = map[int32]string{
	1: "PinMatrixRequestType_Current",
	2: "PinMatrixRequestType_NewFirst",
	3: "PinMatrixRequestType_NewSecond",
}

var PinMatrixRequest_PinMatrixRequestType_value = map[string]int32{
	"PinMatrixRequestType_Current":   1,
	"PinMatrixRequestType_NewFirst":  2,
	"PinMatrixRequestType_NewSecond": 3,
}

func (x PinMatrixRequest_PinMatrixRequestType) Enum() *PinMatrixRequest_PinMatrixRequestType {
	p := new(PinMatrixRequest_PinMatrixRequestType)
	*p = x
	return p
}

func (x PinMatrixRequest_PinMatrixRequestType) String() string {
	return proto.EnumName(PinMatrixRequest_PinMatrixRequestType_name, int32(x))
}

func (x *PinMatrixRequest_PinMatrixRequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PinMatrixRequest_PinMatrixRequestType_value, data, "PinMatrixRequest_PinMatrixRequestType")
	if err != nil {
		return err
	}
	*x = PinMatrixRequest_PinMatrixRequestType(value)
	return nil
}

func (PinMatrixRequest_PinMatrixRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{4, 0}
}

//*
// Response: Success of the previous request
// @end
type Success struct {
	Message              *string  `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Success) Reset()         { *m = Success{} }
func (m *Success) String() string { return proto.CompactTextString(m) }
func (*Success) ProtoMessage()    {}
func (*Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{0}
}

func (m *Success) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Success.Unmarshal(m, b)
}
func (m *Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Success.Marshal(b, m, deterministic)
}
func (m *Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Success.Merge(m, src)
}
func (m *Success) XXX_Size() int {
	return xxx_messageInfo_Success.Size(m)
}
func (m *Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Success proto.InternalMessageInfo

func (m *Success) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

//*
// Response: Failure of the previous request
// @end
type Failure struct {
	Code                 *Failure_FailureType `protobuf:"varint,1,opt,name=code,enum=hw.trezor.messages.common.Failure_FailureType" json:"code,omitempty"`
	Message              *string              `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Failure) Reset()         { *m = Failure{} }
func (m *Failure) String() string { return proto.CompactTextString(m) }
func (*Failure) ProtoMessage()    {}
func (*Failure) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{1}
}

func (m *Failure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Failure.Unmarshal(m, b)
}
func (m *Failure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Failure.Marshal(b, m, deterministic)
}
func (m *Failure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Failure.Merge(m, src)
}
func (m *Failure) XXX_Size() int {
	return xxx_messageInfo_Failure.Size(m)
}
func (m *Failure) XXX_DiscardUnknown() {
	xxx_messageInfo_Failure.DiscardUnknown(m)
}

var xxx_messageInfo_Failure proto.InternalMessageInfo

func (m *Failure) GetCode() Failure_FailureType {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return Failure_Failure_UnexpectedMessage
}

func (m *Failure) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

//*
// Response: Device is waiting for HW button press.
// @auxstart
// @next ButtonAck
type ButtonRequest struct {
	Code                 *ButtonRequest_ButtonRequestType `protobuf:"varint,1,opt,name=code,enum=hw.trezor.messages.common.ButtonRequest_ButtonRequestType" json:"code,omitempty"`
	Data                 *string                          `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ButtonRequest) Reset()         { *m = ButtonRequest{} }
func (m *ButtonRequest) String() string { return proto.CompactTextString(m) }
func (*ButtonRequest) ProtoMessage()    {}
func (*ButtonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{2}
}

func (m *ButtonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButtonRequest.Unmarshal(m, b)
}
func (m *ButtonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButtonRequest.Marshal(b, m, deterministic)
}
func (m *ButtonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButtonRequest.Merge(m, src)
}
func (m *ButtonRequest) XXX_Size() int {
	return xxx_messageInfo_ButtonRequest.Size(m)
}
func (m *ButtonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ButtonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ButtonRequest proto.InternalMessageInfo

func (m *ButtonRequest) GetCode() ButtonRequest_ButtonRequestType {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ButtonRequest_ButtonRequest_Other
}

func (m *ButtonRequest) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

//*
// Request: Computer agrees to wait for HW button press
// @auxend
type ButtonAck struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ButtonAck) Reset()         { *m = ButtonAck{} }
func (m *ButtonAck) String() string { return proto.CompactTextString(m) }
func (*ButtonAck) ProtoMessage()    {}
func (*ButtonAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{3}
}

func (m *ButtonAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ButtonAck.Unmarshal(m, b)
}
func (m *ButtonAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ButtonAck.Marshal(b, m, deterministic)
}
func (m *ButtonAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ButtonAck.Merge(m, src)
}
func (m *ButtonAck) XXX_Size() int {
	return xxx_messageInfo_ButtonAck.Size(m)
}
func (m *ButtonAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ButtonAck.DiscardUnknown(m)
}

var xxx_messageInfo_ButtonAck proto.InternalMessageInfo

//*
// Response: Device is asking computer to show PIN matrix and awaits PIN encoded using this matrix scheme
// @auxstart
// @next PinMatrixAck
type PinMatrixRequest struct {
	Type                 *PinMatrixRequest_PinMatrixRequestType `protobuf:"varint,1,opt,name=type,enum=hw.trezor.messages.common.PinMatrixRequest_PinMatrixRequestType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *PinMatrixRequest) Reset()         { *m = PinMatrixRequest{} }
func (m *PinMatrixRequest) String() string { return proto.CompactTextString(m) }
func (*PinMatrixRequest) ProtoMessage()    {}
func (*PinMatrixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{4}
}

func (m *PinMatrixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PinMatrixRequest.Unmarshal(m, b)
}
func (m *PinMatrixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PinMatrixRequest.Marshal(b, m, deterministic)
}
func (m *PinMatrixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PinMatrixRequest.Merge(m, src)
}
func (m *PinMatrixRequest) XXX_Size() int {
	return xxx_messageInfo_PinMatrixRequest.Size(m)
}
func (m *PinMatrixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PinMatrixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PinMatrixRequest proto.InternalMessageInfo

func (m *PinMatrixRequest) GetType() PinMatrixRequest_PinMatrixRequestType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PinMatrixRequest_PinMatrixRequestType_Current
}

//*
// Request: Computer responds with encoded PIN
// @auxend
type PinMatrixAck struct {
	Pin                  *string  `protobuf:"bytes,1,req,name=pin" json:"pin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PinMatrixAck) Reset()         { *m = PinMatrixAck{} }
func (m *PinMatrixAck) String() string { return proto.CompactTextString(m) }
func (*PinMatrixAck) ProtoMessage()    {}
func (*PinMatrixAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{5}
}

func (m *PinMatrixAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PinMatrixAck.Unmarshal(m, b)
}
func (m *PinMatrixAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PinMatrixAck.Marshal(b, m, deterministic)
}
func (m *PinMatrixAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PinMatrixAck.Merge(m, src)
}
func (m *PinMatrixAck) XXX_Size() int {
	return xxx_messageInfo_PinMatrixAck.Size(m)
}
func (m *PinMatrixAck) XXX_DiscardUnknown() {
	xxx_messageInfo_PinMatrixAck.DiscardUnknown(m)
}

var xxx_messageInfo_PinMatrixAck proto.InternalMessageInfo

func (m *PinMatrixAck) GetPin() string {
	if m != nil && m.Pin != nil {
		return *m.Pin
	}
	return ""
}

//*
// Response: Device awaits encryption passphrase
// @auxstart
// @next PassphraseAck
type PassphraseRequest struct {
	OnDevice             *bool    `protobuf:"varint,1,opt,name=on_device,json=onDevice" json:"on_device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassphraseRequest) Reset()         { *m = PassphraseRequest{} }
func (m *PassphraseRequest) String() string { return proto.CompactTextString(m) }
func (*PassphraseRequest) ProtoMessage()    {}
func (*PassphraseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{6}
}

func (m *PassphraseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassphraseRequest.Unmarshal(m, b)
}
func (m *PassphraseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassphraseRequest.Marshal(b, m, deterministic)
}
func (m *PassphraseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassphraseRequest.Merge(m, src)
}
func (m *PassphraseRequest) XXX_Size() int {
	return xxx_messageInfo_PassphraseRequest.Size(m)
}
func (m *PassphraseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PassphraseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PassphraseRequest proto.InternalMessageInfo

func (m *PassphraseRequest) GetOnDevice() bool {
	if m != nil && m.OnDevice != nil {
		return *m.OnDevice
	}
	return false
}

//*
// Request: Send passphrase back
// @next PassphraseStateRequest
type PassphraseAck struct {
	Passphrase           *string  `protobuf:"bytes,1,opt,name=passphrase" json:"passphrase,omitempty"`
	State                []byte   `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassphraseAck) Reset()         { *m = PassphraseAck{} }
func (m *PassphraseAck) String() string { return proto.CompactTextString(m) }
func (*PassphraseAck) ProtoMessage()    {}
func (*PassphraseAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{7}
}

func (m *PassphraseAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassphraseAck.Unmarshal(m, b)
}
func (m *PassphraseAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassphraseAck.Marshal(b, m, deterministic)
}
func (m *PassphraseAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassphraseAck.Merge(m, src)
}
func (m *PassphraseAck) XXX_Size() int {
	return xxx_messageInfo_PassphraseAck.Size(m)
}
func (m *PassphraseAck) XXX_DiscardUnknown() {
	xxx_messageInfo_PassphraseAck.DiscardUnknown(m)
}

var xxx_messageInfo_PassphraseAck proto.InternalMessageInfo

func (m *PassphraseAck) GetPassphrase() string {
	if m != nil && m.Passphrase != nil {
		return *m.Passphrase
	}
	return ""
}

func (m *PassphraseAck) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

//*
// Response: Device awaits passphrase state
// @next PassphraseStateAck
type PassphraseStateRequest struct {
	State                []byte   `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassphraseStateRequest) Reset()         { *m = PassphraseStateRequest{} }
func (m *PassphraseStateRequest) String() string { return proto.CompactTextString(m) }
func (*PassphraseStateRequest) ProtoMessage()    {}
func (*PassphraseStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{8}
}

func (m *PassphraseStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassphraseStateRequest.Unmarshal(m, b)
}
func (m *PassphraseStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassphraseStateRequest.Marshal(b, m, deterministic)
}
func (m *PassphraseStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassphraseStateRequest.Merge(m, src)
}
func (m *PassphraseStateRequest) XXX_Size() int {
	return xxx_messageInfo_PassphraseStateRequest.Size(m)
}
func (m *PassphraseStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PassphraseStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PassphraseStateRequest proto.InternalMessageInfo

func (m *PassphraseStateRequest) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

//*
// Request: Send passphrase state back
// @auxend
type PassphraseStateAck struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassphraseStateAck) Reset()         { *m = PassphraseStateAck{} }
func (m *PassphraseStateAck) String() string { return proto.CompactTextString(m) }
func (*PassphraseStateAck) ProtoMessage()    {}
func (*PassphraseStateAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{9}
}

func (m *PassphraseStateAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassphraseStateAck.Unmarshal(m, b)
}
func (m *PassphraseStateAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassphraseStateAck.Marshal(b, m, deterministic)
}
func (m *PassphraseStateAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassphraseStateAck.Merge(m, src)
}
func (m *PassphraseStateAck) XXX_Size() int {
	return xxx_messageInfo_PassphraseStateAck.Size(m)
}
func (m *PassphraseStateAck) XXX_DiscardUnknown() {
	xxx_messageInfo_PassphraseStateAck.DiscardUnknown(m)
}

var xxx_messageInfo_PassphraseStateAck proto.InternalMessageInfo

//*
// Structure representing BIP32 (hierarchical deterministic) node
// Used for imports of private key into the device and exporting public key out of device
// @embed
type HDNodeType struct {
	Depth                *uint32  `protobuf:"varint,1,req,name=depth" json:"depth,omitempty"`
	Fingerprint          *uint32  `protobuf:"varint,2,req,name=fingerprint" json:"fingerprint,omitempty"`
	ChildNum             *uint32  `protobuf:"varint,3,req,name=child_num,json=childNum" json:"child_num,omitempty"`
	ChainCode            []byte   `protobuf:"bytes,4,req,name=chain_code,json=chainCode" json:"chain_code,omitempty"`
	PrivateKey           []byte   `protobuf:"bytes,5,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,6,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HDNodeType) Reset()         { *m = HDNodeType{} }
func (m *HDNodeType) String() string { return proto.CompactTextString(m) }
func (*HDNodeType) ProtoMessage()    {}
func (*HDNodeType) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaf30d059fdbc38d, []int{10}
}

func (m *HDNodeType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HDNodeType.Unmarshal(m, b)
}
func (m *HDNodeType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HDNodeType.Marshal(b, m, deterministic)
}
func (m *HDNodeType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HDNodeType.Merge(m, src)
}
func (m *HDNodeType) XXX_Size() int {
	return xxx_messageInfo_HDNodeType.Size(m)
}
func (m *HDNodeType) XXX_DiscardUnknown() {
	xxx_messageInfo_HDNodeType.DiscardUnknown(m)
}

var xxx_messageInfo_HDNodeType proto.InternalMessageInfo

func (m *HDNodeType) GetDepth() uint32 {
	if m != nil && m.Depth != nil {
		return *m.Depth
	}
	return 0
}

func (m *HDNodeType) GetFingerprint() uint32 {
	if m != nil && m.Fingerprint != nil {
		return *m.Fingerprint
	}
	return 0
}

func (m *HDNodeType) GetChildNum() uint32 {
	if m != nil && m.ChildNum != nil {
		return *m.ChildNum
	}
	return 0
}

func (m *HDNodeType) GetChainCode() []byte {
	if m != nil {
		return m.ChainCode
	}
	return nil
}

func (m *HDNodeType) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *HDNodeType) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("hw.trezor.messages.common.Failure_FailureType", Failure_FailureType_name, Failure_FailureType_value)
	proto.RegisterEnum("hw.trezor.messages.common.ButtonRequest_ButtonRequestType", ButtonRequest_ButtonRequestType_name, ButtonRequest_ButtonRequestType_value)
	proto.RegisterEnum("hw.trezor.messages.common.PinMatrixRequest_PinMatrixRequestType", PinMatrixRequest_PinMatrixRequestType_name, PinMatrixRequest_PinMatrixRequestType_value)
	proto.RegisterType((*Success)(nil), "hw.trezor.messages.common.Success")
	proto.RegisterType((*Failure)(nil), "hw.trezor.messages.common.Failure")
	proto.RegisterType((*ButtonRequest)(nil), "hw.trezor.messages.common.ButtonRequest")
	proto.RegisterType((*ButtonAck)(nil), "hw.trezor.messages.common.ButtonAck")
	proto.RegisterType((*PinMatrixRequest)(nil), "hw.trezor.messages.common.PinMatrixRequest")
	proto.RegisterType((*PinMatrixAck)(nil), "hw.trezor.messages.common.PinMatrixAck")
	proto.RegisterType((*PassphraseRequest)(nil), "hw.trezor.messages.common.PassphraseRequest")
	proto.RegisterType((*PassphraseAck)(nil), "hw.trezor.messages.common.PassphraseAck")
	proto.RegisterType((*PassphraseStateRequest)(nil), "hw.trezor.messages.common.PassphraseStateRequest")
	proto.RegisterType((*PassphraseStateAck)(nil), "hw.trezor.messages.common.PassphraseStateAck")
	proto.RegisterType((*HDNodeType)(nil), "hw.trezor.messages.common.HDNodeType")
}

func init() { proto.RegisterFile("messages-common.proto", fileDescriptor_aaf30d059fdbc38d) }

var fileDescriptor_aaf30d059fdbc38d = []byte{
	// 846 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x52, 0x23, 0x37,
	0x10, 0x2e, 0xff, 0x80, 0xed, 0xb6, 0xd9, 0x08, 0xc5, 0x80, 0x09, 0xb0, 0x38, 0xc3, 0x21, 0x5c,
	0xe2, 0x4a, 0xe5, 0x98, 0x53, 0x58, 0x83, 0x2b, 0xd4, 0x16, 0x86, 0x1a, 0xd8, 0xda, 0xa3, 0x4b,
	0xd1, 0xf4, 0x32, 0x2a, 0xcf, 0x48, 0x13, 0x8d, 0x06, 0xf0, 0x5e, 0xf2, 0x6a, 0x79, 0x89, 0xbc,
	0x42, 0xaa, 0x52, 0xb9, 0xe4, 0x11, 0xb6, 0x34, 0x3f, 0x78, 0xc6, 0x66, 0x39, 0xcd, 0xe8, 0xfb,
	0xbe, 0xee, 0x96, 0xba, 0x3f, 0x09, 0x76, 0x42, 0x8c, 0x63, 0x76, 0x8f, 0xf1, 0x8f, 0x5c, 0x85,
	0xa1, 0x92, 0xa3, 0x48, 0x2b, 0xa3, 0xe8, 0xbe, 0xff, 0x38, 0x32, 0x1a, 0x3f, 0x2b, 0x3d, 0x2a,
	0x04, 0xa3, 0x4c, 0xe0, 0x9c, 0x40, 0xeb, 0x36, 0xe1, 0x1c, 0xe3, 0x98, 0x0e, 0xa0, 0x95, 0xb3,
	0x83, 0xda, 0xb0, 0x76, 0xda, 0x71, 0x8b, 0xa5, 0xf3, 0x77, 0x03, 0x5a, 0x13, 0x26, 0x82, 0x44,
	0x23, 0x7d, 0x07, 0x4d, 0xae, 0xbc, 0x4c, 0xf2, 0xe6, 0xe7, 0xd1, 0xe8, 0xab, 0xa9, 0x47, 0x79,
	0x44, 0xf1, 0xbd, 0x5b, 0x44, 0xe8, 0xa6, 0xb1, 0xe5, 0x4a, 0xf5, 0x6a, 0xa5, 0xff, 0xea, 0xd0,
	0x2d, 0xe9, 0xe9, 0x11, 0xec, 0xe7, 0xcb, 0xd9, 0x07, 0x89, 0x4f, 0x11, 0x72, 0x83, 0xde, 0x55,
	0x26, 0x26, 0x35, 0xfa, 0x1d, 0xec, 0x16, 0xf4, 0xbb, 0xc4, 0x18, 0x25, 0x2f, 0x72, 0x09, 0xa9,
	0xd3, 0x1d, 0xd8, 0x2e, 0xb8, 0x73, 0x66, 0xd8, 0x85, 0xd6, 0x4a, 0x93, 0x06, 0x3d, 0x80, 0xbd,
	0x02, 0x3e, 0xe3, 0x46, 0x28, 0x39, 0x66, 0x92, 0x63, 0x10, 0xa0, 0x47, 0x9a, 0x74, 0x0f, 0xbe,
	0x2d, 0xc8, 0x1b, 0xb1, 0x4c, 0xb6, 0x41, 0x07, 0xd0, 0x2f, 0x11, 0xcb, 0x90, 0x4d, 0xba, 0x0b,
	0xb4, 0xc4, 0x5c, 0xca, 0x07, 0x16, 0x08, 0x8f, 0xb4, 0xe8, 0x21, 0x0c, 0x0a, 0x3c, 0x07, 0x6f,
	0xc5, 0xbd, 0x64, 0x26, 0xd1, 0x48, 0xda, 0x95, 0x7c, 0x5a, 0xd9, 0xf6, 0x67, 0xfb, 0xeb, 0x94,
	0x8f, 0x34, 0x55, 0xe6, 0x42, 0xaa, 0xe4, 0xde, 0x9f, 0x24, 0xd2, 0x8b, 0x09, 0xac, 0x70, 0x97,
	0x52, 0x18, 0xc1, 0x02, 0xf1, 0x19, 0x3d, 0xd2, 0x5d, 0xd9, 0xfa, 0x95, 0x88, 0x43, 0x66, 0xb8,
	0x4f, 0x7a, 0x74, 0x1f, 0x76, 0x0a, 0x62, 0x22, 0x74, 0xf8, 0xc8, 0x34, 0x66, 0xb5, 0xb8, 0xf3,
	0x4f, 0x13, 0xb6, 0xb2, 0xbe, 0xb9, 0xf8, 0x47, 0x82, 0xb1, 0xa1, 0xd3, 0xca, 0x74, 0x7f, 0x79,
	0x65, 0xba, 0x95, 0xb8, 0xea, 0xaa, 0x34, 0x69, 0x0a, 0x4d, 0x8f, 0x19, 0x96, 0x8f, 0x39, 0xfd,
	0x77, 0xfe, 0x6f, 0xc0, 0xf6, 0x9a, 0xde, 0xee, 0xbf, 0x02, 0xce, 0xae, 0x8d, 0x8f, 0x9a, 0xd4,
	0xa8, 0x03, 0x6f, 0xab, 0xc4, 0x04, 0xf1, 0xfa, 0x01, 0xf5, 0x9d, 0xaf, 0x31, 0xf6, 0x55, 0x60,
	0x67, 0x7d, 0x0c, 0x07, 0x55, 0xcd, 0x58, 0xc9, 0x4f, 0x42, 0x87, 0xd7, 0x89, 0x89, 0x12, 0x43,
	0x1a, 0xd6, 0x47, 0x55, 0x81, 0x8b, 0x31, 0x9a, 0x73, 0x7c, 0x10, 0x1c, 0x49, 0x73, 0x9d, 0xce,
	0xe3, 0x3f, 0x2a, 0x6d, 0xa7, 0x7f, 0x08, 0x83, 0x2a, 0xfd, 0x51, 0x44, 0x98, 0x07, 0x6f, 0xae,
	0x07, 0xdf, 0x68, 0x65, 0x90, 0x9b, 0x31, 0x0b, 0x02, 0xd2, 0xb2, 0xa3, 0xae, 0xd2, 0xd6, 0x07,
	0x77, 0x4f, 0xa4, 0xbd, 0xbe, 0xeb, 0x62, 0x3e, 0x63, 0x1f, 0xf9, 0x9c, 0x74, 0xec, 0xe8, 0xaa,
	0x82, 0x33, 0xcf, 0xd3, 0x18, 0x5b, 0x2b, 0x1c, 0xc0, 0xde, 0x4a, 0xd1, 0xe4, 0xf7, 0x40, 0xf0,
	0xf7, 0xb8, 0x20, 0x5d, 0x7a, 0x02, 0xc7, 0x55, 0xf2, 0x4a, 0x62, 0xa8, 0xa4, 0xe0, 0xf6, 0x3c,
	0x63, 0x95, 0x48, 0x43, 0x7a, 0xeb, 0xd5, 0x0b, 0xd1, 0xa5, 0xb4, 0x3d, 0xdb, 0xa2, 0x43, 0x38,
	0x5c, 0x29, 0xc1, 0xe2, 0x38, 0xf2, 0x35, 0x8b, 0xd3, 0xbb, 0x49, 0xde, 0xd0, 0x1f, 0xe0, 0xa4,
	0xaa, 0xf8, 0x20, 0xe7, 0x52, 0x3d, 0xca, 0x73, 0xd4, 0xe2, 0x81, 0xd9, 0xcb, 0x75, 0xc3, 0x8c,
	0x4f, 0xbe, 0x71, 0xba, 0xd0, 0xc9, 0x84, 0x67, 0x7c, 0xee, 0xfc, 0x5b, 0x03, 0x62, 0x2d, 0xca,
	0x8c, 0x16, 0x4f, 0x85, 0xf1, 0xee, 0xa0, 0x69, 0x16, 0x51, 0x61, 0xbc, 0x5f, 0x5f, 0x31, 0xde,
	0x6a, 0xe8, 0x1a, 0x90, 0xd9, 0xcf, 0x66, 0x73, 0xfe, 0x84, 0xfe, 0x4b, 0xac, 0x3d, 0xda, 0x4b,
	0xf8, 0x6c, 0x9c, 0x68, 0x8d, 0xd2, 0x90, 0x1a, 0xfd, 0x1e, 0x8e, 0x5e, 0x54, 0x4c, 0xf1, 0x71,
	0x22, 0x74, 0x6c, 0x48, 0xdd, 0x1a, 0xf3, 0x6b, 0x92, 0x5b, 0xe4, 0x4a, 0x7a, 0xa4, 0xe1, 0x0c,
	0xa1, 0xf7, 0xac, 0x39, 0xe3, 0x73, 0x4a, 0xa0, 0x11, 0x09, 0x39, 0xa8, 0x0d, 0xeb, 0xa7, 0x1d,
	0xd7, 0xfe, 0x3a, 0x3f, 0xc1, 0xf6, 0xb2, 0xaf, 0x45, 0x37, 0x0e, 0xa0, 0xa3, 0xe4, 0xcc, 0x4b,
	0x1d, 0x96, 0xb6, 0xa4, 0xed, 0xb6, 0x95, 0xcc, 0x1c, 0xe7, 0x5c, 0xc0, 0xd6, 0x32, 0xc2, 0x26,
	0x7d, 0x0b, 0x10, 0x3d, 0x03, 0xf9, 0xdb, 0x5d, 0x42, 0x68, 0x1f, 0x36, 0x62, 0xc3, 0x4c, 0xf6,
	0xd8, 0xf6, 0xdc, 0x6c, 0xe1, 0x8c, 0x60, 0x77, 0x99, 0xe6, 0xd6, 0x42, 0x45, 0xf5, 0x67, 0x7d,
	0xad, 0xac, 0xef, 0x03, 0x5d, 0xd1, 0xdb, 0x61, 0xfe, 0x55, 0x03, 0xf8, 0xed, 0x7c, 0xaa, 0xbc,
	0xec, 0xbd, 0xee, 0xc3, 0x86, 0x87, 0x91, 0xf1, 0xd3, 0x13, 0x6e, 0xb9, 0xd9, 0x82, 0x0e, 0xa1,
	0xfb, 0x49, 0xc8, 0x7b, 0xd4, 0x91, 0x16, 0xd2, 0x0c, 0xea, 0x29, 0x57, 0x86, 0xec, 0x81, 0xb9,
	0x2f, 0x02, 0x6f, 0x26, 0x93, 0x70, 0xd0, 0x48, 0xf9, 0x76, 0x0a, 0x4c, 0x93, 0x90, 0x1e, 0x01,
	0x70, 0x9f, 0x09, 0x39, 0x4b, 0x9f, 0xa6, 0xe6, 0xb0, 0x7e, 0xda, 0x73, 0x3b, 0x29, 0x32, 0xb6,
	0x6f, 0xcc, 0x31, 0x74, 0xa3, 0xd4, 0x6f, 0x38, 0x9b, 0xe3, 0x62, 0xb0, 0x91, 0x6e, 0x1a, 0x72,
	0xe8, 0x3d, 0x2e, 0x6c, 0x7c, 0x94, 0xde, 0x8e, 0x94, 0xdf, 0x4c, 0xf9, 0x4e, 0x54, 0xdc, 0x97,
	0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x7d, 0x20, 0xa6, 0x35, 0x07, 0x00, 0x00,
}
