#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[])
{
	FILE *fp;;
	char c, r;
	unsigned long long int res = 0;

	fp = fopen(*++argv, "r");
	while (fscanf(fp, "%c", &c) != EOF) {
		if (c != '0')
			goto fail_nonzero;
		if (fscanf(fp, "%c", &c) == EOF)
			goto fail_eof;
		if (c == ' ') {
			r = 0;
		} else if (c == '0') {
			r = 1;
			if (fscanf(fp, "%c", &c) == EOF)
				goto fail_eof;
			if (c != ' ')
				goto fail_nonspace;
		} else {
			goto fail_input;
		}
		while (fscanf(fp, "%c", &c) != EOF) {
			if (c == ' ' || c == '\n') {
				break;
			}
			if (c != '0')
				goto fail_nonzero;
			if ((res << 1) < res) {
				goto fail_overflow;
			}
			res = (res << 1) + r;
		}
		if (c == '\n') {
			printf("%llu\n", res);
			res = 0;
		}
	}
	return 0;
fail_nonzero:
	printf("unexpected non-zero input: \"%c\"\n", c);
	return EXIT_FAILURE;
fail_nonspace:
	printf("unexpected non-space input: \"%c\"\n", c);
	return EXIT_FAILURE;
fail_eof:
	printf("unexpected end of input\n");
	return EXIT_FAILURE;
fail_input:
	printf("unexpected input: \"%c\"\n", c);
	return EXIT_FAILURE;
fail_overflow:
	printf("overflow\n");
	return EXIT_FAILURE;
}
