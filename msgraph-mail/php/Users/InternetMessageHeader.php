<?php

namespace Microsoft\Graph\Users;

use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class InternetMessageHeader 
{
    private array $additionalData;
    
    /** @var string $name Represents the key in a key-value pair. */
    private string $name;
    
    /** @var string $value The value in a key-value pair. */
    private string $value;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, object>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the name property value. Represents the key in a key-value pair.
     * @return string|null
    */
    public function getName(): ?string {
        return $this->name;
    }

    /**
     * Gets the value property value. The value in a key-value pair.
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
        $writer->writeStringValue('name', $this->name);
        $writer->writeStringValue('value', $this->value);
        $writer->writeAdditionalData($this->additionalData);
    }

    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     *  @param array<string,object> $value Value to set for the AdditionalData property.
    */
    public function setAdditionalData(array $value): void {
        $this->additionalData = $value;
    }

    /**
     * Sets the name property value. Represents the key in a key-value pair.
     *  @param string|null $value Value to set for the name property.
    */
    public function setName(string $value): void {
        $this->name = $value;
    }

    /**
     * Sets the value property value. The value in a key-value pair.
     *  @param string|null $value Value to set for the value property.
    */
    public function setValue(string $value): void {
        $this->value = $value;
    }

}
