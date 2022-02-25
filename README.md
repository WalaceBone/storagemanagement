# Storage Management

## How to compile

First, make sure that Go is installed on your computer.

In your terminal or IDE.

To build the project:
```
make
```

To run the project:
```
./simulation map
```

## Sources & Packages

You can find the C4 model in the ***componentDiagram*** folder and the project instructions in ***pdfProject***. The rest of the project files are in the root of the project folder.

No external packages were used.

You can modify the ***map.txt*** file as long as you respect the project guidelines.

## strategy used

To find the best path for our Forklift we used [Breadth First Search Algorithm](https://en.wikipedia.org/wiki/Breadth-first_search).

Then we look for the nearest available depot.
