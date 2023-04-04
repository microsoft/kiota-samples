package utilities

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i30fc20112facbd93862049bcbeffb8938df51d762c36a3f68cc5f482741730d9 "github.com/microsoft/kiota-samples/petstore/go/utilities/store"
    i3c37a61d7272de22ab30735a6fa0fcf0b883a5c82cf706fd10365a3203e1235b "github.com/microsoft/kiota-samples/petstore/go/utilities/user"
    if5191a3b08f2ce6dfdeab2abbf317a97232aa6a8a9ced702da7454d6f82d5050 "github.com/microsoft/kiota-samples/petstore/go/utilities/pet"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://petstore.swagger.io/v2")
    }
    m.pathParameters["baseurl"] = m.requestAdapter.GetBaseUrl()
    return m
}
// Pet the pet property
func (m *ApiClient) Pet()(*if5191a3b08f2ce6dfdeab2abbf317a97232aa6a8a9ced702da7454d6f82d5050.PetRequestBuilder) {
    return if5191a3b08f2ce6dfdeab2abbf317a97232aa6a8a9ced702da7454d6f82d5050.NewPetRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PetById gets an item from the github.com/microsoft/kiota-samples/petstore/go/utilities/.pet.item collection
func (m *ApiClient) PetById(id string)(*if5191a3b08f2ce6dfdeab2abbf317a97232aa6a8a9ced702da7454d6f82d5050.WithPetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["petId"] = id
    }
    return if5191a3b08f2ce6dfdeab2abbf317a97232aa6a8a9ced702da7454d6f82d5050.NewWithPetItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Store the store property
func (m *ApiClient) Store()(*i30fc20112facbd93862049bcbeffb8938df51d762c36a3f68cc5f482741730d9.StoreRequestBuilder) {
    return i30fc20112facbd93862049bcbeffb8938df51d762c36a3f68cc5f482741730d9.NewStoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// User the user property
func (m *ApiClient) User()(*i3c37a61d7272de22ab30735a6fa0fcf0b883a5c82cf706fd10365a3203e1235b.UserRequestBuilder) {
    return i3c37a61d7272de22ab30735a6fa0fcf0b883a5c82cf706fd10365a3203e1235b.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserById gets an item from the github.com/microsoft/kiota-samples/petstore/go/utilities/.user.item collection
func (m *ApiClient) UserById(id string)(*i3c37a61d7272de22ab30735a6fa0fcf0b883a5c82cf706fd10365a3203e1235b.WithUsernameItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["username"] = id
    }
    return i3c37a61d7272de22ab30735a6fa0fcf0b883a5c82cf706fd10365a3203e1235b.NewWithUsernameItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
