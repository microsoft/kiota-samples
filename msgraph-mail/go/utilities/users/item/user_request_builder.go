package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i341158d6b31e193c6856ef08d3a5c0a9d9c9674cd883d99b10844de8d933df7a "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages"
    i9bcbbe222fda57b1c7f212f812b862b760156d2d3a9d1e157ee4f76841199559 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/inferenceclassification"
    ib8edbc02a502b08756d43d6d66bfafac60ac37ce024db3b7e881dc17cea846a8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders"
    ia31bb5deda23f2b146eb943e6d6e68f69c1487606651126c6dbde9f9b4b6ae99 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item"
    if937953a6e06f9e461297be544e9a26dbe87704845fad3ae4b78b2ca32d53222 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item"
)

// UserRequestBuilder builds and executes requests for operations under \users\{user-id}
type UserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewUserRequestBuilderInternal instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserRequestBuilder) {
    m := &UserRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserRequestBuilder instantiates a new UserRequestBuilder and sets the default values.
func NewUserRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *UserRequestBuilder) InferenceClassification()(*i9bcbbe222fda57b1c7f212f812b862b760156d2d3a9d1e157ee4f76841199559.InferenceClassificationRequestBuilder) {
    return i9bcbbe222fda57b1c7f212f812b862b760156d2d3a9d1e157ee4f76841199559.NewInferenceClassificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UserRequestBuilder) MailFolders()(*ib8edbc02a502b08756d43d6d66bfafac60ac37ce024db3b7e881dc17cea846a8.MailFoldersRequestBuilder) {
    return ib8edbc02a502b08756d43d6d66bfafac60ac37ce024db3b7e881dc17cea846a8.NewMailFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MailFoldersById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item collection
func (m *UserRequestBuilder) MailFoldersById(id string)(*if937953a6e06f9e461297be544e9a26dbe87704845fad3ae4b78b2ca32d53222.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id"] = id
    }
    return if937953a6e06f9e461297be544e9a26dbe87704845fad3ae4b78b2ca32d53222.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *UserRequestBuilder) Messages()(*i341158d6b31e193c6856ef08d3a5c0a9d9c9674cd883d99b10844de8d933df7a.MessagesRequestBuilder) {
    return i341158d6b31e193c6856ef08d3a5c0a9d9c9674cd883d99b10844de8d933df7a.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.messages.item collection
func (m *UserRequestBuilder) MessagesById(id string)(*ia31bb5deda23f2b146eb943e6d6e68f69c1487606651126c6dbde9f9b4b6ae99.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return ia31bb5deda23f2b146eb943e6d6e68f69c1487606651126c6dbde9f9b4b6ae99.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
