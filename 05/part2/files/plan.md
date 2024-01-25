# Plan

Brute forcing this would be very inefficient. Instead, convert the maps into a set of functions e.g. 

For a <= x0 < b, x1 = f1(x0),

For c <= x0 < d, x1 = f2(x0),

...

Then we can can compose the functions at each layer.

The equations are basically:

* x1 = x0 + a for x0Min, x0Max > 0.

* x2 = x1 + b for specific ranges of x1

* but x1 = x0 + a, so x2 = x1 + a + b for some new x0Min, x0Max.

We can then split the initial range by working backwards and using mins and maxes.

seed-to-soil map:

			50 98 2

			52 50 48

soil-to-fertilizer map:

			0 15 37

			37 52 2

			39 0 15

x0Min = 98, x0Max= 99,  a = -48.             x1 = x0 - 48 for 98 <= x0 <= 99     (1)
x0Min = 50, x0Max = 97, a = 2                x1 = x0 + 2  for 50 <= x0 <= 98    (2)
x0Min = 1,  x0Max = 49, a = 0                x1 = x0      for 0  <= x0 <= 49    (3)

Then for iteration 2 we have

x1Min 15, x1Max 51, a = -15                  x2 = x1 - 15 for 15 <= x1 <= 53     (4)
x1Min 52, x1Max 53, a = -15
x1Min, 0, x1Max 14, a = 39                   x2 = x1 + 39 for 0  <= x1 <  14     (5)
                                             x2 = x1      for 54 <= x1 <= 99        (6)

Working backwards then, we can sub equations 1, 2, and 3 into equation 4, 5 and 6. 

1. For equation 4, we have, subbing x1 = x0 - 15 into f(x1), an equation for x2 in terms of x0:

|      f(x1)      |         f(x0)          |
|-----------------|------------------------|
|     x1          | x0 - 15                |
| x1 - 48         | x0 - 48 - 15 = x0 - 63 |
| x1 + 2          | x0 + 2 - 15 = x0 - 13  |

2. Then, to calculate our new mins and maxes, sub f(x0) into the x1 ranges. For equation 4 in iteration 2, x2 = x1 - 15 for range 15 <= x1 = 53, sub each of the equations from iteration 1 into the range 15 <= x1 <= 53 to find the range in terms of x0. Then combine the original minima and maxima with the new ones.

   * (1) For x2 = x1 - 63 we have:
      * x0min = 98, x0max = 99 from equation 1. 
      * From equation 4, When does x0 give us 15 <= x1 < 54? When 15 <= x0 - 48 < 54 i.e. when 63 <= x0 < 102.
      * Combining, we use max(63,98) to get an overall minimum of 98. Similarly, for the max, min(99,102) => x1=98,x1 = 99

   * (2) x2 = x0 - 13:
      * with x0Min = 50, x0Max = 97 from equation 2. 
      * when does x0 give us 15 <= x1 < 54? when 15 <= x0 + 2  = 53 or when 13 <= x0 <= 51
      * Combining, we get x0Min = 50, x0Max = 51

   * (3) x2 = x0 - 15:
      * with x0Min = 0, x0Max = 49 from equation 3.
      * 15 <= x0 <= 53  when 30 <= x0 < = 68
      * 15 = x1min, 49 = x1Max 


So overall we get:

| x0Min | x0Max |      f(x1)      |         f(x0)          |
|-------|-------|-----------------|------------------------|
|  15   |  49   |     x1          | x0 - 15                |
|  98   |  99   | x1 - 48         | x0 - 48 - 15 = x0 - 63 |
|  50   |  51   | x1 + 2          | x0 + 2 - 15 = x0 - 13  |

3. Repeat steps 1 and 2 but replace equation 4 with the remaining equations in iteration 2.
4. Repeat steps 1 2 and 3 for layers 3, 4, 5, 6 and 7. Will have x6 in terms of x0 only.

After this, intersect the input ranges with the ranges for the fully composed piecewise function of layer 6. The composed functions are **always** of the form x6 = x0 + a, where a is a constant. This means testing only the lowest possible value of in the intersection is necessary. 

We will then be able to see the lowest possible value outputted by the map.

### Algorithm PseudoCode:

1. Iterate through each map and convert into piecewise form

```
seed-to-soil map:
   50 98 2
   52 50 48
```

gives, using 

x2 = x1 + col1 - col2     for col2 <= x1 <= col2 + col3 - 1

```
f(x1) = {
   x2 = x1 - 48:           98 <= x1
   x2 = x1 + 2:            50 <= x1 <= 97
   x2 = x1:                x1 <= 49
}
```

We will represent this using a slice of structs, ordered by x1Min


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

Represent each layer like this.

fx1, fx2, fx3, x4, fx5, fx6

2. Then iterate through, and create fx2x1 = combined version, fx3x2x1, ... fx6x5x4x3x2x1 as described earlier.. Will have f(x6) in terms of x0 only. 


3. for each of the input ranges, intersect with the ranges for f(x6) in terms of x0. A range has values within another range if any of the following true:
   (a) the min starts in the range
   (b) the max starts in the range
   (c) min and max both outside the range
or conversely, if inputMin > maxInput or inputMax < minInput, then we know the sets are completely disjoint.
If there is an intersection, then test the max of the combined mins

