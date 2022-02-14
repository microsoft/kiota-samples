package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/odata"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
    i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions"
    i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties"
    i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments"
    i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/value"
    iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties"
    i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions/item"
    i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties/item"
    idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments/item"
    if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties/item"
)

// MessageRequestBuilder builds and executes requests for operations under \users\{user-id}\messages\{message-id}
type MessageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MessageRequestBuilderDeleteOptions options for Delete
type MessageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageRequestBuilderGetOptions options for Get
type MessageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MessageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MessageRequestBuilderGetQueryParameters the messages in a mailbox or folder. Read-only. Nullable.
type MessageRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// MessageRequestBuilderPatchOptions options for Patch
type MessageRequestBuilderPatchOptions struct {
    // 
    Body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MessageRequestBuilder) Attachments()(*i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.AttachmentsRequestBuilder) {
    return i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/messages/{message_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageRequestBuilder instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MessageRequestBuilder) Content()(*i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.ContentRequestBuilder) {
    return i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) CreateDeleteRequestInformation(options *MessageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) CreateGetRequestInformation(options *MessageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) CreatePatchRequestInformation(options *MessageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) Delete(options *MessageRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d.CreateErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) Extensions()(*i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.ExtensionsRequestBuilder) {
    return i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) Get(options *MessageRequestBuilderGetOptions)(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.CreateMessageFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message), nil
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.MultiValueExtendedPropertiesRequestBuilder) {
    return iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the messages in a mailbox or folder. Read-only. Nullable.
func (m *MessageRequestBuilder) Patch(options *MessageRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d.CreateErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.SingleValueExtendedPropertiesRequestBuilder) {
    return i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
