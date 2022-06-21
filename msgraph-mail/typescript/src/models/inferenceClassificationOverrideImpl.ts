import {createEmailAddressFromDiscriminatorValue} from './createEmailAddressFromDiscriminatorValue';
import {EmailAddress} from './emailAddress';
import {EmailAddressImpl, EntityImpl} from './index';
import {InferenceClassificationOverride} from './inferenceClassificationOverride';
import {InferenceClassificationType} from './inferenceClassificationType';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class InferenceClassificationOverrideImpl extends EntityImpl implements InferenceClassificationOverride {
    /** The classifyAs property */
    public classifyAs?: InferenceClassificationType | undefined;
    /** The senderEmailAddress property */
    public senderEmailAddress?: EmailAddress | undefined;
    /**
     * Instantiates a new inferenceClassificationOverride and sets the default values.
     * @param inferenceClassificationOverrideParameterValue 
     */
    public constructor(inferenceClassificationOverrideParameterValue?: InferenceClassificationOverride | undefined) {
        super(inferenceClassificationOverrideParameterValue);
        this.classifyAs = inferenceClassificationOverrideParameterValue?.classifyAs;
        this.senderEmailAddress = inferenceClassificationOverrideParameterValue?.senderEmailAddress;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "classifyAs": n => { this.classifyAs = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); },
            "senderEmailAddress": n => { this.senderEmailAddress = n.getObjectValue<EmailAddressImpl>(createEmailAddressFromDiscriminatorValue); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.classifyAs){
            writer.writeEnumValue<InferenceClassificationType>("classifyAs", this.classifyAs);
        }
        if(this.senderEmailAddress){
            writer.writeObjectValue<EmailAddressImpl>("senderEmailAddress", new EmailAddressImpl(this.senderEmailAddress));
        }
    };
}
