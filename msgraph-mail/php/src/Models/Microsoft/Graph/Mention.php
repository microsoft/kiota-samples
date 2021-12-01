<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class Mention extends Entity implements Parsable 
{
    /** @var string|null $application The name of the application where the mention is created. Optional. Not used and defaulted as null for message. */
    private ?string $application;
    
    /** @var string|null $clientReference A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message. */
    private ?string $clientReference;
    
    /** @var EmailAddress|null $createdBy  */
    private ?EmailAddress $createdBy;
    
    /** @var DateTimeOffset|null $createdDateTime The date and time that the mention is created on the client. */
    private ?DateTimeOffset $createdDateTime;
    
    /** @var string|null $deepLink A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message. */
    private ?string $deepLink;
    
    /** @var EmailAddress|null $mentioned  */
    private ?EmailAddress $mentioned;
    
    /** @var string|null $mentionText Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead. */
    private ?string $mentionText;
    
    /** @var DateTimeOffset|null $serverCreatedDateTime The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message. */
    private ?DateTimeOffset $serverCreatedDateTime;
    
    /**
     * Gets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
     * @return string|null
    */
    public function getApplication(): ?string {
        return $this->application;
    }

    /**
     * Gets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
     * @return string|null
    */
    public function getClientReference(): ?string {
        return $this->clientReference;
    }

    /**
     * Gets the createdBy property value. 
     * @return EmailAddress|null
    */
    public function getCreatedBy(): ?EmailAddress {
        return $this->createdBy;
    }

    /**
     * Gets the createdDateTime property value. The date and time that the mention is created on the client.
     * @return DateTimeOffset|null
    */
    public function getCreatedDateTime(): ?DateTimeOffset {
        return $this->createdDateTime;
    }

    /**
     * Gets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
     * @return string|null
    */
    public function getDeepLink(): ?string {
        return $this->deepLink;
    }

    /**
     * Gets the mentioned property value. 
     * @return EmailAddress|null
    */
    public function getMentioned(): ?EmailAddress {
        return $this->mentioned;
    }

    /**
     * Gets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
     * @return string|null
    */
    public function getMentionText(): ?string {
        return $this->mentionText;
    }

    /**
     * Gets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
     * @return DateTimeOffset|null
    */
    public function getServerCreatedDateTime(): ?DateTimeOffset {
        return $this->serverCreatedDateTime;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable>
    */
    public function getFieldDeserializers(): array {
        return array_merge(parent::getFieldDeserializers(), [
            'application' => function (Mention $o, string $n) { $o->setApplication($n); },
            'clientReference' => function (Mention $o, string $n) { $o->setClientReference($n); },
            'createdBy' => function (Mention $o, EmailAddress $n) { $o->setCreatedBy($n); },
            'createdDateTime' => function (Mention $o, DateTimeOffset $n) { $o->setCreatedDateTime($n); },
            'deepLink' => function (Mention $o, string $n) { $o->setDeepLink($n); },
            'mentioned' => function (Mention $o, EmailAddress $n) { $o->setMentioned($n); },
            'mentionText' => function (Mention $o, string $n) { $o->setMentionText($n); },
            'serverCreatedDateTime' => function (Mention $o, DateTimeOffset $n) { $o->setServerCreatedDateTime($n); },
        ]);
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeStringValue('application', $this->application);
        $writer->writeStringValue('clientReference', $this->clientReference);
        $writer->writeObjectValue('createdBy', $this->createdBy);
        $writer->writeObjectValue('createdDateTime', $this->createdDateTime);
        $writer->writeStringValue('deepLink', $this->deepLink);
        $writer->writeObjectValue('mentioned', $this->mentioned);
        $writer->writeStringValue('mentionText', $this->mentionText);
        $writer->writeObjectValue('serverCreatedDateTime', $this->serverCreatedDateTime);
    }

    /**
     * Sets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
     *  @param string|null $value Value to set for the application property.
    */
    public function setApplication(?string $value): void {
        $this->application = $value;
    }

    /**
     * Sets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
     *  @param string|null $value Value to set for the clientReference property.
    */
    public function setClientReference(?string $value): void {
        $this->clientReference = $value;
    }

    /**
     * Sets the createdBy property value. 
     *  @param EmailAddress|null $value Value to set for the createdBy property.
    */
    public function setCreatedBy(?EmailAddress $value): void {
        $this->createdBy = $value;
    }

    /**
     * Sets the createdDateTime property value. The date and time that the mention is created on the client.
     *  @param DateTimeOffset|null $value Value to set for the createdDateTime property.
    */
    public function setCreatedDateTime(?DateTimeOffset $value): void {
        $this->createdDateTime = $value;
    }

    /**
     * Sets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
     *  @param string|null $value Value to set for the deepLink property.
    */
    public function setDeepLink(?string $value): void {
        $this->deepLink = $value;
    }

    /**
     * Sets the mentioned property value. 
     *  @param EmailAddress|null $value Value to set for the mentioned property.
    */
    public function setMentioned(?EmailAddress $value): void {
        $this->mentioned = $value;
    }

    /**
     * Sets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
     *  @param string|null $value Value to set for the mentionText property.
    */
    public function setMentionText(?string $value): void {
        $this->mentionText = $value;
    }

    /**
     * Sets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
     *  @param DateTimeOffset|null $value Value to set for the serverCreatedDateTime property.
    */
    public function setServerCreatedDateTime(?DateTimeOffset $value): void {
        $this->serverCreatedDateTime = $value;
    }

}
