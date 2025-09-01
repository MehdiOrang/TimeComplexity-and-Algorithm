let input = [1,4,4,6,6,7,8,8]
const removeDuplicate = arr => {
  let temp = 0
  for(let i = 0; i< arr.length; i++){
    if(arr[i] === arr[i+1]  || arr[i] === temp){
      temp = arr[i]
      arr.splice(i, 1)
      arr.push(0)
    }
  }
  return arr
}
 console.log(removeDuplicate(input))
