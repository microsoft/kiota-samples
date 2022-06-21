import {createInferenceClassificationOverrideFromDiscriminatorValue} from './createInferenceClassificationOverrideFromDiscriminatorValue';
import {InferenceClassificationOverrideImpl} from './index';
import {InferenceClassificationOverride} from './inferenceClassificationOverride';
import {InferenceClassificationOverrideCollectionResponse} from './inferenceClassificationOverrideCollectionResponse';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class InferenceClassificationOverrideCollectionResponseImpl implements InferenceClassificationOverrideCollectionResponse {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** The nextLink property */
    public nextLink?: string | undefined;
    /** The value property */
    public value?: InferenceClassificationOverride[] | undefined;
    /**
     * Instantiates a new InferenceClassificationOverrideCollectionResponse and sets the default values.
     * @param inferenceClassificationOverrideCollectionResponseParameterValue 
     */
    public constructor(inferenceClassificationOverrideCollectionResponseParameterValue?: InferenceClassificationOverrideCollectionResponse | undefined) {
        this.additionalData = inferenceClassificationOverrideCollectionResponseParameterValue?.additionalData ? inferenceClassificationOverrideCollectionResponseParameterValue?.additionalData! : {};
        this.nextLink = inferenceClassificationOverrideCollectionResponseParameterValue?.nextLink;
        this.value = inferenceClassificationOverrideCollectionResponseParameterValue?.value;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "@odata.nextLink": n => { this.nextLink = n.getStringValue(); },
            "value": n => { this.value = n.getCollectionOfObjectValues<InferenceClassificationOverrideImpl>(createInferenceClassificationOverrideFromDiscriminatorValue); },
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
        if(this.value && this.value.length != 0){        const valueArrValue: InferenceClassificationOverrideImpl[] = []; this.value?.forEach(element => {valueArrValue.push(new InferenceClassificationOverrideImpl(element));});
            writer.writeCollectionOfObjectValues<InferenceClassificationOverrideImpl>("value", valueArrValue);
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
