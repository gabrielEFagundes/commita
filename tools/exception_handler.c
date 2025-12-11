#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "headers/exceptions.h"
#include "auxiliers/ansi_colors.h"

void exception_handler(char *err){
    // here's where I'll have loads of fun ;P
}

char* parse_exception(GitExceptions exc){
    switch(exc){
        case NO_REPO: return
            "%sFatal!\n%s"
            "The directory you're in is %snot%s a git repository!\n"
            "Please, use \"%sgit init%s\" to start a new repository.", 
            RED, RESET, RED, RESET, LIGHT_BLUE, RESET;

        case MISSING_REMOTE_CHANGES: return
            "%sFatal!\n%s"
            "Your %sremote%s repository had recent changes that you %sdo not%s have on your %slocal%s repository!\n"
            "Please, use \"%sgit pull%s\" to pull the remote changes to your local repository.", 
            RED, RESET, BLUE, RESET, RED, RESET, LIGHT_GREEN, RESET, LIGHT_BLUE, RESET;
            
        case NOT_SPECIFIED_REMOTE_BRANCH: return
            "%sFatal!\n%s"
            "You're %snot%s on the same %sbranch%s that you want to commit to!\n"
            "Please, use the flag \"%s-b%s\" when using commita's commit/push system to specify your branch!\n"
            "%sNote that if you don't specify which branch you want to commit, commita assumes that the default branch is \"main\"!%s",
            RED, RESET, RED, RESET, LIGHT_BLUE, RESET, BLUE, RESET, GRAY, RESET;

        case OVERWRITING_CHANGES: return
            "%sFatal!\n%s"
            "Your local changes will be %soverwritten%s by your remote changes!\n"
            "Try %spulling and pushing%s your changes to your remote repository, or %sundo%s your changes and pull your remote.",
            RED, RESET, LIGHT_YELLOW, RESET, LIGHT_BLUE, RESET, PURPLE, RESET;

        case COMMIT_AFTER_MERGE: return
            "%sFatal!\n%s"
            "I'm sorry, but you have %smerge issues%s!\n"
            "Unfortunately, commita %scannot%s solve this issues for you. Please, solve them and then try commiting.",
            RED, RESET, LIGHT_YELLOW, RESET, LIGHT_RED, RESET;

        case MISSING_PERMISSIONS: return
            "%sFatal!\n%s"
            "Commita %scouldn't%s complete your commit.\n"
            "This usually happens when you %sdon't have permission%s to commit to the local repository. Look out for that!",
            RED, RESET, LIGHT_RED, RESET, LIGHT_BLUE, RESET;

        case INTERRUPTED_OPERATION: return
            "%sFatal!\n%s"
            "Seems like git's running %sanother process%s and doing what you're trying to do would interrupt that process!\n"
            "If %sno other%s process of git is running (like any commita's process, just as git pull or push), delete the file index.lock.\n"
            "%sNote that the file is usually located on .git/index.lock%s",
            RED, RESET, BLUE, RESET, LIGHT_BLUE, RESET, GRAY, RESET;

        case NON_EXISTING_BRANCH: return
            "%sFatal!\n%s"
            "The remote branch you're trying to commit to %sdoes not%s exist!\n"
            "Commita doesn't give you this error very often, usually only when the default branch %sis not%s main!",
            RED, RESET, LIGHT_RED, RESET, LIGHT_BLUE, RESET;

        case INVALID_REMOTE_CREDENTIALS: return
            "%sFatal!\n%s"
            "Commita did not recognize you!\n"
            "Use \"%sgit login%s\" and log yourself into a valid account with permissions to push changes to the remote repository!",
            RED, RESET, LIGHT_BLUE, RESET;

        case UNRELATED_HISTORIES: return
            "%sFatal!\n%s"
            "There's no common commits between your %scurrent branch%s and %sremote repository!%s\n"
            "If you're seeing this error, it means commita %scould not%s solve the issue, we're sorry about that!\n"
            "%sFor more help, check out this link: https://www.geeksforgeeks.org/git/how-to-fix-git-refusing-to-merge-unrelated-histories/ %s",
            RED, RESET, BLUE, RESET, LIGHT_PURPLE, RESET, LIGHT_YELLOW, RESET, GRAY, RESET;

        default: return "%sNo error found!%s", LIGHT_GREEN, RESET; // should be if the error is not found or if exc is NULL
    }
}