<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders\Item\SingleValueExtendedProperties\Item;

use Microsoft\Kiota\Abstractions\QueryParameter;

class SingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters 
{
    /**
     * @QueryParameter("%24expand")
     * @var array<string>|null $expand Expand related entities
    */
    public ?array $expand = null;
    
    /**
     * @QueryParameter("%24select")
     * @var array<string>|null $select Select properties to be returned
    */
    public ?array $select = null;
    
}
