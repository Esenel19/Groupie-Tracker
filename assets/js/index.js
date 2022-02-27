
var img = document.getElementsByClassName("image-central")[0]
function chg_img_group(x) {
    if (x == 0) {
        x = "location"
    } else if (x == 1) {
        x = "group"
    } else if (x == 2) {
        x = "relations"
    }
    img.style.backgroundImage = "url(\"../assets/image/image_index/"+x+"-img.jpg"
    img.style.backgroundRepeat = "no-repeat"
    img.style.backgroundSize = "cover";
    img.style.transition = "0.5s"
}
function reset_img() {
    img.style.backgroundImage = "url(\"../assets/image/image_index/concert-img.jpg"
    img.style.backgroundRepeat = "no-repeat"
    img.style.backgroundSize = "cover";
}