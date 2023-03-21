// Given a lowercase string that has alphabetic characters only and no spaces, return the highest value of consonant substrings. Consonants are any letters of the alphabet except "aeiou".

// We shall assign the following values: a = 1, b = 2, c = 3, .... z = 26.

// For example, for the word "zodiacs", let's cross out the vowels. We get: "z o d ia cs"

// -- The consonant substrings are: "z", "d" and "cs" and the values are z = 26, d = 4 and cs = 3 + 19 = 22. The highest is 26.
// solve("zodiacs") = 26

// For the word "strength", solve("strength") = 57
// -- The consonant substrings are: "str" and "ngth" with values "str" = 19 + 20 + 18 = 57 and "ngth" = 14 + 7 + 20 + 8 = 49. The highest is 57.

// For C: do not mutate input.

// More examples in test cases. Good luck!

// If you like this Kata, please try

function solve(s) {
    let obj = {"a": 1,"b": 2,"c": 3,"d": 4,"e": 5, "f": 6,"g": 7,"h": 8,"i": 9,"j": 10,
              "k": 11,"l": 12,"m": 13,"n": 14,"o": 15,"p": 16,"q": 17,"r": 18,"s": 19,
              "t": 20,"u": 21,"v": 22,"w": 23,"x": 24,"y": 25,"z": 26},
        changeString = s.match(/[^aeiou]+/g),
        result = [];
    
    for(const n of changeString){
        let push = 0
            for(const temp of n){
                push += obj[temp]
            }
        result.push(push)
    }
    
    return Math.max(...result);
  };


  console.log(solve("strength"),57)