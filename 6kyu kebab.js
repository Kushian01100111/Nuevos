// Modify the kebabize function so that it converts a camel case string into a kebab case.

// "camelsHaveThreeHumps"  -->  "camels-have-three-humps"
// "camelsHave3Humps"  -->  "camels-have-humps"
// "CAMEL"  -->  "c-a-m-e-l"

// Notes:

//     the returned string should only contain lowercase letters



function kebabize(str) {
    let keb = "";
    for (let i=0; i< str.length; i++){
        let temp = str[i];
        if( temp === undefined){
            return keb;
        } else if (temp !== temp.toUpperCase()){
            keb += temp;
        } else if(!isNaN(Number(temp))){
            continue
        }else{
            if(i === 0){
                keb += temp.toLowerCase();
            }else{
                keb += "-" + temp.toLowerCase();
            }
        }
    }
    return keb;
}

// console.log(kebabize('myCamelCasedString'), 'my-camel-cased-string');  
// console.log(kebabize('myCamelHas3Humps'), 'my-camel-has-humps');
console.log(kebabize('SOS'), 's-o-s');
console.log(kebabize('42'), '');
console.log(kebabize('CodeWars'), 'code-wars');
