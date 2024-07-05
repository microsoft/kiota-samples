package petstore.utilities.user;

import com.microsoft.kiota.BaseRequestBuilder;
import com.microsoft.kiota.BaseRequestConfiguration;
import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import java.io.InputStream;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import petstore.utilities.models.User;
import petstore.utilities.user.createwitharray.CreateWithArrayRequestBuilder;
import petstore.utilities.user.createwithlist.CreateWithListRequestBuilder;
import petstore.utilities.user.item.WithUsernameItemRequestBuilder;
import petstore.utilities.user.login.LoginRequestBuilder;
import petstore.utilities.user.logout.LogoutRequestBuilder;
/**
 * Builds and executes requests for operations under /user
 */
@jakarta.annotation.Generated("com.microsoft.kiota")
public class UserRequestBuilder extends BaseRequestBuilder {
    /**
     * The createWithArray property
     * @return a {@link CreateWithArrayRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public CreateWithArrayRequestBuilder createWithArray() {
        return new CreateWithArrayRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The createWithList property
     * @return a {@link CreateWithListRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public CreateWithListRequestBuilder createWithList() {
        return new CreateWithListRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The login property
     * @return a {@link LoginRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public LoginRequestBuilder login() {
        return new LoginRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * The logout property
     * @return a {@link LogoutRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public LogoutRequestBuilder logout() {
        return new LogoutRequestBuilder(pathParameters, requestAdapter);
    }
    /**
     * Gets an item from the petstore.utilities.user.item collection
     * @param username The name that needs to be fetched. Use user1 for testing. 
     * @return a {@link WithUsernameItemRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public WithUsernameItemRequestBuilder byUsername(@jakarta.annotation.Nonnull final String username) {
        Objects.requireNonNull(username);
        final HashMap<String, Object> urlTplParams = new HashMap<String, Object>(this.pathParameters);
        urlTplParams.put("username", username);
        return new WithUsernameItemRequestBuilder(urlTplParams, requestAdapter);
    }
    /**
     * Instantiates a new {@link UserRequestBuilder} and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public UserRequestBuilder(@jakarta.annotation.Nonnull final HashMap<String, Object> pathParameters, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/user", pathParameters);
    }
    /**
     * Instantiates a new {@link UserRequestBuilder} and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public UserRequestBuilder(@jakarta.annotation.Nonnull final String rawUrl, @jakarta.annotation.Nonnull final RequestAdapter requestAdapter) {
        super(requestAdapter, "{+baseurl}/user", rawUrl);
    }
    /**
     * This can only be done by the logged in user.
     * @param body The request body
     * @return a {@link InputStream}
     */
    @jakarta.annotation.Nullable
    public InputStream post(@jakarta.annotation.Nonnull final User body) {
        return post(body, null);
    }
    /**
     * This can only be done by the logged in user.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a {@link InputStream}
     */
    @jakarta.annotation.Nullable
    public InputStream post(@jakarta.annotation.Nonnull final User body, @jakarta.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = toPostRequestInformation(body, requestConfiguration);
        return this.requestAdapter.sendPrimitive(requestInfo, null, InputStream.class);
    }
    /**
     * This can only be done by the logged in user.
     * @param body The request body
     * @return a {@link RequestInformation}
     */
    @jakarta.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@jakarta.annotation.Nonnull final User body) {
        return toPostRequestInformation(body, null);
    }
    /**
     * This can only be done by the logged in user.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @return a {@link RequestInformation}
     */
    @jakarta.annotation.Nonnull
    public RequestInformation toPostRequestInformation(@jakarta.annotation.Nonnull final User body, @jakarta.annotation.Nullable final java.util.function.Consumer<PostRequestConfiguration> requestConfiguration) {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = new RequestInformation(HttpMethod.POST, urlTemplate, pathParameters);
        requestInfo.configure(requestConfiguration, PostRequestConfiguration::new);
        requestInfo.headers.tryAdd("Accept", "application/json, application/xml");
        requestInfo.setContentFromParsable(requestAdapter, "application/json", body);
        return requestInfo;
    }
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @return a {@link UserRequestBuilder}
     */
    @jakarta.annotation.Nonnull
    public UserRequestBuilder withUrl(@jakarta.annotation.Nonnull final String rawUrl) {
        Objects.requireNonNull(rawUrl);
        return new UserRequestBuilder(rawUrl, requestAdapter);
    }
    /**
     * Configuration for the request such as headers, query parameters, and middleware options.
     */
    @jakarta.annotation.Generated("com.microsoft.kiota")
    public class PostRequestConfiguration extends BaseRequestConfiguration {
    }
}
