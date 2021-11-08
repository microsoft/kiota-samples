<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MentionsPreview implements Parsable 
{
    /** @var array $AdditionalData Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private array $additionalData;
    
    /** @var bool|null $isMentioned True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter. */
    private ?bool $isMentioned;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, mixed>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the isMentioned property value. True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
     * @return bool|null
    */
    public function getIsMentioned(): ?bool {
        return $this->isMentioned;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return  [
            'isMentioned' => function (MentionsPreview $o, bool $n) { $o->setIsMentioned($n); },
        ];
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        $writer->writeBooleanValue('isMentioned', $this->isMentioned);
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
     * Sets the isMentioned property value. True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
     *  @param bool|null $value Value to set for the isMentioned property.
    */
    public function setIsMentioned(?bool $value): void {
        $this->isMentioned = $value;
    }

}
