import {EntityImpl} from './index';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class SingleValueLegacyExtendedPropertyImpl extends EntityImpl implements SingleValueLegacyExtendedProperty {
    /** A property value. */
    public value?: string | undefined;
    /**
     * Instantiates a new singleValueLegacyExtendedProperty and sets the default values.
     * @param singleValueLegacyExtendedPropertyParameterValue 
     */
    public constructor(singleValueLegacyExtendedPropertyParameterValue?: SingleValueLegacyExtendedProperty | undefined) {
        super(singleValueLegacyExtendedPropertyParameterValue);
        this.value = singleValueLegacyExtendedPropertyParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "value": n => { this.value = n.getStringValue(); },
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
            writer.writeStringValue("value", this.value);
        }
    };
}
