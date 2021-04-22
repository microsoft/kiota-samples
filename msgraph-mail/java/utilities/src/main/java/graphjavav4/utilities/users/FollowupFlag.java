package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class FollowupFlag implements Parsable {
    @javax.annotation.Nullable
    public DateTimeTimeZone completedDateTime;
    @javax.annotation.Nullable
    public DateTimeTimeZone dueDateTime;
    @javax.annotation.Nullable
    public FollowupFlagStatus flagStatus;
    @javax.annotation.Nullable
    public DateTimeTimeZone startDateTime;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeObjectValue("completedDateTime", completedDateTime);
        writer.writeObjectValue("dueDateTime", dueDateTime);
        writer.writeEnumValue("flagStatus", flagStatus);
        writer.writeObjectValue("startDateTime", startDateTime);
        writer.writeAdditionalData(this.additionalData);
    }
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    @javax.annotation.Nonnull
    private final Map<String, Object> additionalData = new HashMap<>();
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return additionalData;
    }
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(4);
        fields.put("completedDateTime", (o, n) -> { ((FollowupFlag)o).completedDateTime = n.getObjectValue(DateTimeTimeZone.class); });
        fields.put("dueDateTime", (o, n) -> { ((FollowupFlag)o).dueDateTime = n.getObjectValue(DateTimeTimeZone.class); });
        fields.put("flagStatus", (o, n) -> { ((FollowupFlag)o).flagStatus = n.getEnumValue(FollowupFlagStatus.class); });
        fields.put("startDateTime", (o, n) -> { ((FollowupFlag)o).startDateTime = n.getObjectValue(DateTimeTimeZone.class); });
        return fields;
    }
}
