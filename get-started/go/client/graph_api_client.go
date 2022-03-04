package client

import (
    i1e6e8d869c28dd4f8a4f911410f9299cd7cd087e16dead40558d14efc1888900 "getuser/client/me"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id1ae20a9e00c372d14381acc2299aa946a25894316974139983388e4b11bb84b "github.com/microsoft/kiota/serialization/go/json"
)

// GraphApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphApiClient struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewGraphApiClient instantiates a new GraphApiClient and sets the default values.
func NewGraphApiClient(requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GraphApiClient) {
    m := &GraphApiClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter;
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RegisterDefaultSerializer(func() i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriterFactory { return id1ae20a9e00c372d14381acc2299aa946a25894316974139983388e4b11bb84b.NewJsonSerializationWriterFactory() })
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RegisterDefaultDeserializer(func() i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNodeFactory { return id1ae20a9e00c372d14381acc2299aa946a25894316974139983388e4b11bb84b.NewJsonParseNodeFactory() })
    m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/v1.0")
    return m
}
func (m *GraphApiClient) Me()(*i1e6e8d869c28dd4f8a4f911410f9299cd7cd087e16dead40558d14efc1888900.MeRequestBuilder) {
    return i1e6e8d869c28dd4f8a4f911410f9299cd7cd087e16dead40558d14efc1888900.NewMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
