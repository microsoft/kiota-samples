package petstore.utilities.pet.findbystatus;

import com.microsoft.kiota.serialization.ValuedEnum;
import java.util.Objects;

@jakarta.annotation.Generated("com.microsoft.kiota")
public enum GetStatusQueryParameterType implements ValuedEnum {
    Available("available"),
    Pending("pending"),
    Sold("sold");
    public final String value;
    GetStatusQueryParameterType(final String value) {
        this.value = value;
    }
    @jakarta.annotation.Nonnull
    public String getValue() { return this.value; }
    @jakarta.annotation.Nullable
    public static GetStatusQueryParameterType forValue(@jakarta.annotation.Nonnull final String searchValue) {
        Objects.requireNonNull(searchValue);
        switch(searchValue) {
            case "available": return Available;
            case "pending": return Pending;
            case "sold": return Sold;
            default: return null;
        }
    }
}
