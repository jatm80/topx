Write a program, topX - given a number X and an arbitrarily large file that contains individual
numbers on each line (e.g. 500Gb file), will output the largest X numbers, highest first.

# How to use this:

1) Build the application (requirement Go 1.16+)
   `scripts/build.sh`

2) Generate some numbers:
   `scripts/generate.sh`

3) Run the application:
   `./bin/topx <numbers-file> <X value>`
