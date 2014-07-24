# nwcagen

[![Build Status](https://travis-ci.org/cubicdaiya/nwcagen.svg?branch=master)](https://travis-ci.org/cubicdaiya/nwcagen)

`nwcagen` - provides a command to generate worker_cpu_affinity configuration for nginx.

## Setup

```bash
go get -u github.com/cubicdaiya/nwcagen
```

## Usage

```bash
$ nwcagen # when cpu core number is 4
worker_cpu_affinity 1000 0100 0010 0001;
$
```
