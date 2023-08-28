import {Importance} from './importance';
import {MessageActionFlag} from './messageActionFlag';
import type {Recipient} from './recipient';
import {Sensitivity} from './sensitivity';
import type {SizeRange} from './sizeRange';
import type {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

export interface MessageRulePredicates extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
     */
    bodyContains?: string[] | undefined;
    /**
     * Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
     */
    bodyOrSubjectContains?: string[] | undefined;
    /**
     * Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
     */
    categories?: string[] | undefined;
    /**
     * Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
     */
    fromAddresses?: Recipient[] | undefined;
    /**
     * Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
     */
    hasAttachments?: boolean | undefined;
    /**
     * Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
     */
    headerContains?: string[] | undefined;
    /**
     * The importance property
     */
    importance?: Importance | undefined;
    /**
     * Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
     */
    isApprovalRequest?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
     */
    isAutomaticForward?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
     */
    isAutomaticReply?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
     */
    isEncrypted?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
     */
    isMeetingRequest?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
     */
    isMeetingResponse?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
     */
    isNonDeliveryReport?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
     */
    isPermissionControlled?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
     */
    isReadReceipt?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
     */
    isSigned?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
     */
    isVoicemail?: boolean | undefined;
    /**
     * The messageActionFlag property
     */
    messageActionFlag?: MessageActionFlag | undefined;
    /**
     * Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
     */
    notSentToMe?: boolean | undefined;
    /**
     * Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
     */
    recipientContains?: string[] | undefined;
    /**
     * Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
     */
    senderContains?: string[] | undefined;
    /**
     * The sensitivity property
     */
    sensitivity?: Sensitivity | undefined;
    /**
     * Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
     */
    sentCcMe?: boolean | undefined;
    /**
     * Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
     */
    sentOnlyToMe?: boolean | undefined;
    /**
     * Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
     */
    sentToAddresses?: Recipient[] | undefined;
    /**
     * Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
     */
    sentToMe?: boolean | undefined;
    /**
     * Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
     */
    sentToOrCcMe?: boolean | undefined;
    /**
     * Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
     */
    subjectContains?: string[] | undefined;
    /**
     * The withinSizeRange property
     */
    withinSizeRange?: SizeRange | undefined;
}
