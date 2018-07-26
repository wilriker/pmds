# pmds
Print Move Distance Statistics - Create statistical data on print moves in GCode files

This tool will take an arbitrary amount of GCode files and scans it for print moves. At the end it will display a summary (either for each file or just a total) of the statistical distribution of print moves.

It can be used to determine the "Print move distance" parameter at https://wilriker.github.io/microstep-calculator

Usage
===
Run `pmds -h` to get a short help output
```
Usage of ./pmds:
  -maxMove int
        Maximum distance the longest axis can move in mm (default 300)
  -summary
        Show only summary (this overrules verbose mode)
  -verbose
        Verbose output, i.e. one line for each move
```
