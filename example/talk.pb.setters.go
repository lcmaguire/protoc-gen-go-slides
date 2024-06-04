// Code generated by protoc-gen-go-setters. DO NOT EDIT.
// source: example/talk.proto
package example

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// SetPageSize will set ListTalksRequest.PageSize.
//
// WARNING: no nil checks are performed on the parent message.
func (x *ListTalksRequest) SetPageSize(in int32) {
	x.PageSize = in
}

// SetPageToken will set ListTalksRequest.PageToken.
//
// WARNING: no nil checks are performed on the parent message.
func (x *ListTalksRequest) SetPageToken(in string) {
	x.PageToken = in
}

// SetSpeakerId will set ListTalksRequest.SpeakerId.
//
// WARNING: no nil checks are performed on the parent message.
func (x *ListTalksRequest) SetSpeakerId(in string) {
	x.SpeakerId = in
}

// SetTalks will set ListTalksResponse.Talks.
//
// WARNING: no nil checks are performed on the parent message.
func (x *ListTalksResponse) SetTalks(in []*Talk) {
	x.Talks = in
}

// AppendTalks will append all input values to ListTalksResponse.Talks.
//
// WARNING: no nil checks are performed on the parent message.
func (x *ListTalksResponse) AppendTalks(in ...*Talk) {
	x.Talks = append(x.Talks, in...)
}

// SetNextPageToken will set ListTalksResponse.NextPageToken.
//
// WARNING: no nil checks are performed on the parent message.
func (x *ListTalksResponse) SetNextPageToken(in string) {
	x.NextPageToken = in
}

// SetTalk will set CreateTalksRequest.Talk.
//
// WARNING: no nil checks are performed on the parent message.
func (x *CreateTalksRequest) SetTalk(in *Talk) {
	x.Talk = in
}

// SetTalkId will set Talk.TalkId.
//
// WARNING: no nil checks are performed on the parent message.
func (x *Talk) SetTalkId(in string) {
	x.TalkId = in
}

// SetTalkTime will set Talk.TalkTime.
//
// WARNING: no nil checks are performed on the parent message.
func (x *Talk) SetTalkTime(in *timestamppb.Timestamp) {
	x.TalkTime = in
}

// SetSpeaker will set Talk.Speaker.
//
// WARNING: no nil checks are performed on the parent message.
func (x *Talk) SetSpeaker(in string) {
	x.Speaker = in
}