// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message_MessageType int32

const (
	Message_PING               Message_MessageType = 0
	Message_CHAT               Message_MessageType = 1
	Message_FOLLOW             Message_MessageType = 2
	Message_UNFOLLOW           Message_MessageType = 3
	Message_ORDER              Message_MessageType = 4
	Message_ORDER_REJECT       Message_MessageType = 5
	Message_ORDER_CANCEL       Message_MessageType = 6
	Message_ORDER_CONFIRMATION Message_MessageType = 7
	Message_ORDER_FULFILLMENT  Message_MessageType = 8
	Message_ORDER_COMPLETION   Message_MessageType = 9
	Message_DISPUTE_OPEN       Message_MessageType = 10
	Message_DISPUTE_UPDATE     Message_MessageType = 11
	Message_DISPUTE_CLOSE      Message_MessageType = 12
	Message_REFUND             Message_MessageType = 13
	Message_OFFLINE_ACK        Message_MessageType = 14
	Message_OFFLINE_RELAY      Message_MessageType = 15
	Message_MODERATOR_ADD      Message_MessageType = 16
	Message_MODERATOR_REMOVE   Message_MessageType = 17
	Message_STORE              Message_MessageType = 18
	Message_BLOCK              Message_MessageType = 19
	Message_ERROR              Message_MessageType = 500
)

var Message_MessageType_name = map[int32]string{
	0:   "PING",
	1:   "CHAT",
	2:   "FOLLOW",
	3:   "UNFOLLOW",
	4:   "ORDER",
	5:   "ORDER_REJECT",
	6:   "ORDER_CANCEL",
	7:   "ORDER_CONFIRMATION",
	8:   "ORDER_FULFILLMENT",
	9:   "ORDER_COMPLETION",
	10:  "DISPUTE_OPEN",
	11:  "DISPUTE_UPDATE",
	12:  "DISPUTE_CLOSE",
	13:  "REFUND",
	14:  "OFFLINE_ACK",
	15:  "OFFLINE_RELAY",
	16:  "MODERATOR_ADD",
	17:  "MODERATOR_REMOVE",
	18:  "STORE",
	19:  "BLOCK",
	500: "ERROR",
}
var Message_MessageType_value = map[string]int32{
	"PING":               0,
	"CHAT":               1,
	"FOLLOW":             2,
	"UNFOLLOW":           3,
	"ORDER":              4,
	"ORDER_REJECT":       5,
	"ORDER_CANCEL":       6,
	"ORDER_CONFIRMATION": 7,
	"ORDER_FULFILLMENT":  8,
	"ORDER_COMPLETION":   9,
	"DISPUTE_OPEN":       10,
	"DISPUTE_UPDATE":     11,
	"DISPUTE_CLOSE":      12,
	"REFUND":             13,
	"OFFLINE_ACK":        14,
	"OFFLINE_RELAY":      15,
	"MODERATOR_ADD":      16,
	"MODERATOR_REMOVE":   17,
	"STORE":              18,
	"BLOCK":              19,
	"ERROR":              500,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type Chat_Flag int32

const (
	Chat_MESSAGE Chat_Flag = 0
	Chat_TYPING  Chat_Flag = 1
	Chat_READ    Chat_Flag = 2
)

var Chat_Flag_name = map[int32]string{
	0: "MESSAGE",
	1: "TYPING",
	2: "READ",
}
var Chat_Flag_value = map[string]int32{
	"MESSAGE": 0,
	"TYPING":  1,
	"READ":    2,
}

func (x Chat_Flag) String() string {
	return proto.EnumName(Chat_Flag_name, int32(x))
}
func (Chat_Flag) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2, 0} }

type Message struct {
	MessageType Message_MessageType   `protobuf:"varint,1,opt,name=messageType,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload     *google_protobuf1.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	RequestId   int32                 `protobuf:"varint,3,opt,name=requestId" json:"requestId,omitempty"`
	IsResponse  bool                  `protobuf:"varint,4,opt,name=isResponse" json:"isResponse,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_PING
}

func (m *Message) GetPayload() *google_protobuf1.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequestId() int32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *Message) GetIsResponse() bool {
	if m != nil {
		return m.IsResponse
	}
	return false
}

type Envelope struct {
	Message   *Message `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Pubkey    []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Chat struct {
	MessageId string                     `protobuf:"bytes,1,opt,name=messageId" json:"messageId,omitempty"`
	Subject   string                     `protobuf:"bytes,2,opt,name=subject" json:"subject,omitempty"`
	Message   string                     `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Flag      Chat_Flag                  `protobuf:"varint,5,opt,name=flag,enum=Chat_Flag" json:"flag,omitempty"`
}

func (m *Chat) Reset()                    { *m = Chat{} }
func (m *Chat) String() string            { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()               {}
func (*Chat) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Chat) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Chat) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Chat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Chat) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Chat) GetFlag() Chat_Flag {
	if m != nil {
		return m.Flag
	}
	return Chat_MESSAGE
}

type SignedData struct {
	SenderPubkey   []byte `protobuf:"bytes,1,opt,name=senderPubkey,proto3" json:"senderPubkey,omitempty"`
	SerializedData []byte `protobuf:"bytes,2,opt,name=serializedData,proto3" json:"serializedData,omitempty"`
	Signature      []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedData) Reset()                    { *m = SignedData{} }
func (m *SignedData) String() string            { return proto.CompactTextString(m) }
func (*SignedData) ProtoMessage()               {}
func (*SignedData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *SignedData) GetSenderPubkey() []byte {
	if m != nil {
		return m.SenderPubkey
	}
	return nil
}

func (m *SignedData) GetSerializedData() []byte {
	if m != nil {
		return m.SerializedData
	}
	return nil
}

func (m *SignedData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignedData_Command struct {
	PeerID    string                     `protobuf:"bytes,1,opt,name=peerID" json:"peerID,omitempty"`
	Type      Message_MessageType        `protobuf:"varint,2,opt,name=type,enum=Message_MessageType" json:"type,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *SignedData_Command) Reset()                    { *m = SignedData_Command{} }
func (m *SignedData_Command) String() string            { return proto.CompactTextString(m) }
func (*SignedData_Command) ProtoMessage()               {}
func (*SignedData_Command) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3, 0} }

func (m *SignedData_Command) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *SignedData_Command) GetType() Message_MessageType {
	if m != nil {
		return m.Type
	}
	return Message_PING
}

func (m *SignedData_Command) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type CidList struct {
	Cids []string `protobuf:"bytes,1,rep,name=cids" json:"cids,omitempty"`
}

func (m *CidList) Reset()                    { *m = CidList{} }
func (m *CidList) String() string            { return proto.CompactTextString(m) }
func (*CidList) ProtoMessage()               {}
func (*CidList) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *CidList) GetCids() []string {
	if m != nil {
		return m.Cids
	}
	return nil
}

type Block struct {
	RawData []byte `protobuf:"bytes,1,opt,name=rawData,proto3" json:"rawData,omitempty"`
	Cid     string `protobuf:"bytes,2,opt,name=cid" json:"cid,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *Block) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func (m *Block) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterType((*Chat)(nil), "Chat")
	proto.RegisterType((*SignedData)(nil), "SignedData")
	proto.RegisterType((*SignedData_Command)(nil), "SignedData.Command")
	proto.RegisterType((*CidList)(nil), "CidList")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("Chat_Flag", Chat_Flag_name, Chat_Flag_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xf3, 0x44,
	0x10, 0xad, 0x13, 0x27, 0x4e, 0xc6, 0x69, 0xba, 0x5d, 0x4a, 0x15, 0x2a, 0x28, 0x91, 0x0f, 0x28,
	0x5c, 0x5c, 0x29, 0x95, 0x10, 0x57, 0xd7, 0x5e, 0x17, 0x53, 0xc7, 0x8e, 0x36, 0x0e, 0xa8, 0x5c,
	0x22, 0x27, 0xde, 0x06, 0xd3, 0xc4, 0x36, 0xb1, 0x03, 0x0a, 0x77, 0x7e, 0x02, 0x3f, 0x8f, 0x13,
	0x7f, 0x81, 0x33, 0x42, 0xbb, 0xb6, 0x49, 0x5b, 0xa4, 0x4a, 0xdf, 0x29, 0x33, 0x6f, 0x5e, 0x66,
	0xc6, 0x6f, 0xde, 0xc2, 0xe9, 0x96, 0xe5, 0x79, 0xb8, 0x66, 0x7a, 0xb6, 0x4b, 0x8b, 0xf4, 0xea,
	0x93, 0x75, 0x9a, 0xae, 0x37, 0xec, 0x46, 0x64, 0xcb, 0xfd, 0xd3, 0x4d, 0x98, 0x1c, 0xaa, 0xd2,
	0xe7, 0x6f, 0x4b, 0x45, 0xbc, 0x65, 0x79, 0x11, 0x6e, 0xb3, 0x92, 0xa0, 0xfd, 0x21, 0x83, 0x32,
	0x29, 0xbb, 0xe1, 0xaf, 0x40, 0xad, 0x1a, 0x07, 0x87, 0x8c, 0x0d, 0xa4, 0xa1, 0x34, 0xea, 0x8f,
	0x2f, 0xf4, 0xaa, 0x5c, 0xff, 0xf2, 0x1a, 0x7d, 0x49, 0xc4, 0x3a, 0x28, 0x59, 0x78, 0xd8, 0xa4,
	0x61, 0x34, 0x68, 0x0c, 0xa5, 0x91, 0x3a, 0xbe, 0xd0, 0xcb, 0xb1, 0x7a, 0x3d, 0x56, 0x37, 0x92,
	0x03, 0xad, 0x49, 0xf8, 0x53, 0xe8, 0xee, 0xd8, 0xcf, 0x7b, 0x96, 0x17, 0x4e, 0x34, 0x68, 0x0e,
	0xa5, 0x51, 0x8b, 0x1e, 0x01, 0x7c, 0x0d, 0x10, 0xe7, 0x94, 0xe5, 0x59, 0x9a, 0xe4, 0x6c, 0x20,
	0x0f, 0xa5, 0x51, 0x87, 0xbe, 0x40, 0xb4, 0xbf, 0x1a, 0xa0, 0xbe, 0x58, 0x05, 0x77, 0x40, 0x9e,
	0x3a, 0xde, 0x3d, 0x3a, 0xe1, 0x91, 0xf9, 0x8d, 0x11, 0x20, 0x09, 0x03, 0xb4, 0x6d, 0xdf, 0x75,
	0xfd, 0xef, 0x51, 0x03, 0xf7, 0xa0, 0x33, 0xf7, 0xaa, 0xac, 0x89, 0xbb, 0xd0, 0xf2, 0xa9, 0x45,
	0x28, 0x92, 0x31, 0x82, 0x9e, 0x08, 0x17, 0x94, 0x7c, 0x4b, 0xcc, 0x00, 0xb5, 0x8e, 0x88, 0x69,
	0x78, 0x26, 0x71, 0x51, 0x1b, 0x5f, 0x02, 0xae, 0x10, 0xdf, 0xb3, 0x1d, 0x3a, 0x31, 0x02, 0xc7,
	0xf7, 0x90, 0x82, 0x3f, 0x86, 0xf3, 0x12, 0xb7, 0xe7, 0xae, 0xed, 0xb8, 0xee, 0x84, 0x78, 0x01,
	0xea, 0xe0, 0x0b, 0x40, 0x35, 0x7d, 0x32, 0x75, 0x89, 0x20, 0x77, 0x79, 0x5b, 0xcb, 0x99, 0x4d,
	0xe7, 0x01, 0x59, 0xf8, 0x53, 0xe2, 0x21, 0xc0, 0x18, 0xfa, 0x35, 0x32, 0x9f, 0x5a, 0x46, 0x40,
	0x90, 0x8a, 0xcf, 0xe1, 0xb4, 0xc6, 0x4c, 0xd7, 0x9f, 0x11, 0xd4, 0xe3, 0x9f, 0x41, 0x89, 0x3d,
	0xf7, 0x2c, 0x74, 0x8a, 0xcf, 0x40, 0xf5, 0x6d, 0xdb, 0x75, 0x3c, 0xb2, 0x30, 0xcc, 0x07, 0xd4,
	0xe7, 0xfc, 0x1a, 0xa0, 0xc4, 0x35, 0x1e, 0xd1, 0x19, 0x87, 0x26, 0xbe, 0x45, 0xa8, 0x11, 0xf8,
	0x74, 0x61, 0x58, 0x16, 0x42, 0x7c, 0xa3, 0x23, 0x44, 0xc9, 0xc4, 0xff, 0x8e, 0xa0, 0x73, 0xae,
	0xc2, 0x2c, 0xf0, 0x29, 0x41, 0x98, 0x87, 0x77, 0xae, 0x6f, 0x3e, 0xa0, 0x8f, 0x30, 0x40, 0x8b,
	0x50, 0xea, 0x53, 0xf4, 0x77, 0x53, 0x8b, 0xa0, 0x43, 0x92, 0x5f, 0xd8, 0x26, 0xcd, 0x18, 0xd6,
	0x40, 0xa9, 0xce, 0x2d, 0x3c, 0xa1, 0x8e, 0x3b, 0xb5, 0x17, 0x68, 0x5d, 0xc0, 0x97, 0xd0, 0xce,
	0xf6, 0xcb, 0x67, 0x76, 0x10, 0x16, 0xe8, 0xd1, 0x2a, 0xe3, 0xb7, 0xce, 0xe3, 0x75, 0x12, 0x16,
	0xfb, 0x1d, 0x13, 0xb7, 0xee, 0xd1, 0x23, 0xa0, 0xfd, 0x29, 0x81, 0x6c, 0xfe, 0x18, 0x16, 0x9c,
	0x56, 0x75, 0x72, 0x22, 0x31, 0xa4, 0x4b, 0x8f, 0x00, 0x1e, 0x80, 0x92, 0xef, 0x97, 0x3f, 0xb1,
	0x55, 0x21, 0xba, 0x77, 0x69, 0x9d, 0xf2, 0x4a, 0xbd, 0x5a, 0xb3, 0xac, 0xd4, 0x0b, 0x7d, 0x0d,
	0xdd, 0xff, 0xbc, 0x2e, 0x5c, 0xa4, 0x8e, 0xaf, 0xfe, 0x67, 0xcb, 0xa0, 0x66, 0xd0, 0x23, 0x19,
	0x5f, 0x83, 0xfc, 0xb4, 0x09, 0xd7, 0x83, 0x96, 0xf0, 0x3f, 0xe8, 0x7c, 0x41, 0xdd, 0xde, 0x84,
	0x6b, 0x2a, 0x70, 0xed, 0x4b, 0x90, 0x79, 0x86, 0x55, 0x50, 0x26, 0x64, 0x36, 0x33, 0xee, 0x09,
	0x3a, 0xe1, 0xa7, 0x0a, 0x1e, 0x85, 0x0f, 0x25, 0xee, 0x43, 0x4a, 0x0c, 0x0b, 0x35, 0xb4, 0x7f,
	0x24, 0x80, 0x59, 0xbc, 0x4e, 0x58, 0x64, 0x85, 0x45, 0x88, 0x35, 0xe8, 0xe5, 0x2c, 0x89, 0xd8,
	0x6e, 0x5a, 0x4a, 0x25, 0x09, 0x3d, 0x5e, 0x61, 0xf8, 0x0b, 0xe8, 0xe7, 0x6c, 0x17, 0x87, 0x9b,
	0xf8, 0xb7, 0xf2, 0x5f, 0x95, 0xa0, 0x6f, 0xd0, 0xf7, 0x85, 0xbd, 0xfa, 0x5d, 0x02, 0xc5, 0x4c,
	0xb7, 0xdb, 0x30, 0x89, 0xc4, 0x69, 0x18, 0xdb, 0x39, 0x56, 0x25, 0x6c, 0x95, 0xe1, 0x11, 0xc8,
	0x05, 0x7f, 0xe7, 0x8d, 0x77, 0xde, 0xb9, 0x60, 0xbc, 0xd6, 0xb2, 0xf9, 0x01, 0x5a, 0x6a, 0x9f,
	0x81, 0x62, 0xc6, 0x91, 0x1b, 0xe7, 0x05, 0xc6, 0x20, 0xaf, 0xe2, 0x28, 0x1f, 0x48, 0xc3, 0xe6,
	0xa8, 0x4b, 0x45, 0xac, 0xdd, 0x42, 0xeb, 0x6e, 0x93, 0xae, 0x9e, 0xf9, 0x1d, 0x77, 0xe1, 0xaf,
	0xe2, 0x73, 0x4b, 0x51, 0xea, 0x14, 0x23, 0x68, 0xae, 0xe2, 0xa8, 0xba, 0x3b, 0x0f, 0xef, 0xe4,
	0x1f, 0x1a, 0xd9, 0x72, 0xd9, 0x16, 0x83, 0x6f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xc7,
	0x6a, 0x70, 0x0c, 0x05, 0x00, 0x00,
}
