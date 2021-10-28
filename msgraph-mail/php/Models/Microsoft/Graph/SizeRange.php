<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class SizeRange 
{
    private array $additionalData;
    
    /** @var int $maximumSize The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. */
    private int $maximumSize;
    
    /** @var int $minimumSize The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. */
    private int $minimumSize;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, mixed>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     * @return int|null
    */
    public function getMaximumSize(): ?int {
        return $this->maximumSize;
    }

    /**
     * Gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     * @return int|null
    */
    public function getMinimumSize(): ?int {
        return $this->minimumSize;
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
        $writer->writeObjectValue('maximumSize', $this->maximumSize);
        $writer->writeObjectValue('minimumSize', $this->minimumSize);
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
     * Sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     *  @param int|null $value Value to set for the maximumSize property.
    */
    public function setMaximumSize(?int $value): void {
        $this->maximumSize = $value;
    }

    /**
     * Sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
     *  @param int|null $value Value to set for the minimumSize property.
    */
    public function setMinimumSize(?int $value): void {
        $this->minimumSize = $value;
    }

}
