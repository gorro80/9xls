/*
Usage: xlsfs [options] <filename.xlsx>

Xlsfs starts 9P2000 file server
for the specified <filename.xlsx> document.

The root of filesystem consists of dirs with spreadsheet names.
Each spreadsheet represented as a dir.
The following structure depends on [TBD].

Options:
 -addr="localhost:5640": service listen address
*/
package main
