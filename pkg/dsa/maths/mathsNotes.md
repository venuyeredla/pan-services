## Well ordering property - Every nonempty set of positive integers has a least element. 
Progressions.  
    Sigma notation for Summation,   Pi notation for Product of the sequence numbers.
	Arithmetic A[n+1]= 2*n+d   -  Focus on difference. Sn =1/2(2*a+n*d)
	Geometric A[n]=a * pow(r,n) = a pow(r,0), a pow(r,1) …   Focus on ratios	= Sig{0,n}=a(pow(r,n+1) - 1)/r-1

Factorial are useful for defining Binomial coefficients. fact(0)=1
Binomial(m,k) =m!/(k!* (m-k)!)  => Binomial(m,0)= Binomial(m,m)=1 
					=> Binomial(m,k)= Binomial(m,m-k)

Binomial Theorem :: Let x, y variables and n a positive integer. 
  pow(x+y,n)= Sig{i=0,n} Binomial(n,i) * pow(x,n-i) * pow(y,i)


## Divisibility
  b=a*q+r  Where b is dividend, a is divisor , q- quotient , r- remainder. 0< r< a

 	If r==0 then a divides b written as a | b. 

Some propositions :    a | b , b|c then a | c.   b=a * q1  c=b * q2 => c = a * q1 * q2 


Representation of numbers - Refers Bitwise algorithms 
Every positive integer may be represented as the sum of distinct powers of two. 

Prime numbers 
Numbers of two types, Prime numbers and composite numbers.
A prime is a positive integer greater than 1 that is divisible by no positive integers other than I and itself. 
Ex : n = 2,3,5, 7 ….     n>1    only 1 |n or n | n
A positive integer which is not prime, and which is not equal to l, is called composite. 
Every positive integer greater than one has a prime divisor.
If n is a composite integer, then n has a prime factor not exceeding.. sqrt(n).
Every even positive integer greater than two can be written as the sum of two primes. 

The Fundamental Theorem of Arithmetic - Every positive integer can be written uniquely as a product of primes, with the prime factors in the product written in order of nondecreasing size. 
 240 = 2*2*2*3*5.       a= pow(p1, q1) * pow(p2,q2)…..   b= pow(p1, r1) * pow(p2,r2)…..
gcd(a,b) =pow(p1, min(q1,r1)) * pow(p2, min(q2,r2))
lcm(a,b) =pow(p1, max(q1,r1)) * pow(p2, max(q2,r2))

Finding prime numbers. Consequently divide n by the primes 2,3,5….   sqrt(n) 

GCD - The greatest common divisor of two integers a and b, that are not both zero is the largest integer which divides both a and b.    x | a  and x | b
    
LCM -The least common multiple of two positive integers a and b is the smallest positive integer that is divisible by a and b.   
          Smallest integer x such that. a | x  and b | x 
 a*b = gcd(a,b) * lcm(a,b)

The integers a and b are called relatively prime if a and b have gcd(a, b)=1l. 
If a and b are integers, then a linear combination of a and b is a sum of the form ma+nb, where both m and, n are integers. 
Euclidean Algorithm for GCD.  
    If a > b then   gcd(a,b)=gcd(b,r)    a=b.q+r  

Linear Diophantine equations 
	20x + 50y= 510       i.e solving ax+by=c   This equation has solution only when gcd(a,b)=d  and d | c


Congruences
If a and b are integers, we say that a is congruent to b modulo m if m l (a-b). 

Function.
	Let A and B be nonempty sets. A function f from A to B is an assignment of exactly one element of B to each element of A. We write f(a) = b if b is the unique element of B assigned by the function f to the element a of A. If f is a function from A to B, we write f : A → B. 
Functions are sometimes also called mappings or transformations.

Relation.
Graph 




Functions
Functions arise when a quantity y depends on another one called x. => y=f(x)
Representations:
Verbally
Numerical table of values.
Visually by graph. Graph is set of order pairs. (Domain ,Codomain) {(x,y) | x Of D, y of CD}
 Piecewise functions 

Catalog of essential functions
1. Constant : f(n)= c1+c2 
2. Logarithmic : f(x)=log(x, b), Inverse of exponential, ln-e natural , log-2, log-10  
3. Linear : y=mx+c
4. Super-linear :  f(n) = O(n * log(n))
5. Quadratic (Power) : f(x)=a*pow(x,2)+b.x+c      f(n) = O(pow(n,2)).          
6. Cubic function (Power): f(x)= a * pow(x, 3)+ b * pow(x,2)    => f(n) = O ( pow(n, 3))
7. Exponential (Exponent is variable) : f(x)=pow(a,x)   => Subset is exponent : f(n) = pow (2,x)
8. Factorial : f(n) =n!  => permutations.
9. Polynomial : P(x) = [0,n]sum( a[n] pow(x,n) 
10. Rational : f(x)=P(x)/Q(x) ratio of polynomials 
11. Algebraic : f(x)=sqrt(2*x+1)
12 Trigonometric : sin(x), cos(x)




# Probability 
is helpful in measuring uncertainty. 

Random variable is variable which takes one of possible values randomly. 
Discrete random variable : Occupies discrete set of states. 
 Probability mass function gives the probability of random variable taking a given possible value. 
 	i.e it maps a state/value to it probability. 
          Bernoulli distribution. Multinouli distribution 

Continuous random variable : Takes continuous values called real values.
	
   Probability density function is used describe how the probabilities are distributed when dealing 
Continuous random variables.  
	Gaussian distribution 

Numerical computation :
	Problems arise during rounding of real numbers.
Underflow : Occurs when number near to zero rounded to zero.
Overflow : Numbers with large values are approximated to infinity.
 
Optimization refers to task of minimizing or maximizing of a function f(x) by altering x.
