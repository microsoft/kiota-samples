<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\Extensions;

use Microsoft\Kiota\Abstractions\RequestOption;

class ExtensionsRequestBuilderGetRequestConfiguration 
{
    /**
     * @var array<string, string>|null $headers Request headers
    */
    public ?array $headers = null;
    
    /**
     * @var array<string, RequestOption>|null $options Request options
    */
    public ?array $options = null;
    
    /**
     * @var ExtensionsRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?ExtensionsRequestBuilderGetQueryParameters $queryParameters = null;
    
}
