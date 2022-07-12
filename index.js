const button = document.getElementById('catcreate');
const catholder = document.getElementById('catholder');
button.addEventListener('click', () => {
    var cat = document.createElement("img");
    cat.src = 'cat.gif';
    catholder.appendChild(cat);
    console.log("HUI");
});
