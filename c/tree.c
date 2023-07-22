#include <stdio.h>
#include <stdlib.h>

typedef struct tree {
    int value;
    struct tree *left;
    struct tree *right;
} tree;

void print_tree(tree *root);
void free_memory(tree *root);
void search(tree *root, int value);

int main(void) {
  // Create a binary tree with three nodes: root, left, and right
  tree *root = malloc(sizeof(tree));
  tree *left = malloc(sizeof(tree));
  tree *right = malloc(sizeof(tree));

  // Check if memory allocation was successful for all nodes
  if (root == NULL || left == NULL || right == NULL) {
    // If any allocation failed, free the memory allocated so far
    free_memory(root);

    printf("Error: malloc failed!\n");
    return 1;
  }

  // Set values and connections for the binary tree
  root->value = 2;
  root->left = left;
  root->right = right;

  left->value = 0;
  right->value = 10;

  // Print the binary tree in-order traversal
  print_tree(root);

  // Search for a specific value in the binary tree
  search(root, 0);

  // Free the allocated memory for the binary tree
  free_memory(root);

  return 0;
}

void print_tree(tree *root) {
  // Perform in-order traversal to print the binary tree in ascending order
  if (root == NULL) {
    return;
  }

  print_tree(root->left);
  printf("%d\n", root->value);
  print_tree(root->right);
}

void free_memory(tree *root) {
  // Recursively free the memory of the binary tree nodes
  if (root == NULL) {
    return;
  }

  free_memory(root->left);
  free(root);
  free_memory(root->right);
}

void search(tree *root, int value) {
  // Perform a binary search for a given value in the binary tree
  if (root == NULL) {
    printf("Value not found!\n");
    return;
  }

  if (root->value == value) {
    printf("Found %d\n", value);
    return;
  }

  if (root->value > value) {
    search(root->left, value);
  } else {
    search(root->right, value);
  }
}
