#include <stdio.h>
#include <stdlib.h>

int* createArray(int length);
void printArray(int* array, int length);

int main() {
	int length = 8;
	int* array = createArray(length);
	printArray(array, length);
	return 0;
}

int* createArray(int length) {
	int* array = (int *) malloc(length * sizeof(int));
	int i = 0;
	for (i = 0; i < length; i++) {
		array[i] = i;
	}
	return array;
}

void printArray(int* array, int length) {
	int i = 0;
	for (i = 0; i < length; i++) {
		printf("%d\n", array[i]);
	}
}
