package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class MultiValueLegacyExtendedProperty extends Entity implements Parsable {
    @javax.annotation.Nullable
    public List<Object> value;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfPrimitiveValues("value", value);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("value", (o, n) -> { ((MultiValueLegacyExtendedProperty)o).value = n.getCollectionOfPrimitiveValues(Object.class); });
        return fields;
    }
}
