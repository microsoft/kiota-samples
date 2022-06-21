import {BodyType} from './bodyType';
import {ItemBody} from './itemBody';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class ItemBodyImpl implements ItemBody {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The content of the item. */
    public content?: string | undefined;
    /** The contentType property */
    public contentType?: BodyType | undefined;
    /**
     * Instantiates a new itemBody and sets the default values.
     * @param itemBodyParameterValue 
     */
    public constructor(itemBodyParameterValue?: ItemBody | undefined) {
        this.additionalData = itemBodyParameterValue?.additionalData ? itemBodyParameterValue?.additionalData! : {};
        this.content = itemBodyParameterValue?.content;
        this.contentType = itemBodyParameterValue?.contentType;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "content": n => { this.content = n.getStringValue(); },
            "contentType": n => { this.contentType = n.getEnumValue<BodyType>(BodyType); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.content){
            writer.writeStringValue("content", this.content);
        }
        if(this.contentType){
            writer.writeEnumValue<BodyType>("contentType", this.contentType);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
