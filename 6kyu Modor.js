// Description

// Middle Earth is about to go to war. The forces of good will have many battles with the forces of evil. Different races will certainly be involved. Each race has a certain worth when battling against others. On the side of good we have the following races, with their associated worth:

//     Hobbits: 1
//     Men: 2
//     Elves: 3
//     Dwarves: 3
//     Eagles: 4
//     Wizards: 10

// On the side of evil we have:

//     Orcs: 1
//     Men: 2
//     Wargs: 2
//     Goblins: 2
//     Uruk Hai: 3
//     Trolls: 5
//     Wizards: 10

// Although weather, location, supplies and valor play a part in any battle, if you add up the worth of the side of good and compare it with the worth of the side of evil, the side with the larger worth will tend to win.

// Thus, given the count of each of the races on the side of good, followed by the count of each of the races on the side of evil, determine which side wins.
// Input:

// The function will be given two parameters. Each parameter will be a string of multiple integers separated by a single space. Each string will contain the count of each race on the side of good and evil.

// The first parameter will contain the count of each race on the side of good in the following order:

//     Hobbits, Men, Elves, Dwarves, Eagles, Wizards.

// The second parameter will contain the count of each race on the side of evil in the following order:

//     Orcs, Men, Wargs, Goblins, Uruk Hai, Trolls, Wizards.

// All values are non-negative integers. The resulting sum of the worth for each side will not exceed the limit of a 32-bit integer.
// Output:

// Return "Battle Result: Good triumphs over Evil" if good wins, "Battle Result: Evil eradicates all trace of Good" if evil wins, or "Battle Result: No victor on this battle field" if it ends in a tie.


function goodVsEvil(good, evil){
    let temp = good.split(" ").map(n => Number(n)),
        temp2 = evil.split(" ").map(n => Number(n));
    let goodCount = 0,
        evilCount = 0;
    for(let i = 0; i < temp.length; i++){
        if(i === 0) goodCount += 1*temp[i]
        if(i === 1) goodCount += 2*temp[i]
        if(i === 2) goodCount += 3*temp[i]
        if(i === 3) goodCount += 3*temp[i]
        if(i === 4) goodCount += 4*temp[i]
        if(i === 5) goodCount += 10*temp[i]
    }

    for(let i = 0; i < temp2.length; i++){
        if(i === 0) evilCount += 1*temp2[i]
        if(i === 1) evilCount += 2*temp2[i]
        if(i === 2) evilCount += 2*temp2[i]
        if(i === 3) evilCount += 2*temp2[i]
        if(i === 4) evilCount += 3*temp2[i]
        if(i === 5) evilCount += 5*temp2[i]
        if(i === 6) evilCount += 10*temp2[i]
    }

    if(goodCount > evilCount) return 'Battle Result: Good triumphs over Evil'
    else if(evilCount > goodCount) return 'Battle Result: Evil eradicates all trace of Good'
    else return 'Battle Result: No victor on this battle field'
}


console.log(goodVsEvil('0 0 0 0 0 10', '0 1 1 1 1 0 0'), 'Battle Result: Evil eradicates all trace of Good')
