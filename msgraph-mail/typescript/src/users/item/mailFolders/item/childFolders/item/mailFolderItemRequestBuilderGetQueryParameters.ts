
export interface MailFolderItemRequestBuilderGetQueryParameters {
    /**
     * Expand related entities
     */
    expand?: string[] | undefined;
    /**
     * Include Hidden Folders
     */
    includeHiddenFolders?: string | undefined;
    /**
     * Select properties to be returned
     */
    select?: string[] | undefined;
}
