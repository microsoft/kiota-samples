package graphjavav4.utilities.users.inferenceClassification;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.EmailAddress;
import graphjavav4.utilities.users.Entity;
import graphjavav4.utilities.users.InferenceClassificationType;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class InferenceClassificationOverride extends Entity implements Parsable {
    @javax.annotation.Nullable
    public InferenceClassificationType classifyAs;
    @javax.annotation.Nullable
    public EmailAddress senderEmailAddress;
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("classifyAs", (o, n) -> { ((InferenceClassificationOverride)o).classifyAs = n.getEnumValue(InferenceClassificationType.class); });
            this.put("senderEmailAddress", (o, n) -> { ((InferenceClassificationOverride)o).senderEmailAddress = n.getObjectValue(EmailAddress.class); });
        }};
    }
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeEnumValue("classifyAs", classifyAs);
        writer.writeObjectValue("senderEmailAddress", senderEmailAddress);
    }
}
