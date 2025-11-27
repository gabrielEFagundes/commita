#include <stdio.h>
#include "../tools/tools.h"

int exec_command(const char *branch, char *prefix, int argcount, char *argvec[]){
    if((argcount > 2 || argcount < 2) || argvec[0] != "commita"){
        sprintf("Too many/few arguments\nCorrect indentation: git <prefix> <message>\n\nType commita for help", buffer);
        exit(EXIT_FAILURE);
    }

    char message_buffer[512];

    strcpy(message_buffer, prefix);

    strcat(message_buffer, argvec[1]);

    system("git add .");
    system(strcat("git commit -m ", message_buffer));

    exit(EXIT_SUCCESS);
}

void exec_help(){
    sprintf(buffer, "\nWelcome to Commita\n\n"
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