import {Importance} from '../importance';
import {Recipient} from '../recipient';
import {MessageActionFlag} from './messageActionFlag';
import {Sensitivity} from './sensitivity';
import {SizeRange} from './sizeRange';

export class MessageRulePredicates {
    public bodyContains?: object[] | undefined;
    public bodyOrSubjectContains?: object[] | undefined;
    public categories?: object[] | undefined;
    public fromAddresses?: Recipient[] | undefined;
    public hasAttachments?: boolean | undefined;
    public headerContains?: object[] | undefined;
    public importance?: Importance | undefined;
    public isApprovalRequest?: boolean | undefined;
    public isAutomaticForward?: boolean | undefined;
    public isAutomaticReply?: boolean | undefined;
    public isEncrypted?: boolean | undefined;
    public isMeetingRequest?: boolean | undefined;
    public isMeetingResponse?: boolean | undefined;
    public isNonDeliveryReport?: boolean | undefined;
    public isPermissionControlled?: boolean | undefined;
    public isReadReceipt?: boolean | undefined;
    public isSigned?: boolean | undefined;
    public isVoicemail?: boolean | undefined;
    public messageActionFlag?: MessageActionFlag | undefined;
    public notSentToMe?: boolean | undefined;
    public recipientContains?: object[] | undefined;
    public senderContains?: object[] | undefined;
    public sensitivity?: Sensitivity | undefined;
    public sentCcMe?: boolean | undefined;
    public sentOnlyToMe?: boolean | undefined;
    public sentToAddresses?: Recipient[] | undefined;
    public sentToMe?: boolean | undefined;
    public sentToOrCcMe?: boolean | undefined;
    public subjectContains?: object[] | undefined;
    public withinSizeRange?: SizeRange | undefined;
}
