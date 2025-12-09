#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "../tools.h"

int login(int argc, char *argv[]){
    char *abundance = NULL;
    char message[512];
    int ch;

    if(argc != 4){
        sprintf(message, "\nToo many/few arguments: %d\nCorrect indentation: git <prefix> [-b] [-u] <message>\n\nType commita for help\n", argc);
        printf("%s", message);
        exit(EXIT_FAILURE);
    }

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
                return 1;
            
            default:
                printf("\nWill this be applied locally or globally? Use -l or -g to specify that!\n");
                exit(EXIT_FAILURE);
        }
    }

    snprintf(message, sizeof message,
            "git config %s user.name %s && git config %s user.email %s",
            abundance, argv[2], abundance, argv[3]);

    int status = system(message);

    if(status != 0){
        printf("\nUsername: %s\nEmail: %s\nAbundance: %s", argv[2], argv[3], abundance);
        exit(EXIT_FAILURE);
    }
}

int main(int argc, char *argv[]){
    login(argc, argv);
    return 1;
}