import {createInferenceClassificationOverrideFromDiscriminatorValue} from './createInferenceClassificationOverrideFromDiscriminatorValue';
import {EntityImpl, InferenceClassificationOverrideImpl} from './index';
import {InferenceClassification} from './inferenceClassification';
import {InferenceClassificationOverride} from './inferenceClassificationOverride';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class InferenceClassificationImpl extends EntityImpl implements InferenceClassification {
    /** A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable. */
    public overrides?: InferenceClassificationOverride[] | undefined;
    /**
     * Instantiates a new inferenceClassification and sets the default values.
     * @param inferenceClassificationParameterValue 
     */
    public constructor(inferenceClassificationParameterValue?: InferenceClassification | undefined) {
        super(inferenceClassificationParameterValue);
        this.overrides = inferenceClassificationParameterValue?.overrides;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "overrides": n => { this.overrides = n.getCollectionOfObjectValues<InferenceClassificationOverrideImpl>(createInferenceClassificationOverrideFromDiscriminatorValue); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.overrides && this.overrides.length != 0){        const overridesArrValue: InferenceClassificationOverrideImpl[] = []; this.overrides?.forEach(element => {overridesArrValue.push(new InferenceClassificationOverrideImpl(element));});
            writer.writeCollectionOfObjectValues<InferenceClassificationOverrideImpl>("overrides", overridesArrValue);
        }
    };
}
