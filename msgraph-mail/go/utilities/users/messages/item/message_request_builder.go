package item

import (
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    ib74bdf213c8ecff00ba896023527beb4a5cee0f1561b4ebb235535adb8fd01ba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/attachments/item"
    ia04a7159d2fa11d2263f021f3437a00844488a2ead4b708e8286dfba697c6c7b "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/singleValueExtendedProperties/item"
    ie2165dbef9de3505face24fc57b70c553f65af1630431e50b8c923b24323a3b5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/multiValueExtendedProperties/item"
    i4c213381c08f42149bb0acad01e026154f1d08ef5829ececbd8ea51281fe708b "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/extensions/item"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

type MessageRequestBuilder struct {
    Attachments i0686696817f5336ad23f5048aaf1973675e26317e928255ef44ff8f6cdd744ce.AttachmentsRequestBuilder;
    Content i89225ca370f14a1b7bb2b0941acdb3ba5a42bdbb8774f8e966963a5e7581ed65.ContentRequestBuilder;
    currentPath *string;
    Extensions ib2381cda6aa2598b763d17161ebaa5c941e140eebfa25805cfbb21dbd2ff0805.ExtensionsRequestBuilder;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    MultiValueExtendedProperties iffe5c634b20857a786a409f178841ae1a8aca0c1eeb38128392ae6af2fee4843.MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties i40759593f5e6f2d9981647c7797bc3a3775a679e5a825a2f688c102de1529d72.SingleValueExtendedPropertiesRequestBuilder;
}
