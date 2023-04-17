// Create a function that takes a positive integer and returns the next bigger number that can be formed by rearranging its digits. For example:

//   12 ==> 21
//  513 ==> 531
// 2017 ==> 2071

// If the digits can't be rearranged to form a bigger number, return -1 (or nil in Swift, None in Rust):

//   9 ==> -1
// 111 ==> -1
// 531 ==> -1


function nextBigger(n){
    let numbers = n.toString().split("").map(Number);

    let rearrangePoint = -1;

    for(let i = numbers.length -2; i >= 0; i--){
        if(numbers[i] < numbers[i+1]){
            rearrangePoint= i;
            break 
        }
    }

    if(rearrangePoint === -1) return rearrangePoint;

    // Find the smallest digit to the right of the pivot that is greater than thenumbers
  let changeNum = rearrangePoint + 1;
  for (let i = changeNum + 1; i < numbers.length; i++) {
    if (numbers[i] > numbers[pivotIdx] && numbers[i] < numbers[nextIdx]) {
      changeNum = i;
    }
  }

    let temp = numbers[rearrangePoint];
    numbers[rearrangePoint] = numbers[changeNum];
    numbers[changeNum] =  temp;

    let sortNums= numbers.slice(rearrangePoint + 1).sort((a,b)=> a - b);

    return parseInt(numbers.slice(0, rearrangePoint + 1).concat(sortNums).join(""))
}



console.log(nextBigger(12),21)
console.log(nextBigger(513),531)
console.log(nextBigger(2017),2071)
console.log(nextBigger(414),441)
console.log(nextBigger(144),414)