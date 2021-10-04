<?php

namespace Microsoft\Graph\Users;

use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class SingleValueLegacyExtendedProperty extends Entity 
{
    /** @var string $value A property value. */
    private string $value;
    
    /**
     * Gets the value property value. A property value.
     * @return string|null
    */
    public function getValue(): ?string {
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
     * @param ?ISerializationWriter $writer $writer Serialization writer to use to serialize this model
    */
    public function serialize(?ISerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeStringValue('value', $this->value);
    }

    /**
     * Sets the value property value. A property value.
     *  @param string|null $value Value to set for the value property.
    */
    public function setValue(string $value): void {
        $this->value = $value;
    }

}
