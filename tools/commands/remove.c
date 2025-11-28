#include <stdio.h>
#include "../tools.h"

void exec_git_remove(int argc, char *argv[]){
    exec_command(":fire: rm: ", argc, argv);
}

int main(int argc, char *argv[]){
    exec_git_remove(argc, argv);
}