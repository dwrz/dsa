typedef struct Vector {
	int *data;
	int capacity;
	int size;
} Vector;

void delete_vector(Vector *v);

Vector new_vector(int capacity);

void print_vector(Vector *v);
