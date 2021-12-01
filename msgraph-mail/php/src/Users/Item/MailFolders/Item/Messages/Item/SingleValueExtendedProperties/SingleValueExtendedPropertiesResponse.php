<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\SingleValueExtendedProperties;

use Microsoft\Graph\Models\Microsoft\Graph\SingleValueLegacyExtendedProperty;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class SingleValueExtendedPropertiesResponse implements Parsable 
{
    /** @var array<string, mixed> $AdditionalData Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private array $additionalData;
    
    /** @var string|null $nextLink  */
    private ?string $nextLink;
    
    /** @var array<SingleValueLegacyExtendedProperty>|null $value  */
    private ?array $value;
    
    /**
     * Instantiates a new singleValueExtendedPropertiesResponse and sets the default values.
    */
    public function __construct() {
        $this->additionalData = [];
    }

    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, mixed>
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
     * @return array<SingleValueLegacyExtendedProperty>|null
    */
    public function getValue(): ?array {
        return $this->value;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return  [
            '@odata.nextLink' => function (SingleValueExtendedPropertiesResponse $o, string $n) { $o->setNextLink($n); },
            'value' => function (SingleValueExtendedPropertiesResponse $o, array $n) { $o->setValue($n); },
        ];
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        $writer->writeStringValue('@odata.nextLink', $this->nextLink);
        $writer->writeCollectionOfObjectValues('value', $this->value);
        $writer->writeAdditionalData($this->additionalData);
    }

    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     *  @param array<string,mixed> $value Value to set for the AdditionalData property.
    */
    public function setAdditionalData(?array $value ): void {
        $this->additionalData = $value;
    }

    /**
     * Sets the nextLink property value. 
     *  @param string|null $value Value to set for the nextLink property.
    */
    public function setNextLink(?string $value ): void {
        $this->nextLink = $value;
    }

    /**
     * Sets the value property value. 
     *  @param array<SingleValueLegacyExtendedProperty>|null $value Value to set for the value property.
    */
    public function setValue(?array $value ): void {
        $this->value = $value;
    }

}
