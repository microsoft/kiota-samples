<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;
use Psr\Http\Message\StreamInterface;

class UserConfiguration extends Entity implements Parsable 
{
    /** @var Binary|null $binaryData  */
    private ?Binary $binaryData;
    
    /**
     * Gets the binaryData property value. 
     * @return StreamInterface|null
    */
    public function getBinaryData(): ?StreamInterface {
        return $this->binaryData;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return array_merge(parent::getFieldDeserializers(), [
            'binaryData' => function (UserConfiguration $o, Binary $n) { $o->setBinaryData($n); },
        ]);
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeObjectValue('binaryData', $this->binaryData);
    }

    /**
     * Sets the binaryData property value. 
     *  @param StreamInterface|null $value Value to set for the binaryData property.
    */
    public function setBinaryData(?StreamInterface $value): void {
        $this->binaryData = $value;
    }

}
