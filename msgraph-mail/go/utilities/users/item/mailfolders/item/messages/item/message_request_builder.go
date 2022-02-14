package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/odata"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
    i000de88af33510156072c580aff36ddc934a6dbf385120aee0662cecf92dfee6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/value"
    i165e60b62d06f92d5e6029daee4ae129b1c3d1780dce2631cc18fbe8c6d39c9d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/multivalueextendedproperties"
    i281b692d3f673c995c16b5cc69287c818497fc74562aed972acc168040629f74 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/extensions"
    i833cf7d5250828d0b2288e87035b59c2900b4aef7c2eb2ff5494a38edea8268e "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/singlevalueextendedproperties"
    i9b8e8e1d8c7f8c0fa45b5abe65e2541be081536bd338e75409f4b0ecba2188aa "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/attachments"
    i314a2a68928bac08cd3678f4259eab0e5845a67017fe63526811793a0c7a4615 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/multivalueextendedproperties/item"
    i92b891944e73393404194faa96a7587068f07492bd44347d827c18456b8a2bcf "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/attachments/item"
    ia3b46812df8e43b8356e8bad54f90359cc966c930f9415ac7f61670490ce857f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/singlevalueextendedproperties/item"
    icb423aed2922bb0caf9a0353a5ec5bfeb474588dbd1218008894e86ab49c6e25 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/extensions/item"
)

// MessageRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}
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
// MessageRequestBuilderGetQueryParameters the collection of messages in the mailFolder.
type MessageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
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
func (m *MessageRequestBuilder) Attachments()(*i9b8e8e1d8c7f8c0fa45b5abe65e2541be081536bd338e75409f4b0ecba2188aa.AttachmentsRequestBuilder) {
    return i9b8e8e1d8c7f8c0fa45b5abe65e2541be081536bd338e75409f4b0ecba2188aa.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.messages.item.attachments.item collection
func (m *MessageRequestBuilder) AttachmentsById(id string)(*i92b891944e73393404194faa96a7587068f07492bd44347d827c18456b8a2bcf.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i92b891944e73393404194faa96a7587068f07492bd44347d827c18456b8a2bcf.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMessageRequestBuilderInternal instantiates a new MessageRequestBuilder and sets the default values.
func NewMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}{?select,expand}";
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
func (m *MessageRequestBuilder) Content()(*i000de88af33510156072c580aff36ddc934a6dbf385120aee0662cecf92dfee6.ContentRequestBuilder) {
    return i000de88af33510156072c580aff36ddc934a6dbf385120aee0662cecf92dfee6.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the collection of messages in the mailFolder.
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
// CreateGetRequestInformation the collection of messages in the mailFolder.
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
// CreatePatchRequestInformation the collection of messages in the mailFolder.
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
// Delete the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) Extensions()(*i281b692d3f673c995c16b5cc69287c818497fc74562aed972acc168040629f74.ExtensionsRequestBuilder) {
    return i281b692d3f673c995c16b5cc69287c818497fc74562aed972acc168040629f74.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.messages.item.extensions.item collection
func (m *MessageRequestBuilder) ExtensionsById(id string)(*icb423aed2922bb0caf9a0353a5ec5bfeb474588dbd1218008894e86ab49c6e25.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return icb423aed2922bb0caf9a0353a5ec5bfeb474588dbd1218008894e86ab49c6e25.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(*i165e60b62d06f92d5e6029daee4ae129b1c3d1780dce2631cc18fbe8c6d39c9d.MultiValueExtendedPropertiesRequestBuilder) {
    return i165e60b62d06f92d5e6029daee4ae129b1c3d1780dce2631cc18fbe8c6d39c9d.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i314a2a68928bac08cd3678f4259eab0e5845a67017fe63526811793a0c7a4615.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i314a2a68928bac08cd3678f4259eab0e5845a67017fe63526811793a0c7a4615.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the collection of messages in the mailFolder.
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
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(*i833cf7d5250828d0b2288e87035b59c2900b4aef7c2eb2ff5494a38edea8268e.SingleValueExtendedPropertiesRequestBuilder) {
    return i833cf7d5250828d0b2288e87035b59c2900b4aef7c2eb2ff5494a38edea8268e.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ia3b46812df8e43b8356e8bad54f90359cc966c930f9415ac7f61670490ce857f.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ia3b46812df8e43b8356e8bad54f90359cc966c930f9415ac7f61670490ce857f.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
