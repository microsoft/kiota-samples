package graphjavav4.utilities.models.microsoft.graph;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class InferenceClassificationOverride extends Entity implements Parsable {
    private InferenceClassificationType _classifyAs;
    private EmailAddress _senderEmailAddress;
    /**
     * Instantiates a new inferenceClassificationOverride and sets the default values.
     * @return a void
     */
    public InferenceClassificationOverride() {
        super();
    }
    /**
     * Gets the classifyAs property value. 
     * @return a inferenceClassificationType
     */
    @javax.annotation.Nullable
    public InferenceClassificationType getClassifyAs() {
        return this._classifyAs;
    }
    /**
     * Gets the senderEmailAddress property value. 
     * @return a emailAddress
     */
    @javax.annotation.Nullable
    public EmailAddress getSenderEmailAddress() {
        return this._senderEmailAddress;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("classifyAs", (o, n) -> { ((InferenceClassificationOverride)o).setClassifyAs(n.getEnumValue(InferenceClassificationType.class)); });
            this.put("senderEmailAddress", (o, n) -> { ((InferenceClassificationOverride)o).setSenderEmailAddress(n.getObjectValue(EmailAddress.class)); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeEnumValue("classifyAs", this.getClassifyAs());
        writer.writeObjectValue("senderEmailAddress", this.getSenderEmailAddress());
    }
    /**
     * Sets the classifyAs property value. 
     * @param value Value to set for the classifyAs property.
     * @return a void
     */
    public void setClassifyAs(@javax.annotation.Nullable final InferenceClassificationType value) {
        this._classifyAs = value;
    }
    /**
     * Sets the senderEmailAddress property value. 
     * @param value Value to set for the senderEmailAddress property.
     * @return a void
     */
    public void setSenderEmailAddress(@javax.annotation.Nullable final EmailAddress value) {
        this._senderEmailAddress = value;
    }
}
