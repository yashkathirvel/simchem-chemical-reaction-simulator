# simchem-chemical-reaction-simulator
This is the code repository for the chemical reaction simulator. The goal of this project is to implement a visualization of Andrew’s stochastic simulation of chemical reactions with spatial resolution and single molecule detail.

## usage 

An example of sample input can be found at sample.txt.

To run a simulation of particle, creat a .txt file that describe the whole simulation. 
Please set surface width, time step, generation of the simulation as well as canvas width, scaling factor, and sampling frequency of the gif.
The parameters should be input as order above and separated by a space after "$" symbol.
surface-width time-step generation scaling-factor canvas-width frequency
example: $200.0 1 1 1 1000 1
$400.0 0.01 2000 1 1000 100

Please declare species (i.e. kinds of molecules) that involved in the Simulation.
Avaliable colors: red,blue,green,yellow,pink,cyan
If a species don't exist at beginning, please put zero at "initial-number"
name  radius  mass  diffusion-rate color initial-number
example: #A 3 1 500 red 100
example: #B 3 1 500 blue 100
#A 3 1 500 red 2000
#B 3 1 500 blue 2000

Please declare reactions you wish to simulate.
Zeroth order reaction starts with a "z>"; unimolecular starts with a "u>", bimolecular reaction starts with "b>". No space after these prefix.
if a single molecule vanishes after a reaction, just leave product blank
example: u>A A A 2 #This means unimolecular reaction A->2A with a reaction-constant=2
example: b>A B B B 40000 #This means bimolecular reaction A+B->2B with a reaction-constant=40000
u>A A A 2
u>B 2
b>A B B B 30000

The unit of parameter is not restricted. Units in sample.txt are nanometer and second.

From go/src, run "./go build" from command line.
Then run "./simchem-chemical-reaction-simulator input output" from command line. input is the name of file contains you designated parameters(no .txt suffix needed), 
And output should be the name of generated .gif file.

example: ./simchem-chemical-reaction-simulator sample lotvol

This will return a given simulation for a chemical reaction.
