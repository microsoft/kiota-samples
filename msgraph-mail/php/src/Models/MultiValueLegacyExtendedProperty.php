<?php

namespace Microsoft\Graph\Models;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MultiValueLegacyExtendedProperty extends Entity implements Parsable 
{
    /**
     * @var array<string>|null $value A collection of property values.
    */
    private ?array $value = null;
    
    /**
     * Instantiates a new MultiValueLegacyExtendedProperty and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param ParseNode $parseNode The parse node to use to read the discriminator value and create the object
     * @return MultiValueLegacyExtendedProperty
    */
    public static function createFromDiscriminatorValue(ParseNode $parseNode): MultiValueLegacyExtendedProperty {
        return new MultiValueLegacyExtendedProperty();
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        $o = $this;
        return array_merge(parent::getFieldDeserializers(), [
            'value' => fn(ParseNode $n) => $o->setValue($n->getCollectionOfPrimitiveValues()),
        ]);
    }

    /**
     * Gets the value property value. A collection of property values.
     * @return array<string>|null
    */
    public function getValue(): ?array {
        return $this->value;
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeCollectionOfPrimitiveValues('value', $this->value);
    }

    /**
     * Sets the value property value. A collection of property values.
     *  @param array<string>|null $value Value to set for the value property.
    */
    public function setValue(?array $value ): void {
        $this->value = $value;
    }

}
