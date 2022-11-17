import {EntityImpl} from './index';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MultiValueLegacyExtendedPropertyImpl extends EntityImpl implements MultiValueLegacyExtendedProperty {
    /** A collection of property values. */
    private _value?: string[] | undefined;
    /**
     * Instantiates a new MultiValueLegacyExtendedProperty and sets the default values.
     * @param multiValueLegacyExtendedPropertyParameterValue 
     */
    public constructor(multiValueLegacyExtendedPropertyParameterValue?: MultiValueLegacyExtendedProperty | undefined) {
        super(multiValueLegacyExtendedPropertyParameterValue);
        this._value = multiValueLegacyExtendedPropertyParameterValue?.value;
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
    /**
     * Gets the value property value. A collection of property values.
     * @returns a string
     */
    public get value() {
        return this._value;
    };
    /**
     * Sets the value property value. A collection of property values.
     * @param value Value to set for the value property.
     */
    public set value(value: string[] | undefined) {
        if(value) {
            this._value = value;
        }
    };
}
