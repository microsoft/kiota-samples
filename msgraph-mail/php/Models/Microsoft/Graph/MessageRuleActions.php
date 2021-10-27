<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MessageRuleActions 
{
    private array $additionalData;
    
    /** @var string $assignCategories A list of categories to be assigned to a message. */
    private string $assignCategories;
    
    /** @var string $copyToFolder The ID of a folder that a message is to be copied to. */
    private string $copyToFolder;
    
    /** @var bool $delete Indicates whether a message should be moved to the Deleted Items folder. */
    private bool $delete;
    
    /** @var Recipient $forwardAsAttachmentTo The email addresses of the recipients to which a message should be forwarded as an attachment. */
    private Recipient $forwardAsAttachmentTo;
    
    /** @var Recipient $forwardTo The email addresses of the recipients to which a message should be forwarded. */
    private Recipient $forwardTo;
    
    /** @var bool $markAsRead Indicates whether a message should be marked as read. */
    private bool $markAsRead;
    
    /** @var Importance $markImportance  */
    private Importance $markImportance;
    
    /** @var string $moveToFolder The ID of the folder that a message will be moved to. */
    private string $moveToFolder;
    
    /** @var bool $permanentDelete Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder. */
    private bool $permanentDelete;
    
    /** @var Recipient $redirectTo The email address to which a message should be redirected. */
    private Recipient $redirectTo;
    
    /** @var bool $stopProcessingRules Indicates whether subsequent rules should be evaluated. */
    private bool $stopProcessingRules;
    
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, object>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the assignCategories property value. A list of categories to be assigned to a message.
     * @return string|null
    */
    public function getAssignCategories(): ?string {
        return $this->assignCategories;
    }

    /**
     * Gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
     * @return string|null
    */
    public function getCopyToFolder(): ?string {
        return $this->copyToFolder;
    }

    /**
     * Gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
     * @return bool|null
    */
    public function getDelete(): ?bool {
        return $this->delete;
    }

    /**
     * Gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
     * @return Recipient|null
    */
    public function getForwardAsAttachmentTo(): ?Recipient {
        return $this->forwardAsAttachmentTo;
    }

    /**
     * Gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     * @return Recipient|null
    */
    public function getForwardTo(): ?Recipient {
        return $this->forwardTo;
    }

    /**
     * Gets the markAsRead property value. Indicates whether a message should be marked as read.
     * @return bool|null
    */
    public function getMarkAsRead(): ?bool {
        return $this->markAsRead;
    }

    /**
     * Gets the markImportance property value. 
     * @return Importance|null
    */
    public function getMarkImportance(): ?Importance {
        return $this->markImportance;
    }

    /**
     * Gets the moveToFolder property value. The ID of the folder that a message will be moved to.
     * @return string|null
    */
    public function getMoveToFolder(): ?string {
        return $this->moveToFolder;
    }

    /**
     * Gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     * @return bool|null
    */
    public function getPermanentDelete(): ?bool {
        return $this->permanentDelete;
    }

    /**
     * Gets the redirectTo property value. The email address to which a message should be redirected.
     * @return Recipient|null
    */
    public function getRedirectTo(): ?Recipient {
        return $this->redirectTo;
    }

    /**
     * Gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
     * @return bool|null
    */
    public function getStopProcessingRules(): ?bool {
        return $this->stopProcessingRules;
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
     * @param SerializationWriter $writer $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        $writer->writeCollectionOfObjectValues('assignCategories', $this->assignCategories);
        $writer->writeStringValue('copyToFolder', $this->copyToFolder);
        $writer->writeBooleanValue('delete', $this->delete);
        $writer->writeCollectionOfObjectValues('forwardAsAttachmentTo', $this->forwardAsAttachmentTo);
        $writer->writeCollectionOfObjectValues('forwardTo', $this->forwardTo);
        $writer->writeBooleanValue('markAsRead', $this->markAsRead);
        $writer->writeEnumValue('markImportance', $this->markImportance);
        $writer->writeStringValue('moveToFolder', $this->moveToFolder);
        $writer->writeBooleanValue('permanentDelete', $this->permanentDelete);
        $writer->writeCollectionOfObjectValues('redirectTo', $this->redirectTo);
        $writer->writeBooleanValue('stopProcessingRules', $this->stopProcessingRules);
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
     * Sets the assignCategories property value. A list of categories to be assigned to a message.
     *  @param string|null $value Value to set for the assignCategories property.
    */
    public function setAssignCategories(?string $value): void {
        $this->assignCategories = $value;
    }

    /**
     * Sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
     *  @param string|null $value Value to set for the copyToFolder property.
    */
    public function setCopyToFolder(?string $value): void {
        $this->copyToFolder = $value;
    }

    /**
     * Sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
     *  @param bool|null $value Value to set for the delete property.
    */
    public function setDelete(?bool $value): void {
        $this->delete = $value;
    }

    /**
     * Sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
     *  @param Recipient|null $value Value to set for the forwardAsAttachmentTo property.
    */
    public function setForwardAsAttachmentTo(?Recipient $value): void {
        $this->forwardAsAttachmentTo = $value;
    }

    /**
     * Sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     *  @param Recipient|null $value Value to set for the forwardTo property.
    */
    public function setForwardTo(?Recipient $value): void {
        $this->forwardTo = $value;
    }

    /**
     * Sets the markAsRead property value. Indicates whether a message should be marked as read.
     *  @param bool|null $value Value to set for the markAsRead property.
    */
    public function setMarkAsRead(?bool $value): void {
        $this->markAsRead = $value;
    }

    /**
     * Sets the markImportance property value. 
     *  @param Importance|null $value Value to set for the markImportance property.
    */
    public function setMarkImportance(?Importance $value): void {
        $this->markImportance = $value;
    }

    /**
     * Sets the moveToFolder property value. The ID of the folder that a message will be moved to.
     *  @param string|null $value Value to set for the moveToFolder property.
    */
    public function setMoveToFolder(?string $value): void {
        $this->moveToFolder = $value;
    }

    /**
     * Sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     *  @param bool|null $value Value to set for the permanentDelete property.
    */
    public function setPermanentDelete(?bool $value): void {
        $this->permanentDelete = $value;
    }

    /**
     * Sets the redirectTo property value. The email address to which a message should be redirected.
     *  @param Recipient|null $value Value to set for the redirectTo property.
    */
    public function setRedirectTo(?Recipient $value): void {
        $this->redirectTo = $value;
    }

    /**
     * Sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
     *  @param bool|null $value Value to set for the stopProcessingRules property.
    */
    public function setStopProcessingRules(?bool $value): void {
        $this->stopProcessingRules = $value;
    }

}
