package overrides

import (
    i90ce7b240dbfe27d1efd75544859405d70943a52fa428ec2a2f23fa14d2a50a0 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification"
)

type OverridesResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []i90ce7b240dbfe27d1efd75544859405d70943a52fa428ec2a2f23fa14d2a50a0.InferenceClassificationOverride;
}
