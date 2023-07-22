#include <stdio.h>
#include <stdlib.h>

typedef struct node {
    int number;
    struct node* next;
} node;

void free_linked_list_memory(node* list);

int main(void) {
    node* list = malloc(sizeof(node));
    node* list_item_1 = malloc(sizeof(node));
    node* list_item_2 = malloc(sizeof(node));
    node* list_item_3 = malloc(sizeof(node));

    if (list_item_1 == NULL || list_item_2 == NULL || list_item_3 == NULL) {
        free_linked_list_memory(list);
        printf("Error: malloc failed!\n");
        return 1;
    }

    list->number = 1;
    list->next = list_item_1;

    list_item_1->number = 2;
    list_item_1->next = list_item_2;

    list_item_2->number = 3;
    list_item_2->next = list_item_3;

    list_item_3->number = 4;
    list_item_3->next = NULL;

    for (node* current = list; current != NULL; current = current->next) {
        printf("%d\n", current->number);
    }

    free_linked_list_memory(list);
    return 0;
}

void free_linked_list_memory(node* list) {
    node* current = list;

    while (current != NULL) {
        node* temp = current;
        current = current->next;
        free(temp);
    }
}
