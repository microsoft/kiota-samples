<?php

namespace Microsoft\Graph\Users\Item\InferenceClassification\Overrides;

use Microsoft\Kiota\Abstractions\RequestHeaders;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class OverridesRequestBuilderGetRequestConfiguration 
{
    /**
     * @var RequestHeaders|null $headers Request headers
    */
    public ?RequestHeaders $headers = null;
    
    /**
     * @var array<string, RequestOption>|null $options Request options
    */
    public ?array $options = null;
    
    /**
     * @var OverridesRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?OverridesRequestBuilderGetQueryParameters $queryParameters = null;
    
}
