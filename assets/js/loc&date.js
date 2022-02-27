var divs = document.getElementsByClassName("artist-block");
var prev = document.querySelector("#prev");
var next = document.querySelector("#next");
var regex = /(?<=#).*/gm; //regex
var all_inputs = window.location.href;
var input = all_inputs.match(regex);
var mess = document.querySelector(".message");

if (divs.length >= 1) {
  if (input != null) {
    var index = input - 1; //0 et 51
    prev.style.display = "none";
    next.style.display = "none";
  } else {
    var index = 0; //0 et 51
  }

  if (divs.length == 1) {
    prev.style.display = "none";
    next.style.display = "none";
  }

  var indexbase = index;
  divs[index].style.display = "initial";

  prev.addEventListener(
    "click",
    function (event) {
      if (divs.length != 1) {
        divs[index].style.display = "none";
        if (index == 0) {
          index = divs.length - 1;
        } else {
          index--;
        }
        divs[index].style.display = "block";
      }
    },
    false
  );

  next.addEventListener(
    "click",
    function (event) {
      if (divs.length != 1) {
        divs[index].style.display = "none";
        if (index == divs.length - 1) {
          index = 0;
        } else {
          index++;
        }
        divs[index].style.display = "block";
      }
    },
    false
  );
} else {
  mess.style.display = "flex";
  prev.style.display = "none";
  next.style.display = "none";
  document
    .querySelector(".page-locAndDate")
    .setAttribute("style", "min-height: auto;");
}
