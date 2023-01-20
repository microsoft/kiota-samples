<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders\Item;

use Microsoft\Kiota\Abstractions\RequestHeaders;
use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class MailFolderItemRequestBuilderGetRequestConfiguration 
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
     * @var MailFolderItemRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?MailFolderItemRequestBuilderGetQueryParameters $queryParameters = null;
    
}
