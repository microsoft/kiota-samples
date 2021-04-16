package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class SizeRange implements Parsable {
    /** The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    @javax.annotation.Nullable
    public Integer maximumSize;
    /** The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    @javax.annotation.Nullable
    public Integer minimumSize;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeIntegerValue("maximumSize", maximumSize);
        writer.writeIntegerValue("minimumSize", minimumSize);
    }
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("maximumSize", (o, n) -> { ((SizeRange)o).maximumSize = n.getIntegerValue(); });
        fields.put("minimumSize", (o, n) -> { ((SizeRange)o).minimumSize = n.getIntegerValue(); });
        return fields;
    }
}
