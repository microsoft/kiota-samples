<?php

namespace Microsoft\Graph\Models;

use DateTime;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Microsoft\Kiota\Abstractions\Serialization\SerializationWriter;
use Psr\Http\Message\StreamInterface;

class Message extends OutlookItem implements Parsable 
{
    /**
     * @var array<Attachment>|null $attachments The fileAttachment and itemAttachment attachments for the message.
    */
    private ?array $attachments = null;
    
    /**
     * @var array<Recipient>|null $bccRecipients The Bcc: recipients for the message.
    */
    private ?array $bccRecipients = null;
    
    /**
     * @var ItemBody|null $body The body property
    */
    private ?ItemBody $body = null;
    
    /**
     * @var string|null $bodyPreview The first 255 characters of the message body. It is in text format.
    */
    private ?string $bodyPreview = null;
    
    /**
     * @var array<Recipient>|null $ccRecipients The Cc: recipients for the message.
    */
    private ?array $ccRecipients = null;
    
    /**
     * @var string|null $conversationId The ID of the conversation the email belongs to.
    */
    private ?string $conversationId = null;
    
    /**
     * @var StreamInterface|null $conversationIndex Indicates the position of the message within the conversation.
    */
    private ?StreamInterface $conversationIndex = null;
    
    /**
     * @var array<Extension>|null $extensions The collection of open extensions defined for the message. Nullable.
    */
    private ?array $extensions = null;
    
    /**
     * @var FollowupFlag|null $flag The flag property
    */
    private ?FollowupFlag $flag = null;
    
    /**
     * @var Recipient|null $from The from property
    */
    private ?Recipient $from = null;
    
    /**
     * @var bool|null $hasAttachments Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
    */
    private ?bool $hasAttachments = null;
    
    /**
     * @var Importance|null $importance The importance property
    */
    private ?Importance $importance = null;
    
    /**
     * @var InferenceClassificationType|null $inferenceClassification The inferenceClassification property
    */
    private ?InferenceClassificationType $inferenceClassification = null;
    
    /**
     * @var array<InternetMessageHeader>|null $internetMessageHeaders A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
    */
    private ?array $internetMessageHeaders = null;
    
    /**
     * @var string|null $internetMessageId The message ID in the format specified by RFC2822.
    */
    private ?string $internetMessageId = null;
    
    /**
     * @var bool|null $isDeliveryReceiptRequested Indicates whether a read receipt is requested for the message.
    */
    private ?bool $isDeliveryReceiptRequested = null;
    
    /**
     * @var bool|null $isDraft Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
    */
    private ?bool $isDraft = null;
    
    /**
     * @var bool|null $isRead Indicates whether the message has been read.
    */
    private ?bool $isRead = null;
    
    /**
     * @var bool|null $isReadReceiptRequested Indicates whether a read receipt is requested for the message.
    */
    private ?bool $isReadReceiptRequested = null;
    
    /**
     * @var array<MultiValueLegacyExtendedProperty>|null $multiValueExtendedProperties The collection of multi-value extended properties defined for the message. Nullable.
    */
    private ?array $multiValueExtendedProperties = null;
    
    /**
     * @var string|null $parentFolderId The unique identifier for the message's parent mailFolder.
    */
    private ?string $parentFolderId = null;
    
    /**
     * @var DateTime|null $receivedDateTime The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    */
    private ?DateTime $receivedDateTime = null;
    
    /**
     * @var array<Recipient>|null $replyTo The email addresses to use when replying.
    */
    private ?array $replyTo = null;
    
    /**
     * @var Recipient|null $sender The sender property
    */
    private ?Recipient $sender = null;
    
    /**
     * @var DateTime|null $sentDateTime The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    */
    private ?DateTime $sentDateTime = null;
    
    /**
     * @var array<SingleValueLegacyExtendedProperty>|null $singleValueExtendedProperties The collection of single-value extended properties defined for the message. Nullable.
    */
    private ?array $singleValueExtendedProperties = null;
    
    /**
     * @var string|null $subject The subject of the message.
    */
    private ?string $subject = null;
    
    /**
     * @var array<Recipient>|null $toRecipients The To: recipients for the message.
    */
    private ?array $toRecipients = null;
    
    /**
     * @var ItemBody|null $uniqueBody The uniqueBody property
    */
    private ?ItemBody $uniqueBody = null;
    
    /**
     * @var string|null $webLink The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, the browser shows the message in the Outlook on the web review pane.The message opens in the browser if you are signed in to your mailbox via Outlook on the web. You are prompted to sign in if you are not already signed in with the browser.This URL cannot be accessed from within an iFrame.
    */
    private ?string $webLink = null;
    
    /**
     * Instantiates a new Message and sets the default values.
    */
    public function __construct() {
        parent::__construct();
    }

    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param ParseNode $parseNode The parse node to use to read the discriminator value and create the object
     * @return Message
    */
    public static function createFromDiscriminatorValue(ParseNode $parseNode): Message {
        return new Message();
    }

    /**
     * Gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @return array<Attachment>|null
    */
    public function getAttachments(): ?array {
        return $this->attachments;
    }

    /**
     * Gets the bccRecipients property value. The Bcc: recipients for the message.
     * @return array<Recipient>|null
    */
    public function getBccRecipients(): ?array {
        return $this->bccRecipients;
    }

    /**
     * Gets the body property value. The body property
     * @return ItemBody|null
    */
    public function getBody(): ?ItemBody {
        return $this->body;
    }

    /**
     * Gets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
     * @return string|null
    */
    public function getBodyPreview(): ?string {
        return $this->bodyPreview;
    }

    /**
     * Gets the ccRecipients property value. The Cc: recipients for the message.
     * @return array<Recipient>|null
    */
    public function getCcRecipients(): ?array {
        return $this->ccRecipients;
    }

    /**
     * Gets the conversationId property value. The ID of the conversation the email belongs to.
     * @return string|null
    */
    public function getConversationId(): ?string {
        return $this->conversationId;
    }

    /**
     * Gets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @return StreamInterface|null
    */
    public function getConversationIndex(): ?StreamInterface {
        return $this->conversationIndex;
    }

    /**
     * Gets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @return array<Extension>|null
    */
    public function getExtensions(): ?array {
        return $this->extensions;
    }

    /**
     * The deserialization information for the current model
     * @return array<string, callable(ParseNode): void>
    */
    public function getFieldDeserializers(): array {
        $o = $this;
        return array_merge(parent::getFieldDeserializers(), [
            'attachments' => fn(ParseNode $n) => $o->setAttachments($n->getCollectionOfObjectValues([Attachment::class, 'createFromDiscriminatorValue'])),
            'bccRecipients' => fn(ParseNode $n) => $o->setBccRecipients($n->getCollectionOfObjectValues([Recipient::class, 'createFromDiscriminatorValue'])),
            'body' => fn(ParseNode $n) => $o->setBody($n->getObjectValue([ItemBody::class, 'createFromDiscriminatorValue'])),
            'bodyPreview' => fn(ParseNode $n) => $o->setBodyPreview($n->getStringValue()),
            'ccRecipients' => fn(ParseNode $n) => $o->setCcRecipients($n->getCollectionOfObjectValues([Recipient::class, 'createFromDiscriminatorValue'])),
            'conversationId' => fn(ParseNode $n) => $o->setConversationId($n->getStringValue()),
            'conversationIndex' => fn(ParseNode $n) => $o->setConversationIndex($n->getBinaryContent()),
            'extensions' => fn(ParseNode $n) => $o->setExtensions($n->getCollectionOfObjectValues([Extension::class, 'createFromDiscriminatorValue'])),
            'flag' => fn(ParseNode $n) => $o->setFlag($n->getObjectValue([FollowupFlag::class, 'createFromDiscriminatorValue'])),
            'from' => fn(ParseNode $n) => $o->setFrom($n->getObjectValue([Recipient::class, 'createFromDiscriminatorValue'])),
            'hasAttachments' => fn(ParseNode $n) => $o->setHasAttachments($n->getBooleanValue()),
            'importance' => fn(ParseNode $n) => $o->setImportance($n->getEnumValue(Importance::class)),
            'inferenceClassification' => fn(ParseNode $n) => $o->setInferenceClassification($n->getEnumValue(InferenceClassificationType::class)),
            'internetMessageHeaders' => fn(ParseNode $n) => $o->setInternetMessageHeaders($n->getCollectionOfObjectValues([InternetMessageHeader::class, 'createFromDiscriminatorValue'])),
            'internetMessageId' => fn(ParseNode $n) => $o->setInternetMessageId($n->getStringValue()),
            'isDeliveryReceiptRequested' => fn(ParseNode $n) => $o->setIsDeliveryReceiptRequested($n->getBooleanValue()),
            'isDraft' => fn(ParseNode $n) => $o->setIsDraft($n->getBooleanValue()),
            'isRead' => fn(ParseNode $n) => $o->setIsRead($n->getBooleanValue()),
            'isReadReceiptRequested' => fn(ParseNode $n) => $o->setIsReadReceiptRequested($n->getBooleanValue()),
            'multiValueExtendedProperties' => fn(ParseNode $n) => $o->setMultiValueExtendedProperties($n->getCollectionOfObjectValues([MultiValueLegacyExtendedProperty::class, 'createFromDiscriminatorValue'])),
            'parentFolderId' => fn(ParseNode $n) => $o->setParentFolderId($n->getStringValue()),
            'receivedDateTime' => fn(ParseNode $n) => $o->setReceivedDateTime($n->getDateTimeValue()),
            'replyTo' => fn(ParseNode $n) => $o->setReplyTo($n->getCollectionOfObjectValues([Recipient::class, 'createFromDiscriminatorValue'])),
            'sender' => fn(ParseNode $n) => $o->setSender($n->getObjectValue([Recipient::class, 'createFromDiscriminatorValue'])),
            'sentDateTime' => fn(ParseNode $n) => $o->setSentDateTime($n->getDateTimeValue()),
            'singleValueExtendedProperties' => fn(ParseNode $n) => $o->setSingleValueExtendedProperties($n->getCollectionOfObjectValues([SingleValueLegacyExtendedProperty::class, 'createFromDiscriminatorValue'])),
            'subject' => fn(ParseNode $n) => $o->setSubject($n->getStringValue()),
            'toRecipients' => fn(ParseNode $n) => $o->setToRecipients($n->getCollectionOfObjectValues([Recipient::class, 'createFromDiscriminatorValue'])),
            'uniqueBody' => fn(ParseNode $n) => $o->setUniqueBody($n->getObjectValue([ItemBody::class, 'createFromDiscriminatorValue'])),
            'webLink' => fn(ParseNode $n) => $o->setWebLink($n->getStringValue()),
        ]);
    }

    /**
     * Gets the flag property value. The flag property
     * @return FollowupFlag|null
    */
    public function getFlag(): ?FollowupFlag {
        return $this->flag;
    }

    /**
     * Gets the from property value. The from property
     * @return Recipient|null
    */
    public function getFrom(): ?Recipient {
        return $this->from;
    }

    /**
     * Gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @return bool|null
    */
    public function getHasAttachments(): ?bool {
        return $this->hasAttachments;
    }

    /**
     * Gets the importance property value. The importance property
     * @return Importance|null
    */
    public function getImportance(): ?Importance {
        return $this->importance;
    }

    /**
     * Gets the inferenceClassification property value. The inferenceClassification property
     * @return InferenceClassificationType|null
    */
    public function getInferenceClassification(): ?InferenceClassificationType {
        return $this->inferenceClassification;
    }

    /**
     * Gets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
     * @return array<InternetMessageHeader>|null
    */
    public function getInternetMessageHeaders(): ?array {
        return $this->internetMessageHeaders;
    }

    /**
     * Gets the internetMessageId property value. The message ID in the format specified by RFC2822.
     * @return string|null
    */
    public function getInternetMessageId(): ?string {
        return $this->internetMessageId;
    }

    /**
     * Gets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @return bool|null
    */
    public function getIsDeliveryReceiptRequested(): ?bool {
        return $this->isDeliveryReceiptRequested;
    }

    /**
     * Gets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
     * @return bool|null
    */
    public function getIsDraft(): ?bool {
        return $this->isDraft;
    }

    /**
     * Gets the isRead property value. Indicates whether the message has been read.
     * @return bool|null
    */
    public function getIsRead(): ?bool {
        return $this->isRead;
    }

    /**
     * Gets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @return bool|null
    */
    public function getIsReadReceiptRequested(): ?bool {
        return $this->isReadReceiptRequested;
    }

    /**
     * Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @return array<MultiValueLegacyExtendedProperty>|null
    */
    public function getMultiValueExtendedProperties(): ?array {
        return $this->multiValueExtendedProperties;
    }

    /**
     * Gets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
     * @return string|null
    */
    public function getParentFolderId(): ?string {
        return $this->parentFolderId;
    }

    /**
     * Gets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @return DateTime|null
    */
    public function getReceivedDateTime(): ?DateTime {
        return $this->receivedDateTime;
    }

    /**
     * Gets the replyTo property value. The email addresses to use when replying.
     * @return array<Recipient>|null
    */
    public function getReplyTo(): ?array {
        return $this->replyTo;
    }

    /**
     * Gets the sender property value. The sender property
     * @return Recipient|null
    */
    public function getSender(): ?Recipient {
        return $this->sender;
    }

    /**
     * Gets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @return DateTime|null
    */
    public function getSentDateTime(): ?DateTime {
        return $this->sentDateTime;
    }

    /**
     * Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @return array<SingleValueLegacyExtendedProperty>|null
    */
    public function getSingleValueExtendedProperties(): ?array {
        return $this->singleValueExtendedProperties;
    }

    /**
     * Gets the subject property value. The subject of the message.
     * @return string|null
    */
    public function getSubject(): ?string {
        return $this->subject;
    }

    /**
     * Gets the toRecipients property value. The To: recipients for the message.
     * @return array<Recipient>|null
    */
    public function getToRecipients(): ?array {
        return $this->toRecipients;
    }

    /**
     * Gets the uniqueBody property value. The uniqueBody property
     * @return ItemBody|null
    */
    public function getUniqueBody(): ?ItemBody {
        return $this->uniqueBody;
    }

    /**
     * Gets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, the browser shows the message in the Outlook on the web review pane.The message opens in the browser if you are signed in to your mailbox via Outlook on the web. You are prompted to sign in if you are not already signed in with the browser.This URL cannot be accessed from within an iFrame.
     * @return string|null
    */
    public function getWebLink(): ?string {
        return $this->webLink;
    }

    /**
     * Serializes information the current object
     * @param SerializationWriter $writer Serialization writer to use to serialize this model
    */
    public function serialize(SerializationWriter $writer): void {
        parent::serialize($writer);
        $writer->writeCollectionOfObjectValues('attachments', $this->getAttachments());
        $writer->writeCollectionOfObjectValues('bccRecipients', $this->getBccRecipients());
        $writer->writeObjectValue('body', $this->getBody());
        $writer->writeStringValue('bodyPreview', $this->getBodyPreview());
        $writer->writeCollectionOfObjectValues('ccRecipients', $this->getCcRecipients());
        $writer->writeStringValue('conversationId', $this->getConversationId());
        $writer->writeBinaryContent('conversationIndex', $this->getConversationIndex());
        $writer->writeCollectionOfObjectValues('extensions', $this->getExtensions());
        $writer->writeObjectValue('flag', $this->getFlag());
        $writer->writeObjectValue('from', $this->getFrom());
        $writer->writeBooleanValue('hasAttachments', $this->getHasAttachments());
        $writer->writeEnumValue('importance', $this->getImportance());
        $writer->writeEnumValue('inferenceClassification', $this->getInferenceClassification());
        $writer->writeCollectionOfObjectValues('internetMessageHeaders', $this->getInternetMessageHeaders());
        $writer->writeStringValue('internetMessageId', $this->getInternetMessageId());
        $writer->writeBooleanValue('isDeliveryReceiptRequested', $this->getIsDeliveryReceiptRequested());
        $writer->writeBooleanValue('isDraft', $this->getIsDraft());
        $writer->writeBooleanValue('isRead', $this->getIsRead());
        $writer->writeBooleanValue('isReadReceiptRequested', $this->getIsReadReceiptRequested());
        $writer->writeCollectionOfObjectValues('multiValueExtendedProperties', $this->getMultiValueExtendedProperties());
        $writer->writeStringValue('parentFolderId', $this->getParentFolderId());
        $writer->writeDateTimeValue('receivedDateTime', $this->getReceivedDateTime());
        $writer->writeCollectionOfObjectValues('replyTo', $this->getReplyTo());
        $writer->writeObjectValue('sender', $this->getSender());
        $writer->writeDateTimeValue('sentDateTime', $this->getSentDateTime());
        $writer->writeCollectionOfObjectValues('singleValueExtendedProperties', $this->getSingleValueExtendedProperties());
        $writer->writeStringValue('subject', $this->getSubject());
        $writer->writeCollectionOfObjectValues('toRecipients', $this->getToRecipients());
        $writer->writeObjectValue('uniqueBody', $this->getUniqueBody());
        $writer->writeStringValue('webLink', $this->getWebLink());
    }

    /**
     * Sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @param array<Attachment>|null $value Value to set for the attachments property.
    */
    public function setAttachments(?array $value): void {
        $this->attachments = $value;
    }

    /**
     * Sets the bccRecipients property value. The Bcc: recipients for the message.
     * @param array<Recipient>|null $value Value to set for the bccRecipients property.
    */
    public function setBccRecipients(?array $value): void {
        $this->bccRecipients = $value;
    }

    /**
     * Sets the body property value. The body property
     * @param ItemBody|null $value Value to set for the body property.
    */
    public function setBody(?ItemBody $value): void {
        $this->body = $value;
    }

    /**
     * Sets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
     * @param string|null $value Value to set for the bodyPreview property.
    */
    public function setBodyPreview(?string $value): void {
        $this->bodyPreview = $value;
    }

    /**
     * Sets the ccRecipients property value. The Cc: recipients for the message.
     * @param array<Recipient>|null $value Value to set for the ccRecipients property.
    */
    public function setCcRecipients(?array $value): void {
        $this->ccRecipients = $value;
    }

    /**
     * Sets the conversationId property value. The ID of the conversation the email belongs to.
     * @param string|null $value Value to set for the conversationId property.
    */
    public function setConversationId(?string $value): void {
        $this->conversationId = $value;
    }

    /**
     * Sets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @param StreamInterface|null $value Value to set for the conversationIndex property.
    */
    public function setConversationIndex(?StreamInterface $value): void {
        $this->conversationIndex = $value;
    }

    /**
     * Sets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @param array<Extension>|null $value Value to set for the extensions property.
    */
    public function setExtensions(?array $value): void {
        $this->extensions = $value;
    }

    /**
     * Sets the flag property value. The flag property
     * @param FollowupFlag|null $value Value to set for the flag property.
    */
    public function setFlag(?FollowupFlag $value): void {
        $this->flag = $value;
    }

    /**
     * Sets the from property value. The from property
     * @param Recipient|null $value Value to set for the from property.
    */
    public function setFrom(?Recipient $value): void {
        $this->from = $value;
    }

    /**
     * Sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @param bool|null $value Value to set for the hasAttachments property.
    */
    public function setHasAttachments(?bool $value): void {
        $this->hasAttachments = $value;
    }

    /**
     * Sets the importance property value. The importance property
     * @param Importance|null $value Value to set for the importance property.
    */
    public function setImportance(?Importance $value): void {
        $this->importance = $value;
    }

    /**
     * Sets the inferenceClassification property value. The inferenceClassification property
     * @param InferenceClassificationType|null $value Value to set for the inferenceClassification property.
    */
    public function setInferenceClassification(?InferenceClassificationType $value): void {
        $this->inferenceClassification = $value;
    }

    /**
     * Sets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
     * @param array<InternetMessageHeader>|null $value Value to set for the internetMessageHeaders property.
    */
    public function setInternetMessageHeaders(?array $value): void {
        $this->internetMessageHeaders = $value;
    }

    /**
     * Sets the internetMessageId property value. The message ID in the format specified by RFC2822.
     * @param string|null $value Value to set for the internetMessageId property.
    */
    public function setInternetMessageId(?string $value): void {
        $this->internetMessageId = $value;
    }

    /**
     * Sets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @param bool|null $value Value to set for the isDeliveryReceiptRequested property.
    */
    public function setIsDeliveryReceiptRequested(?bool $value): void {
        $this->isDeliveryReceiptRequested = $value;
    }

    /**
     * Sets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
     * @param bool|null $value Value to set for the isDraft property.
    */
    public function setIsDraft(?bool $value): void {
        $this->isDraft = $value;
    }

    /**
     * Sets the isRead property value. Indicates whether the message has been read.
     * @param bool|null $value Value to set for the isRead property.
    */
    public function setIsRead(?bool $value): void {
        $this->isRead = $value;
    }

    /**
     * Sets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @param bool|null $value Value to set for the isReadReceiptRequested property.
    */
    public function setIsReadReceiptRequested(?bool $value): void {
        $this->isReadReceiptRequested = $value;
    }

    /**
     * Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @param array<MultiValueLegacyExtendedProperty>|null $value Value to set for the multiValueExtendedProperties property.
    */
    public function setMultiValueExtendedProperties(?array $value): void {
        $this->multiValueExtendedProperties = $value;
    }

    /**
     * Sets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
     * @param string|null $value Value to set for the parentFolderId property.
    */
    public function setParentFolderId(?string $value): void {
        $this->parentFolderId = $value;
    }

    /**
     * Sets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @param DateTime|null $value Value to set for the receivedDateTime property.
    */
    public function setReceivedDateTime(?DateTime $value): void {
        $this->receivedDateTime = $value;
    }

    /**
     * Sets the replyTo property value. The email addresses to use when replying.
     * @param array<Recipient>|null $value Value to set for the replyTo property.
    */
    public function setReplyTo(?array $value): void {
        $this->replyTo = $value;
    }

    /**
     * Sets the sender property value. The sender property
     * @param Recipient|null $value Value to set for the sender property.
    */
    public function setSender(?Recipient $value): void {
        $this->sender = $value;
    }

    /**
     * Sets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @param DateTime|null $value Value to set for the sentDateTime property.
    */
    public function setSentDateTime(?DateTime $value): void {
        $this->sentDateTime = $value;
    }

    /**
     * Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @param array<SingleValueLegacyExtendedProperty>|null $value Value to set for the singleValueExtendedProperties property.
    */
    public function setSingleValueExtendedProperties(?array $value): void {
        $this->singleValueExtendedProperties = $value;
    }

    /**
     * Sets the subject property value. The subject of the message.
     * @param string|null $value Value to set for the subject property.
    */
    public function setSubject(?string $value): void {
        $this->subject = $value;
    }

    /**
     * Sets the toRecipients property value. The To: recipients for the message.
     * @param array<Recipient>|null $value Value to set for the toRecipients property.
    */
    public function setToRecipients(?array $value): void {
        $this->toRecipients = $value;
    }

    /**
     * Sets the uniqueBody property value. The uniqueBody property
     * @param ItemBody|null $value Value to set for the uniqueBody property.
    */
    public function setUniqueBody(?ItemBody $value): void {
        $this->uniqueBody = $value;
    }

    /**
     * Sets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, the browser shows the message in the Outlook on the web review pane.The message opens in the browser if you are signed in to your mailbox via Outlook on the web. You are prompted to sign in if you are not already signed in with the browser.This URL cannot be accessed from within an iFrame.
     * @param string|null $value Value to set for the webLink property.
    */
    public function setWebLink(?string $value): void {
        $this->webLink = $value;
    }

}
