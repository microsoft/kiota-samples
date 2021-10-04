<?php

namespace Microsoft\Graph\Users;

use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class Recipient 
{
    private array $additionalData;
    
    /** @var EmailAddress $emailAddress  */
    private EmailAddress $emailAddress;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, object>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the emailAddress property value. 
     * @return EmailAddress|null
    */
    public function getEmailAddress(): ?EmailAddress {
        return $this->emailAddress;
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
        $writer->writeObjectValue('emailAddress', $this->emailAddress);
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
     * Sets the emailAddress property value. 
     *  @param EmailAddress|null $value Value to set for the emailAddress property.
    */
    public function setEmailAddress(EmailAddress $value): void {
        $this->emailAddress = $value;
    }

}
