import type {OutlookItem} from './outlookItem';
import {serializeEntity} from './serializeEntity';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeOutlookItem(writer: SerializationWriter, outlookItem: OutlookItem | undefined = {} as OutlookItem) : void {
        serializeEntity(writer, outlookItem)
        writer.writeCollectionOfPrimitiveValues<string>("categories", outlookItem.categories);
        writer.writeStringValue("changeKey", outlookItem.changeKey);
        writer.writeDateValue("createdDateTime", outlookItem.createdDateTime);
        writer.writeDateValue("lastModifiedDateTime", outlookItem.lastModifiedDateTime);
}
