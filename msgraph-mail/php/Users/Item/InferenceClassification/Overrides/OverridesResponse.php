<?php

namespace Microsoft\Graph\Users\Item\InferenceClassification\Overrides;

use Microsoft\Graph\Models\Microsoft\Graph\InferenceClassificationOverride;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class OverridesResponse 
{
    private array $additionalData;
    
    /** @var string $nextLink  */
    private string $nextLink;
    
    /** @var InferenceClassificationOverride $value  */
    private InferenceClassificationOverride $value;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, object>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the nextLink property value. 
     * @return string|null
    */
    public function getNextLink(): ?string {
        return $this->nextLink;
    }

    /**
     * Gets the value property value. 
     * @return InferenceClassificationOverride|null
    */
    public function getValue(): ?InferenceClassificationOverride {
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
     * @param SerializationWriter $writer $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        $writer->writeStringValue('@odata.nextLink', $this->nextLink);
        $writer->writeCollectionOfObjectValues('value', $this->value);
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
     * Sets the nextLink property value. 
     *  @param string|null $value Value to set for the nextLink property.
    */
    public function setNextLink(?string $value): void {
        $this->nextLink = $value;
    }

    /**
     * Sets the value property value. 
     *  @param InferenceClassificationOverride|null $value Value to set for the value property.
    */
    public function setValue(?InferenceClassificationOverride $value): void {
        $this->value = $value;
    }

}
