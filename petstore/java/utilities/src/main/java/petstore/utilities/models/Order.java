package petstore.utilities.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
@jakarta.annotation.Generated("com.microsoft.kiota")
public class Order implements AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    private Map<String, Object> additionalData;
    /**
     * The complete property
     */
    private Boolean complete;
    /**
     * The id property
     */
    private Long id;
    /**
     * The petId property
     */
    private Long petId;
    /**
     * The quantity property
     */
    private Integer quantity;
    /**
     * The shipDate property
     */
    private OffsetDateTime shipDate;
    /**
     * Order Status
     */
    private OrderStatus status;
    /**
     * Instantiates a new {@link Order} and sets the default values.
     */
    public Order() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a {@link Order}
     */
    @jakarta.annotation.Nonnull
    public static Order createFromDiscriminatorValue(@jakarta.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new Order();
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
     * Gets the complete property value. The complete property
     * @return a {@link Boolean}
     */
    @jakarta.annotation.Nullable
    public Boolean getComplete() {
        return this.complete;
    }
    /**
     * The deserialization information for the current model
     * @return a {@link Map<String, java.util.function.Consumer<ParseNode>>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(6);
        deserializerMap.put("complete", (n) -> { this.setComplete(n.getBooleanValue()); });
        deserializerMap.put("id", (n) -> { this.setId(n.getLongValue()); });
        deserializerMap.put("petId", (n) -> { this.setPetId(n.getLongValue()); });
        deserializerMap.put("quantity", (n) -> { this.setQuantity(n.getIntegerValue()); });
        deserializerMap.put("shipDate", (n) -> { this.setShipDate(n.getOffsetDateTimeValue()); });
        deserializerMap.put("status", (n) -> { this.setStatus(n.getEnumValue(OrderStatus::forValue)); });
        return deserializerMap;
    }
    /**
     * Gets the id property value. The id property
     * @return a {@link Long}
     */
    @jakarta.annotation.Nullable
    public Long getId() {
        return this.id;
    }
    /**
     * Gets the petId property value. The petId property
     * @return a {@link Long}
     */
    @jakarta.annotation.Nullable
    public Long getPetId() {
        return this.petId;
    }
    /**
     * Gets the quantity property value. The quantity property
     * @return a {@link Integer}
     */
    @jakarta.annotation.Nullable
    public Integer getQuantity() {
        return this.quantity;
    }
    /**
     * Gets the shipDate property value. The shipDate property
     * @return a {@link OffsetDateTime}
     */
    @jakarta.annotation.Nullable
    public OffsetDateTime getShipDate() {
        return this.shipDate;
    }
    /**
     * Gets the status property value. Order Status
     * @return a {@link OrderStatus}
     */
    @jakarta.annotation.Nullable
    public OrderStatus getStatus() {
        return this.status;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public void serialize(@jakarta.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeBooleanValue("complete", this.getComplete());
        writer.writeLongValue("id", this.getId());
        writer.writeLongValue("petId", this.getPetId());
        writer.writeIntegerValue("quantity", this.getQuantity());
        writer.writeOffsetDateTimeValue("shipDate", this.getShipDate());
        writer.writeEnumValue("status", this.getStatus());
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
     * Sets the complete property value. The complete property
     * @param value Value to set for the complete property.
     */
    public void setComplete(@jakarta.annotation.Nullable final Boolean value) {
        this.complete = value;
    }
    /**
     * Sets the id property value. The id property
     * @param value Value to set for the id property.
     */
    public void setId(@jakarta.annotation.Nullable final Long value) {
        this.id = value;
    }
    /**
     * Sets the petId property value. The petId property
     * @param value Value to set for the petId property.
     */
    public void setPetId(@jakarta.annotation.Nullable final Long value) {
        this.petId = value;
    }
    /**
     * Sets the quantity property value. The quantity property
     * @param value Value to set for the quantity property.
     */
    public void setQuantity(@jakarta.annotation.Nullable final Integer value) {
        this.quantity = value;
    }
    /**
     * Sets the shipDate property value. The shipDate property
     * @param value Value to set for the shipDate property.
     */
    public void setShipDate(@jakarta.annotation.Nullable final OffsetDateTime value) {
        this.shipDate = value;
    }
    /**
     * Sets the status property value. Order Status
     * @param value Value to set for the status property.
     */
    public void setStatus(@jakarta.annotation.Nullable final OrderStatus value) {
        this.status = value;
    }
}
