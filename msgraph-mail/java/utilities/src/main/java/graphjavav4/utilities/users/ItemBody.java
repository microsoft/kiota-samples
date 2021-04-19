package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
public class ItemBody implements Parsable {
    @javax.annotation.Nullable
    public String content;
    @javax.annotation.Nullable
    public BodyType contentType;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("content", content);
        writer.writeObjectValue("contentType", contentType);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("content", (o, n) -> { ((ItemBody)o).content = n.getStringValue(); });
        fields.put("contentType", (o, n) -> { ((ItemBody)o).contentType = n.getObjectValue(BodyType.class); });
        return fields;
    }
}
