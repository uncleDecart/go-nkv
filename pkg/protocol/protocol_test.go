package protocol_test

import (
	sdk "github.com/nkval/go-nkv/pkg/protocol"
	"github.com/onsi/gomega"
	"testing"
)

func TestUnmarshalRequest(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	input := "GET 12345 client1 key1"
	command, err := sdk.UnmarshalRequest(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(command).To(gomega.BeEquivalentTo(
		&sdk.Request{
			Request:   sdk.RequestGet,
			RequestID: "12345",
			ClientID:  "client1",
			Key:       "key1",
		},
	))

	input = "PUT 12345 client1 key1 YmF6aW5nYQo="
	command, err = sdk.UnmarshalRequest(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(command).To(gomega.BeEquivalentTo(
		&sdk.Request{
			Request:   sdk.RequestPut,
			RequestID: "12345",
			ClientID:  "client1",
			Key:       "key1",
			Data:      []byte{98, 97, 122, 105, 110, 103, 97, 10},
		},
	))

	input = "DEL 12345 client1 key1"
	command, err = sdk.UnmarshalRequest(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(command).To(gomega.BeEquivalentTo(
		&sdk.Request{
			Request:   sdk.RequestDel,
			RequestID: "12345",
			ClientID:  "client1",
			Key:       "key1",
		},
	))

	input = "SUB 12345 client1 key1"
	command, err = sdk.UnmarshalRequest(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(command).To(gomega.BeEquivalentTo(
		&sdk.Request{
			Request:   sdk.RequestSub,
			RequestID: "12345",
			ClientID:  "client1",
			Key:       "key1",
		},
	))

	input = "UNSUB 12345 client1 key1"
	command, err = sdk.UnmarshalRequest(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(command).To(gomega.BeEquivalentTo(
		&sdk.Request{
			Request:   sdk.RequestUnsub,
			RequestID: "12345",
			ClientID:  "client1",
			Key:       "key1",
		},
	))
}

func TestMarshalRequest(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	command := &sdk.Request{
		Request:   sdk.RequestGet,
		RequestID: "12345",
		ClientID:  "client1",
		Key:       "key1",
	}
	expected := "GET 12345 client1 key1"
	g.Expect(sdk.MarshalRequest(command), expected)

	command = &sdk.Request{
		Request:   sdk.RequestPut,
		RequestID: "12345",
		ClientID:  "client1",
		Key:       "key1",
		Data:      []byte{98, 97, 122, 105, 110, 103, 97, 10},
	}
	expected = "PUT 12345 client1 key1 YmF6aW5nYQo="
	g.Expect(sdk.MarshalRequest(command), expected)

	command = &sdk.Request{
		Request:   sdk.RequestDel,
		RequestID: "12345",
		ClientID:  "client1",
		Key:       "key1",
	}
	expected = "DEL 12345 client1 key1"
	g.Expect(sdk.MarshalRequest(command), expected)

	command = &sdk.Request{
		Request:   sdk.RequestSub,
		RequestID: "12345",
		ClientID:  "client1",
		Key:       "key1",
	}
	expected = "SUB 12345 client1 key1"
	g.Expect(sdk.MarshalRequest(command), expected)

	command = &sdk.Request{
		Request:   sdk.RequestUnsub,
		RequestID: "12345",
		ClientID:  "client1",
		Key:       "key1",
	}
	expected = "UNSUB 12345 client1 key1"
	g.Expect(sdk.MarshalRequest(command), expected)
}

func TestMarsalResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	resp := &sdk.Response{
		RequestID: "12345",
		Status:    false,
	}
	expected := "12345 FAILED"
	g.Expect(sdk.MarshalResponse(resp), expected)

	resp = &sdk.Response{
		RequestID: "12345",
		Status:    true,
		Data:      []byte{98, 97, 122, 105, 110, 103, 97, 10},
	}
	expected = "12345 OK YmF6aW5nYQo="
	g.Expect(sdk.MarshalResponse(resp), expected)
}

func TestUnmarshalResponse(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	input := "12345 FAILED"
	resp, err := sdk.UnmarshalResponse(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(resp).To(gomega.BeEquivalentTo(
		&sdk.Response{
			RequestID: "12345",
			Status:    false,
			Data:      []byte{},
		},
	))

	input = "12345 OK YmF6aW5nYQo="
	resp, err = sdk.UnmarshalResponse(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(resp).To(gomega.BeEquivalentTo(
		&sdk.Response{
			RequestID: "12345",
			Status:    true,
			Data:      []byte{98, 97, 122, 105, 110, 103, 97, 10},
		},
	))
}

func TestMarshalNotification(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	n := &sdk.Notification{
		Type: sdk.NotificationHello,
		Key:  "12345",
	}
	expected := "HELLO 12345"
	g.Expect(sdk.MarshalNotification(n), expected)

	n = &sdk.Notification{
		Type: sdk.NotificationUpdate,
		Key:  "12345",
		Data: []byte{98, 97, 122, 105, 110, 103, 97, 10},
	}
	expected = "UPDATE 12345 YmF6aW5nYQo="
	g.Expect(sdk.MarshalNotification(n), expected)

	n = &sdk.Notification{
		Type: sdk.NotificationClose,
		Key:  "12345",
	}
	expected = "CLOSE 12345"
	g.Expect(sdk.MarshalNotification(n), expected)
}

func TestUnmarshalNotification(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	input := "HELLO 12345"
	resp, err := sdk.UnmarshalNotification(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(resp).To(gomega.BeEquivalentTo(
		&sdk.Notification{
			Type: sdk.NotificationHello,
			Key:  "12345",
		},
	))

	input = "UPDATE 12345 YmF6aW5nYQo="
	resp, err = sdk.UnmarshalNotification(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(resp).To(gomega.BeEquivalentTo(
		&sdk.Notification{
			Type: sdk.NotificationUpdate,
			Key:  "12345",
			Data: []byte{98, 97, 122, 105, 110, 103, 97, 10},
		},
	))

	input = "CLOSE 12345"
	resp, err = sdk.UnmarshalNotification(input)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(resp).To(gomega.BeEquivalentTo(
		&sdk.Notification{
			Type: sdk.NotificationClose,
			Key:  "12345",
		},
	))
}
