package utilities

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id1ae20a9e00c372d14381acc2299aa946a25894316974139983388e4b11bb84b "github.com/microsoft/kiota/serialization/go/json"
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    i993da4875956788f567345c5c97cefaefa8f33555fde8bb29d1fd49e374ccccd "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter;
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RegisterDefaultSerializer(func() i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriterFactory { return id1ae20a9e00c372d14381acc2299aa946a25894316974139983388e4b11bb84b.NewJsonSerializationWriterFactory() })
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RegisterDefaultDeserializer(func() i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNodeFactory { return id1ae20a9e00c372d14381acc2299aa946a25894316974139983388e4b11bb84b.NewJsonParseNodeFactory() })
    m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/v1.0")
    return m
}
func (m *ApiClient) Users()(*ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.UsersRequestBuilder) {
    return ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item collection
func (m *ApiClient) UsersById(id string)(*i993da4875956788f567345c5c97cefaefa8f33555fde8bb29d1fd49e374ccccd.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user_id"] = id
    }
    return i993da4875956788f567345c5c97cefaefa8f33555fde8bb29d1fd49e374ccccd.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
