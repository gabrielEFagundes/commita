#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <ctype.h>
#include <string.h>
#include "headers/tools.h"

char buffer[1000];

int exec_command(char *prefix, int argc, char *argv[])
{
    if(argc != 1){
        sprintf(buffer, "\nToo many/few arguments: %d\nCorrect indentation: git <prefix> [-b] [-u] <message>\n\nType commita for help\n", argc);
        printf("%s", buffer);
        exit(EXIT_FAILURE);
    }
    
    char message_buffer[512], commit[512], addOrigin[512], checkpush[512], origin[512];
    char *branch = NULL, *url = NULL;

    int ch;
    while((ch = getopt(argc, argv, "b:u:")) != -1){
        switch(ch){
            case 'b':
                branch = optarg;
                snprintf(checkpush, sizeof checkpush,
                        "git checkout -b %s || git checkout %s && " // I have to rewrite this, or else it won't work on powershell/cmd
                        "git push -u origin %s",
                        branch, branch, branch);
                break;

            case 'u':
                url = optarg;
                snprintf(addOrigin, sizeof addOrigin,
                        "git remote add origin %s && ",
                        url);
                break;

            case '?':
                printf("Unknown flag parsed. Type 'commita' or 'commita --help' for help");
                return 1;
        }
    }

    sprintf(message_buffer, "%s %s", prefix, argv[optind]);
    snprintf(commit, sizeof commit,
            "git add -A && "
            "git commit -m \"%s\" && ",
            message_buffer);

    if(branch == NULL){
        snprintf(checkpush, sizeof checkpush,
                "git checkout -b main || git checkout main && "
                "git push -u origin main");
    }

    snprintf(origin, sizeof origin, "%s %s %s", addOrigin, commit, checkpush);

    int statusCode = system(origin);

    if(statusCode != 0){
        sprintf(buffer, "\nStatus Code: %d\nCommand: %s\n\nUrl: %s\nBranch: %s", statusCode, origin, url, branch);
        printf("%s", buffer);
        exit(EXIT_FAILURE);
    }

    return 0;
}