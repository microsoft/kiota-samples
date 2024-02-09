<?php

namespace Microsoft\Graph\Models;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MessageRule extends Entity implements Parsable 
{
    /**
     * @var MessageRuleActions|null $actions The actions property
    */
    private ?MessageRuleActions $actions = null;
    
    /**
     * @var MessageRulePredicates|null $conditions The conditions property
    */
    private ?MessageRulePredicates $conditions = null;
    
    /**
     * @var string|null $displayName The display name of the rule.
    */
    private ?string $displayName = null;
    
    /**
     * @var MessageRulePredicates|null $exceptions The exceptions property
    */
    private ?MessageRulePredicates $exceptions = null;
    
    /**
     * @var bool|null $hasError Indicates whether the rule is in an error condition. Read-only.
    */
    private ?bool $hasError = null;
    
    /**
     * @var bool|null $isEnabled Indicates whether the rule is enabled to be applied to messages.
    */
    private ?bool $isEnabled = null;
    
    /**
     * @var bool|null $isReadOnly Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
    */
    private ?bool $isReadOnly = null;
    
    /**
     * @var int|null $sequence Indicates the order in which the rule is executed, among other rules.
    */
    private ?int $sequence = null;
    
    /**
     * Instantiates a new MessageRule and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param ParseNode $parseNode The parse node to use to read the discriminator value and create the object
     * @return MessageRule
    */
    public static function createFromDiscriminatorValue(ParseNode $parseNode): MessageRule {
        return new MessageRule();
    }

    /**
     * Gets the actions property value. The actions property
     * @return MessageRuleActions|null
    */
    public function getActions(): ?MessageRuleActions {
        return $this->actions;
    }

    /**
     * Gets the conditions property value. The conditions property
     * @return MessageRulePredicates|null
    */
    public function getConditions(): ?MessageRulePredicates {
        return $this->conditions;
    }

    /**
     * Gets the displayName property value. The display name of the rule.
     * @return string|null
    */
    public function getDisplayName(): ?string {
        return $this->displayName;
    }

    /**
     * Gets the exceptions property value. The exceptions property
     * @return MessageRulePredicates|null
    */
    public function getExceptions(): ?MessageRulePredicates {
        return $this->exceptions;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable(ParseNode): void>
    */
    public function getFieldDeserializers(): array {
        $o = $this;
        return array_merge(parent::getFieldDeserializers(), [
            'actions' => fn(ParseNode $n) => $o->setActions($n->getObjectValue([MessageRuleActions::class, 'createFromDiscriminatorValue'])),
            'conditions' => fn(ParseNode $n) => $o->setConditions($n->getObjectValue([MessageRulePredicates::class, 'createFromDiscriminatorValue'])),
            'displayName' => fn(ParseNode $n) => $o->setDisplayName($n->getStringValue()),
            'exceptions' => fn(ParseNode $n) => $o->setExceptions($n->getObjectValue([MessageRulePredicates::class, 'createFromDiscriminatorValue'])),
            'hasError' => fn(ParseNode $n) => $o->setHasError($n->getBooleanValue()),
            'isEnabled' => fn(ParseNode $n) => $o->setIsEnabled($n->getBooleanValue()),
            'isReadOnly' => fn(ParseNode $n) => $o->setIsReadOnly($n->getBooleanValue()),
            'sequence' => fn(ParseNode $n) => $o->setSequence($n->getIntegerValue()),
        ]);
    }

    /**
     * Gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     * @return bool|null
    */
    public function getHasError(): ?bool {
        return $this->hasError;
    }

    /**
     * Gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     * @return bool|null
    */
    public function getIsEnabled(): ?bool {
        return $this->isEnabled;
    }

    /**
     * Gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     * @return bool|null
    */
    public function getIsReadOnly(): ?bool {
        return $this->isReadOnly;
    }

    /**
     * Gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     * @return int|null
    */
    public function getSequence(): ?int {
        return $this->sequence;
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeObjectValue('actions', $this->getActions());
        $writer->writeObjectValue('conditions', $this->getConditions());
        $writer->writeStringValue('displayName', $this->getDisplayName());
        $writer->writeObjectValue('exceptions', $this->getExceptions());
        $writer->writeBooleanValue('hasError', $this->getHasError());
        $writer->writeBooleanValue('isEnabled', $this->getIsEnabled());
        $writer->writeBooleanValue('isReadOnly', $this->getIsReadOnly());
        $writer->writeIntegerValue('sequence', $this->getSequence());
    }

    /**
     * Sets the actions property value. The actions property
     * @param MessageRuleActions|null $value Value to set for the actions property.
    */
    public function setActions(?MessageRuleActions $value): void {
        $this->actions = $value;
    }

    /**
     * Sets the conditions property value. The conditions property
     * @param MessageRulePredicates|null $value Value to set for the conditions property.
    */
    public function setConditions(?MessageRulePredicates $value): void {
        $this->conditions = $value;
    }

    /**
     * Sets the displayName property value. The display name of the rule.
     * @param string|null $value Value to set for the displayName property.
    */
    public function setDisplayName(?string $value): void {
        $this->displayName = $value;
    }

    /**
     * Sets the exceptions property value. The exceptions property
     * @param MessageRulePredicates|null $value Value to set for the exceptions property.
    */
    public function setExceptions(?MessageRulePredicates $value): void {
        $this->exceptions = $value;
    }

    /**
     * Sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     * @param bool|null $value Value to set for the hasError property.
    */
    public function setHasError(?bool $value): void {
        $this->hasError = $value;
    }

    /**
     * Sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     * @param bool|null $value Value to set for the isEnabled property.
    */
    public function setIsEnabled(?bool $value): void {
        $this->isEnabled = $value;
    }

    /**
     * Sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     * @param bool|null $value Value to set for the isReadOnly property.
    */
    public function setIsReadOnly(?bool $value): void {
        $this->isReadOnly = $value;
    }

    /**
     * Sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     * @param int|null $value Value to set for the sequence property.
    */
    public function setSequence(?int $value): void {
        $this->sequence = $value;
    }

}
