import {EntityImpl} from './index';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MultiValueLegacyExtendedPropertyImpl extends EntityImpl implements MultiValueLegacyExtendedProperty {
    /** A collection of property values. */
    public value?: string[] | undefined;
    /**
     * Instantiates a new multiValueLegacyExtendedProperty and sets the default values.
     * @param multiValueLegacyExtendedPropertyParameterValue 
     */
    public constructor(multiValueLegacyExtendedPropertyParameterValue?: MultiValueLegacyExtendedProperty | undefined) {
        super(multiValueLegacyExtendedPropertyParameterValue);
        this.value = multiValueLegacyExtendedPropertyParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "value": n => { this.value = n.getCollectionOfPrimitiveValues<string>(); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.value){
            writer.writeCollectionOfPrimitiveValues<string>("value", this.value);
        }
    };
}
