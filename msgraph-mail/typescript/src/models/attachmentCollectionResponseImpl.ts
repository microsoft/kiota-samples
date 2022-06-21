import {Attachment} from './attachment';
import {AttachmentCollectionResponse} from './attachmentCollectionResponse';
import {createAttachmentFromDiscriminatorValue} from './createAttachmentFromDiscriminatorValue';
import {AttachmentImpl} from './index';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class AttachmentCollectionResponseImpl implements AttachmentCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: Attachment[] | undefined;
    /**
     * Instantiates a new AttachmentCollectionResponse and sets the default values.
     * @param attachmentCollectionResponseParameterValue 
     */
    public constructor(attachmentCollectionResponseParameterValue?: AttachmentCollectionResponse | undefined) {
        this.additionalData = attachmentCollectionResponseParameterValue?.additionalData ? attachmentCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = attachmentCollectionResponseParameterValue?.nextLink;
        this.value = attachmentCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<AttachmentImpl>(createAttachmentFromDiscriminatorValue); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.nextLink){
            writer.writeStringValue("@odata.nextLink", this.nextLink);
        }
        if(this.value && this.value.length != 0){        const valueArrValue: AttachmentImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new AttachmentImpl(element));});
            writer.writeCollectionOfObjectValues<AttachmentImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
