package item

import (
	i90ce7b240dbfe27d1efd75544859405d70943a52fa428ec2a2f23fa14d2a50a0 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification"
	iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders"
	i0534d8fa67c49f8a9de6c0e5b380f4de8def1eb3520f2cbc0e054e2f0d8b4ddb "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/item"
	ia21502713d6ef2ca89be92f9aca3f2bcfbbf7da87a376cf2ab6bdd9dc78f7c88 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages"
	i86f521d0387d65d750eeee005d9fc2d9c62ea2a68edef6c5c22a58dcd07526ba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/item"
	ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

type UserRequestBuilder struct {
	currentPath string
	httpCore    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore
	isRawUrl    bool
	pathSegment string
}

func NewUserRequestBuilder(currentPath string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore, isRawUrl bool) *UserRequestBuilder {
	m := &UserRequestBuilder{}
	m.pathSegment = ""
	m.httpCore = httpCore
	m.currentPath = currentPath
	m.isRawUrl = isRawUrl
	return m
}
func (m *UserRequestBuilder) InferenceClassification() i90ce7b240dbfe27d1efd75544859405d70943a52fa428ec2a2f23fa14d2a50a0.InferenceClassificationRequestBuilder {
	return *i90ce7b240dbfe27d1efd75544859405d70943a52fa428ec2a2f23fa14d2a50a0.NewInferenceClassificationRequestBuilder(m.currentPath+m.pathSegment, m.httpCore, false)
}
func (m *UserRequestBuilder) MailFolders() iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a.MailFoldersRequestBuilder {
	return *iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a.NewMailFoldersRequestBuilder(m.currentPath+m.pathSegment, m.httpCore, false)
}
func (m *UserRequestBuilder) MailFoldersById(id string) i0534d8fa67c49f8a9de6c0e5b380f4de8def1eb3520f2cbc0e054e2f0d8b4ddb.MailFolderRequestBuilder {
	return *i0534d8fa67c49f8a9de6c0e5b380f4de8def1eb3520f2cbc0e054e2f0d8b4ddb.NewMailFolderRequestBuilder(m.currentPath+m.pathSegment+"/mailFolders/"+id, m.httpCore, false)
}
func (m *UserRequestBuilder) Messages() ia21502713d6ef2ca89be92f9aca3f2bcfbbf7da87a376cf2ab6bdd9dc78f7c88.MessagesRequestBuilder {
	return *ia21502713d6ef2ca89be92f9aca3f2bcfbbf7da87a376cf2ab6bdd9dc78f7c88.NewMessagesRequestBuilder(m.currentPath+m.pathSegment, m.httpCore, false)
}
func (m *UserRequestBuilder) MessagesById(id string) i86f521d0387d65d750eeee005d9fc2d9c62ea2a68edef6c5c22a58dcd07526ba.MessageRequestBuilder {
	return *i86f521d0387d65d750eeee005d9fc2d9c62ea2a68edef6c5c22a58dcd07526ba.NewMessageRequestBuilder(m.currentPath+m.pathSegment+"/messages/"+id, m.httpCore, false)
}
