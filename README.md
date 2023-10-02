# Mandatory exercise set 1
###### Adam Rose - adjr@itu.dk - Oktober 2, 2023

## Exercise 1
Having Bob's public key $pk_1$, knowing the generator $g$ and the group $p$, as well as the message she wants to send, $m$, all that remains for Alice is to do the following:
- Choose a random number: $y∈Z_p^*$ (`19`)
- Compute the shared secret: $s = pk_1^y \ mod \ p$ (`22`)
- Compute her own public key: $pk_2 = g^y \ mod \ p$ (`25`)
- Encrypt the message: $m_{encr} = s * m$ (`22`)
- Send the encrypted message together with her public key to Bob (`28, 31-33`)

All of this is implemented in the file `nodes/alice.go`. Each of the above points also refer to at what lines of code they are implemented.
Since we are dealing with very big numbers, i used the `math/big` package to do the calculations. To this end, I implemented a helper function in `calc/calc.go` to avoid having to deal with type-casting. 

Finally, Bob is able to decrypt the message, which is done at line `19` in `nodes/bob.go`.
Running the program shows the following output, which shows that the encryption and decryption is successful:
```
=================================================
Assigment part 1, Alice and Bob exchange messages
=================================================
Alice : i'm encrypting the message 2000 with y = 5043
Alice : the encrypted message is 2276000
Alice : my own public key is 2316
Alice : broadcasting message
r1 : received message from Alice with pub_key 2316 and value 2276000
r1 : broadcasting message
Bob : received message from Alice with pub_key 2316 and value 2276000
Bob : decrypting received message
Bob : decrypted message is 2000
```
## Exercise 2
Since the size of the group $p$ is relatively small, it seemed like a good approach to simply brute force Bob's private key. This is done by looping through all the possible values of $x$, which are $0 ≤ x ≤ p$, until an $x$ is found that satisfies the condition $g^x \ mod \ p = pk$.
This is implemented in the file `nodes/eve.go` at lines `18-23`. 

Running the program shows that Eve successfully finds Bob's private key, and is able to decrypt Alice's message:
```
============================================
Assigment part 2, Eve finds Bobs private key
============================================
Eve : Finding Bob's private key
Alice : i'm encrypting the message 2000 with y = 5027
Alice : the encrypted message is 9676000
Alice : my own public key is 1615
Alice : broadcasting message
r1 : received message from Alice with pub_key 1615 and value 9676000    
r1 : broadcasting message
Bob : received message from Alice with pub_key 1615 and value 9676000   
Bob : decrypting received message
Bob : decrypted message is 2000
Eve : Bob's private key is 66
Eve : intercepted message from Alice with pub_key 1615 and value 9676000
Eve : decrypting!
Eve : The decrypted message is 2000
```
## Exercise 3
Using channels I simulated an adversary intercepting and modifying the message. In ElGamal, if the ciphertext is multiplied, it will remain so even after decrypting. This is since the encryption itself happens through multiplication, i.e. $s * m * f$ is the same as $s * (m * f)$, where $s$ is the shared secret, $m$ is the message, and $f$ is the factor that the adversary modifies the message with.

This is implemented in line `14` of `nodes/weave.go`. The output from the program shows that Weave is able to successfully modify the message so it reads '4000' instead of '2000':
```
=======================================================
Assigment part 3, Weave intercepts and modifies message
=======================================================
Alice : i'm encrypting the message 2000 with y = 438   
Alice : the encrypted message is 5088000
Alice : my own public key is 605
Alice : broadcasting message
r1 : received message from Alice with pub_key 605 and value 5088000
r1 : broadcasting message
Weave : received message from Alice with pub_key 605 and value 5088000
Weave : modifying message to have value 10176000
r2 : received message from Alice with pub_key 605 and value 10176000
r2 : broadcasting message
Bob : received message from Alice with pub_key 605 and value 10176000
Bob : decrypting received message
Bob : decrypted message is 4000
```

## How to run the code

The code is run by executing `go run main.go`. 
`time.Sleep()`'s are in place to prevent the program from terminating too early, so it takes perhaps 12 seconds to complete. 