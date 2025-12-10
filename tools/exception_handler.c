#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "headers/exceptions.h"

void exception_handler(char *err){
    // here's where the I'll have loads of fun ;P
}

void parse_exception(GitExceptions exc){
    switch(exc){
        case NO_REPO:
            // todo
            break;

        case MISSING_REMOTE_CHANGES:
            // todo
            break;

        case NOT_SPECIFIED_REMOTE_BRANCH:
            // todo
            break;

        case OVERWRITING_CHANGES:
            // todo
            break;

        case COMMIT_AFTER_MERGE:
            // todo
            break;

        case MISSING_PERMISSIONS:
            // todo
            break;

        case INTERRUPTED_OPERATION:
            // todo
            break;

        case NON_EXISTING_BRANCH:
            // todo
            break;

        case INVALID_REMOTE_CREDENTIALS:
            // todo
            break;

        case UNRELATED_HISTORIES:
            // todo
            break;

        default:
            // should be if the error is not found or if exc is NULL
            break;
    }
}