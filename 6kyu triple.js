// tripledouble(num1,num2)

// which takes numbers num1 and num2 and returns 1 if there is a straight triple of a number at any place in num1 and also a straight double of the same number in num2.

// If this isn't the case, return 0
// Examples

// tripledouble(451999277, 41177722899) == 1 // num1 has straight triple 999s and 
//                                           // num2 has straight double 99s

// tripledouble(1222345, 12345) == 0 // num1 has straight triple 2s but num2 has only a single 2

// tripledouble(12345, 12345) == 0

// tripledouble(666789, 12345667) == 1

function tripledouble(num1, num2) {
    let temp1 = num1.toString().split(""),
        temp2 = num2.toString().split(""),
        result = [false, false],
        curr = "",
        count = 0;
  for(let i = 0; i < temp1.length; i++){
    let temp = curr
    curr = temp1[i]
    if(count === 2){
      count = 0;
      result[0] = [true, temp]
      curr = "";
      break
    }
    else if (temp === curr) count++
  }
  
  for(let i = 0; i <= temp2.length; i++){
    let temp = curr
    curr = temp2[i]
    if(count === 1 && temp !== result[0][1]) count=0
    else if(count === 1 && temp === result[0][1]){
        count = 0;
        result[1] = [true, temp]
        curr = "";
      break
    }
    else if (temp === curr) count++
  }
  
  return result[0][0] && result[1][0] ? 1 : 0
}

console.log(tripledouble(451999277,41177722899),1)
console.log(tripledouble(1222345, 12345),0)
console.log(tripledouble(12345, 12345),0)
console.log(tripledouble(666789, 12345667),1)

