package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
public class FollowupFlag implements Parsable {
    @javax.annotation.Nullable
    public DateTimeTimeZone completedDateTime;
    @javax.annotation.Nullable
    public DateTimeTimeZone dueDateTime;
    @javax.annotation.Nullable
    public FollowupFlagStatus flagStatus;
    @javax.annotation.Nullable
    public DateTimeTimeZone startDateTime;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeObjectValue("completedDateTime", completedDateTime);
        writer.writeObjectValue("dueDateTime", dueDateTime);
        writer.writeObjectValue("flagStatus", flagStatus);
        writer.writeObjectValue("startDateTime", startDateTime);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(4);
        fields.put("completedDateTime", (o, n) -> { ((FollowupFlag)o).completedDateTime = n.getObjectValue(DateTimeTimeZone.class); });
        fields.put("dueDateTime", (o, n) -> { ((FollowupFlag)o).dueDateTime = n.getObjectValue(DateTimeTimeZone.class); });
        fields.put("flagStatus", (o, n) -> { ((FollowupFlag)o).flagStatus = n.getObjectValue(FollowupFlagStatus.class); });
        fields.put("startDateTime", (o, n) -> { ((FollowupFlag)o).startDateTime = n.getObjectValue(DateTimeTimeZone.class); });
        return fields;
    }
}
