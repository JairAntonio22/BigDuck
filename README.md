
# The BigDuck programming language compiler

BigDuck is language aimed for the developement of mathematical models commonly
used in Machine-Learning and Data Science. Therefore this language will include
integer and floating point arithmetic, vector and matrix operations, and some
basic utilities for reading and writing `.csv` files. All this with the purpose
to make it easier for the user to work within the Machine-Learning and Data
Science fields.

# Example

```
#| Prime sieve algorithm |#
proc main()
    var primes [100]int;
    var k, i, prime_count int;
    var is_prime bool;
{
    primes[0] <- 2;
    primes[1] <- 3;
    prime_count <- 2;

    loop k <- 5; prime_count < 100; k <- k + 2 {
        is_prime <- true;

        loop i <- 0; i < prime_count; i <- i + 1 {
            if mod(k, primes[i]) = 0 {
                is_prime <- false;
                break;
            }
        }

        if is_prime {
            primes[prime_count] <- k;
            prime_count <- prime_count + 1;
        }
    }

    print("The first", n, "primes are");

    loop i <- 0; i < 100; i <- i + 1 {
        print(primes[i]);
    }
}
```
