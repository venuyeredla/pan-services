# Integer reprsentations 

The Bitwise Algorithms is used to perform operations at the bit-level or to manipulate bits in different ways. The bitwise operations are found to be much faster and are sometimes used to improve the efficiency of a program.

Operations with bits are used in Data compression (data is compressed by converting it from one representation to another, to reduce the space) ,Exclusive-Or Encryption (an algorithm to encrypt the data for safety issues). 

In order to encode, decode or compress files we have to extract the data at bit level. Bitwise Operations are faster and closer to the system and sometimes optimise the program to a good level.


Integers can be represented with given bases. ex. Binary-2,  Octal-8,  Decimal-10,  Hex - 16 
Representation is n= Summantion(a[k],pow(b,k)) 0<=k<=r >>  b>1 & n>=0
Decimal :  (965)10= 9 .pow(10,2) +6 .pow(10,1) +5.pow(10,0) 
Binary : (101011111)2 =1·pow(2,8) + 0·pow(2,7) +1·pow(2,6) + 0·pow(2,5) +1.pow(2,4) + 1·pow(2,3) +1·pow(2,2) +1·pow(2,1) +1·pow(2,0) = 351. 

Octal  (7016)8 =7·83 +0·82 +1·8+6=3598. 

Constructing base b expansion.
  Procedure : conver(n, b>1) 
  	q:=n, k:=0
	while q != 0
		ak=q mod b
		q= q div b 
		k++	
	return ak….ao
Algorithms for integer operations.

 Addition
The algorithms for performing operations with integers using their binary expansions are extremely important in computer arithmetic. 
a = (an−1an−2 . . . a1a0)2, b = (bn−1bn−2 . . . b1b0)2, 
a0 + b0 = c0 · 2 + s0,     a1 +b1 +c0 =c1 ·2+s1, 
7 + 6 = 0111 + 0110 = 1101



Bitwise operators: 
  Flipping :   ~8(1000) -> 7(0111) 
  Left shift:  x<<1 <=> x*2. 
  Right Shift: x>>1 <=> x/2. 
Binary - Operates on equal length binary patterns
AND(&)  Two bits should be 1 otherwise 0. Can be used to extract last bit.
OR (|) -  Both bits are 0 result is 0, else 1.
XOR(^) - Exclusive OR. i.e Both doesn’t have to be same. Doing 3 times gives the same number. 
		 
Problems
1. Checking given number is power of 2.
     Flip number bits will be equal to (x-1)    Return  (x && !(x & (x-1)))
2. Getting last bit of the number. 
	Use & operator ex. 7(0111) & 1 ==1 then1 or else 0
2. Count the number of ones in binary representation. 
          While (n) { n =n&(n-1)  count ++

Find the element that appears once.
  1) Using Map with key and count as value. 
  2) 


Find even or odd number.
 * 
 * Extracting last bit can be done with  & operator.
 * Checking given number is power of 2 using only binary operators
 * Count the number of ones in binary expression
 * Multiply a given Integer with 3.5 without using %, /, *. 
 * 
 * Checking given number is power of 2.