# Part 1

Very straight forward. I can just split the input ByteSlice into smaller byteslices by splitting on each comma in the input.

After that just perfrom the algorthim for each input and sum.

"Increase the current value by the ASCII code you just determined.
Set the current value to itself multiplied by 17.
Set the current value to the remainder of dividing itself by 256."

# Part 2

* Basically, if have a '-' e.g. cm-, find the box with cm in and remove it
* If have a '=', replace existing in place or if not existing, put at back e.g. cm=7 will replace cm=9 or if no cm=9 just add at back. 
* How to know which box in the first place? The hash algorithm from before but only on the letters e.g. hash(rn) = 0 for rn=1

256 boxes will be filled.

After this, calculate focusing power.

The focusing power of a single lens is the result of multiplying together:

One plus the box number of the lens in question.
The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
The focal length of the lens.
At the end of the above example, the focusing power of each lens is as follows:

rn: 1 (box 0) * 1 (first slot) * 1 (focal length) = 1
cm: 1 (box 0) * 2 (second slot) * 2 (focal length) = 4
ot: 4 (box 3) * 1 (first slot) * 7 (focal length) = 28
ab: 4 (box 3) * 2 (second slot) * 5 (focal length) = 40
pc: 4 (box 3) * 3 (third slot) * 6 (focal length) = 72

so just iterate thrpugh the boxes to calculate tjis,

## Plan: 

1. Split slice on comma like in part 1
2. For each of the new slices, find what comes before = or -. First time find the result of the alg, add to a map for quicker lookup of HASHMAP alg value. []byte("-") = 45. []byte("=") = 
3. Use a map inside the box number mapping byteSlice to `{focalLength, slot int}` 
4. 