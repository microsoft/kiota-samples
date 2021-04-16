package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class ItemBody implements Parsable {
    @javax.annotation.Nullable
    public String content;
    @javax.annotation.Nullable
    public BodyType contentType;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("content", content);
        writer.writeEnumValue("contentType", contentType);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("content", (o, n) -> { ((ItemBody)o).content = n.getStringValue(); });
        fields.put("contentType", (o, n) -> { ((ItemBody)o).contentType = n.getEnumValue(BodyType.class); });
        return fields;
    }
}
