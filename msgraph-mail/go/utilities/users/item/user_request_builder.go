package item

import (
    url "net/url"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i90ce7b240dbfe27d1efd75544859405d70943a52fa428ec2a2f23fa14d2a50a0 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification"
    ia21502713d6ef2ca89be92f9aca3f2bcfbbf7da87a376cf2ab6bdd9dc78f7c88 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages"
    iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders"
    i0534d8fa67c49f8a9de6c0e5b380f4de8def1eb3520f2cbc0e054e2f0d8b4ddb "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/item"
    i86f521d0387d65d750eeee005d9fc2d9c62ea2a68edef6c5c22a58dcd07526ba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/item"
)

type UserRequestBuilder struct {
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    InferenceClassification *i90ce7b240dbfe27d1efd75544859405d70943a52fa428ec2a2f23fa14d2a50a0.InferenceClassificationRequestBuilder;
    MailFolders *iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a.MailFoldersRequestBuilder;
    Messages *ia21502713d6ef2ca89be92f9aca3f2bcfbbf7da87a376cf2ab6bdd9dc78f7c88.MessagesRequestBuilder;
    pathSegment *string;
}
func NewUserRequestBuilder(currentPath *string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*UserRequestBuilder) {
    return nil
}
func (m *UserRequestBuilder) MailFolders(id *string)(i0534d8fa67c49f8a9de6c0e5b380f4de8def1eb3520f2cbc0e054e2f0d8b4ddb.MailFolderRequestBuilder) {
    return i0534d8fa67c49f8a9de6c0e5b380f4de8def1eb3520f2cbc0e054e2f0d8b4ddb.MailFolderRequestBuilder.NewMailFolderRequestBuilder(*m.currentPath + *m.pathSegment + "/mailFolders/" + id, m.httpCore);
}
func (m *UserRequestBuilder) Messages(id *string)(i86f521d0387d65d750eeee005d9fc2d9c62ea2a68edef6c5c22a58dcd07526ba.MessageRequestBuilder) {
    return i86f521d0387d65d750eeee005d9fc2d9c62ea2a68edef6c5c22a58dcd07526ba.MessageRequestBuilder.NewMessageRequestBuilder(*m.currentPath + *m.pathSegment + "/messages/" + id, m.httpCore);
}
