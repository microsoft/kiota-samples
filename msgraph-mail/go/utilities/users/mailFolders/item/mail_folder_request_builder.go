package item

import (
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    i46b27c2daa5ab011537ca6c6735bbeb91b703573a7e8811db086394f28a7c08d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/singleValueExtendedProperties/item"
    i8c9f45a1979d17736366410c1baec3d7f74b39c26bfbd40c780cc81b23f5b339 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/item"
    ic702e2bf4a1e6b8b51abfcbc68be683ecb77f09fd7abb67ebb3b5c78aed52b74 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/multiValueExtendedProperties/item"
    i95222651871325313cff161967280a9b01f5d7367b2d6a065b9e14980afcbc8e "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/childFolders/item"
    iaae208e1698c8160302fd5dba8bd6c3477bd40752beebb79543261a210f982cc "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messageRules/item"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

type MailFolderRequestBuilder struct {
    ChildFolders ibd13cbcfdb84c75317c8a2882bb01a9f7f51b0bfe8138026b3d0416273111c4f.ChildFoldersRequestBuilder;
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    MessageRules i9e312e033b16a91798bcf5116b4d245cf441f0bf87b2306548ff0e2d5a3718d3.MessageRulesRequestBuilder;
    Messages if8618aabf6f4cc05c99484a216fa20689c48344080f5d6188f8b0a323306cf5c.MessagesRequestBuilder;
    MultiValueExtendedProperties ica98007d3c5deed2de76bd7cfc10d581874ff36b8e492eee4c91119577a657a6.MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties i3ac9377585b061690c98820e39e5a54c9650864b2f9e476a0e8b4e2022464de6.SingleValueExtendedPropertiesRequestBuilder;
}
