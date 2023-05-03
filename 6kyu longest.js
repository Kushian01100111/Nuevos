// For a given string s find the character c (or C) with longest consecutive repetition and return:

// [c, l]

// where l (or L) is the length of the repetition. If there are two or more characters with the same l return the first in order of appearance.

// For empty string return:

// ["", 0]

// Happy coding! :)



function longestRepetition(s) {
    let repetition = {},
        current = "",
        currentNum = 0;

    for(let i = 0; i < s.length; i++){
        if(!repetition[s[i]] || s[i] !== current){
            repetition[" " + s[i]] = 1
            current = s[i]
        }else if(s[i] === s[i-1]){
            repetition[s[i]] += 1
        }
    }
    

    let result = ["",0],
        currentMax = 0;

    for(const current in repetition){
        let amount =repetition[current];
        if(amount > currentMax){
            currentMax = amount;
            result = [current, currentMax]
        }else if(amount === currentMax){
            continue
        }
    }

    return result;
  }


  function longestRepetition(s) {
    let currentChar = "",
        currentCount = 0,
        maxChar = "",
        maxCount = 0;
    
    for(let i = 0; i < s.length; i++){
      if(s[i] === currentChar){
        currentCount++;
      }else{
        if(currentCount > maxCount){
          maxChar = currentChar;
          maxCount = currentCount;
        }
        currentChar = s[i];
        currentCount = 1;
      }
    }
    
    if(currentCount > maxCount){
      maxChar = currentChar;
      maxCount = currentCount;
    }
    
    return [maxChar, maxCount];
  }
  


  console.log(longestRepetition('nuvgi90gho1a6djzpy74xvmmj0bb4818c69emhtyuj9ntc6e7ke3s5req4xvdfra34z621f7681g39sm2ga9g3a6emq69ux6stwrtfafhv4v6p3a5w9u0b81mxll3b9j652vrgzl3ape2fg1eueu7q3us1x294fhy3el4in9wkpbu0ont8r8xlx10ce0r8subrxwud52l43sbmnkqqcrcvycjqc6cckf4cmvkbgv1tubssdzk4xe9z51oy9u7moitod4rzg5o922odvg9uj0nbdirwfhkmy408t4jtrfe49647yamzh72gf9kvkxxydocqo3w93dhcdew58zeriz9c6llk3q2fj0kssp357bem4ncu2whnyb9g34vomcznjkmb6erklbsv1tjn4yqmtf4nkast3tgx4l0vh9t9coqmwkobmtxi2eow7b07sd6n27utbl4e8due6l5ivi1y8dez8d8qgxx1ih2aqt7qy8108gx9df4i0u'), [ 'm', 2 ])
  console.log(longestRepetition("cbdeuuu900"), ["u", 3] )
