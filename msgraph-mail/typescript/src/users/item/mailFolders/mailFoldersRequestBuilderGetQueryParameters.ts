
<<<<<<< HEAD
export interface MailFoldersRequestBuilderGetQueryParameters {
=======
/** Get the mail folder collection directly under the root folder of the signed-in user. The returned collection includes any mail search folders directly under the root. By default, this operation does not return hidden folders. Use a query parameter _includeHiddenFolders_ to include them in the response. */
export class MailFoldersRequestBuilderGetQueryParameters {
>>>>>>> main
    /** Include count of items */
    count?: boolean | undefined;
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
