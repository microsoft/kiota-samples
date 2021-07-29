package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    iab1a0dd9a491da948f96274b48b251c29156f5ab5cdef57a533c314762a96cba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/childfolders"
    id4ccc2d90f7e823c6e1c9e4a9a1f0e784a26e3fd6a68c038bda15d61728d0e46 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messagerules"
    i17b176b6a1414dd086eba0883aee3fc49c21b71511308a01a6732762ead9b8e8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages"
    i660191d0d3ef7133b6f3e0a4532d18c73c2f52a7222a034c4492d6ec6ef3fc3e "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/multivalueextendedproperties"
    ib4b1a269396ed151935b7d4ec590d14cf3ed072057ab6475f11863830766d145 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/singlevalueextendedproperties"
    i0dcb0d8c6a911f67209ea659b254ba6713191c318f2799cf8a9383d821d87061 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/childfolders/item"
    i5d2f606000ccb668c7a688148beea132a38f81fbf9cf32e175421151435d1b29 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messagerules/item"
    ifdebf055cf7726bc6ca41ec4be0e381235fd73f507038a2ea27cd08a6fa09506 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/item"
    id2dc360a4b1db0a33958a93a105ee47a087c7d3853efa23c73dbf41c05cd738e "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/multivalueextendedproperties/item"
    ic93bf2d2dca930088dbc375148726cf159640b4b7220bf5ee92920b4e2022faa "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/singlevalueextendedproperties/item"
)

type MailFolderRequestBuilder struct {
    ChildFolders iab1a0dd9a491da948f96274b48b251c29156f5ab5cdef57a533c314762a96cba.ChildFoldersRequestBuilder;
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    MessageRules id4ccc2d90f7e823c6e1c9e4a9a1f0e784a26e3fd6a68c038bda15d61728d0e46.MessageRulesRequestBuilder;
    Messages i17b176b6a1414dd086eba0883aee3fc49c21b71511308a01a6732762ead9b8e8.MessagesRequestBuilder;
    MultiValueExtendedProperties i660191d0d3ef7133b6f3e0a4532d18c73c2f52a7222a034c4492d6ec6ef3fc3e.MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties ib4b1a269396ed151935b7d4ec590d14cf3ed072057ab6475f11863830766d145.SingleValueExtendedPropertiesRequestBuilder;
}
