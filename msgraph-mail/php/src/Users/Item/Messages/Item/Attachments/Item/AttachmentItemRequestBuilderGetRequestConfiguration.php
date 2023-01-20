<?php

namespace Microsoft\Graph\Users\Item\Messages\Item\Attachments\Item;

use Microsoft\Kiota\Abstractions\RequestHeaders;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class AttachmentItemRequestBuilderGetRequestConfiguration 
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
     * @var AttachmentItemRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?AttachmentItemRequestBuilderGetQueryParameters $queryParameters = null;
    
}
