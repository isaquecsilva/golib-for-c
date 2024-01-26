#include <stdio.h>
#include <stdlib.h>
#include "../build/lib/gobase64.h"

int main()
{
	// Encode
	char* encoded = CharPtrToB64("hello world!");
	printf("encoded: %s\n", encoded);
	
	// Decode
	char* decoded = B64ToCharPtr(encoded);
	printf("decoded: %s\n", decoded);

	free(encoded);
	free(decoded);
	return 0;
}