package graphjavav4.utilities.users.inferenceClassification.overrides;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.inferenceClassification.InferenceClassificationOverride;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class OverridesResponse implements Parsable {
    @javax.annotation.Nullable
    public List<InferenceClassificationOverride> value;
    @javax.annotation.Nullable
    public String nextLink;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfObjectValues("value", value);
        writer.writeStringValue("nextLink", nextLink);
    }
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("value", (o, n) -> { ((OverridesResponse)o).value = n.getCollectionOfObjectValues(InferenceClassificationOverride.class); });
        fields.put("nextLink", (o, n) -> { ((OverridesResponse)o).nextLink = n.getStringValue(); });
        return fields;
    }
}
