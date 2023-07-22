#include <stdio.h>

void pointer_value_and_addresses();
void pointer_arithmetics();

int main(void) {
  // Call the function to demonstrate pointer value and memory addresses
  pointer_value_and_addresses();
  printf("\n");
  
  // Call the function to demonstrate pointer arithmetic with arrays
  pointer_arithmetics();
  
  return 0;
}

void pointer_value_and_addresses() {
  // This function demonstrates how pointers work and their relationship to memory addresses.
  printf("POINTER VALUE AND ADDRESSES: \n");

  int n = 50;
  printf("Initial value of N: %i\n", n);
  
  int *p = &n;
  printf("Value of N through P: %i\n", *p);
  
  printf("Memory address of N: %p\n", (void*)&n);
  printf("Memory address of P: %p\n", (void*)&p);
  
  printf("P points to the memory address: %p\n", (void*)p);

  *p = 100;
  printf("N changed using P. N now equals: %i\n", n);
}

void pointer_arithmetics() {
  // This function demonstrates pointer arithmetic with arrays to access and modify array elements.
  printf("POINTER ARITHMETICS: \n");

  int a[6] = {1, 2, 3, 4, 5, 0};
  int *b = a; 

  printf("Initial values of Array 'a': \n");
  for(int i = 0; i < 5; i++){
    printf("%i ", a[i]);
  }
  printf("\n");

  *b = 11;
  *(b + 1) = 22;
  *(b + 2) = 33;
  *(b + 3) = 44;
  *(b + 4) = 55;

  printf("After changing Array 'a' using Pointer arithmetic: \n");
  for(int i = 0; i < 5; i++){
    printf("%i ", *(b + i));
  }
  printf("\n");
}
