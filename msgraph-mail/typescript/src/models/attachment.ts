import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export interface Attachment extends Entity, Parsable {
    /**
     * The MIME type.
     */
    contentType?: string | undefined;
    /**
     * true if the attachment is an inline attachment; otherwise, false.
     */
    isInline?: boolean | undefined;
    /**
     * The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     */
    lastModifiedDateTime?: Date | undefined;
    /**
     * The attachment's file name.
     */
    name?: string | undefined;
    /**
     * The length of the attachment in bytes.
     */
    size?: number | undefined;
}
export function createAttachmentFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoAttachment;
}
export function deserializeIntoAttachment(attachment: Attachment | undefined = {} as Attachment) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(attachment),
        "contentType": n => { attachment.contentType = n.getStringValue(); },
        "isInline": n => { attachment.isInline = n.getBooleanValue(); },
        "lastModifiedDateTime": n => { attachment.lastModifiedDateTime = n.getDateValue(); },
        "name": n => { attachment.name = n.getStringValue(); },
        "size": n => { attachment.size = n.getNumberValue(); },
    }
}
export function serializeAttachment(writer: SerializationWriter, attachment: Attachment | undefined = {} as Attachment) : void {
        serializeEntity(writer, attachment)
        writer.writeStringValue("contentType", attachment.contentType);
        writer.writeBooleanValue("isInline", attachment.isInline);
        writer.writeDateValue("lastModifiedDateTime", attachment.lastModifiedDateTime);
        writer.writeStringValue("name", attachment.name);
        writer.writeNumberValue("size", attachment.size);
}
