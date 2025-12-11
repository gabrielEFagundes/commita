#include <stdio.h>

/// @brief Handles exceptions and errors that Git might return
///
/// Prints a user-friendly error message
///
/// @param err Parsed error message, used by the algorithm to find the exact error
void exception_handler(char *err);

typedef enum {
    // fatal: not a git repository
    NO_REPO,
    // error: failed to push some refs
    MISSING_REMOTE_CHANGES,
    // fatal: The current branch has no upstream branch
    NOT_SPECIFIED_REMOTE_BRANCH,
    // error: Your local changes would be overwritten by merge
    OVERWRITING_CHANGES,
    // fatal: Cannot do a partial commit during a merge
    COMMIT_AFTER_MERGE,
    // error: commit failed - cannot update HEAD
    MISSING_PERMISSIONS,
    // fatal: Unable to create 'path/.git/index.lock': File exists
    INTERRUPTED_OPERATION,
    // error: Remote ref does not exist
    NON_EXISTING_BRANCH,
    // fatal: Authentication failed
    INVALID_REMOTE_CREDENTIALS,
    // fatal: refusing to merge unrelated histories
    UNRELATED_HISTORIES
} GitExceptions;

/// @brief Used to parse an exception and return a more user-friendly error message.
///
/// @param exc Exception to be parsed, has to be part of GitExceptions enum
char* parse_exception(GitExceptions exc);