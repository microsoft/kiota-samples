<?php

namespace Microsoft\Graph\Models\Microsoft\Graph;

use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;

class MailFolder extends Entity implements Parsable 
{
    /** @var int $childFolderCount The number of immediate child mailFolders in the current mailFolder. */
    private int $childFolderCount;
    
    /** @var array $childFolders The collection of child folders in the mailFolder. */
    private array $childFolders;
    
    /** @var string $displayName The mailFolder's display name. */
    private string $displayName;
    
    /** @var bool $isHidden Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders. */
    private bool $isHidden;
    
    /** @var array $messageRules The collection of rules that apply to the user's Inbox folder. */
    private array $messageRules;
    
    /** @var array $messages The collection of messages in the mailFolder. */
    private array $messages;
    
    /** @var array $multiValueExtendedProperties The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable. */
    private array $multiValueExtendedProperties;
    
    /** @var string $parentFolderId The unique identifier for the mailFolder's parent mailFolder. */
    private string $parentFolderId;
    
    /** @var array $singleValueExtendedProperties The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable. */
    private array $singleValueExtendedProperties;
    
    /** @var int $totalItemCount The number of items in the mailFolder. */
    private int $totalItemCount;
    
    /** @var int $unreadItemCount The number of items in the mailFolder marked as unread. */
    private int $unreadItemCount;
    
    /**
     * Gets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
     * @return int|null
    */
    public function getChildFolderCount(): ?int {
        return $this->childFolderCount;
    }

    /**
     * Gets the childFolders property value. The collection of child folders in the mailFolder.
     * @return array|null
    */
    public function getChildFolders(): ?array {
        return $this->childFolders;
    }

    /**
     * Gets the displayName property value. The mailFolder's display name.
     * @return string|null
    */
    public function getDisplayName(): ?string {
        return $this->displayName;
    }

    /**
     * Gets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
     * @return bool|null
    */
    public function getIsHidden(): ?bool {
        return $this->isHidden;
    }

    /**
     * Gets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
     * @return array|null
    */
    public function getMessageRules(): ?array {
        return $this->messageRules;
    }

    /**
     * Gets the messages property value. The collection of messages in the mailFolder.
     * @return array|null
    */
    public function getMessages(): ?array {
        return $this->messages;
    }

    /**
     * Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @return array|null
    */
    public function getMultiValueExtendedProperties(): ?array {
        return $this->multiValueExtendedProperties;
    }

    /**
     * Gets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
     * @return string|null
    */
    public function getParentFolderId(): ?string {
        return $this->parentFolderId;
    }

    /**
     * Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @return array|null
    */
    public function getSingleValueExtendedProperties(): ?array {
        return $this->singleValueExtendedProperties;
    }

    /**
     * Gets the totalItemCount property value. The number of items in the mailFolder.
     * @return int|null
    */
    public function getTotalItemCount(): ?int {
        return $this->totalItemCount;
    }

    /**
     * Gets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
     * @return int|null
    */
    public function getUnreadItemCount(): ?int {
        return $this->unreadItemCount;
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
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeObjectValue('childFolderCount', $this->childFolderCount);
        $writer->writeCollectionOfObjectValues('childFolders', $this->childFolders);
        $writer->writeStringValue('displayName', $this->displayName);
        $writer->writeBooleanValue('isHidden', $this->isHidden);
        $writer->writeCollectionOfObjectValues('messageRules', $this->messageRules);
        $writer->writeCollectionOfObjectValues('messages', $this->messages);
        $writer->writeCollectionOfObjectValues('multiValueExtendedProperties', $this->multiValueExtendedProperties);
        $writer->writeStringValue('parentFolderId', $this->parentFolderId);
        $writer->writeCollectionOfObjectValues('singleValueExtendedProperties', $this->singleValueExtendedProperties);
        $writer->writeObjectValue('totalItemCount', $this->totalItemCount);
        $writer->writeObjectValue('unreadItemCount', $this->unreadItemCount);
    }

    /**
     * Sets the childFolderCount property value. The number of immediate child mailFolders in the current mailFolder.
     *  @param int|null $value Value to set for the childFolderCount property.
    */
    public function setChildFolderCount(?int $value): void {
        $this->childFolderCount = $value;
    }

    /**
     * Sets the childFolders property value. The collection of child folders in the mailFolder.
     *  @param array|null $value Value to set for the childFolders property.
    */
    public function setChildFolders(?array $value): void {
        $this->childFolders = $value;
    }

    /**
     * Sets the displayName property value. The mailFolder's display name.
     *  @param string|null $value Value to set for the displayName property.
    */
    public function setDisplayName(?string $value): void {
        $this->displayName = $value;
    }

    /**
     * Sets the isHidden property value. Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
     *  @param bool|null $value Value to set for the isHidden property.
    */
    public function setIsHidden(?bool $value): void {
        $this->isHidden = $value;
    }

    /**
     * Sets the messageRules property value. The collection of rules that apply to the user's Inbox folder.
     *  @param array|null $value Value to set for the messageRules property.
    */
    public function setMessageRules(?array $value): void {
        $this->messageRules = $value;
    }

    /**
     * Sets the messages property value. The collection of messages in the mailFolder.
     *  @param array|null $value Value to set for the messages property.
    */
    public function setMessages(?array $value): void {
        $this->messages = $value;
    }

    /**
     * Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     *  @param array|null $value Value to set for the multiValueExtendedProperties property.
    */
    public function setMultiValueExtendedProperties(?array $value): void {
        $this->multiValueExtendedProperties = $value;
    }

    /**
     * Sets the parentFolderId property value. The unique identifier for the mailFolder's parent mailFolder.
     *  @param string|null $value Value to set for the parentFolderId property.
    */
    public function setParentFolderId(?string $value): void {
        $this->parentFolderId = $value;
    }

    /**
     * Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     *  @param array|null $value Value to set for the singleValueExtendedProperties property.
    */
    public function setSingleValueExtendedProperties(?array $value): void {
        $this->singleValueExtendedProperties = $value;
    }

    /**
     * Sets the totalItemCount property value. The number of items in the mailFolder.
     *  @param int|null $value Value to set for the totalItemCount property.
    */
    public function setTotalItemCount(?int $value): void {
        $this->totalItemCount = $value;
    }

    /**
     * Sets the unreadItemCount property value. The number of items in the mailFolder marked as unread.
     *  @param int|null $value Value to set for the unreadItemCount property.
    */
    public function setUnreadItemCount(?int $value): void {
        $this->unreadItemCount = $value;
    }

}
