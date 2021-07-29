package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    i0686696817f5336ad23f5048aaf1973675e26317e928255ef44ff8f6cdd744ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/attachments"
    i7a8f25a40ec18f53dd1fe13983249e861ca16f093020bf1d1f77bdea2b68285f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/content"
    ib2381cda6aa2598b763d17161ebaa5c941e140eebfa25805cfbb21dbd2ff0805 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/extensions"
    id35afca47a5e5a75ec80ed3fe0e8a5bddf19271c23cfbd9f6678773404139756 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/multivalueextendedproperties"
    i514af90a8d1241ee2aec5dae95b1b4028641b8511c05b3b7710d33898668b1c0 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/singlevalueextendedproperties"
    ib74bdf213c8ecff00ba896023527beb4a5cee0f1561b4ebb235535adb8fd01ba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/attachments/item"
    i4c213381c08f42149bb0acad01e026154f1d08ef5829ececbd8ea51281fe708b "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/extensions/item"
    i54bf43874c83bde999c0160122d8c20bde3e0dfabeec784a4ae151188b4686c4 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/multivalueextendedproperties/item"
    i4dc594f64060f7179afa645e9dc31eac49b9bb535bf1df6fe899682087dc226c "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/singlevalueextendedproperties/item"
)

type MessageRequestBuilder struct {
    Attachments i0686696817f5336ad23f5048aaf1973675e26317e928255ef44ff8f6cdd744ce.AttachmentsRequestBuilder;
    Content i7a8f25a40ec18f53dd1fe13983249e861ca16f093020bf1d1f77bdea2b68285f.ContentRequestBuilder;
    currentPath *string;
    Extensions ib2381cda6aa2598b763d17161ebaa5c941e140eebfa25805cfbb21dbd2ff0805.ExtensionsRequestBuilder;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    MultiValueExtendedProperties id35afca47a5e5a75ec80ed3fe0e8a5bddf19271c23cfbd9f6678773404139756.MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties i514af90a8d1241ee2aec5dae95b1b4028641b8511c05b3b7710d33898668b1c0.SingleValueExtendedPropertiesRequestBuilder;
}
