package graphjavav4.utilities.users;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.RequestAdapter;
import graphjavav4.utilities.users.item.UserItemRequestBuilder;
import java.util.HashMap;
import java.util.Objects;
/**
 * Builds and executes requests for operations under /users
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class UsersRequestBuilder extends BaseRequestBuilder {
    /**
     * Gets an item from the graphjavav4.utilities.users.item collection
     * @param userId The unique identifier of user
     * @return a {@link UserItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public UserItemRequestBuilder byUserId(@jakarta.annotation.Nonnull final String userId) {
        Objects.requireNonNull(userId);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("user%2Did", userId);
        return new UserItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new {@link UsersRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public UsersRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/users", pathParameters);
    }
    /**
     * Instantiates a new {@link UsersRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public UsersRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/users", rawUrl);
    }
}
