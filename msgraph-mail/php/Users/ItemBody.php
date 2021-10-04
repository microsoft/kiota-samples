<?php

namespace Microsoft\Graph\Users;

use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class ItemBody 
{
    private array $additionalData;
    
    /** @var string $content The content of the item. */
    private string $content;
    
    /** @var BodyType $contentType  */
    private BodyType $contentType;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, object>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the content property value. The content of the item.
     * @return string|null
    */
    public function getContent(): ?string {
        return $this->content;
    }

    /**
     * Gets the contentType property value. 
     * @return BodyType|null
    */
    public function getContentType(): ?BodyType {
        return $this->contentType;
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
        $writer->writeStringValue('content', $this->content);
        $writer->writeEnumValue('contentType', $this->contentType);
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
     * Sets the content property value. The content of the item.
     *  @param string|null $value Value to set for the content property.
    */
    public function setContent(string $value): void {
        $this->content = $value;
    }

    /**
     * Sets the contentType property value. 
     *  @param BodyType|null $value Value to set for the contentType property.
    */
    public function setContentType(BodyType $value): void {
        $this->contentType = $value;
    }

}
