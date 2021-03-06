#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	int a, i = 0, ibs = 32;
	int *ib = malloc(ibs * sizeof(int));
	char c;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%d", &a) != EOF) {
		if (i == ibs) {
			ibs += ibs / 2;
			ib = realloc(ib, ibs * sizeof(int));
		}
		ib[i++] = a;
		if ((c = getc(fp)) == ';') {
			fscanf(fp, "%d", &a);
			int n = 0, j, k;
			for (j = 0; j < i - 1 && 2 * ib[j] < a; j++)
				for (k = i - 1; k > j && ib[j] + ib[k] >= a; k--) {
					i--;
					if (ib[j] + ib[k] == a) {
						if (n++ > 0)
							putchar(';');
						printf("%d,%d", ib[j], ib[k]);
					}
				}
			if (n == 0)
				printf("NULL");
			putchar('\n');
			i = 0;
		}
	}
	free(ib);
	return 0;
}
