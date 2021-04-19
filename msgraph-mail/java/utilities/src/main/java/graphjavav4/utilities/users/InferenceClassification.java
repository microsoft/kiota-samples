package graphjavav4.utilities.users;

import java.util.Objects;
import graphjavav4.utilities.users.inferenceClassification.InferenceClassificationOverride;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class InferenceClassification extends Entity implements Parsable {
    @javax.annotation.Nullable
    public List<InferenceClassificationOverride> overrides;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfObjectValues("overrides", overrides);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("overrides", (o, n) -> { ((InferenceClassification)o).overrides = n.getCollectionOfObjectValues(InferenceClassificationOverride.class); });
        return fields;
    }
}
