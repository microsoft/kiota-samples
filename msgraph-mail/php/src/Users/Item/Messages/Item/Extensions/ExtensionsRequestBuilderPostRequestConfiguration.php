<?php

namespace Microsoft\Graph\Users\Item\Messages\Item\Extensions;

use Microsoft\Kiota\Abstractions\RequestOption;

class ExtensionsRequestBuilderPostRequestConfiguration 
{
    /**
     * @var array<string, string>|null $headers Request headers
    */
    public ?array $headers = null;
    
    /**
     * @var array<string, RequestOption>|null $options Request options
    */
    public ?array $options = null;
    
}
