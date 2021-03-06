// Code generated by protoc-gen-go.
// source: responses.proto
// DO NOT EDIT!

package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Response that contains a result set.
type ResultSetResponse struct {
	ConnectionId string     `protobuf:"bytes,1,opt,name=connection_id,json=connectionId" json:"connection_id,omitempty"`
	StatementId  uint32     `protobuf:"varint,2,opt,name=statement_id,json=statementId" json:"statement_id,omitempty"`
	OwnStatement bool       `protobuf:"varint,3,opt,name=own_statement,json=ownStatement" json:"own_statement,omitempty"`
	Signature    *Signature `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	FirstFrame   *Frame     `protobuf:"bytes,5,opt,name=first_frame,json=firstFrame" json:"first_frame,omitempty"`
	UpdateCount  uint64     `protobuf:"varint,6,opt,name=update_count,json=updateCount" json:"update_count,omitempty"`
	// with no signature nor other data.
	Metadata *RpcMetadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *ResultSetResponse) Reset()                    { *m = ResultSetResponse{} }
func (m *ResultSetResponse) String() string            { return proto.CompactTextString(m) }
func (*ResultSetResponse) ProtoMessage()               {}
func (*ResultSetResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ResultSetResponse) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *ResultSetResponse) GetFirstFrame() *Frame {
	if m != nil {
		return m.FirstFrame
	}
	return nil
}

func (m *ResultSetResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to PrepareAndExecuteRequest
type ExecuteResponse struct {
	Results          []*ResultSetResponse `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	MissingStatement bool                 `protobuf:"varint,2,opt,name=missing_statement,json=missingStatement" json:"missing_statement,omitempty"`
	Metadata         *RpcMetadata         `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *ExecuteResponse) Reset()                    { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()               {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ExecuteResponse) GetResults() []*ResultSetResponse {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ExecuteResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to PrepareRequest
type PrepareResponse struct {
	Statement *StatementHandle `protobuf:"bytes,1,opt,name=statement" json:"statement,omitempty"`
	Metadata  *RpcMetadata     `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *PrepareResponse) Reset()                    { *m = PrepareResponse{} }
func (m *PrepareResponse) String() string            { return proto.CompactTextString(m) }
func (*PrepareResponse) ProtoMessage()               {}
func (*PrepareResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *PrepareResponse) GetStatement() *StatementHandle {
	if m != nil {
		return m.Statement
	}
	return nil
}

func (m *PrepareResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to FetchRequest
type FetchResponse struct {
	Frame            *Frame       `protobuf:"bytes,1,opt,name=frame" json:"frame,omitempty"`
	MissingStatement bool         `protobuf:"varint,2,opt,name=missing_statement,json=missingStatement" json:"missing_statement,omitempty"`
	MissingResults   bool         `protobuf:"varint,3,opt,name=missing_results,json=missingResults" json:"missing_results,omitempty"`
	Metadata         *RpcMetadata `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *FetchResponse) Reset()                    { *m = FetchResponse{} }
func (m *FetchResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchResponse) ProtoMessage()               {}
func (*FetchResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *FetchResponse) GetFrame() *Frame {
	if m != nil {
		return m.Frame
	}
	return nil
}

func (m *FetchResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to CreateStatementRequest
type CreateStatementResponse struct {
	ConnectionId string       `protobuf:"bytes,1,opt,name=connection_id,json=connectionId" json:"connection_id,omitempty"`
	StatementId  uint32       `protobuf:"varint,2,opt,name=statement_id,json=statementId" json:"statement_id,omitempty"`
	Metadata     *RpcMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *CreateStatementResponse) Reset()                    { *m = CreateStatementResponse{} }
func (m *CreateStatementResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateStatementResponse) ProtoMessage()               {}
func (*CreateStatementResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *CreateStatementResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to CloseStatementRequest
type CloseStatementResponse struct {
	Metadata *RpcMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *CloseStatementResponse) Reset()                    { *m = CloseStatementResponse{} }
func (m *CloseStatementResponse) String() string            { return proto.CompactTextString(m) }
func (*CloseStatementResponse) ProtoMessage()               {}
func (*CloseStatementResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *CloseStatementResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to OpenConnectionRequest {
type OpenConnectionResponse struct {
	Metadata *RpcMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *OpenConnectionResponse) Reset()                    { *m = OpenConnectionResponse{} }
func (m *OpenConnectionResponse) String() string            { return proto.CompactTextString(m) }
func (*OpenConnectionResponse) ProtoMessage()               {}
func (*OpenConnectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *OpenConnectionResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to CloseConnectionRequest {
type CloseConnectionResponse struct {
	Metadata *RpcMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *CloseConnectionResponse) Reset()                    { *m = CloseConnectionResponse{} }
func (m *CloseConnectionResponse) String() string            { return proto.CompactTextString(m) }
func (*CloseConnectionResponse) ProtoMessage()               {}
func (*CloseConnectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *CloseConnectionResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response to ConnectionSyncRequest
type ConnectionSyncResponse struct {
	ConnProps *ConnectionProperties `protobuf:"bytes,1,opt,name=conn_props,json=connProps" json:"conn_props,omitempty"`
	Metadata  *RpcMetadata          `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *ConnectionSyncResponse) Reset()                    { *m = ConnectionSyncResponse{} }
func (m *ConnectionSyncResponse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionSyncResponse) ProtoMessage()               {}
func (*ConnectionSyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *ConnectionSyncResponse) GetConnProps() *ConnectionProperties {
	if m != nil {
		return m.ConnProps
	}
	return nil
}

func (m *ConnectionSyncResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type DatabasePropertyElement struct {
	Key      *DatabaseProperty `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value    *TypedValue       `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Metadata *RpcMetadata      `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *DatabasePropertyElement) Reset()                    { *m = DatabasePropertyElement{} }
func (m *DatabasePropertyElement) String() string            { return proto.CompactTextString(m) }
func (*DatabasePropertyElement) ProtoMessage()               {}
func (*DatabasePropertyElement) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *DatabasePropertyElement) GetKey() *DatabaseProperty {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *DatabasePropertyElement) GetValue() *TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DatabasePropertyElement) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Response for Meta#getDatabaseProperties()
type DatabasePropertyResponse struct {
	Props    []*DatabasePropertyElement `protobuf:"bytes,1,rep,name=props" json:"props,omitempty"`
	Metadata *RpcMetadata               `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *DatabasePropertyResponse) Reset()                    { *m = DatabasePropertyResponse{} }
func (m *DatabasePropertyResponse) String() string            { return proto.CompactTextString(m) }
func (*DatabasePropertyResponse) ProtoMessage()               {}
func (*DatabasePropertyResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *DatabasePropertyResponse) GetProps() []*DatabasePropertyElement {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *DatabasePropertyResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Send contextual information about some error over the wire from the server.
type ErrorResponse struct {
	Exceptions    []string     `protobuf:"bytes,1,rep,name=exceptions" json:"exceptions,omitempty"`
	HasExceptions bool         `protobuf:"varint,7,opt,name=has_exceptions,json=hasExceptions" json:"has_exceptions,omitempty"`
	ErrorMessage  string       `protobuf:"bytes,2,opt,name=error_message,json=errorMessage" json:"error_message,omitempty"`
	Severity      Severity     `protobuf:"varint,3,opt,name=severity,enum=Severity" json:"severity,omitempty"`
	ErrorCode     uint32       `protobuf:"varint,4,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	SqlState      string       `protobuf:"bytes,5,opt,name=sql_state,json=sqlState" json:"sql_state,omitempty"`
	Metadata      *RpcMetadata `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *ErrorResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type SyncResultsResponse struct {
	MissingStatement bool         `protobuf:"varint,1,opt,name=missing_statement,json=missingStatement" json:"missing_statement,omitempty"`
	MoreResults      bool         `protobuf:"varint,2,opt,name=more_results,json=moreResults" json:"more_results,omitempty"`
	Metadata         *RpcMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *SyncResultsResponse) Reset()                    { *m = SyncResultsResponse{} }
func (m *SyncResultsResponse) String() string            { return proto.CompactTextString(m) }
func (*SyncResultsResponse) ProtoMessage()               {}
func (*SyncResultsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *SyncResultsResponse) GetMetadata() *RpcMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Generic metadata for the server to return with each response.
type RpcMetadata struct {
	ServerAddress string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress" json:"server_address,omitempty"`
}

func (m *RpcMetadata) Reset()                    { *m = RpcMetadata{} }
func (m *RpcMetadata) String() string            { return proto.CompactTextString(m) }
func (*RpcMetadata) ProtoMessage()               {}
func (*RpcMetadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

// Response to a commit request
type CommitResponse struct {
}

func (m *CommitResponse) Reset()                    { *m = CommitResponse{} }
func (m *CommitResponse) String() string            { return proto.CompactTextString(m) }
func (*CommitResponse) ProtoMessage()               {}
func (*CommitResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

// Response to a rollback request
type RollbackResponse struct {
}

func (m *RollbackResponse) Reset()                    { *m = RollbackResponse{} }
func (m *RollbackResponse) String() string            { return proto.CompactTextString(m) }
func (*RollbackResponse) ProtoMessage()               {}
func (*RollbackResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func init() {
	proto.RegisterType((*ResultSetResponse)(nil), "ResultSetResponse")
	proto.RegisterType((*ExecuteResponse)(nil), "ExecuteResponse")
	proto.RegisterType((*PrepareResponse)(nil), "PrepareResponse")
	proto.RegisterType((*FetchResponse)(nil), "FetchResponse")
	proto.RegisterType((*CreateStatementResponse)(nil), "CreateStatementResponse")
	proto.RegisterType((*CloseStatementResponse)(nil), "CloseStatementResponse")
	proto.RegisterType((*OpenConnectionResponse)(nil), "OpenConnectionResponse")
	proto.RegisterType((*CloseConnectionResponse)(nil), "CloseConnectionResponse")
	proto.RegisterType((*ConnectionSyncResponse)(nil), "ConnectionSyncResponse")
	proto.RegisterType((*DatabasePropertyElement)(nil), "DatabasePropertyElement")
	proto.RegisterType((*DatabasePropertyResponse)(nil), "DatabasePropertyResponse")
	proto.RegisterType((*ErrorResponse)(nil), "ErrorResponse")
	proto.RegisterType((*SyncResultsResponse)(nil), "SyncResultsResponse")
	proto.RegisterType((*RpcMetadata)(nil), "RpcMetadata")
	proto.RegisterType((*CommitResponse)(nil), "CommitResponse")
	proto.RegisterType((*RollbackResponse)(nil), "RollbackResponse")
}

var fileDescriptor2 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6e, 0xdb, 0x38,
	0x14, 0x85, 0xec, 0xd8, 0x89, 0xaf, 0x9f, 0xd1, 0x60, 0x66, 0x8c, 0x79, 0x21, 0xe3, 0x60, 0x90,
	0x00, 0x33, 0xd0, 0x22, 0x93, 0x1f, 0x68, 0x5c, 0x07, 0xed, 0x22, 0x68, 0x40, 0x17, 0xdd, 0x0a,
	0x8c, 0x74, 0xe3, 0x08, 0xd1, 0x2b, 0x24, 0x9d, 0xc6, 0x7f, 0x50, 0xa0, 0x8b, 0xee, 0xba, 0xee,
	0x0f, 0xf4, 0x1f, 0x7b, 0x29, 0xca, 0xb4, 0xd3, 0x38, 0x6e, 0xdd, 0x76, 0x65, 0xeb, 0xdc, 0xc3,
	0x7b, 0xae, 0x0e, 0xc9, 0x23, 0xe8, 0x0a, 0x94, 0x79, 0x96, 0x4a, 0x94, 0x5e, 0x2e, 0x32, 0x95,
	0xfd, 0xd6, 0x0a, 0xb2, 0x24, 0xc9, 0x52, 0xf3, 0x34, 0xf8, 0x50, 0x81, 0x5d, 0x86, 0x72, 0x1a,
	0xab, 0x31, 0x2a, 0x56, 0x52, 0xdd, 0x7d, 0x68, 0x07, 0x59, 0x9a, 0x62, 0xa0, 0xa2, 0x2c, 0xf5,
	0xa3, 0xb0, 0xef, 0xec, 0x39, 0x87, 0x0d, 0xd6, 0x5a, 0x80, 0xcf, 0x43, 0xf7, 0x6f, 0x68, 0x49,
	0xc5, 0x15, 0x26, 0x98, 0x2a, 0xcd, 0xa9, 0x10, 0xa7, 0xcd, 0x9a, 0x16, 0x23, 0x0a, 0xf5, 0xc9,
	0x5e, 0xa7, 0xbe, 0x85, 0xfa, 0x55, 0xe2, 0xec, 0xb0, 0x16, 0x81, 0xe3, 0x39, 0xe6, 0x1e, 0x42,
	0x43, 0x46, 0x93, 0x94, 0xab, 0xa9, 0xc0, 0xfe, 0x16, 0x11, 0x9a, 0x47, 0xe0, 0x8d, 0xe7, 0x08,
	0x5b, 0x14, 0xdd, 0x03, 0x68, 0x5e, 0x46, 0x42, 0x2a, 0xff, 0x52, 0xf0, 0x04, 0xfb, 0xb5, 0x82,
	0x5b, 0xf7, 0x4e, 0xf5, 0x13, 0x83, 0xa2, 0x54, 0xfc, 0xd7, 0xa3, 0x4d, 0xf3, 0x90, 0x04, 0xfc,
	0x20, 0x9b, 0x92, 0x6c, 0x9d, 0x98, 0x5b, 0xac, 0x69, 0xb0, 0xa1, 0x86, 0x48, 0x75, 0x27, 0x41,
	0xc5, 0x09, 0xe0, 0xfd, 0xed, 0xa2, 0x51, 0xcb, 0x63, 0x79, 0x70, 0x56, 0x62, 0xcc, 0x56, 0x07,
	0xef, 0x1d, 0xe8, 0x8e, 0xee, 0x30, 0x98, 0x2a, 0xb4, 0x06, 0xfd, 0x07, 0xdb, 0xa2, 0x70, 0x4d,
	0x92, 0x35, 0x55, 0x5a, 0xec, 0x7a, 0x0f, 0x5c, 0x64, 0x73, 0x8a, 0xfb, 0x2f, 0xec, 0x26, 0x91,
	0x94, 0x51, 0x3a, 0x59, 0xb2, 0xa2, 0x52, 0x58, 0xd1, 0x2b, 0x0b, 0xcb, 0x76, 0x2c, 0x06, 0xab,
	0xae, 0x1d, 0xec, 0x1a, 0xba, 0xe7, 0x02, 0x73, 0x2e, 0x16, 0x73, 0x79, 0xe4, 0xa5, 0x55, 0x70,
	0x8a, 0xd5, 0x3d, 0xcf, 0xf6, 0x7e, 0xc6, 0xd3, 0x30, 0xd6, 0x8e, 0xae, 0x14, 0xab, 0xac, 0x15,
	0xfb, 0xe8, 0x40, 0xfb, 0x14, 0x55, 0x70, 0x65, 0xb5, 0xfe, 0x80, 0x9a, 0xd9, 0x07, 0xe7, 0xde,
	0x3e, 0x18, 0x70, 0xb3, 0x77, 0x3e, 0x80, 0xee, 0x9c, 0x3c, 0xb7, 0xd5, 0x9c, 0x94, 0x4e, 0x09,
	0xb3, 0xd2, 0xc9, 0xe5, 0x79, 0xb7, 0xd6, 0xce, 0xfb, 0xd6, 0x81, 0x5f, 0x87, 0x02, 0x49, 0xc1,
	0xca, 0xfc, 0xf0, 0xe3, 0xfd, 0xf5, 0x5b, 0x75, 0x02, 0xbf, 0x0c, 0xe3, 0x4c, 0xae, 0x98, 0x65,
	0xb9, 0x87, 0xf3, 0xa5, 0x1e, 0x2f, 0x72, 0x4c, 0x87, 0x76, 0xc8, 0x6f, 0xe8, 0x31, 0x24, 0x53,
	0xf4, 0x1c, 0xdf, 0xd5, 0xe4, 0x8e, 0x5e, 0xc6, 0xae, 0x1f, 0xcf, 0xd2, 0xc0, 0xf6, 0x38, 0x06,
	0xd0, 0x1e, 0xfa, 0x94, 0x2d, 0xb9, 0x2c, 0xbb, 0xfc, 0xec, 0x2d, 0xc8, 0xe7, 0x84, 0xa3, 0x50,
	0x11, 0x4a, 0xd6, 0xd0, 0x44, 0xfd, 0x2c, 0x37, 0x38, 0x84, 0x7a, 0x53, 0x9f, 0xd2, 0x9f, 0x0b,
	0x2e, 0xb1, 0xec, 0x35, 0x1b, 0xc5, 0xe6, 0x0c, 0xed, 0x43, 0xf5, 0x1a, 0x67, 0xa5, 0xe8, 0xae,
	0xf7, 0x39, 0x8d, 0xe9, 0x2a, 0x6d, 0x6a, 0xed, 0x96, 0xc7, 0x53, 0x2c, 0x75, 0x9a, 0xde, 0xcb,
	0x59, 0x8e, 0xe1, 0x2b, 0x0d, 0x31, 0x53, 0xd9, 0x60, 0x53, 0x15, 0xf4, 0x1f, 0xa8, 0x2c, 0x2e,
	0x62, 0x6d, 0x6e, 0x82, 0x8e, 0x87, 0xbe, 0xf7, 0xc8, 0xd8, 0xcc, 0xd0, 0x36, 0xf0, 0xe0, 0x4d,
	0x05, 0xda, 0x23, 0x21, 0x32, 0x61, 0xb5, 0xfe, 0x02, 0xc0, 0xbb, 0x00, 0x73, 0xed, 0xb0, 0x11,
	0x6c, 0xb0, 0x25, 0xc4, 0xfd, 0x07, 0x3a, 0x57, 0x5c, 0xfa, 0x4b, 0x9c, 0xed, 0xe2, 0x72, 0xb5,
	0x09, 0x1d, 0x2d, 0x68, 0x74, 0x2b, 0x50, 0xf7, 0xf5, 0x13, 0x94, 0x92, 0x4f, 0x8c, 0x47, 0x74,
	0x2b, 0x0a, 0xf0, 0xcc, 0x60, 0xd4, 0x6b, 0x47, 0xe2, 0x2d, 0x8a, 0x48, 0xcd, 0x0a, 0x77, 0x3a,
	0x47, 0x0d, 0x6f, 0x5c, 0x02, 0xcc, 0x96, 0xdc, 0x3f, 0x69, 0xa4, 0xa2, 0x57, 0x90, 0x85, 0x26,
	0xd4, 0xdb, 0xac, 0x51, 0x20, 0x43, 0x02, 0xdc, 0xdf, 0x29, 0xa6, 0x6e, 0x62, 0x13, 0x0c, 0x45,
	0x8c, 0x37, 0x68, 0xed, 0x4d, 0x5c, 0xdc, 0x8e, 0x7b, 0x56, 0xd4, 0xd7, 0x5a, 0xf1, 0xce, 0x81,
	0x9f, 0xca, 0xf3, 0xa7, 0xd3, 0xc1, 0x1a, 0xb2, 0x32, 0x7b, 0x9c, 0x47, 0xb2, 0x87, 0xee, 0x79,
	0x92, 0x09, 0xb4, 0xc1, 0x63, 0x32, 0xaa, 0xa9, 0xb1, 0x55, 0xa9, 0xb3, 0xfe, 0x48, 0x1c, 0x43,
	0x73, 0xa9, 0xa0, 0x9d, 0x97, 0x28, 0xc8, 0x13, 0x9f, 0x87, 0x21, 0x09, 0xc8, 0x32, 0x69, 0xda,
	0x06, 0x7d, 0x62, 0xc0, 0x41, 0x0f, 0x3a, 0x43, 0xfa, 0x28, 0x47, 0x36, 0x15, 0x06, 0x2e, 0xf4,
	0x58, 0x16, 0xc7, 0x17, 0x3c, 0xb8, 0x9e, 0x63, 0x27, 0x03, 0xd8, 0xcb, 0xc4, 0xc4, 0xe3, 0x39,
	0x0f, 0xae, 0xd0, 0x0b, 0x78, 0x1c, 0x44, 0x0a, 0x3d, 0x7e, 0xcb, 0x55, 0x14, 0x70, 0xf3, 0x39,
	0xbf, 0xa8, 0x17, 0x3f, 0xff, 0x7f, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xd0, 0x5a, 0x8c, 0xf6,
	0x07, 0x00, 0x00,
}
