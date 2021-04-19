package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
public class InternetMessageHeader implements Parsable {
    @javax.annotation.Nullable
    public String name;
    @javax.annotation.Nullable
    public String value;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("name", name);
        writer.writeStringValue("value", value);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("name", (o, n) -> { ((InternetMessageHeader)o).name = n.getStringValue(); });
        fields.put("value", (o, n) -> { ((InternetMessageHeader)o).value = n.getStringValue(); });
        return fields;
    }
}
