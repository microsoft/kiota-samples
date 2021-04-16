package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class DateTimeTimeZone implements Parsable {
    /** A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.  */
    @javax.annotation.Nullable
    public String dateTime;
    /** Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.  */
    @javax.annotation.Nullable
    public String timeZone;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("dateTime", dateTime);
        writer.writeStringValue("timeZone", timeZone);
    }
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("dateTime", (o, n) -> { ((DateTimeTimeZone)o).dateTime = n.getStringValue(); });
        fields.put("timeZone", (o, n) -> { ((DateTimeTimeZone)o).timeZone = n.getStringValue(); });
        return fields;
    }
}
