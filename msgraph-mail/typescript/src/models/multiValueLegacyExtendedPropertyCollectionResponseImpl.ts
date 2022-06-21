import {createMultiValueLegacyExtendedPropertyFromDiscriminatorValue} from './createMultiValueLegacyExtendedPropertyFromDiscriminatorValue';
import {MultiValueLegacyExtendedPropertyImpl} from './index';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {MultiValueLegacyExtendedPropertyCollectionResponse} from './multiValueLegacyExtendedPropertyCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MultiValueLegacyExtendedPropertyCollectionResponseImpl implements MultiValueLegacyExtendedPropertyCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: MultiValueLegacyExtendedProperty[] | undefined;
    /**
     * Instantiates a new MultiValueLegacyExtendedPropertyCollectionResponse and sets the default values.
     * @param multiValueLegacyExtendedPropertyCollectionResponseParameterValue 
     */
    public constructor(multiValueLegacyExtendedPropertyCollectionResponseParameterValue?: MultiValueLegacyExtendedPropertyCollectionResponse | undefined) {
        this.additionalData = multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.additionalData ? multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.nextLink;
        this.value = multiValueLegacyExtendedPropertyCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>(createMultiValueLegacyExtendedPropertyFromDiscriminatorValue); },
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
        if(this.value && this.value.length != 0){        const valueArrValue: MultiValueLegacyExtendedPropertyImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new MultiValueLegacyExtendedPropertyImpl(element));});
            writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
