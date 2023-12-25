# Plan

Brute forcing this would be very inefficient.

instead, convert the maps into a set of functions e.g. 

for a <= x < b, y = f1(x)
for c <= x < d, y = f2(x)

then can compose the functions

my equations at map 1 are basically

y1 = x1 +a1 for specific ranges of x where x is the input.

then y2 = x2 +a2 for specific ranges of x2

but x2 = x1 + a1

so y2 = x1 + a1 + a2 

Then if required i can split the initial range by working backwards and using mins and maxes

i can start out with map 1 as 3 different structs for the below:

seed-to-soil map:
			50 98 2
			52 50 48

soil-to-fertilizer map:
			0 15 37
			37 52 2
			39 0 15

x1Min = 98, x1Max= 99,  a = -48.             x2 = x1 - 48 for 98 <= x1          (1)
x1Min = 50, x1Max = 97, a = 2                x2 = x1 + 2  for 50 <= x1 <  98    (2)
x1Min = 1,  x1Max = 49, a = 0                x2 = x1      for 0  <  x1 <= 49    (3)

Then for iteration 2 i have

x2Min 15, x2Max 51, a = -15                  x3 = x2 - 15 for 15 <= x2 < 54
x2Min 52, x2Max 53, a = -15
x2Min, 0, x2Max 14, a = 39                   x3 = x2 + 39 for 1  <= x2 < 15 
                                             x3 = x2      for 55 <= x2 
Working backwards then

we want to know how we get the boundary inputs for the second equation OUT OF the first equation

15 <= x2 < 54 when:

* 15 < x1 <= 49 (x2 = x1) (15 to 49)     => x3 = x1 - 15 for  for x1 between 15 and 49
* 98 <= x1 < 100 (x2 = x1 - 48) (50,51)  => x3 = x1 - 63 fr x1 = 50, 51
* 50, 51, (x2 = x1 + 2) (52,53)          => x3 = x1 - 13 for x1 - 52, 53

A method for this?



for equation in iteration 2, e.g. first one is x3 = x2 - 15 for 15 <= x2 < 54, sub each of the equations from iteration 2 and then find x1 for which 15 <= x2 < 54
The three equations for getting x2, combined with x3 = x2 - 15 


give us

seed-to-soil map:
   50 98 2
   52 50 48

x1Min = 98, x1Max= 99,  a = -48.             x2 = x1 - 48 for 98 <= x1          (1)
x1Min = 50, x1Max = 97, a = 2                x2 = x1 + 2  for 50 <= x1 <  98    (2)
x1Min = 1,  x1Max = 49, a = 0                x2 = x1      for 0  <  x1 <= 49    (3)

(1) x3 = x1 - 63 with x1min = 98, x1max = 99 from equation 1: when does x1 give us 15 <=x2 < 54? when 15 <= x1 -48 < 54 or when 63 <= x1 < 102 => max(63,98), min(99,102) => x1=98,x1 = 99
(2) x3 = x1 - 13 with x1Min = 50, x1Max = 97 from equation 2. when does x1 give us 15 <= x2 < 54? when 15 <= x1 + 2  < 54 = when 13 <= x1 < 52
(3) x3 = x1 - 15 with x1Min = 1, x1Max = 49.  15 <= x1 < 54  when 30 <= x1 <  69 -->        15 = x1min, 49 = x1Max 

so we get 15-49 included, 98, 99, 63 <= x1 <97

Therefore: Algorithm is

If we define f(x1): x2 = x1 + a for  x1Min < x1 <= x1Max 

then f(x2): x3 = x2 + b for x2Min < x2 <= x2Max

we: 

if fm is a single part of the piecewise function fM, and fn of fN with m > n,

1. Subbed each f(x1) into each f(x2)

```go
for _, fm := range fM {
   for _, fn := range fN {
      // combined = fn + a + b 
   }
}
```

2. subbed f(x1) into minima/maxima

```go
for _, fm := range fM {
   for _, fn := range fN {
      // combined = fn + a + b 
      // min < fn + a <= max
   }
}
```

3. we now have 2 max and min bound by initial x1 minima/maxima (x1Min and x2Min)and the new boundaries. f(x2) = f(x1) + a + b for max(x2BoundaryMinInTermsOfx1, x1Min) < x1 <= min(x2BoundaryMaxInTermsOfx1, x1Max)

```go
for _, fm := range fM {
   for _, fn := range fN {
      // combined = fn + a + b 
      // max(x2BoundaryMinInTermsOfx1, x1Min) < x1 <= min(x2BoundaryMaxInTermsOfx1, x1Max)
   }
}
```

### Plan:

1. Iterate through each piecewise equation and convert into piecewise form

```
seed-to-soil map:
   50 98 2
   52 50 48
```

gives, using 

x2 = x1 + col1 - col2     for col2 <= x1 <= col2 + col3 - 1

f(x1) = {
   x2 = x1 - 48:           98 <= x1
   x2 = x1 + 2:            50 <= x1 <= 97
   x2 = x1:                x1 <= 49
}

can represent this using a slice of structs, ordered by x1Min


```
{
   transform: 0
   x1Min: nil
   x1Max: 49
}
{
   transform: 2
   x1Min: 50
   x1Max: 97
}
{
   transform: -48
   x1Min: 98
   x1Max: nil
},
```


for 0 transform does it need adding? could say if range not found after search assume no map and add into new arr on iteration.
Represent each layer like this.

fx1, fx2, fx3, x4, fx5, fx6

2. Then iterate through, and create fx2x1 = combined version, fx3x2x1, ... fx6x5x4x3x2x1. Dynamic naming? or could just use slices indexing

3. Order fx6x5x4x3x2x1 by x1Min + transform and for each min get the number closest to it from the inputs.





