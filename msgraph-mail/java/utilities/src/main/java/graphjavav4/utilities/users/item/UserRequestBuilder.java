package graphjavav4.utilities.users.item;

import com.microsoft.kiota.RequestAdapter;
import graphjavav4.utilities.users.item.inferenceClassification.InferenceClassificationRequestBuilder;
import graphjavav4.utilities.users.item.mailFolders.item.MailFolderRequestBuilder;
import graphjavav4.utilities.users.item.mailFolders.MailFoldersRequestBuilder;
import graphjavav4.utilities.users.item.messages.item.MessageRequestBuilder;
import graphjavav4.utilities.users.item.messages.MessagesRequestBuilder;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}  */
public class UserRequestBuilder {
    /** Current path for the request  */
    private final String currentPath;
    @javax.annotation.Nonnull
    public InferenceClassificationRequestBuilder inferenceClassification() {
        return new InferenceClassificationRequestBuilder(currentPath + pathSegment, requestAdapter, false);
    }
    /** Whether the current path is a raw URL  */
    private final boolean isRawUrl;
    @javax.annotation.Nonnull
    public MailFoldersRequestBuilder mailFolders() {
        return new MailFoldersRequestBuilder(currentPath + pathSegment, requestAdapter, false);
    }
    @javax.annotation.Nonnull
    public MessagesRequestBuilder messages() {
        return new MessagesRequestBuilder(currentPath + pathSegment, requestAdapter, false);
    }
    /** Path segment to use to build the URL for the current request builder  */
    private final String pathSegment;
    /** The http core service to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param requestAdapter The http core service to use to execute the requests.
     * @return a void
     */
    public UserRequestBuilder(final String currentPath, final RequestAdapter requestAdapter) {
        this(currentPath, requestAdapter, true);
    }
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param currentPath Current path for the request
     * @param isRawUrl Whether the current path is a raw URL
     * @param requestAdapter The http core service to use to execute the requests.
     * @return a void
     */
    public UserRequestBuilder(@javax.annotation.Nonnull final String currentPath, @javax.annotation.Nonnull final RequestAdapter requestAdapter, final boolean isRawUrl) {
        Objects.requireNonNull(currentPath);
        Objects.requireNonNull(requestAdapter);
        this.pathSegment = "";
        this.requestAdapter = requestAdapter;
        this.currentPath = currentPath;
        this.isRawUrl = isRawUrl;
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.mailFolders.item collection
     * @param id Unique identifier of the item
     * @return a mailFolderRequestBuilder
     */
    @javax.annotation.Nonnull
    public MailFolderRequestBuilder mailFolders(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MailFolderRequestBuilder(currentPath + pathSegment + "/mailFolders/" + id, requestAdapter, false);
    }
    /**
     * Gets an item from the graphjavav4.utilities.users.item.messages.item collection
     * @param id Unique identifier of the item
     * @return a messageRequestBuilder
     */
    @javax.annotation.Nonnull
    public MessageRequestBuilder messages(@javax.annotation.Nonnull final String id) {
        Objects.requireNonNull(id);
        return new MessageRequestBuilder(currentPath + pathSegment + "/messages/" + id, requestAdapter, false);
    }
}
