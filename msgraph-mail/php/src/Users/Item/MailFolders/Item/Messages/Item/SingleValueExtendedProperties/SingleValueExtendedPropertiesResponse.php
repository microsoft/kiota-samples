<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\SingleValueExtendedProperties;

use Microsoft\Graph\Models\Microsoft\Graph\SingleValueLegacyExtendedProperty;
use Microsoft\Kiota\Abstractions\Serialization\AdditionalDataHolder;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class SingleValueExtendedPropertiesResponse implements AdditionalDataHolder, Parsable 
{
    /** @var array<string, mixed> $AdditionalData Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private array $additionalData;
    
    /** @var string|null $nextLink The nextLink property */
    private ?string $nextLink = null;
    
    /** @var array<SingleValueLegacyExtendedProperty>|null $value The value property */
    private ?array $value = null;
    
    /**
     * Instantiates a new singleValueExtendedPropertiesResponse and sets the default values.
    */
    public function __construct() {
        $this->additionalData = [];
    }

    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param ParseNode $parseNode The parse node to use to read the discriminator value and create the object
     * @return SingleValueExtendedPropertiesResponse
    */
    public static function createFromDiscriminatorValue(ParseNode $parseNode): SingleValueExtendedPropertiesResponse {
        return new SingleValueExtendedPropertiesResponse();
    }

    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, mixed>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        $currentObject = $this;
        return  [
            '@odata.nextLink' => function (ParseNode $n) use ($currentObject) { $currentObject->setNextLink($n->getStringValue()); },
            'value' => function (ParseNode $n) use ($currentObject) { $currentObject->setValue($n->getCollectionOfObjectValues(array(SingleValueLegacyExtendedProperty::class, 'createFromDiscriminatorValue'))); },
        ];
    }

    /**
     * Gets the @odata.nextLink property value. The nextLink property
     * @return string|null
    */
    public function getNextLink(): ?string {
        return $this->nextLink;
    }

    /**
     * Gets the value property value. The value property
     * @return array<SingleValueLegacyExtendedProperty>|null
    */
    public function getValue(): ?array {
        return $this->value;
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
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     *  @param array<string,mixed> $value Value to set for the AdditionalData property.
    */
    public function setAdditionalData(?array $value ): void {
        $this->additionalData = $value;
    }

    /**
     * Sets the @odata.nextLink property value. The nextLink property
     *  @param string|null $value Value to set for the nextLink property.
    */
    public function setNextLink(?string $value ): void {
        $this->nextLink = $value;
    }

    /**
     * Sets the value property value. The value property
     *  @param array<SingleValueLegacyExtendedProperty>|null $value Value to set for the value property.
    */
    public function setValue(?array $value ): void {
        $this->value = $value;
    }

}
