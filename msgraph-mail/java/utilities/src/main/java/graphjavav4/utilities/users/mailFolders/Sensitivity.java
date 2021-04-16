package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.ValuedEnum;

public enum Sensitivity implements ValuedEnum {
    Normal("normal"),
    Personal("personal"),
    Private("private"),
    Confidential("confidential");
    public final String value;
    Sensitivity(final String value) {
        this.value = value;
    }
    @javax.annotation.Nonnull
    public String getValue() { return this.value; }
    @javax.annotation.Nullable
    public static Sensitivity forValue(@javax.annotation.Nonnull final String searchValue) {
        switch(searchValue) {
            case "normal": return Normal;
            case "personal": return Personal;
            case "private": return Private;
            case "confidential": return Confidential;
            default: return null;
        }
    }
}
