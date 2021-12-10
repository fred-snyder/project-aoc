# Notes part 2

Life support rating = OxygenGenRating * CO2ScrubberRating

Filtering out values until only one remains

1. Start consider just the first bit
2. Keep only numbers selected by the bit criteria for the type of rating value for which your are searching. Discrad numbers that do not match the criteria.
3. If you only have one number left, stop. THIS IS THE RATING VALUE FOR THE CURRENT SEARCH

## Find oxygen generator rating

Loop over all the numbers: Determine the most common value in the current bit position. (So, 0 or 1)
Only keep the LINES that meet the criterium
Start over with the lines that are left from the previous operation

If 0 and 1 have the same number of occurance. Keep the numbers that are __**1**__ in that position (discard numbers that are 0 in that position). And continue

STOP UNTIL 1 NUMBER IS LEFT

## Find CO2 scrubber rating

Determine the LEAST common value. So least 0, or least 1
Only keep lines that have the result of previous operation at that index
If 0 and 1 are equally commo keep the __**0**__.

STOP UNTIL 1 NUMBER IS LEFT

## Example

Consider the numbers and create a new slice of lines. Remove the lines that do not meet the criterium
Analyse the numbers as described above.

## Solution

Convert the two binary numbers to intigers and multiply them together.

