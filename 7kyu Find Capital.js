// Instructions

// Write a function that takes a single string (word) as argument. The function must return an ordered list containing the indexes of all capital letters in the string.
// Example (Input --> Output)

// "CodEWaRs" --> [0,3,4,6]



var capitals = function (word) {
	var indexes = [];
  for(let i=0; i < word.length; i++){
    if(word.charAt(i) === word.charAt(i).toUpperCase()) indexes.push(i)
  }
  return indexes
};