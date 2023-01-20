<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\SingleValueExtendedProperties\Item;

use Microsoft\Kiota\Abstractions\RequestHeaders;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class SingleValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration 
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
     * @var SingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?SingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters $queryParameters = null;
    
}
