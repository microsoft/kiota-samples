<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\Messages;

use Microsoft\Kiota\Abstractions\RequestOption;

class MessagesRequestBuilderPostRequestConfiguration 
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
