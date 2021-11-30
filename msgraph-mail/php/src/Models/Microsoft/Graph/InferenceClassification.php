<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class InferenceClassification extends Entity implements Parsable 
{
    /** @var array<InferenceClassificationOverride>|null $overrides A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable. */
    private ?array $overrides;
    
    /**
     * Instantiates a new inferenceClassification and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Gets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @return array<InferenceClassificationOverride>|null
    */
    public function getOverrides(): ?array {
        return $this->overrides;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return array_merge(parent::getFieldDeserializers(), [
            'overrides' => function (InferenceClassification $o, array $n) { $o->setOverrides($n); },
        ]);
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeCollectionOfObjectValues('overrides', $this->overrides);
    }

    /**
     * Sets the overrides property value. A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     *  @param array|null $value Value to set for the overrides property.
    */
    public function setOverrides(?array $value): void {
        $this->overrides = $value;
    }

}
