package inferenceclassification

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    id37525301f612e8f9f28658e2ceeb7e0dbca43868161cbacdbd261efc2c213df "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification/overrides"
    ic84e53bbe0d89365c60cb7ce42953129b0a25b34d24d19255e83087a1152ecff "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification/overrides/item"
)

type InferenceClassificationRequestBuilder struct {
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    Overrides id37525301f612e8f9f28658e2ceeb7e0dbca43868161cbacdbd261efc2c213df.OverridesRequestBuilder;
    pathSegment *string;
}
