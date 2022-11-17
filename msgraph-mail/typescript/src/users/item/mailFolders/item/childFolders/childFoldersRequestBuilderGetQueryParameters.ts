
<<<<<<< HEAD
export interface ChildFoldersRequestBuilderGetQueryParameters {
=======
/** Get the folder collection under the specified folder. You can use the `.../me/mailFolders` shortcut to get the top-level folder collection and navigate to another folder. By default, this operation does not return hidden folders. Use a query parameter _includeHiddenFolders_ to include them in the response. */
export class ChildFoldersRequestBuilderGetQueryParameters {
>>>>>>> main
    /** Include count of items */
    count?: boolean | undefined;
    /** Expand related entities */
    expand?: string[] | undefined;
    /** Filter items by property values */
    filter?: string | undefined;
    /** Order items by property values */
    orderby?: string[] | undefined;
    /** Select properties to be returned */
    select?: string[] | undefined;
    /** Skip the first n items */
    skip?: number | undefined;
    /** Show only the first n items */
    top?: number | undefined;
}
