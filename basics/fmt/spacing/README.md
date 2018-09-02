
## Spacing
When you print out byte slices using the %x verb it comes out as one giant string of hex numbers. You can delimit the bytes with a space by using the space flag (‘ ‘).

For example, formatting []byte{1,2,3,4} with and without the space flag:
```
%x  ➡ "01020304"
% x ➡ "01 02 03 04"
```

## Width & precision
We can make formatting more useful by adding various flags to the verb. This is especially important for floating-point numbers where you typically need to round them to a specific number of decimal places.

The precision can be specified by adding a period and a number after the % sign. For example, we can use %.2f to specify two decimal places of precision so formatting 100.567 would print as “100.57”. Note that the second decimal place is rounded.

The width specifies the total number of characters your formatted string will take up. If your formatted value is less than width then it will pad with spaces. This is useful when you’re printing tabular data and you want fields to line up. For example, we can add to our previous format and set the width to 8 by adding the number before the decimal place: %8.2f. Printing 100.567 with this format will return “••100.57” (where • is a space).

We can map this out in a table to show how it works for various widths and precisions:
```
%8.0f ➡ "     101"
%8.1f ➡ "   100.6"
%8.2f ➡ "  100.57"
%8.3f ➡ " 100.567"
```

## Left alignment
In our previous example our values were right-aligned. This works well for financial applications where you may want the decimal places lined up on the right. However, if you want to left-align your fields you can use the “-” flag:
```
%-8.0f ➡ "101     "
%-8.1f ➡ "100.6   "
%-8.2f ➡ "100.57  "
%-8.3f ➡ "100.567 "
```

## Zero padding
Sometimes you want to pad using zeros instead of spaces. For instance, you may need to generate fixed-width strings from an number. We can use the zero (‘0’) flag to do this. Printing the number 123 with an 8-byte width and padded with zeros looks like this:
```
%08d ➡ "00000123"
```