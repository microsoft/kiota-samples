<?php

namespace Microsoft\Graph\Models;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class InferenceClassificationOverride extends Entity implements Parsable 
{
    /**
     * @var InferenceClassificationType|null $classifyAs The classifyAs property
    */
    private ?InferenceClassificationType $classifyAs = null;
    
    /**
     * @var EmailAddress|null $senderEmailAddress The senderEmailAddress property
    */
    private ?EmailAddress $senderEmailAddress = null;
    
    /**
     * Instantiates a new InferenceClassificationOverride and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param ParseNode $parseNode The parse node to use to read the discriminator value and create the object
     * @return InferenceClassificationOverride
    */
    public static function createFromDiscriminatorValue(ParseNode $parseNode): InferenceClassificationOverride {
        return new InferenceClassificationOverride();
    }

    /**
     * Gets the classifyAs property value. The classifyAs property
     * @return InferenceClassificationType|null
    */
    public function getClassifyAs(): ?InferenceClassificationType {
        return $this->classifyAs;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable(ParseNode): void>
    */
    public function getFieldDeserializers(): array {
        $o = $this;
        return array_merge(parent::getFieldDeserializers(), [
            'classifyAs' => fn(ParseNode $n) => $o->setClassifyAs($n->getEnumValue(InferenceClassificationType::class)),
            'senderEmailAddress' => fn(ParseNode $n) => $o->setSenderEmailAddress($n->getObjectValue([EmailAddress::class, 'createFromDiscriminatorValue'])),
        ]);
    }

    /**
     * Gets the senderEmailAddress property value. The senderEmailAddress property
     * @return EmailAddress|null
    */
    public function getSenderEmailAddress(): ?EmailAddress {
        return $this->senderEmailAddress;
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeEnumValue('classifyAs', $this->getClassifyAs());
        $writer->writeObjectValue('senderEmailAddress', $this->getSenderEmailAddress());
    }

    /**
     * Sets the classifyAs property value. The classifyAs property
     * @param InferenceClassificationType|null $value Value to set for the classifyAs property.
    */
    public function setClassifyAs(?InferenceClassificationType $value): void {
        $this->classifyAs = $value;
    }

    /**
     * Sets the senderEmailAddress property value. The senderEmailAddress property
     * @param EmailAddress|null $value Value to set for the senderEmailAddress property.
    */
    public function setSenderEmailAddress(?EmailAddress $value): void {
        $this->senderEmailAddress = $value;
    }

}
