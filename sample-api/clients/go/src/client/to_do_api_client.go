package client

import (
	i8de5759b6015dc65a49e22f6e69c66b78c8d9b53bf0a58e5b8b1f403b48f4a47 "todoclient/client/todoitems"

	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
	i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
	i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
)

// ToDoApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ToDoApiClient struct {
	// Path parameters for the request
	pathParameters map[string]string
	// The request adapter to use to execute the requests.
	requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
	// Url template to use to build the URL for the current request builder
	urlTemplate string
}

// NewToDoApiClient instantiates a new ToDoApiClient and sets the default values.
func NewToDoApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ToDoApiClient {
	m := &ToDoApiClient{}
	m.pathParameters = make(map[string]string)
	m.urlTemplate = "{+baseurl}"
	m.requestAdapter = requestAdapter
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory {
		return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory()
	})
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory {
		return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory()
	})
	if m.requestAdapter.GetBaseUrl() == "" {
		m.requestAdapter.SetBaseUrl("https://localhost:7206")
	}
	m.pathParameters["baseurl"] = m.requestAdapter.GetBaseUrl()
	return m
}

// ToDoItems provides operations to manage the collection of ToDoItem entities.
func (m *ToDoApiClient) ToDoItems() *i8de5759b6015dc65a49e22f6e69c66b78c8d9b53bf0a58e5b8b1f403b48f4a47.ToDoItemsRequestBuilder {
	return i8de5759b6015dc65a49e22f6e69c66b78c8d9b53bf0a58e5b8b1f403b48f4a47.NewToDoItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}

// ToDoItemsById provides operations to manage the collection of ToDoItem entities.
func (m *ToDoApiClient) ToDoItemsById(id string) *i8de5759b6015dc65a49e22f6e69c66b78c8d9b53bf0a58e5b8b1f403b48f4a47.ToDoItemItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.pathParameters {
		urlTplParams[idx] = item
	}
	if id != "" {
		urlTplParams["ToDoItem%2Did"] = id
	}
	return i8de5759b6015dc65a49e22f6e69c66b78c8d9b53bf0a58e5b8b1f403b48f4a47.NewToDoItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
