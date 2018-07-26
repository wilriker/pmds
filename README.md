# pmds
Print Move Distance Statistics - Create statistical data on print moves in GCode files

This tool will take an arbitrary amount of GCode files and scans it for print moves. At the end it will display a summary (either for each file or just a total) of the statistical distribution of print moves.

It can be used to determine the "Print move distance" parameter at https://wilriker.github.io/microstep-calculator

Usage
---
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

Sample Output
---
```
$ ./pmds --summary *.gcode
Summary
Shortest Print Move: 0.0009mm
Average Print Move: 1.8779mm
Longest Print Move: 179.5543mm
Percentiles:
  1% of print moves are <= 0.0127mm
  2% of print moves are <= 0.021mm
  3% of print moves are <= 0.0408mm
  4% of print moves are <= 0.0524mm
  5% of print moves are <= 0.0726mm
  6% of print moves are <= 0.0971mm
  7% of print moves are <= 0.1212mm
  8% of print moves are <= 0.1414mm
  9% of print moves are <= 0.157mm
 10% of print moves are <= 0.1745mm
 15% of print moves are <= 0.2047mm
 20% of print moves are <= 0.2254mm
 25% of print moves are <= 0.2614mm
 30% of print moves are <= 0.286mm
 35% of print moves are <= 0.3179mm
 40% of print moves are <= 0.3431mm
 45% of print moves are <= 0.3779mm
 50% of print moves are <= 0.4254mm
 55% of print moves are <= 0.4801mm
 60% of print moves are <= 0.5532mm
 65% of print moves are <= 0.6434mm
 70% of print moves are <= 0.7581mm
 75% of print moves are <= 0.9123mm
 80% of print moves are <= 1.0678mm
 85% of print moves are <= 1.6035mm
 90% of print moves are <= 2.9938mm
 95% of print moves are <= 5.5499mm
100% of print moves are <= 179.5543mm
```
