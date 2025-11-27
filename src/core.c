#include <stdio.h>
#include "../tools/tools.h"

char buffer[1000];

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
        "   remove      :fire: remove: <<message>>\n\n"
        "The syntax for most commands are:\n\n"
        "git [prefix] <<commit message>> <<branch>> <<repository url>>\n\n"
        "   prefix      the type of commit message you want"
        "   message     your commit's message\n"
        "   branch      your repo's branch, can also be a new one\n"
        "   url         your repo's url\n\n");

    printf("%s", buffer);
}

int main(int argc, char *argv[]){
    if(strcmp(argv[1], "commita") == 0){
        exec_help();
    }
}