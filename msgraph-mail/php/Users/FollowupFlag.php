<?php

namespace Microsoft\Graph\Users;

use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class FollowupFlag 
{
    private array $additionalData;
    
    /** @var DateTimeTimeZone $completedDateTime  */
    private DateTimeTimeZone $completedDateTime;
    
    /** @var DateTimeTimeZone $dueDateTime  */
    private DateTimeTimeZone $dueDateTime;
    
    /** @var FollowupFlagStatus $flagStatus  */
    private FollowupFlagStatus $flagStatus;
    
    /** @var DateTimeTimeZone $startDateTime  */
    private DateTimeTimeZone $startDateTime;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, object>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the completedDateTime property value. 
     * @return DateTimeTimeZone|null
    */
    public function getCompletedDateTime(): ?DateTimeTimeZone {
        return $this->completedDateTime;
    }

    /**
     * Gets the dueDateTime property value. 
     * @return DateTimeTimeZone|null
    */
    public function getDueDateTime(): ?DateTimeTimeZone {
        return $this->dueDateTime;
    }

    /**
     * Gets the flagStatus property value. 
     * @return FollowupFlagStatus|null
    */
    public function getFlagStatus(): ?FollowupFlagStatus {
        return $this->flagStatus;
    }

    /**
     * Gets the startDateTime property value. 
     * @return DateTimeTimeZone|null
    */
    public function getStartDateTime(): ?DateTimeTimeZone {
        return $this->startDateTime;
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
        $writer->writeObjectValue('completedDateTime', $this->completedDateTime);
        $writer->writeObjectValue('dueDateTime', $this->dueDateTime);
        $writer->writeEnumValue('flagStatus', $this->flagStatus);
        $writer->writeObjectValue('startDateTime', $this->startDateTime);
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
     * Sets the completedDateTime property value. 
     *  @param DateTimeTimeZone|null $value Value to set for the completedDateTime property.
    */
    public function setCompletedDateTime(DateTimeTimeZone $value): void {
        $this->completedDateTime = $value;
    }

    /**
     * Sets the dueDateTime property value. 
     *  @param DateTimeTimeZone|null $value Value to set for the dueDateTime property.
    */
    public function setDueDateTime(DateTimeTimeZone $value): void {
        $this->dueDateTime = $value;
    }

    /**
     * Sets the flagStatus property value. 
     *  @param FollowupFlagStatus|null $value Value to set for the flagStatus property.
    */
    public function setFlagStatus(FollowupFlagStatus $value): void {
        $this->flagStatus = $value;
    }

    /**
     * Sets the startDateTime property value. 
     *  @param DateTimeTimeZone|null $value Value to set for the startDateTime property.
    */
    public function setStartDateTime(DateTimeTimeZone $value): void {
        $this->startDateTime = $value;
    }

}
