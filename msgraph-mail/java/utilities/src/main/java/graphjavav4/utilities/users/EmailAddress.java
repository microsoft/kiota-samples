package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class EmailAddress implements Parsable {
    @javax.annotation.Nullable
    public String address;
    @javax.annotation.Nullable
    public String name;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("address", address);
        writer.writeStringValue("name", name);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("address", (o, n) -> { ((EmailAddress)o).address = n.getStringValue(); });
        fields.put("name", (o, n) -> { ((EmailAddress)o).name = n.getStringValue(); });
        return fields;
    }
}
