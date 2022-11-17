import {AttachmentImpl, EntityImpl, ExtensionImpl, InferenceClassificationImpl, InferenceClassificationOverrideImpl, MailFolderImpl, MessageImpl, MessageRuleImpl, MultiValueLegacyExtendedPropertyImpl, OutlookItemImpl, SingleValueLegacyExtendedPropertyImpl} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createEntityFromDiscriminatorValue(parseNode: ParseNode | undefined) : EntityImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new EntityImpl();
}
