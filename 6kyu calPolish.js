// Your job is to create a calculator which evaluates expressions in Reverse Polish notation.

// For example expression 5 1 2 + 4 * + 3 - (which is equivalent to 5 + ((1 + 2) * 4) - 3 in normal notation) should evaluate to 14.

// For your convenience, the input is formatted such that a space is provided between every token.

// Empty expression should evaluate to 0.

// Valid operations are +, -, *, /.

// You may assume that there won't be exceptional situations (like stack underflow or division by zero).

function calc(expr) {
    let nums = expr.split(" ")
        arr = [],
        result = 0;
  
  for(let i=0;i<nums.length; i++){
    let temp = nums[i]
    if(["+","-","*","/"].includes(temp)){
        let se = arr.pop(),
            fi = arr.pop();
        
        switch(temp){
            case "+": result = fi+se; break;
            case "-": result = fi-se; break;
            case "*": result = fi*se; break;
            case "/": result = fi/se; break;
        }
        arr.push(result)
      } else if (nums.length === 1){
        result = Number(temp);
      }else{
        arr.push(Number(temp))
      }
  }
  return result
}

console.log(calc("5 1 2 + 4 * + 3 -"),14)
console.log(calc("3 4 +"),7)
console.log(calc("3 4 -"),-1)   
console.log(calc("3 4 *"),12)   
console.log(calc("3 4 /"),0.75)
console.log(calc("3 4 5 * -"),-17)
console.log(calc("3 4 + 5 *"),35)
console.log(calc("3.5"),3.5)