#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "headers/exceptions.h"

void exception_handler(char *err){
    // here's where I'll have loads of fun ;P
}

void parse_exception(GitExceptions exc){
    switch(exc){
        case NO_REPO: return "";

        case MISSING_REMOTE_CHANGES: return "";

        case NOT_SPECIFIED_REMOTE_BRANCH: return "";

        case OVERWRITING_CHANGES: return "";

        case COMMIT_AFTER_MERGE: return "";

        case MISSING_PERMISSIONS: return "";

        case INTERRUPTED_OPERATION: return "";

        case NON_EXISTING_BRANCH: return "";

        case INVALID_REMOTE_CREDENTIALS: return "";

        case UNRELATED_HISTORIES: return "";

        default: return ""; // should be if the error is not found or if exc is NULL
    }
}