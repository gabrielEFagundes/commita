#include <stdio.h>
#include "../tools/tools.h"

void exec_help()
{
    sprintf(buffer, 
        "\nCommita - v0.1.0"
        "\nWelcome to Commita!\n\n"
        "All commands starts with 'git'\n\n"
        "   feat        :sparkles: feat: <<message>>\n"
        "   fix         :bug: fix: <<message>>\n"
        "   chore       :art: chore: <<message>>\n"
        "   docs        :memo: docs: <<message>>\n"
        "   refactor    :recycle: refactor: <<message>>\n"
        "   remove      :fire: remove: <<message>>\n\n");

    printf("%s", buffer);
}

int main(int argc, char *argv[]){
    if(strcmp(argv[1], "commita") == 0){
        exec_help();
    }
}