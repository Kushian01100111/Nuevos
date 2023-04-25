// Write a function that accepts a string, and returns true if it is in the form of a phone number.
// Assume that any integer from 0-9 in any of the spots will produce a valid phone number.

// Only worry about the following format:
// (123) 456-7890 (don't forget the space after the close parentheses)

// Examples:

// "(123) 456-7890"  => true
// "(1111)555 2345"  => false
// "(098) 123 4567"  => false

function validPhoneNumber(phoneNumber){
    let firstSection = phoneNumber.split(" ")[0].replace(/[()]/g, ""),
        secundSection = phoneNumber.split("-")[0].split(" ")[1],
        treeSection = phoneNumber.split("-")[1];

    if(firstSection && secundSection && treeSection){
        return firstSection.length === 3 && secundSection.length === 3 && treeSection.length === 4 ? true: false;
    }else{
        return false
    }
}


console.log(validPhoneNumber("(123) 456 7890"), true)
