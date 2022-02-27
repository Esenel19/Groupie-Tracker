document.getElementById("artists").classList.add("active");
var div = document.getElementById('content');
var divs = div.getElementsByClassName('ArtistsName');
var input = document.getElementById("inputArtists");
Reg = "";

input.addEventListener('keyup', function(event) {
    if(this.value != "") {
        reg = new RegExp('.*' + this.value + '.*', 'gi') ;
    for (var i = 0; i < divs.length; i += 1) {
        if(reg.test(divs[i].id)) {
            divs[i].style.display = "block";
        } else {
            divs[i].style.display = "none";
        }
      }
    } else {
        for (var i = 0; i < divs.length; i += 1) {
                divs[i].style.display = "block";
          }
    }
    
}, false);