// Pair of gloves
// Winter is coming, you must prepare your ski holidays. The objective of this kata is to determine the number of pair of gloves you can constitute from the gloves you have in your drawer.

// Given an array describing the color of each glove, return the number of pairs you can constitute, assuming that only gloves of the same color can form pairs.

// Examples:
// input = ["red", "green", "red", "blue", "blue"]
// result = 2 (1 red pair + 1 blue pair)

// input = ["red", "red", "red", "red", "red", "red"]
// result = 3 (3 red pairs)


    function numberOfPairs(gloves){
        let numOfGloves = {},
            pairs = 0;
    gloves.forEach((n)=>{
        if(numOfGloves[n]) {numOfGloves[n] = numOfGloves[n]+1;}
        else numOfGloves[n]= numOfGloves[n]||1;
    })
    
    for(const [color, amount] of Object.entries(numOfGloves)){
        if(amount % 2=== 0){
        pairs += (amount/2)
        }else if(amount % 2 === 1 && amount > 1){
            pairs += ((amount-1)/2)
        }
    }
    
    return pairs
    }

 console.log(numberOfPairs(['red', 'red', 'red', 'red', 'red', 'red']), 3);
    console.log(numberOfPairs(['red', 'green', 'blue', 'blue', 'blue']), 1);
    console.log(numberOfPairs(['gray', 'black', 'purple', 'purple', 'gray', 'black']), 3);