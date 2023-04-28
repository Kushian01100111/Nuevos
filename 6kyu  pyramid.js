// Write a function that when given a number >= 0, returns an Array of ascending length subarrays.

// pyramid(0) => [ ]
// pyramid(1) => [ [1] ]
// pyramid(2) => [ [1], [1, 1] ]
// pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]

// Note: the subarrays should be filled with 1s


function pyramid(n) {
    let result = [],
        temp = "1".repeat(n).split("").map(Number);
    for(let i = n; i > 0; i--){
        if(i === n) result.push(temp);
        else {
            let temp = result[result.length-1].slice(-i)
            result.unshift(temp)
        }
    }
    return result
  }


  console.log(pyramid(0), [])
  console.log(pyramid(1), [[1]])
  console.log(pyramid(2), [[1], [1, 1]])
  console.log(pyramid(3), [[1], [1, 1], [1, 1, 1]])
