import {Attachment} from './attachment';
import {serializeEntity} from './serializeEntity';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeAttachment(writer: SerializationWriter, attachment: Attachment | undefined = {} as Attachment) : void {
        serializeEntity(writer, attachment)
        writer.writeStringValue("contentType", attachment.contentType);
        writer.writeBooleanValue("isInline", attachment.isInline);
        writer.writeDateValue("lastModifiedDateTime", attachment.lastModifiedDateTime);
        writer.writeStringValue("name", attachment.name);
        writer.writeNumberValue("size", attachment.size);
}
