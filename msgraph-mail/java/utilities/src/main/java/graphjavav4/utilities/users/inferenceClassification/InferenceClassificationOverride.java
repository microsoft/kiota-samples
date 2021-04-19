package graphjavav4.utilities.users.inferenceClassification;

import java.util.Objects;
import graphjavav4.utilities.users.InferenceClassificationType;
import graphjavav4.utilities.users.EmailAddress;
import graphjavav4.utilities.users.Entity;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
public class InferenceClassificationOverride extends Entity implements Parsable {
    @javax.annotation.Nullable
    public InferenceClassificationType classifyAs;
    @javax.annotation.Nullable
    public EmailAddress senderEmailAddress;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeObjectValue("classifyAs", classifyAs);
        writer.writeObjectValue("senderEmailAddress", senderEmailAddress);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("classifyAs", (o, n) -> { ((InferenceClassificationOverride)o).classifyAs = n.getObjectValue(InferenceClassificationType.class); });
        fields.put("senderEmailAddress", (o, n) -> { ((InferenceClassificationOverride)o).senderEmailAddress = n.getObjectValue(EmailAddress.class); });
        return fields;
    }
}
