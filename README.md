
# The BigDuck programming language compiler

The BigDuck programming language is aimed for the developement mathematical and
scientific computations, numerical methods, and some basic statistics. Therefore
this language includes integer and floating point arithmetic, trigonometric and
trascendental functions, other commonly used math operations, and vector
operations to support some statistical functions.


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

# Dependencies

- Go
- ANTLR

# Instalation

To install the compiler run the following command.

```
make install
```

This will compile the compiler and copy it into the `/usr/local/bin/` directory. 
