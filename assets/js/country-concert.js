document.getElementById("location").classList.add("active")

    // Boutons pour voir ou cacher les cartes
    var showMap = document.getElementsByClassName("show-map")
    var hideMap = document.getElementsByClassName("hide-map")
    var count = 0
    for (let i = 0 ; i <= 51 ; i++) {
      if (showMap[i] != null) {
        showMap[i].addEventListener("click", () => {
          document.getElementsByClassName("info-artist")[i].classList.add("activate")
          showMap[i].classList.remove("activate")
          hideMap[i].classList.add("activate")
        })
        hideMap[i].addEventListener("click", () => {
          document.getElementsByClassName("info-artist")[i].classList.remove("activate")
          showMap[i].classList.add("activate")
          hideMap[i].classList.remove("activate")
        })
      }
      if (document.getElementById(i+1) != null) {
        var x = document.getElementById(i+1).childElementCount;
        document.getElementsByClassName("number")[count].innerHTML = x
        count++
      }
    }