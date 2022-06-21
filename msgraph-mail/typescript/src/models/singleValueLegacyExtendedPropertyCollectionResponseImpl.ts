import {createSingleValueLegacyExtendedPropertyFromDiscriminatorValue} from './createSingleValueLegacyExtendedPropertyFromDiscriminatorValue';
import {SingleValueLegacyExtendedPropertyImpl} from './index';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedPropertyCollectionResponse} from './singleValueLegacyExtendedPropertyCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class SingleValueLegacyExtendedPropertyCollectionResponseImpl implements SingleValueLegacyExtendedPropertyCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: SingleValueLegacyExtendedProperty[] | undefined;
    /**
     * Instantiates a new SingleValueLegacyExtendedPropertyCollectionResponse and sets the default values.
     * @param singleValueLegacyExtendedPropertyCollectionResponseParameterValue 
     */
    public constructor(singleValueLegacyExtendedPropertyCollectionResponseParameterValue?: SingleValueLegacyExtendedPropertyCollectionResponse | undefined) {
        this.additionalData = singleValueLegacyExtendedPropertyCollectionResponseParameterValue?.additionalData ? singleValueLegacyExtendedPropertyCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = singleValueLegacyExtendedPropertyCollectionResponseParameterValue?.nextLink;
        this.value = singleValueLegacyExtendedPropertyCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<SingleValueLegacyExtendedPropertyImpl>(createSingleValueLegacyExtendedPropertyFromDiscriminatorValue); },
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
        if(this.value && this.value.length != 0){        const valueArrValue: SingleValueLegacyExtendedPropertyImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new SingleValueLegacyExtendedPropertyImpl(element));});
            writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedPropertyImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
