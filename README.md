# go-blockchain

A simple Blockchain written in Golang. This is a learning exercise.

## Add a block

To add a block, with the data `XYZ`:

`go-blockchain add -data "XYZ"`

## View the blockchain

You can print out the entire blockchain:

`go-blockchain print`

## Details

- Bolt DB is used, and is stored in the "db" file in the dir the program is executed
