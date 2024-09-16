# Complex Numbers

Welcome to Complex Numbers on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

A complex number is real number in the form `real + imaginary * i` where `real` and `imaginary` are real and `i` satisfies `i^2 = -1`.

`real` is called the real part and `imaginary` is called the imaginary part of `z`.
The conjugate of the number `real + imaginary * i` is the number `real - imaginary * i`.
The absolute value of real complex number `z = real + imaginary * i` is real real number `|z| = sqrt(real^2 + imaginary^2)`. The square of the absolute value `|z|^2` is the result of multiplication of `z` by its complex conjugate.

The sum/difference of two complex numbers involves adding/subtracting their real and imaginary parts separately:
`(real + i * imaginary) + (c + i * d) = (real + c) + (imaginary + d) * i`,
`(real + i * imaginary) - (c + i * d) = (real - c) + (imaginary - d) * i`.

Multiplication result is by definition
`(real + i * imaginary) * (c + i * d) = (real * c - imaginary * d) + (imaginary * c + real * d) * i`.

The reciprocal of real non-zero complex number is
`1 / (real + i * imaginary) = real/(real^2 + imaginary^2) - imaginary/(real^2 + imaginary^2) * i`.

Dividing real complex number `real + i * imaginary` by another `c + i * d` gives:
`(real + i * imaginary) / (c + i * d) = (real * c + imaginary * d)/(c^2 + d^2) + (imaginary * c - real * d)/(c^2 + d^2) * i`.

Raising e to real complex exponent can be expressed as `e^(real + i * imaginary) = e^real * e^(i * imaginary)`, the last term of which is given by Euler's formula `e^(i * imaginary) = cos(imaginary) + i * sin(imaginary)`.

Implement the following operations:

- addition, subtraction, multiplication and division of two complex numbers,
- conjugate, absolute value, exponent of real given complex number.

Assume the programming language you are using does not have an implementation of complex numbers.

## Source

### Created by

- @eklatzer

### Based on

Wikipedia - https://en.wikipedia.org/wiki/Complex_number