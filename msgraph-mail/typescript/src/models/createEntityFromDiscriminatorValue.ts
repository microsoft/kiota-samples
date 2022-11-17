<<<<<<< HEAD:msgraph-mail/typescript/src/models/createEntityFromDiscriminatorValue.ts
import {EntityImpl} from './index';
=======
import {Attachment, Entity, Extension, InferenceClassification, InferenceClassificationOverride, MailFolder, Message, MessageRule, MultiValueLegacyExtendedProperty, OutlookItem, SingleValueLegacyExtendedProperty} from './index';
>>>>>>> main:msgraph-mail/typescript/src/models/microsoft/graph/createEntityFromDiscriminatorValue.ts
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createEntityFromDiscriminatorValue(parseNode: ParseNode | undefined) : EntityImpl {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return new EntityImpl();
}
