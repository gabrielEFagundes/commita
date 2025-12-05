#include <stdio.h>
#include <unistd.h>
#include <getopt.h>
#include "../tools/tools.h"

char buffer[1000];

static struct option long_options[] = {
    {"version", no_argument, NULL, 'v'},
    {"help", no_argument, NULL, 'h'},
    {NULL, 0, NULL, 0}
};

void show_help()
{
    sprintf(buffer, 
        "\nCommita - v1.1.1"
        "\nWelcome to Commita!\n\n"
        "All commands starts with 'git'\n\n"
        "   feat        :sparkles: feat: <<message>>\n"
        "   fix         :bug: fix: <<message>>\n"
        "   chore       :art: chore: <<message>>\n"
        "   docs        :memo: docs: <<message>>\n"
        "   refactor    :recycle: refactor: <<message>>\n"
        "   remove      :fire: rm: <<message>>\n\n"
        "The syntax for most commands are:\n\n"
        "git [prefix] -u <<repo's url>> -b <<branch>> <<message>>\n\n"
        "   prefix      the type of commit message you want\n"
        "   -b          (optional) your desired branch, uses the default 'main' branch if not parsed\n"
        "   -u          (optional) your repo's url, only parse on your first commit or if not parsed yet\n"
        "   message     your commit message\n\n");

    printf("%s", buffer);
}

void show_version(){
    sprintf(buffer, "\nCommita - v1.1.1\n");
    printf("%s", buffer);
}

void wrong_opt(){
    sprintf(buffer, "\nUnknown flag parsed, the available options are:\n\n"
                    "   -v     Version\n"
                    "   -h     Help\n");
    printf("%s", buffer);
}

int main(int argc, char *argv[]){
    int ch;
    while((ch = getopt_long(argc, argv, "hv", long_options, 0))){
        switch(ch){
            case 'h':
                show_help();
                exit(EXIT_SUCCESS);

            case 'v':
                show_version();
                exit(EXIT_SUCCESS);

            case '?':
                wrong_opt();
                exit(EXIT_FAILURE);

            default:
                show_help();
                exit(EXIT_SUCCESS);
        }
    }
    return 0;
}