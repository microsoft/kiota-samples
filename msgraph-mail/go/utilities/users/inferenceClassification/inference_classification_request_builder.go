package inferenceclassification

import (
    ia2b8f3935f8c022d86b38c994d4daf8b11bb39477ea39d8eda69a08652743dc0 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification/overrides"
    id74555dc918972503a06d88b0ba575fc80d906b0b8c3efebe6a207392ee1a32d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification/overrides/item"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

type InferenceClassificationRequestBuilder struct {
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    Overrides ia2b8f3935f8c022d86b38c994d4daf8b11bb39477ea39d8eda69a08652743dc0.OverridesRequestBuilder;
    pathSegment *string;
}
