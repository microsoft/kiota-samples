import {deserializeIntoAttachment} from './deserializeIntoAttachment';
import {deserializeIntoEntity} from './deserializeIntoEntity';
import {deserializeIntoExtension} from './deserializeIntoExtension';
import {deserializeIntoInferenceClassification} from './deserializeIntoInferenceClassification';
import {deserializeIntoInferenceClassificationOverride} from './deserializeIntoInferenceClassificationOverride';
import {deserializeIntoMailFolder} from './deserializeIntoMailFolder';
import {deserializeIntoMessage} from './deserializeIntoMessage';
import {deserializeIntoMessageRule} from './deserializeIntoMessageRule';
import {deserializeIntoMultiValueLegacyExtendedProperty} from './deserializeIntoMultiValueLegacyExtendedProperty';
import {deserializeIntoOutlookItem} from './deserializeIntoOutlookItem';
import {deserializeIntoSingleValueLegacyExtendedProperty} from './deserializeIntoSingleValueLegacyExtendedProperty';
import {Entity} from './index';
import {ParseNode} from '@microsoft/kiota-abstractions';

export function createEntityFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoEntity;
}
