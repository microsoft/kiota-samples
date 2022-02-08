<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MultiValueLegacyExtendedProperty extends Entity 
{
    /** @var array<string>|null $value A collection of property values. */
    private ?array $value;
    
    /**
     * Instantiates a new multiValueLegacyExtendedProperty and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Gets the value property value. A collection of property values.
     * @return array<string>|null
    */
    public function getValue(): ?array {
        return $this->value;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return array_merge(parent::getFieldDeserializers(), [
            'value' => function (self $o, ParseNode $n) { $o->setValue($n->getCollectionOfPrimitiveValues()); },
        ]);
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeCollectionOfNonParsableObjectValues('value', $this->value);
    }

    /**
     * Sets the value property value. A collection of property values.
     *  @param array<string>|null $value Value to set for the value property.
    */
    public function setValue(?array $value ): void {
        $this->value = $value;
    }

}
