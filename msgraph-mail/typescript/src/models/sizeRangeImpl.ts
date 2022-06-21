import {SizeRange} from './sizeRange';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class SizeRangeImpl implements SizeRange {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. */
    public maximumSize?: number | undefined;
    /** The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply. */
    public minimumSize?: number | undefined;
    /**
     * Instantiates a new sizeRange and sets the default values.
     * @param sizeRangeParameterValue 
     */
    public constructor(sizeRangeParameterValue?: SizeRange | undefined) {
        this.additionalData = sizeRangeParameterValue?.additionalData ? sizeRangeParameterValue?.additionalData! : {};
        this.maximumSize = sizeRangeParameterValue?.maximumSize;
        this.minimumSize = sizeRangeParameterValue?.minimumSize;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "maximumSize": n => { this.maximumSize = n.getNumberValue(); },
            "minimumSize": n => { this.minimumSize = n.getNumberValue(); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.maximumSize){
            writer.writeNumberValue("maximumSize", this.maximumSize);
        }
        if(this.minimumSize){
            writer.writeNumberValue("minimumSize", this.minimumSize);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
