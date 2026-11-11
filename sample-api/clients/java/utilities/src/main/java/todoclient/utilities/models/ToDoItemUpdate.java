package todoclient.utilities.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
@jakarta.annotation.Generated("com.microsoft.kiota")
public class ToDoItemUpdate implements AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    private Map<String, Object> additionalData;
    /**
     * The dueDate property
     */
    private OffsetDateTime dueDate;
    /**
     * The isComplete property
     */
    private Boolean isComplete;
    /**
     * The name property
     */
    private String name;
    /**
     * The priority property
     */
    private Integer priority;
    /**
     * The type property
     */
    private String type;
    /**
     * Instantiates a new {@link ToDoItemUpdate} and sets the default values.
     */
    public ToDoItemUpdate() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a {@link ToDoItemUpdate}
     */
    @jakarta.annotation.Nonnull
    public static ToDoItemUpdate createFromDiscriminatorValue(@jakarta.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new ToDoItemUpdate();
    }
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a {@link Map<String, Object>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this.additionalData;
    }
    /**
     * Gets the dueDate property value. The dueDate property
     * @return a {@link OffsetDateTime}
     */
    @jakarta.annotation.Nullable
    public OffsetDateTime getDueDate() {
        return this.dueDate;
    }
    /**
     * The deserialization information for the current model
     * @return a {@link Map<String, java.util.function.Consumer<ParseNode>>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(5);
        deserializerMap.put("dueDate", (n) -> { this.setDueDate(n.getOffsetDateTimeValue()); });
        deserializerMap.put("isComplete", (n) -> { this.setIsComplete(n.getBooleanValue()); });
        deserializerMap.put("name", (n) -> { this.setName(n.getStringValue()); });
        deserializerMap.put("priority", (n) -> { this.setPriority(n.getIntegerValue()); });
        deserializerMap.put("type", (n) -> { this.setType(n.getStringValue()); });
        return deserializerMap;
    }
    /**
     * Gets the isComplete property value. The isComplete property
     * @return a {@link Boolean}
     */
    @jakarta.annotation.Nullable
    public Boolean getIsComplete() {
        return this.isComplete;
    }
    /**
     * Gets the name property value. The name property
     * @return a {@link String}
     */
    @jakarta.annotation.Nullable
    public String getName() {
        return this.name;
    }
    /**
     * Gets the priority property value. The priority property
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getPriority() {
        return this.priority;
    }
    /**
     * Gets the type property value. The type property
     * @return a {@link String}
     */
    @jakarta.annotation.Nullable
    public String getType() {
        return this.type;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public void serialize(@jakarta.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeOffsetDateTimeValue("dueDate", this.getDueDate());
        writer.writeBooleanValue("isComplete", this.getIsComplete());
        writer.writeStringValue("name", this.getName());
        writer.writeIntegerValue("priority", this.getPriority());
        writer.writeStringValue("type", this.getType());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     */
    public void setAdditionalData(@jakarta.annotation.Nullable final Map<String, Object> value) {
        this.additionalData = value;
    }
    /**
     * Sets the dueDate property value. The dueDate property
     * @param value Value to set for the dueDate property.
     */
    public void setDueDate(@jakarta.annotation.Nullable final OffsetDateTime value) {
        this.dueDate = value;
    }
    /**
     * Sets the isComplete property value. The isComplete property
     * @param value Value to set for the isComplete property.
     */
    public void setIsComplete(@jakarta.annotation.Nullable final Boolean value) {
        this.isComplete = value;
    }
    /**
     * Sets the name property value. The name property
     * @param value Value to set for the name property.
     */
    public void setName(@jakarta.annotation.Nullable final String value) {
        this.name = value;
    }
    /**
     * Sets the priority property value. The priority property
     * @param value Value to set for the priority property.
     */
    public void setPriority(@jakarta.annotation.Nullable final Integer value) {
        this.priority = value;
    }
    /**
     * Sets the type property value. The type property
     * @param value Value to set for the type property.
     */
    public void setType(@jakarta.annotation.Nullable final String value) {
        this.type = value;
    }
}
