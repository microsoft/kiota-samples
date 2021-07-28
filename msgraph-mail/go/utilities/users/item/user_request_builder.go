package item

import (
    i07bf331452fadd0ef76c41057329da8cd4d22a38d2f83761adf79ad6459573e3 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/item"
    i86f521d0387d65d750eeee005d9fc2d9c62ea2a68edef6c5c22a58dcd07526ba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/item"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

type UserRequestBuilder struct {
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    InferenceClassification i0effd602dd086476132a16986fbc6cd414dbacebaf8ee7743cfd3d7986503d77.InferenceClassificationRequestBuilder;
    MailFolders ia254bc7eebe487216682ef10c3748ccc8d561219b8bde857ab0c71fc56c4a3a7.MailFoldersRequestBuilder;
    Messages ia21502713d6ef2ca89be92f9aca3f2bcfbbf7da87a376cf2ab6bdd9dc78f7c88.MessagesRequestBuilder;
    pathSegment *string;
}
