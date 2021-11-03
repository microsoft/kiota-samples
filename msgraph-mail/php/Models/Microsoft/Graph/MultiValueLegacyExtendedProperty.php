<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MultiValueLegacyExtendedProperty extends Entity implements Parsable 
{
    /** @var array $value A collection of property values. */
    private array $value;
    
    /**
     * Gets the value property value. A collection of property values.
     * @return array|null
    */
    public function getValue(): ?array {
        return $this->value;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        echo('This is the body of the deserializer.');
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        if (is_null($writer)) {
            throw new \Exception('$writer cannot be null');
        }
        parent::serialize($writer);
        $writer->writeCollectionOfObjectValues('value', $this->value);
    }

    /**
     * Sets the value property value. A collection of property values.
     *  @param array|null $value Value to set for the value property.
    */
    public function setValue(?array $value): void {
        $this->value = $value;
    }

}
