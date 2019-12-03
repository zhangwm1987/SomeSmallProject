#include<stdio.h>
#include<stdlib.h>
#include <unistd.h>

typedef struct node {
    int val;
    struct node* next;
} node_t;

node_t* newNode(int data) {
    node_t *node = malloc(sizeof(node_t));
    node->val = data;
    node->next = NULL;
    return node;
}

void main(){
    node_t *head = newNode(0);
    node_t *tmp = head;
    int i = 0;

    for (i = 0; i < 5432; i++) {
        tmp = head;
        while(tmp) {
            if ((tmp->val + 1) < 10) {
                tmp->val = tmp->val + 1;
                tmp = tmp->next;
                break;
            } else {
                tmp->val = 0;
                if (tmp->next == NULL) {
                    tmp->next = newNode(0);
                }
                tmp = tmp->next;
            }
        }
    }
    
    printf("#############\n");
    while(head) {
        printf("%d ", head->val);
        head = head->next;
    }
}


