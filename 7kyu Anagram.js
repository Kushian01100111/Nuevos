// An anagram is the result of rearranging the letters of a word to produce a new word (see wikipedia).

// Note: anagrams are case insensitive

// Complete the function to return true if the two arguments given are anagrams of each other; return false otherwise.
// Examples

//     "foefet" is an anagram of "toffee"

//     "Buckethead" is an anagram of "DeathCubeK"



// write the function isAnagram
var isAnagram = function(test, original) {
    const frequencyMap = {};

    const addToMap = (str, map) => {
      for (const char of str) {
        const lowerChar = char.toLowerCase();
        map[lowerChar] = (map[lowerChar] || 0) + 1;
      }
    }
  
    addToMap(original, frequencyMap);
    addToMap(test, frequencyMap);
  
    for (const char in frequencyMap) {
      if (frequencyMap[char] % 2 !== 0) {
        return false;
      }
    }
  
    return true;
  };



  console.log(isAnagram("QDOTpYTGItYsfHaqUP", "YPtQsUIGpOFaqTHYDT"), true)