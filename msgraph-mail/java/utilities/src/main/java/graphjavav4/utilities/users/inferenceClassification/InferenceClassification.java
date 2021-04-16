package graphjavav4.utilities.users.inferenceClassification;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.Entity;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
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
