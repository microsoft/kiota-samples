package petstore.utilities.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
@jakarta.annotation.Generated("com.microsoft.kiota")
public class Pet implements AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    private Map<String, Object> additionalData;
    /**
     * The category property
     */
    private Category category;
    /**
     * The id property
     */
    private Long id;
    /**
     * The name property
     */
    private String name;
    /**
     * The photoUrls property
     */
    private java.util.List<String> photoUrls;
    /**
     * pet status in the store
     */
    private PetStatus status;
    /**
     * The tags property
     */
    private java.util.List<Tag> tags;
    /**
     * Instantiates a new {@link Pet} and sets the default values.
     */
    public Pet() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a {@link Pet}
     */
    @jakarta.annotation.Nonnull
    public static Pet createFromDiscriminatorValue(@jakarta.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new Pet();
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
     * Gets the category property value. The category property
     * @return a {@link Category}
     */
    @jakarta.annotation.Nullable
    public Category getCategory() {
        return this.category;
    }
    /**
     * The deserialization information for the current model
     * @return a {@link Map<String, java.util.function.Consumer<ParseNode>>}
     */
    @jakarta.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(6);
        deserializerMap.put("category", (n) -> { this.setCategory(n.getObjectValue(Category::createFromDiscriminatorValue)); });
        deserializerMap.put("id", (n) -> { this.setId(n.getLongValue()); });
        deserializerMap.put("name", (n) -> { this.setName(n.getStringValue()); });
        deserializerMap.put("photoUrls", (n) -> { this.setPhotoUrls(n.getCollectionOfPrimitiveValues(String.class)); });
        deserializerMap.put("status", (n) -> { this.setStatus(n.getEnumValue(PetStatus::forValue)); });
        deserializerMap.put("tags", (n) -> { this.setTags(n.getCollectionOfObjectValues(Tag::createFromDiscriminatorValue)); });
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
     * Gets the name property value. The name property
     * @return a {@link String}
     */
    @jakarta.annotation.Nullable
    public String getName() {
        return this.name;
    }
    /**
     * Gets the photoUrls property value. The photoUrls property
     * @return a {@link java.util.List<String>}
     */
    @jakarta.annotation.Nullable
    public java.util.List<String> getPhotoUrls() {
        return this.photoUrls;
    }
    /**
     * Gets the status property value. pet status in the store
     * @return a {@link PetStatus}
     */
    @jakarta.annotation.Nullable
    public PetStatus getStatus() {
        return this.status;
    }
    /**
     * Gets the tags property value. The tags property
     * @return a {@link java.util.List<Tag>}
     */
    @jakarta.annotation.Nullable
    public java.util.List<Tag> getTags() {
        return this.tags;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public void serialize(@jakarta.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeObjectValue("category", this.getCategory());
        writer.writeLongValue("id", this.getId());
        writer.writeStringValue("name", this.getName());
        writer.writeCollectionOfPrimitiveValues("photoUrls", this.getPhotoUrls());
        writer.writeEnumValue("status", this.getStatus());
        writer.writeCollectionOfObjectValues("tags", this.getTags());
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
     * Sets the category property value. The category property
     * @param value Value to set for the category property.
     */
    public void setCategory(@jakarta.annotation.Nullable final Category value) {
        this.category = value;
    }
    /**
     * Sets the id property value. The id property
     * @param value Value to set for the id property.
     */
    public void setId(@jakarta.annotation.Nullable final Long value) {
        this.id = value;
    }
    /**
     * Sets the name property value. The name property
     * @param value Value to set for the name property.
     */
    public void setName(@jakarta.annotation.Nullable final String value) {
        this.name = value;
    }
    /**
     * Sets the photoUrls property value. The photoUrls property
     * @param value Value to set for the photoUrls property.
     */
    public void setPhotoUrls(@jakarta.annotation.Nullable final java.util.List<String> value) {
        this.photoUrls = value;
    }
    /**
     * Sets the status property value. pet status in the store
     * @param value Value to set for the status property.
     */
    public void setStatus(@jakarta.annotation.Nullable final PetStatus value) {
        this.status = value;
    }
    /**
     * Sets the tags property value. The tags property
     * @param value Value to set for the tags property.
     */
    public void setTags(@jakarta.annotation.Nullable final java.util.List<Tag> value) {
        this.tags = value;
    }
}
