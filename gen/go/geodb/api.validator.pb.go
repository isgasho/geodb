// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *KeysRequest) Validate() error {
	return nil
}
func (this *KeysResponse) Validate() error {
	return nil
}
func (this *StreamRequest) Validate() error {
	return nil
}
func (this *StreamResponse) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}
func (this *SetRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SetResponse) Validate() error {
	return nil
}
func (this *GetRequest) Validate() error {
	return nil
}
func (this *GetResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetRegexRequest) Validate() error {
	return nil
}
func (this *GetRegexResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SeekRequest) Validate() error {
	return nil
}
func (this *SeekResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *DeleteResponse) Validate() error {
	return nil
}
func (this *PingRequest) Validate() error {
	return nil
}
func (this *PingResponse) Validate() error {
	return nil
}
func (this *Point) Validate() error {
	return nil
}
func (this *Object) Validate() error {
	if this.Point != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Point); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Point", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Event) Validate() error {
	if this.Point != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Point); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Point", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
