// Assume "#" is like a backspace in string. This means that string "a#bc#d" actually is "bd"

// Your task is to process a string with "#" symbols.
// Examples

// "abc#d##c"      ==>  "ac"
// "abc##d######"  ==>  ""
// "#######"       ==>  ""
// ""              ==>  ""


function cleanString(str) {
    let stack = [];
  
    for (let i = 0; i < str.length; i++) {
      if (str[i] === "#") {
        stack.pop();
      } else {
        stack.push(str[i]);
      }
    }
  
    return stack.join("");
  }
