package graphjavav4.utilities.users.mailFolders;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
public class SizeRange implements Parsable {
    @javax.annotation.Nullable
    public Integer maximumSize;
    @javax.annotation.Nullable
    public Integer minimumSize;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeIntegerValue("maximumSize", maximumSize);
        writer.writeIntegerValue("minimumSize", minimumSize);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("maximumSize", (o, n) -> { ((SizeRange)o).maximumSize = n.getIntegerValue(); });
        fields.put("minimumSize", (o, n) -> { ((SizeRange)o).minimumSize = n.getIntegerValue(); });
        return fields;
    }
}
