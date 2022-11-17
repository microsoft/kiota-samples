<<<<<<< HEAD
import {Entity} from './entity';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface Attachment extends Entity, Partial<Parsable> {
    /** The MIME type. */
    contentType?: string | undefined;
    /** true if the attachment is an inline attachment; otherwise, false. */
    isInline?: boolean | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z */
    lastModifiedDateTime?: Date | undefined;
    /** The display name of the attachment. This does not need to be the actual file name. */
    name?: string | undefined;
    /** The length of the attachment in bytes. */
    size?: number | undefined;
=======
import {Attachment} from './attachment';
import {EntityImpl} from './index';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class AttachmentImpl extends EntityImpl implements Attachment {
    /** The MIME type. */
    private _contentType?: string | undefined;
    /** true if the attachment is an inline attachment; otherwise, false. */
    private _isInline?: boolean | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z */
    private _lastModifiedDateTime?: Date | undefined;
    /** The display name of the attachment. This does not need to be the actual file name. */
    private _name?: string | undefined;
    /** The length of the attachment in bytes. */
    private _size?: number | undefined;
    /**
<<<<<<< HEAD:msgraph-mail/typescript/src/models/attachmentImpl.ts
     * Instantiates a new attachment and sets the default values.
     * @param attachmentParameterValue 
=======
     * Instantiates a new Attachment and sets the default values.
>>>>>>> main:msgraph-mail/typescript/src/models/attachment.ts
     */
    public constructor(attachmentParameterValue?: Attachment | undefined) {
        super(attachmentParameterValue);
        this._contentType = attachmentParameterValue?.contentType;
        this._isInline = attachmentParameterValue?.isInline;
        this._lastModifiedDateTime = attachmentParameterValue?.lastModifiedDateTime;
        this._name = attachmentParameterValue?.name;
        this._size = attachmentParameterValue?.size;
    };
    /**
     * Gets the contentType property value. The MIME type.
     * @returns a string
     */
    public get contentType() {
        return this._contentType;
    };
    /**
     * Sets the contentType property value. The MIME type.
     * @param value Value to set for the contentType property.
     */
    public set contentType(value: string | undefined) {
        if(value) {
            this._contentType = value;
        }
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "contentType": n => { this.contentType = n.getStringValue(); },
            "isInline": n => { this.isInline = n.getBooleanValue(); },
            "lastModifiedDateTime": n => { this.lastModifiedDateTime = n.getDateValue(); },
            "name": n => { this.name = n.getStringValue(); },
            "size": n => { this.size = n.getNumberValue(); },
        };
    };
    /**
     * Gets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
     * @returns a boolean
     */
    public get isInline() {
        return this._isInline;
    };
    /**
     * Sets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
     * @param value Value to set for the isInline property.
     */
    public set isInline(value: boolean | undefined) {
        if(value) {
            this._isInline = value;
        }
    };
    /**
     * Gets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @returns a Date
     */
    public get lastModifiedDateTime() {
        return this._lastModifiedDateTime;
    };
    /**
     * Sets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     * @param value Value to set for the lastModifiedDateTime property.
     */
    public set lastModifiedDateTime(value: Date | undefined) {
        if(value) {
            this._lastModifiedDateTime = value;
        }
    };
    /**
     * Gets the name property value. The display name of the attachment. This does not need to be the actual file name.
     * @returns a string
     */
    public get name() {
        return this._name;
    };
    /**
     * Sets the name property value. The display name of the attachment. This does not need to be the actual file name.
     * @param value Value to set for the name property.
     */
    public set name(value: string | undefined) {
        if(value) {
            this._name = value;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.contentType){
            writer.writeStringValue("contentType", this.contentType);
        }
        if(this.isInline){
            writer.writeBooleanValue("isInline", this.isInline);
        }
        if(this.lastModifiedDateTime){
            writer.writeDateValue("lastModifiedDateTime", this.lastModifiedDateTime);
        }
        if(this.name){
            writer.writeStringValue("name", this.name);
        }
        if(this.size){
            writer.writeNumberValue("size", this.size);
        }
    };
    /**
     * Gets the size property value. The length of the attachment in bytes.
     * @returns a integer
     */
    public get size() {
        return this._size;
    };
    /**
     * Sets the size property value. The length of the attachment in bytes.
     * @param value Value to set for the size property.
     */
    public set size(value: number | undefined) {
        if(value) {
            this._size = value;
        }
    };
>>>>>>> main
}
