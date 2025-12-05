#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "../tools.h"

void login(int argc, char *argv[]){
    char *abundance = NULL;
    char message[512];
    int ch;

    while((ch = getopt(argc, argv, "lg")) != -1){
        switch(ch){
            case 'l':
                abundance = "--local";
                break;

            case 'g':
                abundance = "--global";
                break;

            case '?':
                printf("\nUnknown flag parsed. Type 'commita' or 'commita --help' for help\n");
                break;
            
            default:
                printf("\nWill this be applied locally or globally? Use -l or -g to specify that!\n");
                break;
        }
    }

    snprintf(message, sizeof message,
            "git config %s user.name %s && git config %s user.email %s",
            abundance, argv[optind], argv[optind+1]);

    printf("%s", message);
}

int main(int argc, char *argv[]){
    login(argc, argv);
}

/*
    git config (--local or --global) user.name (name) && git config (--local or --global) user.email (email)
*/