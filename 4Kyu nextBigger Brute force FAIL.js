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

    let rearrangePoint = -1,
        changeNum = 0;

    for(let i = numbers.length -2; i >= 0; i--){
        if(numbers[i] < numbers[i+1]){
            rearrangePoint= i;
            changeNum = i + 1;
            break 
        }
    }

    if(rearrangePoint === -1) return rearrangePoint;


    let temp = numbers[rearrangePoint];
    numbers[rearrangePoint] = numbers[changeNum];
    numbers[changeNum] =  temp;

    return parseInt(numbers.map(n => n.toString()).join(""))
}



console.log(nextBigger(12),21)
console.log(nextBigger(513),531)
console.log(nextBigger(2017),2071)
console.log(nextBigger(414),441)
console.log(nextBigger(144),414)