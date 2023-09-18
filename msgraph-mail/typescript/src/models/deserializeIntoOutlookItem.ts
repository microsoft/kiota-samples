// tslint:disable
// eslint-disable
// Generated by Microsoft Kiota
import { deserializeIntoEntity } from './deserializeIntoEntity';
import { type OutlookItem } from './outlookItem';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function deserializeIntoOutlookItem(outlookItem: OutlookItem | undefined = {} as OutlookItem) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(outlookItem),
        "categories": n => { outlookItem.categories = n.getCollectionOfPrimitiveValues<string>(); },
        "changeKey": n => { outlookItem.changeKey = n.getStringValue(); },
        "createdDateTime": n => { outlookItem.createdDateTime = n.getDateValue(); },
        "lastModifiedDateTime": n => { outlookItem.lastModifiedDateTime = n.getDateValue(); },
    }
}
// tslint:enable
// eslint-enable
