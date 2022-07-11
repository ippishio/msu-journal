// let button = document.getElementById('but');
// button.addEventListener('click',()=>button.disabled=true);
let check = document.getElementById('pass');
let input = document.getElementsByClassName('butr');
//input.type = "text";


check.addEventListener('change',()=>{
    if(check.checked = true){
        input.setAttribute('type','text')

    }
       

    
})