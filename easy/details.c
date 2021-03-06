#include <stdio.h>

int main(int argc, char *argv[])
{
	FILE *fp;
	char c, p = '\n';
	int m = 10, dots = 0;

	fp = fopen(*++argv, "r");
	while ((c = getc(fp)) != EOF || p != '\n') {
		switch (c) {
		case '\n':
		case EOF:
			if (dots < m)
				m = dots;
			printf("%d\n", m);
			p = '\n';
			m = 10;
			dots = 0;
			break;
		case ',':
			if (dots < m)
				m = dots;
			dots = 0;
			p = c;
			break;
		case '.':
			dots++;
		case 'X':
			p = c;
			break;
		case 'Y':
			if (p == 'X')
				m = 0;
			p = c;
		}
	}
	return 0;
}
