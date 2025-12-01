#include <stdio.h>
#include <stdlib.h>
#include "tools.h"

char buffer[1000];

int exec_command(char *prefix, int argc, char *argv[])
{
    if(argc > 4 || argc < 4){
        sprintf(buffer, "\nToo many/few arguments\nCorrect indentation: git <prefix> <message> <repo_url> <branch>\n\nType commita for help\n");
        printf("%s", buffer);
        exit(EXIT_FAILURE);
    }
    
    char message_buffer[512];
    char origin[512];

    sprintf(message_buffer, "%s %s", prefix, argv[1]);

    snprintf(origin, sizeof origin, 
            "git add -A && "
            "git commit -m %s && "
            "git remote remove origin && "
            "git remote add origin %s && "
            "git checkout -b %s || git checkout %s &&" // I have to rewrite this, or else it won't work on powershell/cmd
            "git push -u origin %s",
            message_buffer, argv[2], argv[3], argv[3], argv[3]);

    int statusCode = system(origin);

    if(statusCode != 0){
        sprintf(buffer, "\nStatus Code: %d", statusCode);
        printf("%s", buffer);
        exit(EXIT_FAILURE);
    }

    return 0;
}