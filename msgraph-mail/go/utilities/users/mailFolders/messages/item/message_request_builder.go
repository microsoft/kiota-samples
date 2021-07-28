package item

import (
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    ia99cfa9415d76594f23f9203f6c0781910192430e0495749b24ae578cca977a4 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/attachments/item"
    i9daf6d054961a1fa93e65ba348ad7cbae234b91c63b2ae47964102d0bdd78434 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/singleValueExtendedProperties/item"
    i4cfb67aa6e97472c4a12053fedbe772214c671eba23ba62c7476328b062f3aa5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/multiValueExtendedProperties/item"
    i9d883d8086a9391a6ca4c2bcac59d80d05afcf2e21ceceab6e9f27729e18972a "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/extensions/item"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

type MessageRequestBuilder struct {
    Attachments ifbd8fe607eece29e08924ff3d5f24e8cfdf6b4e4d1d5ac56cfafa62346d7ed90.AttachmentsRequestBuilder;
    Content i64cad71c14a88e79fc689b46bf8255a4b26ed69bdd5b273d9b16fd296f644c84.ContentRequestBuilder;
    currentPath *string;
    Extensions i525295fc3ac0012f64a4caf4407e886671d2bc7ed80a3ff0346a3cd56405fa61.ExtensionsRequestBuilder;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    MultiValueExtendedProperties i7e2f2cf336b45fd1920a0f7bd75639edf8c3ffa5809c90223ecef87a4eabee66.MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties ia986e72d952128f3a048e23bf8c62132c7e6f1368ea463fc4fb3f236ee77c836.SingleValueExtendedPropertiesRequestBuilder;
}
