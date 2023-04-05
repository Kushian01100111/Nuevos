// Your task, is to create N×N multiplication table, of size provided in parameter.

// For example, when given size is 3:

// 1 2 3
// 2 4 6
// 3 6 9

// For the given example, the return value should be:

// [[1,2,3],[2,4,6],[3,6,9]]

multiplicationTable = function(size) {
    let resultArr = [];
    for(let i = 1 ; i <= size; i++){
      let temp = new Array(size)
      for(let x = 1; x <= temp.length; x++){
        temp[x-1] = i*x
      }
      resultArr.push(temp)
    }
    return resultArr
}

console.log(multiplicationTable(3), [[1,2,3], [2,4,6], [3,6,9]])
