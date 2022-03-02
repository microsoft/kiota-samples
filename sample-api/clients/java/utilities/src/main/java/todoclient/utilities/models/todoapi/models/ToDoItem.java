package todoclient.utilities.models.todoapi.models;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class ToDoItem implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    private Map<String, Object> _additionalData;
    private OffsetDateTime _dueDate;
    private String _id;
    private Boolean _isComplete;
    private String _name;
    private Integer _priority;
    private String _type;
    /**
     * Instantiates a new ToDoItem and sets the default values.
     * @return a void
     */
    public ToDoItem() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this._additionalData;
    }
    /**
     * Gets the dueDate property value. 
     * @return a OffsetDateTime
     */
    @javax.annotation.Nullable
    public OffsetDateTime getDueDate() {
        return this._dueDate;
    }
    /**
     * Gets the id property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getId() {
        return this._id;
    }
    /**
     * Gets the isComplete property value. 
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsComplete() {
        return this._isComplete;
    }
    /**
     * Gets the name property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getName() {
        return this._name;
    }
    /**
     * Gets the priority property value. 
     * @return a integer
     */
    @javax.annotation.Nullable
    public Integer getPriority() {
        return this._priority;
    }
    /**
     * Gets the type property value. 
     * @return a string
     */
    @javax.annotation.Nullable
    public String getType() {
        return this._type;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(6) {{
            this.put("dueDate", (o, n) -> { ((ToDoItem)o).setDueDate(n.getOffsetDateTimeValue()); });
            this.put("id", (o, n) -> { ((ToDoItem)o).setId(n.getStringValue()); });
            this.put("isComplete", (o, n) -> { ((ToDoItem)o).setIsComplete(n.getBooleanValue()); });
            this.put("name", (o, n) -> { ((ToDoItem)o).setName(n.getStringValue()); });
            this.put("priority", (o, n) -> { ((ToDoItem)o).setPriority(n.getIntegerValue()); });
            this.put("type", (o, n) -> { ((ToDoItem)o).setType(n.getStringValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeOffsetDateTimeValue("dueDate", this.getDueDate());
        writer.writeStringValue("id", this.getId());
        writer.writeBooleanValue("isComplete", this.getIsComplete());
        writer.writeStringValue("name", this.getName());
        writer.writeIntegerValue("priority", this.getPriority());
        writer.writeStringValue("type", this.getType());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    public void setAdditionalData(@javax.annotation.Nullable final Map<String, Object> value) {
        this._additionalData = value;
    }
    /**
     * Sets the dueDate property value. 
     * @param value Value to set for the dueDate property.
     * @return a void
     */
    public void setDueDate(@javax.annotation.Nullable final OffsetDateTime value) {
        this._dueDate = value;
    }
    /**
     * Sets the id property value. 
     * @param value Value to set for the id property.
     * @return a void
     */
    public void setId(@javax.annotation.Nullable final String value) {
        this._id = value;
    }
    /**
     * Sets the isComplete property value. 
     * @param value Value to set for the isComplete property.
     * @return a void
     */
    public void setIsComplete(@javax.annotation.Nullable final Boolean value) {
        this._isComplete = value;
    }
    /**
     * Sets the name property value. 
     * @param value Value to set for the name property.
     * @return a void
     */
    public void setName(@javax.annotation.Nullable final String value) {
        this._name = value;
    }
    /**
     * Sets the priority property value. 
     * @param value Value to set for the priority property.
     * @return a void
     */
    public void setPriority(@javax.annotation.Nullable final Integer value) {
        this._priority = value;
    }
    /**
     * Sets the type property value. 
     * @param value Value to set for the type property.
     * @return a void
     */
    public void setType(@javax.annotation.Nullable final String value) {
        this._type = value;
    }
}
