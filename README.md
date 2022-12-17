# simchem-chemical-reaction-simulator
This is the code repository for the chemical reaction simulator.

To run a simulation of particle, creat a .txt file that describe the whole simulation. Guide on writing such a file could be found in sample.txt
The unit of parameter is not restricted. Units in sample.txt are nanometer and second.

From go/src, run "./go build" from command line.
Then run "./simchem-chemical-reaction-simulator input output" from command line. input is the name of file contains you designated parameters(no .txt suffix needed), 
And output should be the name of generated .gif file.

example: ./simchem-chemical-reaction-simulator sample lotvol
