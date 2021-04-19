package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
public class DateTimeTimeZone implements Parsable {
    @javax.annotation.Nullable
    public String dateTime;
    @javax.annotation.Nullable
    public String timeZone;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("dateTime", dateTime);
        writer.writeStringValue("timeZone", timeZone);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("dateTime", (o, n) -> { ((DateTimeTimeZone)o).dateTime = n.getStringValue(); });
        fields.put("timeZone", (o, n) -> { ((DateTimeTimeZone)o).timeZone = n.getStringValue(); });
        return fields;
    }
}
