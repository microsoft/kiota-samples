<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MessageRuleActions implements Parsable 
{
    /** @var array $AdditionalData Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    private array $additionalData;
    
    /** @var array<string>|null $assignCategories A list of categories to be assigned to a message. */
    private ?array $assignCategories;
    
    /** @var string|null $copyToFolder The ID of a folder that a message is to be copied to. */
    private ?string $copyToFolder;
    
    /** @var bool|null $delete Indicates whether a message should be moved to the Deleted Items folder. */
    private ?bool $delete;
    
    /** @var array<Recipient>|null $forwardAsAttachmentTo The email addresses of the recipients to which a message should be forwarded as an attachment. */
    private ?array $forwardAsAttachmentTo;
    
    /** @var array<Recipient>|null $forwardTo The email addresses of the recipients to which a message should be forwarded. */
    private ?array $forwardTo;
    
    /** @var bool|null $markAsRead Indicates whether a message should be marked as read. */
    private ?bool $markAsRead;
    
    /** @var Importance|null $markImportance  */
    private ?Importance $markImportance;
    
    /** @var string|null $moveToFolder The ID of the folder that a message will be moved to. */
    private ?string $moveToFolder;
    
    /** @var bool|null $permanentDelete Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder. */
    private ?bool $permanentDelete;
    
    /** @var array<Recipient>|null $redirectTo The email address to which a message should be redirected. */
    private ?array $redirectTo;
    
    /** @var bool|null $stopProcessingRules Indicates whether subsequent rules should be evaluated. */
    private ?bool $stopProcessingRules;
    
    /**
     * Instantiates a new messageRuleActions and sets the default values.
    */
    public function __construct() {
        $this->additionalData = [];
    }

    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return array<string, mixed>
    */
    public function getAdditionalData(): array {
        return $this->additionalData;
    }

    /**
     * Gets the assignCategories property value. A list of categories to be assigned to a message.
     * @return array<string>|null
    */
    public function getAssignCategories(): ?array {
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
     * @return array<Recipient>|null
    */
    public function getForwardAsAttachmentTo(): ?array {
        return $this->forwardAsAttachmentTo;
    }

    /**
     * Gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     * @return array<Recipient>|null
    */
    public function getForwardTo(): ?array {
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
     * @return array<Recipient>|null
    */
    public function getRedirectTo(): ?array {
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
        return  [
            'assignCategories' => function (MessageRuleActions $o, array $n) { $o->setAssignCategories($n); },
            'copyToFolder' => function (MessageRuleActions $o, string $n) { $o->setCopyToFolder($n); },
            'delete' => function (MessageRuleActions $o, bool $n) { $o->setDelete($n); },
            'forwardAsAttachmentTo' => function (MessageRuleActions $o, array $n) { $o->setForwardAsAttachmentTo($n); },
            'forwardTo' => function (MessageRuleActions $o, array $n) { $o->setForwardTo($n); },
            'markAsRead' => function (MessageRuleActions $o, bool $n) { $o->setMarkAsRead($n); },
            'markImportance' => function (MessageRuleActions $o, Importance $n) { $o->setMarkImportance($n); },
            'moveToFolder' => function (MessageRuleActions $o, string $n) { $o->setMoveToFolder($n); },
            'permanentDelete' => function (MessageRuleActions $o, bool $n) { $o->setPermanentDelete($n); },
            'redirectTo' => function (MessageRuleActions $o, array $n) { $o->setRedirectTo($n); },
            'stopProcessingRules' => function (MessageRuleActions $o, bool $n) { $o->setStopProcessingRules($n); },
        ];
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
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
     *  @param array<string,mixed> $value Value to set for the AdditionalData property.
    */
    public function setAdditionalData(array $value): void {
        $this->additionalData = $value;
    }

    /**
     * Sets the assignCategories property value. A list of categories to be assigned to a message.
     *  @param array|null $value Value to set for the assignCategories property.
    */
    public function setAssignCategories(?array $value = null): void {
        $this->assignCategories = $value;
    }

    /**
     * Sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
     *  @param string|null $value Value to set for the copyToFolder property.
    */
    public function setCopyToFolder(?string $value = null): void {
        $this->copyToFolder = $value;
    }

    /**
     * Sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
     *  @param bool|null $value Value to set for the delete property.
    */
    public function setDelete(?bool $value = null): void {
        $this->delete = $value;
    }

    /**
     * Sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
     *  @param array|null $value Value to set for the forwardAsAttachmentTo property.
    */
    public function setForwardAsAttachmentTo(?array $value = null): void {
        $this->forwardAsAttachmentTo = $value;
    }

    /**
     * Sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     *  @param array|null $value Value to set for the forwardTo property.
    */
    public function setForwardTo(?array $value = null): void {
        $this->forwardTo = $value;
    }

    /**
     * Sets the markAsRead property value. Indicates whether a message should be marked as read.
     *  @param bool|null $value Value to set for the markAsRead property.
    */
    public function setMarkAsRead(?bool $value = null): void {
        $this->markAsRead = $value;
    }

    /**
     * Sets the markImportance property value. 
     *  @param Importance|null $value Value to set for the markImportance property.
    */
    public function setMarkImportance(?Importance $value = null): void {
        $this->markImportance = $value;
    }

    /**
     * Sets the moveToFolder property value. The ID of the folder that a message will be moved to.
     *  @param string|null $value Value to set for the moveToFolder property.
    */
    public function setMoveToFolder(?string $value = null): void {
        $this->moveToFolder = $value;
    }

    /**
     * Sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     *  @param bool|null $value Value to set for the permanentDelete property.
    */
    public function setPermanentDelete(?bool $value = null): void {
        $this->permanentDelete = $value;
    }

    /**
     * Sets the redirectTo property value. The email address to which a message should be redirected.
     *  @param array|null $value Value to set for the redirectTo property.
    */
    public function setRedirectTo(?array $value = null): void {
        $this->redirectTo = $value;
    }

    /**
     * Sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
     *  @param bool|null $value Value to set for the stopProcessingRules property.
    */
    public function setStopProcessingRules(?bool $value = null): void {
        $this->stopProcessingRules = $value;
    }

}
