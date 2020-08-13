package shmsg

import (
	"encoding/base64"

	"google.golang.org/protobuf/proto"
)

// UrlEncodeMessage encodes Message as a string, which is safe to be used as part of an URL
func UrlEncodeMessage(msg *Message) (string, error) {
	out, err := proto.Marshal(msg)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(out), nil
}

// UrlDecodeMessage decodes a Message from the given string
func UrlDecodeMessage(encoded string) (*Message, error) {
	msg := Message{}
	out, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(out, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
