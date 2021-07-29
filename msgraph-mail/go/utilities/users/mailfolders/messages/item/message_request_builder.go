package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    i3a8d456234d570904652042e0d4c0c5c4510aca0788f80c075c5daa8f534b477 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/attachments"
    iaeab5810035e3d7ff3e4901893330c349ce3db065ba7b9e0dcaadaf7479be003 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/content"
    i1e8f5e1b0ab12ae689692569cf00409a7ffcf8119e0c731d1aab8fd8c24b9c9a "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/extensions"
    ic1605e2a56328ce22ef4db9ad89f54800a8f837cf7ba24dd304d6ddd4f629093 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/multivalueextendedproperties"
    i5f4f3512bc76131a2bfb807e4ef99b24c83995d2d107aba0fd68cf3b07f8b750 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/singlevalueextendedproperties"
    iace81820e7196eb3ec506adf50a909fd751a256f8329891ef2cbbdd1d8fe9b2d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/attachments/item"
    ibce1fc49ff1dfb26493978b2a56f0be047bae8336ff396cb9deb7e0715c92eca "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/extensions/item"
    i99bbed137b62cf613096d6395b85791513d9e404f617b07c9675049bc9df09f8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/multivalueextendedproperties/item"
    idf8de8c43ef15f0fdc414d98d3e0cf730d22735ef1baafa82a410f3c5cbf2e90 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/singlevalueextendedproperties/item"
)

type MessageRequestBuilder struct {
    Attachments i3a8d456234d570904652042e0d4c0c5c4510aca0788f80c075c5daa8f534b477.AttachmentsRequestBuilder;
    Content iaeab5810035e3d7ff3e4901893330c349ce3db065ba7b9e0dcaadaf7479be003.ContentRequestBuilder;
    currentPath *string;
    Extensions i1e8f5e1b0ab12ae689692569cf00409a7ffcf8119e0c731d1aab8fd8c24b9c9a.ExtensionsRequestBuilder;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    MultiValueExtendedProperties ic1605e2a56328ce22ef4db9ad89f54800a8f837cf7ba24dd304d6ddd4f629093.MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties i5f4f3512bc76131a2bfb807e4ef99b24c83995d2d107aba0fd68cf3b07f8b750.SingleValueExtendedPropertiesRequestBuilder;
}
